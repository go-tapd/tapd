package webhook

import (
	"context"
	"errors"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// Dispatcher is a dispatcher for webhook events.
type Dispatcher struct {
	storyCreateListeners []StoryCreateListener
	storyUpdateListeners []StoryUpdateListener
	bugCreateListeners   []BugCreateListener
}

type Option func(*Dispatcher)

func WithRegisters(listeners ...any) Option {
	return func(d *Dispatcher) {
		d.Registers(listeners...)
	}
}

// NewDispatcher returns a new Dispatcher instance.
func NewDispatcher(opts ...Option) *Dispatcher {
	dispatcher := &Dispatcher{}
	for _, opt := range opts {
		opt(dispatcher)
	}
	return dispatcher
}

func (d *Dispatcher) Registers(listeners ...any) {
	for _, listener := range listeners {
		if l, ok := listener.(StoryCreateListener); ok {
			d.RegisterStoryCreateListener(l)
		}

		if l, ok := listener.(StoryUpdateListener); ok {
			d.RegisterStoryUpdateListener(l)
		}

		if l, ok := listener.(BugCreateListener); ok {
			d.RegisterBugCreateListener(l)
		}

		// todo: add other listeners
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, event any) error {
	switch e := event.(type) {
	case *StoryCreateEvent:
		return d.processStoryCreate(ctx, e)
	case *StoryUpdateEvent:
		return d.processStoryUpdate(ctx, e)
	case *BugCreateEvent:
		return d.processBugCreate(ctx, e)
	default:
		return errors.New("tapd: webhook dispatcher unsupported event")
	}
}

func (d *Dispatcher) DispatchPayload(ctx context.Context, payload []byte) error {
	_, event, err := ParseWebhookEvent(payload)
	if err != nil {
		return err
	}
	return d.Dispatch(ctx, event)
}

func (d *Dispatcher) DispatchRequest(req *http.Request) error {
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return d.DispatchPayload(req.Context(), payload)
}

func (d *Dispatcher) RegisterStoryCreateListener(listeners ...StoryCreateListener) {
	d.storyCreateListeners = append(d.storyCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryUpdateListener(listeners ...StoryUpdateListener) {
	d.storyUpdateListeners = append(d.storyUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCreateListener(listeners ...BugCreateListener) {
	d.bugCreateListeners = append(d.bugCreateListeners, listeners...)
}

func processEvent[L, E any](ctx context.Context, listeners []L, event E, fn func(L, E) error) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range listeners {
		listener := listener
		eg.Go(func() error {
			return fn(listener, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processStoryCreate(ctx context.Context, event *StoryCreateEvent) error {
	return processEvent(ctx, d.storyCreateListeners, event, func(listener StoryCreateListener, event *StoryCreateEvent) error {
		return listener.OnStoryCreate(ctx, event)
	})
}

func (d *Dispatcher) processStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error {
	return processEvent(ctx, d.storyUpdateListeners, event, func(listener StoryUpdateListener, event *StoryUpdateEvent) error {
		return listener.OnStoryUpdate(ctx, event)
	})
}

func (d *Dispatcher) processBugCreate(ctx context.Context, event *BugCreateEvent) error {
	return processEvent(ctx, d.bugCreateListeners, event, func(listener BugCreateListener, event *BugCreateEvent) error {
		return listener.OnBugCreate(ctx, event)
	})
}
