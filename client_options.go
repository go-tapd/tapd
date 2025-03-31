package tapd

import "net/http"

type ClientOption func(*Client) error

// WithBaseURL sets the baseURL for the client
func WithBaseURL(urlStr string) ClientOption {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

// WithBasicAuth sets the clientID and clientSecret for the client
func WithBasicAuth(clientID, clientSecret string) ClientOption {
	return func(c *Client) error {
		c.clientID = clientID
		c.clientSecret = clientSecret
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}
