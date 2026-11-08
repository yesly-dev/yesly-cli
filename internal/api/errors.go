package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Error is a non-2xx answer from the API.
type Error struct {
	Status    int
	Type      string
	Code      string
	Param     string
	Message   string
	RequestID string
	// Body is the answer as received, usually {"error": {...}}.
	Body []byte
}

func (e *Error) Error() string {
	var details []string
	details = append(details, fmt.Sprintf("status %d", e.Status))
	if e.Type != "" {
		details = append(details, "type "+e.Type)
	}
	if e.Code != "" {
		details = append(details, "code "+e.Code)
	}
	if e.Param != "" {
		details = append(details, "param "+e.Param)
	}
	if e.RequestID != "" {
		details = append(details, "request "+e.RequestID)
	}
	return fmt.Sprintf("%s (%s)", e.Message, strings.Join(details, ", "))
}

// ConnectionError is a request that got no answer.
type ConnectionError struct {
	Message string
	Err     error
}

func (e *ConnectionError) Error() string { return e.Message + ": " + e.Err.Error() }
func (e *ConnectionError) Unwrap() error { return e.Err }

func errorFrom(status int, body []byte, h http.Header) *Error {
	var env struct {
		Error struct {
			Type      string `json:"type"`
			Code      string `json:"code"`
			Param     string `json:"param"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	_ = json.Unmarshal(body, &env)
	e := &Error{
		Status:    status,
		Type:      env.Error.Type,
		Code:      env.Error.Code,
		Param:     env.Error.Param,
		Message:   env.Error.Message,
		RequestID: env.Error.RequestID,
		Body:      body,
	}
	if e.RequestID == "" {
		e.RequestID = h.Get("Request-Id")
	}
	if e.Message == "" {
		e.Message = fmt.Sprintf("the Yesly API answered %d", status)
	}
	return e
}
