package tapd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	apiClientID     = "tapd-client-id"
	apiClientSecret = "tapd-client-secret"
	apiAccessToken  = "tapd-access-token"
	successResponse = `{
  "status": 1,
  "data": {},
  "info": "success"
}`
)

var ctx = context.Background()

func createServerClient(t *testing.T, handler http.Handler) (*httptest.Server, *Client) {
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	client, err := NewClient(
		apiClientID, apiClientSecret,
		WithBaseURL(srv.URL),
		WithHTTPClient(NewRetryableHTTPClient(
			WithRetryableHTTPClientLogger(log.New(os.Stderr, "", log.LstdFlags)),
		)),
	)
	assert.NoError(t, err)

	return srv, client
}

func loadData(t *testing.T, filepath string) []byte {
	content, err := os.ReadFile(filepath)
	assert.NoError(t, err)
	return content
}

func TestClient_BasicAuth(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/basic-auth", r.URL.Path)

		// check basic auth
		clientID, clientSecret, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, apiClientID, clientID)
		assert.Equal(t, apiClientSecret, clientSecret)

		// nolint:errcheck
		fmt.Fprint(w, successResponse)
	}))
	assert.Equal(t, authTypeBasic, client.authType)

	req, err := client.NewRequest(ctx, http.MethodGet, "__/basic-auth", nil, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_PATAuth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/basic-auth", r.URL.Path)

		// check Authorization header for PAT
		authHeader := r.Header.Get("Authorization")
		assert.NotEmpty(t, authHeader)
		assert.Equal(t, "Bearer "+apiAccessToken, authHeader)

		// nolint:errcheck
		fmt.Fprint(w, successResponse)
	}))
	t.Cleanup(srv.Close)

	client, err := NewPATClient(apiAccessToken, WithBaseURL(srv.URL))
	assert.NoError(t, err)

	// Check if the client is using PAT authentication
	req, err := client.NewRequest(ctx, http.MethodGet, "__/basic-auth", nil, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_ErrorResponse(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/error-response", r.URL.Path)

		// nolint:errcheck
		fmt.Fprint(w, `{
  "status": 0,
  "data": {},
  "info": "error"
}`)
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/error-response", nil, nil)
	assert.NoError(t, err)

	_, err = client.Do(req, nil)
	assert.Error(t, err)
	assert.True(t, IsErrorResponse(err))
}

func TestClient_NormalRequest(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/normal-request", r.URL.Path)

		// check header
		assert.Equal(t, defaultUserAgent, r.Header.Get("User-Agent"))

		// check basic auth
		clientID, clientSecret, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, apiClientID, clientID)
		assert.Equal(t, apiClientSecret, clientSecret)

		fmt.Fprint(w, successResponse) // nolint:errcheck
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/normal-request", nil, nil)
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_WithRequestOption(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/__/request-option", r.URL.Path)

		// check header
		assert.Equal(t, "header-value", r.Header.Get("header-name"))
		assert.Equal(t, "headers-value", r.Header.Get("headers-name"))
		assert.Equal(t, "func-value", r.Header.Get("func-name"))
		assert.Contains(t, r.Header.Values("func-name-2"), "func-value-2")
		assert.Contains(t, r.Header.Values("func-name-2"), "func-value-3")
		assert.Equal(t, "test-user-agent", r.Header.Get("User-Agent"))

		// check basic auth
		clientID, clientSecret, ok := r.BasicAuth()
		assert.True(t, ok)
		assert.Equal(t, "test-client-id", clientID)
		assert.Equal(t, "test-client-secret", clientSecret)

		fmt.Fprint(w, successResponse) // nolint:errcheck
	}))

	req, err := client.NewRequest(ctx, http.MethodGet, "__/request-option", nil, []RequestOption{
		WithRequestBasicAuth("test-client-id", "test-client-secret"),
		WithRequestHeader("header-name", "header-value"),
		WithRequestHeaders(map[string]string{
			"headers-name": "headers-value",
		}),
		WithRequestHeaderFunc(func(header http.Header) {
			header.Set("func-name", "func-value")
			header.Add("func-name-2", "func-value-2")
			header.Add("func-name-2", "func-value-3")
		}),
		WithRequestUserAgent("test-user-agent"),
	})
	assert.NoError(t, err)

	resp, err := client.Do(req, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestClient_WithRequestOption_WithAuth(t *testing.T) {
	var (
		assertBasicAuthFunc = func(expectedClientID, expectedClientSecret string) func(t *testing.T, r *http.Request) {
			return func(t *testing.T, r *http.Request) {
				clientID, clientSecret, ok := r.BasicAuth()
				assert.True(t, ok)
				assert.Equal(t, expectedClientID, clientID)
				assert.Equal(t, expectedClientSecret, clientSecret)
			}
		}
		assertPATAuthFunc = func(expectedAccessToken string) func(t *testing.T, r *http.Request) {
			return func(t *testing.T, r *http.Request) {
				authHeader := r.Header.Get("Authorization")
				assert.NotEmpty(t, authHeader)
				assert.Equal(t, "Bearer "+expectedAccessToken, authHeader)
			}
		}

		createClientBasicAuthFunc = func(clientID, clientSecret string) func(srv *httptest.Server) (*Client, error) {
			return func(srv *httptest.Server) (*Client, error) {
				return NewClient(clientID, clientSecret, WithBaseURL(srv.URL))
			}
		}
		createClientPATAuthFunc = func(accessToken string) func(srv *httptest.Server) (*Client, error) {
			return func(srv *httptest.Server) (*Client, error) {
				return NewPATClient(accessToken, WithBaseURL(srv.URL))
			}
		}
	)

	tests := []struct {
		name             string
		createClientFunc func(srv *httptest.Server) (*Client, error)
		requestOption    RequestOption
		wantFunc         func(t *testing.T, r *http.Request)
	}{
		{
			"client is created with Basic Auth, but request option is nil",
			createClientBasicAuthFunc("client-a", "secret-a"),
			nil,
			assertBasicAuthFunc("client-a", "secret-a"),
		},
		{
			"client is created with PAT Auth, but request option is nil",
			createClientPATAuthFunc("access-token-a"),
			nil,
			assertPATAuthFunc("access-token-a"),
		},
		{
			"Use Basic Auth when client is created with Basic Auth",
			createClientBasicAuthFunc("client-a", "secret-a"),
			WithRequestBasicAuth("client-b", "secret-b"),
			assertBasicAuthFunc("client-b", "secret-b"),
		},
		{
			"Use Basic Auth when client is created with PAT Auth",
			createClientBasicAuthFunc("client-a", "secret-a"),
			WithRequestAccessToken("access-token-b"),
			assertPATAuthFunc("access-token-b"),
		},
		{
			"Use PAT Auth when client is created with PAT Auth",
			createClientPATAuthFunc("access-token-a"),
			WithRequestAccessToken("access-token-b"),
			assertPATAuthFunc("access-token-b"),
		},
		{
			"Use PAT Auth when client is created with Basic Auth",
			createClientPATAuthFunc("access-token-a"),
			WithRequestBasicAuth("client-b", "secret-b"),
			assertBasicAuthFunc("client-b", "secret-b"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				tt.wantFunc(t, r)
				fmt.Fprint(w, successResponse) // nolint:errcheck
			}))

			client, err := tt.createClientFunc(srv)
			require.NoError(t, err)

			_, _, _ = client.StoryService.GetStories(
				ctx, nil,
				tt.requestOption,
			)
		})
	}
}
