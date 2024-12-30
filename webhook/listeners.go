package webhook

import "context"

type (
	StoryCreateListener interface {
		OnStoryCreate(ctx context.Context, event *StoryCreateEvent) error
	}

	StoryUpdateListener interface {
		OnStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error
	}

	TaskUpdateListener interface {
		OnTaskUpdate(ctx context.Context, event *TaskUpdateEvent) error
	}

	StoryCommentAddListener interface {
		OnStoryCommentAdd(ctx context.Context, event *StoryCommentAddEvent) error
	}

	BugCreateListener interface {
		OnBugCreate(ctx context.Context, event *BugCreateEvent) error
	}

	BugUpdateListener interface {
		OnBugUpdate(ctx context.Context, event *BugUpdateEvent) error
	}

	BugCommentAddListener interface {
		OnBugCommentAdd(ctx context.Context, event *BugCommentAddEvent) error
	}

	BugCommentUpdateListener interface {
		OnBugCommentUpdate(ctx context.Context, event *BugCommentUpdateEvent) error
	}
)
