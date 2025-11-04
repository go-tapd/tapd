# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Go-Tapd-SDK** is a Go client library for accessing the TAPD API, a Chinese Agile project management platform. The project is in active development but marked as non-stable (not recommended for production use).

- **Module**: `github.com/go-tapd/tapd`
- **Go Version**: 1.24.0
- **Architecture**: Clean architecture with service-oriented design

## Development Commands

### Building and Testing
```bash
# Run tests with race detection
make test

# Clean up dependencies
make go-mod-tidy

# Run linter (includes go-mod-tidy)
make lint

# Fix lint issues automatically
make lint-fix

# Check if working tree is clean (used in CI)
make check-clean-work
```

### Running Specific Tests
```bash
# Run all tests
go test ./... -race

# Run unit tests only (excluding integration tests)
go test . -race

# Run integration/production tests (requires real API credentials)
go test ./tests/... -race
```

## Code Architecture

### Client Structure
The SDK follows a **client-centric architecture** with a main `Client` struct that orchestrates all services:

```go
type Client struct {
    baseURL       *url.URL
    authType      authType  // basic or PAT
    clientID      string    // Basic auth
    clientSecret  string    // Basic auth
    accessToken   string    // PAT auth
    userAgent     string
    httpClient    *http.Client

    // Service registry (13 total services)
    StoryService      StoryService
    BugService        BugService
    IterationService  IterationService
    TaskService       TaskService
    // ... other services
}
```

### Authentication Methods
1. **Basic Authentication**: `tapd.NewClient(clientID, clientSecret)`
2. **Personal Access Token**: `tapd.NewPATClient(accessToken)`

### Service Pattern
Each service follows a consistent interface/implementation pattern:
```go
type StoryService interface {
    CreateStory(ctx context.Context, req *CreateStoryRequest) (*Story, *Response, error)
    GetStories(ctx context.Context, req *GetStoriesRequest) ([]Story, *Response, error)
    // ... other methods
}

type storyService struct {
    client *Client
}

func NewStoryService(client *Client) StoryService {
    return &storyService{client: client}
}
```

### Key Services (13 total)
- **StoryService** (largest: 1740 lines) - Story/requirement management
- **BugService** (801 lines) - Bug/defect tracking
- **TaskService** (709 lines) - Task management
- **IterationService** (566 lines) - Iteration/sprint management
- **CommentService** (205 lines) - Comment handling
- **AttachmentService** (187 lines) - File attachments
- **TimesheetService** (242 lines) - Time tracking
- **WorkspaceService** (161 lines) - Workspace/project management
- **LabelService** (179 lines) - Tag/label management
- **ReportService** (100 lines) - Project reports
- **MeasureService** (80 lines) - Metrics
- **UserService** (59 lines) - User management
- **WorkflowService** (77 lines) - Workflow status management

### Special Types

#### Multi[T] - Comma-separated Values
For API parameters that accept comma-separated values:
```go
// Creates "1,2,3"
ID: tapd.NewMulti[int64](1, 2, 3)
Fields: tapd.NewMulti[string]("name", "description", "status")
```

#### Enum[T] - Pipe-separated Values
For API parameters that accept pipe-separated enum values:
```go
// Creates "1|2|3"
Status: tapd.NewEnum[string]("active", "pending")
```

### Webhook System
Comprehensive webhook dispatcher in `webhook/` directory:
```go
dispatcher := webhook.NewDispatcher(
    webhook.WithRegisters(&StoryUpdateListener{}),
)

// Event types supported:
// - story::create/update/delete
// - task::create/update/delete
// - bug::create/update/delete
// - iteration::create/update/delete
// - story_comment::add/update/delete
// - task_comment::add/update/delete
// - bug_comment::add/update/delete
```

## Development Patterns

### Request/Response Pattern
- **Request parameters**: Always pointer types (`*string`, `*int64`, etc.)
- **Response fields**: Non-pointer types, use pointer types for nullable fields
- **Functional options**: Used for client and request configuration

### Error Handling
All service methods return `(result, *Response, error)` triple where:
- `result`: The main response data
- `*Response`: HTTP response metadata with pagination info
- `error`: Any error that occurred

### Testing Strategy
1. **Unit Tests**: Same-package tests (`*_test.go`) with `t.Parallel()`
2. **Integration Tests**: Separate `tests/` directory for real API calls
3. **Mock Testing**: Interface-based design enables easy mocking

### Code Quality
- **20+ linters** enabled via golangci-lint
- **Custom revive rules** for code style
- **Race detection** enabled in all tests
- **gofumpt + goimports** for formatting

## File Organization

```
├── client*.go           # Client initialization and configuration
├── api_*.go            # Service implementations (13 files)
├── api_types.go        # Generic types (Multi[T], Enum[T])
├── request.go          # Request options and builders
├── response.go         # HTTP response handling
├── helpers.go          # Utility functions
├── webhook/            # Webhook handling (12 files)
├── tests/              # Integration tests (15 files)
├── .golangci.yml       # Linting configuration
├── Makefile           # Build automation
└── features.md        # API implementation roadmap
```

## API Coverage

The SDK provides extensive TAPD API coverage with ongoing development. See `features.md` for detailed implementation status. Key areas:

- ✅ **Stories, Bugs, Tasks, Iterations**: Full CRUD + associations
- ✅ **Comments, Attachments**: Complete support
- ✅ **Workspaces, Labels, Reports**: Core functionality
- ✅ **Webhooks**: Comprehensive event handling
- 🔄 **Testing, Releases, Wiki**: Partial implementation
- ⏳ **Many specialized APIs**: Available for contribution

## Contributing

When adding new features:

1. **Follow Service Pattern**: Create interface + implementation pair
2. **Use Generic Types**: Leverage `Multi[T]` and `Enum[T]` appropriately
3. **Maintain Test Coverage**: Add unit tests + integration tests
4. **Update features.md**: Mark implemented APIs
5. **Webhook Support**: Follow existing listener pattern for events
6. **Chinese Comments**: Consider adding Chinese documentation alongside English

## Important Notes

- **Non-stable**: Not recommended for production use
- **Go 1.24+ Required**: Uses modern Go features including generics
- **API Documentation**: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/
- **Active Development**: Regular commits and CI/CD
- **Community Contributions**: Welcome via issues and PRs