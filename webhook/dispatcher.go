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
	storyCreateListeners      []StoryCreateListener
	storyUpdateListeners      []StoryUpdateListener
	taskUpdateListeners       []TaskUpdateListener
	storyCommentAddListeners  []StoryCommentAddListener
	bugCreateListeners        []BugCreateListener
	bugUpdateListeners        []BugUpdateListener
	bugCommentAddListeners    []BugCommentAddListener
	bugCommentUpdateListeners []BugCommentUpdateListener
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

		if l, ok := listener.(TaskUpdateListener); ok {
			d.RegisterTaskUpdateListener(l)
		}

		if l, ok := listener.(StoryCommentAddListener); ok {
			d.RegisterStoryCommentAddListener(l)
		}

		if l, ok := listener.(BugCreateListener); ok {
			d.RegisterBugCreateListener(l)
		}

		if l, ok := listener.(BugUpdateListener); ok {
			d.RegisterBugUpdateListener(l)
		}

		if l, ok := listener.(BugCommentAddListener); ok {
			d.RegisterBugCommentAddListener(l)
		}

		if l, ok := listener.(BugCommentUpdateListener); ok {
			d.RegisterBugCommentUpdateListener(l)
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
	case *TaskUpdateEvent:
		return d.processTaskUpdate(ctx, e)
	case *StoryCommentAddEvent:
		return d.processStoryCommentAdd(ctx, e)
	case *BugCreateEvent:
		return d.processBugCreate(ctx, e)
	case *BugUpdateEvent:
		return d.processBugUpdate(ctx, e)
	case *BugCommentAddEvent:
		return d.processBugCommentAdd(ctx, e)
	case *BugCommentUpdateEvent:
		return d.processBugCommentUpdate(ctx, e)
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

func (d *Dispatcher) RegisterTaskUpdateListener(listeners ...TaskUpdateListener) {
	d.taskUpdateListeners = append(d.taskUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryCommentAddListener(listeners ...StoryCommentAddListener) {
	d.storyCommentAddListeners = append(d.storyCommentAddListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCreateListener(listeners ...BugCreateListener) {
	d.bugCreateListeners = append(d.bugCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugUpdateListener(listeners ...BugUpdateListener) {
	d.bugUpdateListeners = append(d.bugUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCommentAddListener(listeners ...BugCommentAddListener) {
	d.bugCommentAddListeners = append(d.bugCommentAddListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCommentUpdateListener(listeners ...BugCommentUpdateListener) {
	d.bugCommentUpdateListeners = append(d.bugCommentUpdateListeners, listeners...)
}

func (d *Dispatcher) processStoryCreate(ctx context.Context, event *StoryCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyCreateListeners {
		eg.Go(func() error {
			return listener.OnStoryCreate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyUpdateListeners {
		eg.Go(func() error {
			return listener.OnStoryUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processTaskUpdate(ctx context.Context, event *TaskUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskUpdateListeners {
		eg.Go(func() error {
			return listener.OnTaskUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processStoryCommentAdd(ctx context.Context, event *StoryCommentAddEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyCommentAddListeners {
		eg.Go(func() error {
			return listener.OnStoryCommentAdd(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processBugCreate(ctx context.Context, event *BugCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugCreateListeners {
		eg.Go(func() error {
			return listener.OnBugCreate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processBugUpdate(ctx context.Context, event *BugUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugUpdateListeners {
		eg.Go(func() error {
			return listener.OnBugUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processBugCommentAdd(ctx context.Context, event *BugCommentAddEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugCommentAddListeners {
		eg.Go(func() error {
			return listener.OnBugCommentAdd(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processBugCommentUpdate(ctx context.Context, event *BugCommentUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugCommentUpdateListeners {
		eg.Go(func() error {
			return listener.OnBugCommentUpdate(ctx, event)
		})
	}
	return eg.Wait()
}
