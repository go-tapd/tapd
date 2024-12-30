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
	// 需求/任务/缺陷类
	storyCreateListeners []StoryCreateListener
	storyUpdateListeners []StoryUpdateListener
	storyDeleteListeners []StoryDeleteListener
	taskCreateListeners  []TaskCreateListener
	taskUpdateListeners  []TaskUpdateListener
	taskDeleteListeners  []TaskDeleteListener
	bugCreateListeners   []BugCreateListener
	bugUpdateListeners   []BugUpdateListener
	bugDeleteListeners   []BugDeleteListener

	// 评论类：需求/任务/缺陷
	storyCommentAddListeners    []StoryCommentAddListener
	storyCommentUpdateListeners []StoryCommentUpdateListener
	storyCommentDeleteListeners []StoryCommentDeleteListener
	taskCommentAddListeners     []TaskCommentAddListener
	taskCommentUpdateListeners  []TaskCommentUpdateListener
	taskCommentDeleteListeners  []TaskCommentDeleteListener
	bugCommentAddListeners      []BugCommentAddListener
	bugCommentUpdateListeners   []BugCommentUpdateListener
	bugCommentDeleteListeners   []BugCommentDeleteListener

	// 迭代
	iterationCreateListeners []IterationCreateListener
	iterationUpdateListeners []IterationUpdateListener
	iterationDeleteListeners []IterationDeleteListener
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

		if l, ok := listener.(StoryDeleteListener); ok {
			d.RegisterStoryDeleteListener(l)
		}

		if l, ok := listener.(TaskCreateListener); ok {
			d.RegisterTaskCreateListener(l)
		}

		if l, ok := listener.(TaskUpdateListener); ok {
			d.RegisterTaskUpdateListener(l)
		}

		if l, ok := listener.(TaskDeleteListener); ok {
			d.RegisterTaskDeleteListener(l)
		}

		if l, ok := listener.(BugCreateListener); ok {
			d.RegisterBugCreateListener(l)
		}

		if l, ok := listener.(BugUpdateListener); ok {
			d.RegisterBugUpdateListener(l)
		}

		if l, ok := listener.(BugDeleteListener); ok {
			d.RegisterBugDeleteListener(l)
		}

		if l, ok := listener.(StoryCommentAddListener); ok {
			d.RegisterStoryCommentAddListener(l)
		}

		if l, ok := listener.(StoryCommentUpdateListener); ok {
			d.RegisterStoryCommentUpdateListener(l)
		}

		if l, ok := listener.(StoryCommentDeleteListener); ok {
			d.RegisterStoryCommentDeleteListener(l)
		}

		if l, ok := listener.(TaskCommentAddListener); ok {
			d.RegisterTaskCommentAddListener(l)
		}

		if l, ok := listener.(TaskCommentUpdateListener); ok {
			d.RegisterTaskCommentUpdateListener(l)
		}

		if l, ok := listener.(TaskCommentDeleteListener); ok {
			d.RegisterTaskCommentDeleteListener(l)
		}

		if l, ok := listener.(BugCommentAddListener); ok {
			d.RegisterBugCommentAddListener(l)
		}

		if l, ok := listener.(BugCommentUpdateListener); ok {
			d.RegisterBugCommentUpdateListener(l)
		}

		if l, ok := listener.(BugCommentDeleteListener); ok {
			d.RegisterBugCommentDeleteListener(l)
		}

		if l, ok := listener.(IterationCreateListener); ok {
			d.RegisterIterationCreateListener(l)
		}

		if l, ok := listener.(IterationUpdateListener); ok {
			d.RegisterIterationUpdateListener(l)
		}

		if l, ok := listener.(IterationDeleteListener); ok {
			d.RegisterIterationDeleteListener(l)
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
	case *StoryDeleteEvent:
		return d.processStoryDelete(ctx, e)
	case *TaskCreateEvent:
		return d.processTaskCreate(ctx, e)
	case *TaskUpdateEvent:
		return d.processTaskUpdate(ctx, e)
	case *TaskDeleteEvent:
		return d.processTaskDelete(ctx, e)
	case *BugCreateEvent:
		return d.processBugCreate(ctx, e)
	case *BugUpdateEvent:
		return d.processBugUpdate(ctx, e)
	case *BugDeleteEvent:
		return d.processBugDelete(ctx, e)
	case *StoryCommentAddEvent:
		return d.processStoryCommentAdd(ctx, e)
	case *StoryCommentUpdateEvent:
		return d.processStoryCommentUpdate(ctx, e)
	case *StoryCommentDeleteEvent:
		return d.processStoryCommentDelete(ctx, e)
	case *TaskCommentAddEvent:
		return d.processTaskCommentAdd(ctx, e)
	case *TaskCommentUpdateEvent:
		return d.processTaskCommentUpdate(ctx, e)
	case *TaskCommentDeleteEvent:
		return d.processTaskCommentDelete(ctx, e)
	case *BugCommentAddEvent:
		return d.processBugCommentAdd(ctx, e)
	case *BugCommentUpdateEvent:
		return d.processBugCommentUpdate(ctx, e)
	case *BugCommentDeleteEvent:
		return d.processBugCommentDelete(ctx, e)
	case *IterationCreateEvent:
		return d.processIterationCreate(ctx, e)
	case *IterationUpdateEvent:
		return d.processIterationUpdate(ctx, e)
	case *IterationDeleteEvent:
		return d.processIterationDelete(ctx, e)
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

type dispatchRequestOptions struct {
	ctx context.Context
}

type DispatchRequestOption func(*dispatchRequestOptions)

func WithDispatchRequestContext(ctx context.Context) DispatchRequestOption {
	return func(o *dispatchRequestOptions) {
		o.ctx = ctx
	}
}

func (d *Dispatcher) DispatchRequest(req *http.Request, opts ...DispatchRequestOption) error {
	o := &dispatchRequestOptions{
		ctx: req.Context(),
	}
	for _, opt := range opts {
		opt(o)
	}

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return d.DispatchPayload(o.ctx, payload)
}

func (d *Dispatcher) RegisterStoryCreateListener(listeners ...StoryCreateListener) {
	d.storyCreateListeners = append(d.storyCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryUpdateListener(listeners ...StoryUpdateListener) {
	d.storyUpdateListeners = append(d.storyUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryDeleteListener(listeners ...StoryDeleteListener) {
	d.storyDeleteListeners = append(d.storyDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskCreateListener(listeners ...TaskCreateListener) {
	d.taskCreateListeners = append(d.taskCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskUpdateListener(listeners ...TaskUpdateListener) {
	d.taskUpdateListeners = append(d.taskUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskDeleteListener(listeners ...TaskDeleteListener) {
	d.taskDeleteListeners = append(d.taskDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCreateListener(listeners ...BugCreateListener) {
	d.bugCreateListeners = append(d.bugCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugUpdateListener(listeners ...BugUpdateListener) {
	d.bugUpdateListeners = append(d.bugUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugDeleteListener(listeners ...BugDeleteListener) {
	d.bugDeleteListeners = append(d.bugDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryCommentAddListener(listeners ...StoryCommentAddListener) {
	d.storyCommentAddListeners = append(d.storyCommentAddListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryCommentUpdateListener(listeners ...StoryCommentUpdateListener) {
	d.storyCommentUpdateListeners = append(d.storyCommentUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterStoryCommentDeleteListener(listeners ...StoryCommentDeleteListener) {
	d.storyCommentDeleteListeners = append(d.storyCommentDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskCommentAddListener(listeners ...TaskCommentAddListener) {
	d.taskCommentAddListeners = append(d.taskCommentAddListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskCommentUpdateListener(listeners ...TaskCommentUpdateListener) {
	d.taskCommentUpdateListeners = append(d.taskCommentUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterTaskCommentDeleteListener(listeners ...TaskCommentDeleteListener) {
	d.taskCommentDeleteListeners = append(d.taskCommentDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCommentAddListener(listeners ...BugCommentAddListener) {
	d.bugCommentAddListeners = append(d.bugCommentAddListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCommentUpdateListener(listeners ...BugCommentUpdateListener) {
	d.bugCommentUpdateListeners = append(d.bugCommentUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterBugCommentDeleteListener(listeners ...BugCommentDeleteListener) {
	d.bugCommentDeleteListeners = append(d.bugCommentDeleteListeners, listeners...)
}

func (d *Dispatcher) RegisterIterationCreateListener(listeners ...IterationCreateListener) {
	d.iterationCreateListeners = append(d.iterationCreateListeners, listeners...)
}

func (d *Dispatcher) RegisterIterationUpdateListener(listeners ...IterationUpdateListener) {
	d.iterationUpdateListeners = append(d.iterationUpdateListeners, listeners...)
}

func (d *Dispatcher) RegisterIterationDeleteListener(listeners ...IterationDeleteListener) {
	d.iterationDeleteListeners = append(d.iterationDeleteListeners, listeners...)
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

func (d *Dispatcher) processStoryDelete(ctx context.Context, event *StoryDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyDeleteListeners {
		eg.Go(func() error {
			return listener.OnStoryDelete(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processTaskCreate(ctx context.Context, event *TaskCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskCreateListeners {
		eg.Go(func() error {
			return listener.OnTaskCreate(ctx, event)
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

func (d *Dispatcher) processTaskDelete(ctx context.Context, event *TaskDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskDeleteListeners {
		eg.Go(func() error {
			return listener.OnTaskDelete(ctx, event)
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

func (d *Dispatcher) processBugDelete(ctx context.Context, event *BugDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugDeleteListeners {
		eg.Go(func() error {
			return listener.OnBugDelete(ctx, event)
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

func (d *Dispatcher) processStoryCommentUpdate(ctx context.Context, event *StoryCommentUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyCommentUpdateListeners {
		eg.Go(func() error {
			return listener.OnStoryCommentUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processStoryCommentDelete(ctx context.Context, event *StoryCommentDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.storyCommentDeleteListeners {
		eg.Go(func() error {
			return listener.OnStoryCommentDelete(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processTaskCommentAdd(ctx context.Context, event *TaskCommentAddEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskCommentAddListeners {
		eg.Go(func() error {
			return listener.OnTaskCommentAdd(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processTaskCommentUpdate(ctx context.Context, event *TaskCommentUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskCommentUpdateListeners {
		eg.Go(func() error {
			return listener.OnTaskCommentUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processTaskCommentDelete(ctx context.Context, event *TaskCommentDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.taskCommentDeleteListeners {
		eg.Go(func() error {
			return listener.OnTaskCommentDelete(ctx, event)
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

func (d *Dispatcher) processBugCommentDelete(ctx context.Context, event *BugCommentDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.bugCommentDeleteListeners {
		eg.Go(func() error {
			return listener.OnBugCommentDelete(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processIterationCreate(ctx context.Context, event *IterationCreateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.iterationCreateListeners {
		eg.Go(func() error {
			return listener.OnIterationCreate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processIterationUpdate(ctx context.Context, event *IterationUpdateEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.iterationUpdateListeners {
		eg.Go(func() error {
			return listener.OnIterationUpdate(ctx, event)
		})
	}
	return eg.Wait()
}

func (d *Dispatcher) processIterationDelete(ctx context.Context, event *IterationDeleteEvent) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, listener := range d.iterationDeleteListeners {
		eg.Go(func() error {
			return listener.OnIterationDelete(ctx, event)
		})
	}
	return eg.Wait()
}
