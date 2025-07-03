package sendgrid

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const baseURLPath string = "/v3"

var (
	ErrIncorrectResponse = errors.New("response is incorrect")
)

// setup sets up a test HTTP server along with a sendgrid.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	return setupWithPath()
}

// setupWithPath sets up a test HTTP server along with a sendgrid.Client with the path.
func setupWithPath() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the sendgrid client being tested and is
	// configured to use test server.
	client = New(
		"test-token",
		OptionSubuser("dummy"),
		OptionBaseURL(server.URL+baseURLPath),
		OptionHTTPClient(&http.Client{}),
		OptionDebug(false),
		OptionLog(log.New(os.Stderr, "kenzo0107/sendgrid", log.LstdFlags|log.Lshortfile)),
	)

	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestDebugfAndDebugln_WithDebugEnabled(t *testing.T) {
	client := New("test-api-key", OptionDebug(true), OptionLog(log.New(io.Discard, "", 0)))

	// Test that these don't panic when debug is enabled
	client.Debugf("test %s", "message")
	client.Debugln("test", "message")
}

func TestDebugfAndDebugln_WithDebugDisabled(t *testing.T) {
	client := New("test-api-key", OptionDebug(false))

	// Test that these don't panic when debug is disabled
	client.Debugf("test %s", "message")
	client.Debugln("test", "message")
}

func TestClient_Do_WithInvalidResponse(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v3/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `invalid json`)
	})

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	err = client.Do(context.Background(), req, &result)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
}

func TestClient_Do_WithNilContext(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Do(nil, req, nil)
	assert.Error(t, err)
}

func TestClient_Do_WithContextCancelled(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v3/test", func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	})

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err = client.Do(ctx, req, nil)
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}

func TestClient_Do_WithResponseWriter(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/v3/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"key":"value"}`)
	})

	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	err = client.Do(context.Background(), req, &buf)
	assert.NoError(t, err)
	assert.Equal(t, `{"key":"value"}`, buf.String())
}

func TestAddOptions_WithInvalidStruct(t *testing.T) {
	client := New("test-api-key")

	// Test with a struct that contains invalid types for URL encoding
	invalidOpts := struct {
		Field chan int `url:"field"`
	}{
		Field: make(chan int),
	}

	_, err := client.AddOptions("/test", invalidOpts)
	assert.NoError(t, err) // go-querystring ignores unsupported types, so no error is expected
}

func TestAddOptions_MalformedURL(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	opts := &StatsOptions{
		StartDate: "2025-01-01",
	}

	_, err := client.AddOptions(":invalid-url", opts)
	assert.Error(t, err)
}

func TestNewRequest_MalformedURLStr(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	// Malformed urlStr
	_, err := client.NewRequest("GET", ":invalid-url", nil)
	assert.Error(t, err)
}

func TestNewRequest_WithInvalidBody(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	// Test with a body that cannot be JSON marshaled
	invalidBody := make(chan int)

	_, err := client.NewRequest("POST", "/test", invalidBody)
	assert.Error(t, err)
}

func TestNewRequest_WithTrailingSlashInBaseURL(t *testing.T) {
	client := New("test-api-key", OptionBaseURL("http://localhost:8080/"))
	req, err := client.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/v3/test", req.URL.String())
}

func TestNewRequest_WithoutTrailingSlashInBaseURL(t *testing.T) {
	client := New("test-api-key", OptionBaseURL("http://localhost:8080"))
	req, err := client.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/v3/test", req.URL.String())
}

func TestNewRequest_WithSubuser(t *testing.T) {
	client := New("test-api-key", OptionSubuser("test-subuser"))
	req, err := client.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)
	assert.Equal(t, "test-subuser", req.Header.Get("On-Behalf-Of"))
}

func TestNewRequest_WithBody(t *testing.T) {
	client := New("test-api-key")
	body := map[string]string{"key": "value"}
	req, err := client.NewRequest("POST", "/test", body)
	assert.NoError(t, err)
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"))

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(req.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\"key\":\"value\"}\n", buf.String())
}
