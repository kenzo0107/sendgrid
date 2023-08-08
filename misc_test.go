package sendgrid

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestErrorResponse(t *testing.T) {
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
