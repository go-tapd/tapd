package tapd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL   = "https://api.tapd.cn/"
	defaultUserAgent = "go-tapd"
)

var defaultHTTPClient = NewRetryableHTTPClient()

type Client struct {
	// baseURL for API requests.
	baseURL *url.URL

	// clientID, clientSecret for basic authentication.
	clientID, clientSecret string

	// userAgent used for HTTP requests
	userAgent string

	// httpClient is the HTTP client used to communicate with the API.
	httpClient *http.Client

	// services used for talking to different parts of the Tapd API.
	StoryService      *StoryService
	BugService        BugService
	IterationService  IterationService
	TaskService       *TaskService
	CommentService    CommentService
	ReportService     ReportService
	AttachmentService AttachmentService
	TimesheetService  TimesheetService
	WorkspaceService  WorkspaceService
	LabelService      LabelService
	MeasureService    MeasureService
	UserService       UserService
	WorkflowService   WorkflowService
	SettingService    SettingService
}

// NewClient returns a new Tapd API client.
// Alias for NewBasicAuthClient.
func NewClient(clientID, clientSecret string, opts ...ClientOption) (*Client, error) {
	return NewBasicAuthClient(clientID, clientSecret, opts...)
}

// NewBasicAuthClient returns a new Tapd API client with basic authentication.
func NewBasicAuthClient(clientID, clientSecret string, opts ...ClientOption) (*Client, error) {
	return newClient(append(opts,
		WithBasicAuth(clientID, clientSecret))...)
}

// newClient returns a new Tapd API client.
func newClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		userAgent:  defaultUserAgent,
		httpClient: defaultHTTPClient,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	// setup
	if err := c.setup(); err != nil {
		return nil, err
	}

	// services
	c.StoryService = &StoryService{client: c}
	c.BugService = NewBugService(c)
	c.IterationService = NewIterationService(c)
	c.TaskService = &TaskService{client: c}
	c.CommentService = NewCommentService(c)
	c.ReportService = NewReportService(c)
	c.AttachmentService = NewAttachmentService(c)
	c.TimesheetService = NewTimesheetService(c)
	c.WorkspaceService = NewWorkspaceService(c)
	c.LabelService = NewLabelService(c)
	c.MeasureService = NewMeasureService(c)
	c.UserService = NewUserService(c)
	c.WorkflowService = NewWorkflowService(c)
	c.SettingService = NewSettingService(c)

	return c, nil
}

// setup sets up the client for API requests.
func (c *Client) setup() error {
	if c.baseURL == nil {
		if err := c.setBaseURL(defaultBaseURL); err != nil {
			return err
		}
	}

	return nil
}

// setBaseURL sets the base URL for API requests to a custom endpoint.
func (c *Client) setBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	c.baseURL = baseURL

	return nil
}

func (c *Client) NewRequest(ctx context.Context, method, path string, data any, opts []RequestOption) (*http.Request, error) { //nolint:lll
	u := *c.baseURL
	unescaped, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + unescaped

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.userAgent != "" {
		reqHeaders.Set("User-Agent", c.userAgent)
	}

	var body io.Reader
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = io.NopCloser(bytes.NewReader(b))
		}
	case data != nil:
		q, err := query.Values(data)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// basic auth
	if c.clientID != "" && c.clientSecret != "" {
		req.SetBasicAuth(c.clientID, c.clientSecret)
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	// Apply request options
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v any) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()              // nolint:errcheck
	defer io.Copy(io.Discard, resp.Body) // nolint:errcheck

	// decode response body
	var rawBody RawBody
	if err := json.NewDecoder(resp.Body).Decode(&rawBody); err != nil {
		return nil, err
	}

	// debug mode
	// body, _ := json.Marshal(rawBody)
	// fmt.Println(string(body))
	// spew.Dump(rawBody)

	// check status
	if rawBody.Status != 1 {
		return nil, &ErrorResponse{
			response: resp,
			rawBody:  &rawBody,
			err:      errors.New(rawBody.Info),
		}
	}

	if v != nil {
		if err := json.Unmarshal(rawBody.Data, v); err != nil {
			return nil, err
		}
	}

	return newResponse(resp), err
}
