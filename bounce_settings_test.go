package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetBounceSettings(t *testing.T) {
	tt := []struct {
		name       string
		statusCode int
		response   string
		expected   *BounceSettings
		expectErr  bool
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			response:   `{"soft_bounces": 7}`,
			expected: &BounceSettings{
				SoftBouncePurgeDays: 7,
			},
			expectErr: false,
		},
		{
			name:       "not found - return default",
			statusCode: http.StatusNotFound,
			response:   `{}`,
			expected: &BounceSettings{
				SoftBouncePurgeDays: 7,
			},
			expectErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				w.WriteHeader(tc.statusCode)
				fmt.Fprint(w, tc.response)
			})

			got, err := client.GetBounceSettings(context.Background())
			if tc.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

func TestClient_UpdateBounceSettings(t *testing.T) {
	tt := []struct {
		name       string
		input      *InputUpdateBounceSettings
		statusCode int
		response   string
		expected   *BounceSettings
		expectErr  bool
	}{
		{
			name: "success",
			input: &InputUpdateBounceSettings{
				SoftBouncePurgeDays: 14,
			},
			statusCode: http.StatusOK,
			response:   `{"soft_bounces": 14}`,
			expected: &BounceSettings{
				SoftBouncePurgeDays: 14,
			},
			expectErr: false,
		},
		{
			name: "not found - return input",
			input: &InputUpdateBounceSettings{
				SoftBouncePurgeDays: 30,
			},
			statusCode: http.StatusNotFound,
			response:   `{}`,
			expected: &BounceSettings{
				SoftBouncePurgeDays: 30,
			},
			expectErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			defer teardown()

			mux.HandleFunc("/mail_settings/bounce_purge", func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodPatch, r.Method)
				w.WriteHeader(tc.statusCode)
				fmt.Fprint(w, tc.response)
			})

			got, err := client.UpdateBounceSettings(context.Background(), tc.input)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}