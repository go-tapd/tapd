package tapd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Response represents an API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response.
func newResponse(httpResp *http.Response) *Response {
	return &Response{Response: httpResp}
}

// RawBody represents a raw body.
type RawBody struct {
	Status int             `json:"status"`
	Data   json.RawMessage `json:"data"`
	Info   string          `json:"info"`
}

// ErrorResponse represents a tapd error response.
type ErrorResponse struct {
	response *http.Response
	rawBody  *RawBody
	err      error
}

func (e *ErrorResponse) Error() string {
	if e.rawBody != nil {
		return fmt.Sprintf("code: %d, info: %s", e.rawBody.Status, e.rawBody.Info)
	}

	if e.response != nil {
		return fmt.Sprintf("status code: %d, err: %v", e.response.StatusCode, e.err)
	}

	return e.err.Error()
}

func (e *ErrorResponse) Unwrap() error {
	return e.err
}

func IsErrorResponse(err error) bool {
	var e *ErrorResponse
	return errors.As(err, &e)
}

// CountResponse represents the response of count.
type CountResponse struct {
	Count int `json:"count"`
}
