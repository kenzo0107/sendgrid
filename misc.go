package sendgrid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/pkg/errors"
)

// ErrorResponse is sendgrid error response
type ErrorResponse struct {
	Errors []*Error `json:"errors"`
}

// Error is sendgrid error
type Error struct {
	Field   *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Errs : error
func (t ErrorResponse) Errs() error {
	s := []string{}
	for _, err := range t.Errors {
		var msg strings.Builder
		if err.Field != nil {
			msg.WriteString("field: ")
			msg.WriteString(*err.Field)
			msg.WriteString(", ")
		}
		msg.WriteString("message: ")
		msg.WriteString(*err.Message)
		s = append(s, msg.String())
	}

	if len(s) == 0 {
		return nil
	}

	return errors.New(strings.Join(s, ", "))
}

// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("sendgrid server error: %s", t.Status)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

func checkStatusCode(resp *http.Response, d debug) error {
	// return no error if response returns status code 2xx
	if resp.StatusCode/100 == 2 {
		return nil
	}

	if err := logResponse(resp, d); err != nil {
		return err
	}

	errorResponse := new(ErrorResponse)
	if err := newJSONParser(errorResponse)(resp); err == nil {
		return errorResponse.Errs()
	}

	return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
}

type responseParser func(*http.Response) error

func newJSONParser(dst interface{}) responseParser {
	return func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(dst)
	}
}

func logResponse(resp *http.Response, d debug) error {
	if d.Debug() {
		text, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
		d.Debugln(string(text))
	}

	return nil
}
