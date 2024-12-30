package webhook

import "context"

// 需求/任务/缺陷类
type (
	StoryCreateListener interface {
		OnStoryCreate(ctx context.Context, event *StoryCreateEvent) error
	}

	StoryUpdateListener interface {
		OnStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error
	}

	StoryDeleteListener interface {
		OnStoryDelete(ctx context.Context, event *StoryDeleteEvent) error
	}

	TaskCreateListener interface {
		OnTaskCreate(ctx context.Context, event *TaskCreateEvent) error
	}

	TaskUpdateListener interface {
		OnTaskUpdate(ctx context.Context, event *TaskUpdateEvent) error
	}

	TaskDeleteListener interface {
		OnTaskDelete(ctx context.Context, event *TaskDeleteEvent) error
	}

	BugCreateListener interface {
		OnBugCreate(ctx context.Context, event *BugCreateEvent) error
	}

	BugUpdateListener interface {
		OnBugUpdate(ctx context.Context, event *BugUpdateEvent) error
	}

	BugDeleteListener interface {
		OnBugDelete(ctx context.Context, event *BugDeleteEvent) error
	}
)

type (
	StoryCommentAddListener interface {
		OnStoryCommentAdd(ctx context.Context, event *StoryCommentAddEvent) error
	}

	StoryCommentUpdateListener interface {
		OnStoryCommentUpdate(ctx context.Context, event *StoryCommentUpdateEvent) error
	}

	StoryCommentDeleteListener interface {
		OnStoryCommentDelete(ctx context.Context, event *StoryCommentDeleteEvent) error
	}

	TaskCommentAddListener interface {
		OnTaskCommentAdd(ctx context.Context, event *TaskCommentAddEvent) error
	}

	TaskCommentUpdateListener interface {
		OnTaskCommentUpdate(ctx context.Context, event *TaskCommentUpdateEvent) error
	}

	TaskCommentDeleteListener interface {
		OnTaskCommentDelete(ctx context.Context, event *TaskCommentDeleteEvent) error
	}

	BugCommentAddListener interface {
		OnBugCommentAdd(ctx context.Context, event *BugCommentAddEvent) error
	}

	BugCommentUpdateListener interface {
		OnBugCommentUpdate(ctx context.Context, event *BugCommentUpdateEvent) error
	}

	BugCommentDeleteListener interface {
		OnBugCommentDelete(ctx context.Context, event *BugCommentDeleteEvent) error
	}
)

type (
	IterationCreateListener interface {
		OnIterationCreate(ctx context.Context, event *IterationCreateEvent) error
	}

	IterationUpdateListener interface {
		OnIterationUpdate(ctx context.Context, event *IterationUpdateEvent) error
	}

	IterationDeleteListener interface {
		OnIterationDelete(ctx context.Context, event *IterationDeleteEvent) error
	}
)
