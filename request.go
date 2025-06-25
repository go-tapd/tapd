package tapd

import (
	"net/http"
)

type RequestOption func(*http.Request) error

func WithRequestBasicAuth(clientID, clientSecret string) RequestOption {
	return func(req *http.Request) error {
		// Note: It just happens that the Header has the same name(Authorization), so the assignment can be overwritten.
		req.SetBasicAuth(clientID, clientSecret)
		return nil
	}
}

func WithRequestAccessToken(accessToken string) RequestOption {
	return func(req *http.Request) error {
		// Note: It just happens that the Header has the same name(Authorization), so the assignment can be overwritten.
		req.Header.Set("Authorization", "Bearer "+accessToken)
		return nil
	}
}

func WithRequestHeader(name, value string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set(name, value)
		return nil
	}
}

func WithRequestHeaders(headers map[string]string) RequestOption {
	return func(req *http.Request) error {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
		return nil
	}
}

func WithRequestHeaderFunc(fn func(http.Header)) RequestOption {
	return func(req *http.Request) error {
		fn(req.Header)
		return nil
	}
}

func WithRequestUserAgent(userAgent string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set("User-Agent", userAgent)
		return nil
	}
}
