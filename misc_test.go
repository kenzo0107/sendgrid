package sendgrid

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrorResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/templates/d-12345abcde/versions/aaaaaa-bbbb-0000-0000-aaaaaaaaa", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, `{"error": "You cannot switch editors once a dynamic template version has been created."}`); err != nil {
			t.Fatal(err)
		}
	})

	client.debug = true
	client.httpclient = &http.Client{}
	client.log = log.New(os.Stdout, "sendgrid: ", log.Lshortfile|log.LstdFlags)

	client.Debugf("%s", "test")
	client.Debugln("test")

	if _, err := client.UpdateTemplateVersion(context.TODO(), "d-12345abcde", "aaaaaa-bbbb-0000-0000-aaaaaaaaa", &InputUpdateTemplateVersion{
		Editor: "code",
	}); err == nil {
		t.Fatal("expected an error but got none", err)
	}
}

func TestErrorsResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusNotFound)
		if _, err := fmt.Fprint(w, `{"errors":[{"message": "teammate does not exis"}]}`); err != nil {
			t.Fatal(err)
		}
	})

	client.debug = true
	client.httpclient = &http.Client{}
	client.log = log.New(os.Stdout, "sendgrid: ", log.Lshortfile|log.LstdFlags)

	client.Debugf("%s", "test")
	client.Debugln("test")

	if _, err := client.GetTeammate(context.TODO(), "dummy"); err == nil {
		t.Fatal("expected an error but got none", err)
	}
}

func TestStatusUnAuthorized(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
	})

	_, err := client.GetTeammate(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none", err)
	}
}

func TestErrorResponseErr(t *testing.T) {
	// Test ErrorResponse.Err() with error
	errorResp := ErrorResponse{Error: "test error message"}
	err := errorResp.Err()
	if err == nil {
		t.Fatal("expected an error but got none")
	}
	if err.Error() != "test error message" {
		t.Fatalf("expected 'test error message', got '%s'", err.Error())
	}

	// Test ErrorResponse.Err() with no error
	emptyErrorResp := ErrorResponse{Error: ""}
	err = emptyErrorResp.Err()
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
}

func TestErrorsResponseErrs(t *testing.T) {
	// Test ErrorsResponse.Errs() with errors
	field := "email"
	message := "is required"
	errorsResp := ErrorsResponse{
		Errors: []*Error{
			{Field: &field, Message: &message},
		},
	}
	err := errorsResp.Errs()
	if err == nil {
		t.Fatal("expected an error but got none")
	}
	expected := "field: email, message: is required"
	if err.Error() != expected {
		t.Fatalf("expected '%s', got '%s'", expected, err.Error())
	}

	// Test ErrorsResponse.Errs() with no errors
	emptyErrorsResp := ErrorsResponse{Errors: []*Error{}}
	err = emptyErrorsResp.Errs()
	if err != nil {
		t.Fatalf("expected no error but got: %v", err)
	}
}

func TestStatusCodeError(t *testing.T) {
	statusErr := statusCodeError{
		Code:   404,
		Status: "404 Not Found",
	}

	// Test Error() method
	expected := "sendgrid server error: 404 Not Found"
	if statusErr.Error() != expected {
		t.Fatalf("expected '%s', got '%s'", expected, statusErr.Error())
	}

	// Test HTTPStatusCode() method
	if statusErr.HTTPStatusCode() != 404 {
		t.Fatalf("expected 404, got %d", statusErr.HTTPStatusCode())
	}
}

func TestRateLimitedError(t *testing.T) {
	retryAfter := time.Minute * 5
	rateLimitErr := &RateLimitedError{RetryAfter: retryAfter}

	expected := "sendgrid rate limit exceeded, retry after 5m0s"
	if rateLimitErr.Error() != expected {
		t.Fatalf("expected '%s', got '%s'", expected, rateLimitErr.Error())
	}
}

func TestRateLimitHandling(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", fmt.Sprintf("%d", time.Now().Add(1*time.Minute).Unix()))
		w.WriteHeader(http.StatusTooManyRequests)
	})

	_, err := client.GetTeammate(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	rateLimitErr, ok := err.(*RateLimitedError)
	if !ok {
		t.Fatalf("expected RateLimitedError, got %T", err)
	}

	if rateLimitErr.RetryAfter <= 0 {
		t.Fatal("expected positive retry after duration")
	}
}

func TestCheckStatusCode_InvalidRateLimitHeader(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/teammates/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Reset", "invalid-timestamp")
		w.WriteHeader(http.StatusTooManyRequests)
	})

	_, err := client.GetTeammate(context.TODO(), "dummy")
	if err == nil {
		t.Fatal("expected an error but got none")
	}

	// Should get a parsing error, not a RateLimitedError
	_, ok := err.(*RateLimitedError)
	if ok {
		t.Fatal("expected parsing error, got RateLimitedError")
	}
}

func TestLogResponse_Error(t *testing.T) {
	client := New("test-api-key", OptionDebug(true), OptionLog(log.New(io.Discard, "", 0)))
	resp := &http.Response{
		Body: io.NopCloser(&errorReader{}),
	}
	err := logResponse(resp, client)
	assert.Error(t, err)
}

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("test error")
}