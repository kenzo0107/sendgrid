package sendgrid

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuppressionListOptions(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	opts := &SuppressionListOptions{
		StartTime: 1609459200, // 2021-01-01 00:00:00 UTC
		EndTime:   1612137600, // 2021-02-01 00:00:00 UTC
		Limit:     50,
		Offset:    10,
		Email:     "test@example.com",
	}

	path := "/suppression/bounces"
	fullPath, err := client.AddOptions(path, opts)
	assert.NoError(t, err)
	assert.Contains(t, fullPath, "start_time=1609459200")
	assert.Contains(t, fullPath, "end_time=1612137600")
	assert.Contains(t, fullPath, "limit=50")
	assert.Contains(t, fullPath, "offset=10")
	assert.Contains(t, fullPath, "email=test%40example.com")
}

func TestInputDeleteSuppressions(t *testing.T) {
	// Test with specific emails
	input1 := &InputDeleteSuppressions{
		Emails: []string{"test1@example.com", "test2@example.com"},
	}
	assert.Len(t, input1.Emails, 2)
	assert.False(t, input1.DeleteAll)

	// Test with delete all flag
	input2 := &InputDeleteSuppressions{
		DeleteAll: true,
	}
	assert.True(t, input2.DeleteAll)
	assert.Empty(t, input2.Emails)
}

func TestBounceStruct(t *testing.T) {
	bounce := Bounce{
		Created: 1609459200,
		Email:   "test@example.com",
		Reason:  "550 5.1.1 User unknown",
		Status:  "5.1.1",
	}

	assert.Equal(t, int64(1609459200), bounce.Created)
	assert.Equal(t, "test@example.com", bounce.Email)
	assert.Equal(t, "550 5.1.1 User unknown", bounce.Reason)
	assert.Equal(t, "5.1.1", bounce.Status)
}

func TestBlockStruct(t *testing.T) {
	block := Block{
		Created: 1609459200,
		Email:   "test@example.com",
		Reason:  "IP temporarily blocked",
	}

	assert.Equal(t, int64(1609459200), block.Created)
	assert.Equal(t, "test@example.com", block.Email)
	assert.Equal(t, "IP temporarily blocked", block.Reason)
}

func TestSpamReportStruct(t *testing.T) {
	spamReport := SpamReport{
		Created: 1609459200,
		Email:   "test@example.com",
		IP:      "192.168.1.1",
	}

	assert.Equal(t, int64(1609459200), spamReport.Created)
	assert.Equal(t, "test@example.com", spamReport.Email)
	assert.Equal(t, "192.168.1.1", spamReport.IP)
}

func TestInvalidEmailStruct(t *testing.T) {
	invalidEmail := InvalidEmail{
		Created: 1609459200,
		Email:   "invalid@example.com",
		Reason:  "Mail domain mentioned in email address is unknown",
	}

	assert.Equal(t, int64(1609459200), invalidEmail.Created)
	assert.Equal(t, "invalid@example.com", invalidEmail.Email)
	assert.Equal(t, "Mail domain mentioned in email address is unknown", invalidEmail.Reason)
}

func TestGetBounces(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/bounces", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","reason":"550 5.1.1 User unknown","status":"5.1.1"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	bounces, err := client.GetBounces(ctx, nil)
	assert.NoError(t, err)
	assert.Len(t, bounces, 1)
	assert.Equal(t, "test@example.com", bounces[0].Email)
	assert.Equal(t, "550 5.1.1 User unknown", bounces[0].Reason)
}

func TestGetBounces_WithOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/bounces", r.URL.Path)
		assert.Equal(t, "1609459200", r.URL.Query().Get("start_time"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	opts := &SuppressionListOptions{
		StartTime: 1609459200,
	}

	_, err := client.GetBounces(ctx, opts)
	assert.NoError(t, err)
}

func TestGetBounces_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBounces(ctx, nil)
	assert.Error(t, err)
}

func TestGetBounce(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/bounces/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","reason":"550 5.1.1 User unknown","status":"5.1.1"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	bounce, err := client.GetBounce(ctx, "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", bounce.Email)
}

func TestGetBounce_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBounce(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetBounce_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBounce(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestDeleteBounces(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/bounces", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteBounces(ctx, input)
	assert.NoError(t, err)
}

func TestDeleteBounces_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteBounces(ctx, input)
	assert.Error(t, err)
}

func TestDeleteBounce(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/bounces/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteBounce(ctx, "test@example.com")
	assert.NoError(t, err)
}

func TestDeleteBounce_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteBounce(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetBlocks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/blocks", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","reason":"IP temporarily blocked"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	blocks, err := client.GetBlocks(ctx, nil)
	assert.NoError(t, err)
	assert.Len(t, blocks, 1)
	assert.Equal(t, "test@example.com", blocks[0].Email)
	assert.Equal(t, "IP temporarily blocked", blocks[0].Reason)
}

func TestGetBlocks_WithOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/blocks", r.URL.Path)
		assert.Equal(t, "1609459200", r.URL.Query().Get("start_time"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	opts := &SuppressionListOptions{
		StartTime: 1609459200,
	}

	_, err := client.GetBlocks(ctx, opts)
	assert.NoError(t, err)
}

func TestGetBlocks_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBlocks(ctx, nil)
	assert.Error(t, err)
}

func TestGetBlock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/blocks/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","reason":"IP temporarily blocked"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	block, err := client.GetBlock(ctx, "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", block.Email)
}

func TestGetBlock_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBlock(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetBlock_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetBlock(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestDeleteBlocks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/blocks", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteBlocks(ctx, input)
	assert.NoError(t, err)
}

func TestDeleteBlocks_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteBlocks(ctx, input)
	assert.Error(t, err)
}

func TestDeleteBlock(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/blocks/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteBlock(ctx, "test@example.com")
	assert.NoError(t, err)
}

func TestDeleteBlock_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteBlock(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetSpamReports(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/spam_reports", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","ip":"192.168.1.1"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	reports, err := client.GetSpamReports(ctx, nil)
	assert.NoError(t, err)
	assert.Len(t, reports, 1)
	assert.Equal(t, "test@example.com", reports[0].Email)
	assert.Equal(t, "192.168.1.1", reports[0].IP)
}

func TestGetSpamReports_WithOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/spam_reports", r.URL.Path)
		assert.Equal(t, "1609459200", r.URL.Query().Get("start_time"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	opts := &SuppressionListOptions{
		StartTime: 1609459200,
	}

	_, err := client.GetSpamReports(ctx, opts)
	assert.NoError(t, err)
}

func TestGetSpamReports_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetSpamReports(ctx, nil)
	assert.Error(t, err)
}

func TestGetSpamReport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/spam_reports/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"test@example.com","ip":"192.168.1.1"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	report, err := client.GetSpamReport(ctx, "test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", report.Email)
}

func TestGetSpamReport_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetSpamReport(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetSpamReport_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetSpamReport(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestDeleteSpamReports(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/spam_reports", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteSpamReports(ctx, input)
	assert.NoError(t, err)
}

func TestDeleteSpamReports_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"test@example.com"},
	}

	err := client.DeleteSpamReports(ctx, input)
	assert.Error(t, err)
}

func TestDeleteSpamReport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/spam_reports/test@example.com", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteSpamReport(ctx, "test@example.com")
	assert.NoError(t, err)
}

func TestDeleteSpamReport_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteSpamReport(ctx, "test@example.com")
	assert.Error(t, err)
}

func TestGetInvalidEmails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/invalid_emails", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"invalid@example.com","reason":"Mail domain mentioned in email address is unknown"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	invalidEmails, err := client.GetInvalidEmails(ctx, nil)
	assert.NoError(t, err)
	assert.Len(t, invalidEmails, 1)
	assert.Equal(t, "invalid@example.com", invalidEmails[0].Email)
	assert.Equal(t, "Mail domain mentioned in email address is unknown", invalidEmails[0].Reason)
}

func TestGetInvalidEmails_WithOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/invalid_emails", r.URL.Path)
		assert.Equal(t, "1609459200", r.URL.Query().Get("start_time"))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	opts := &SuppressionListOptions{
		StartTime: 1609459200,
	}

	_, err := client.GetInvalidEmails(ctx, opts)
	assert.NoError(t, err)
}

func TestGetInvalidEmails_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetInvalidEmails(ctx, nil)
	assert.Error(t, err)
}

func TestGetInvalidEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/suppression/invalid_emails/invalid@example.com", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"created":1609459200,"email":"invalid@example.com","reason":"Mail domain mentioned in email address is unknown"}]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	invalidEmail, err := client.GetInvalidEmail(ctx, "invalid@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "invalid@example.com", invalidEmail.Email)
}

func TestGetInvalidEmail_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[]`)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetInvalidEmail(ctx, "invalid@example.com")
	assert.Error(t, err)
}

func TestGetInvalidEmail_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	_, err := client.GetInvalidEmail(ctx, "invalid@example.com")
	assert.Error(t, err)
}

func TestDeleteInvalidEmails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/invalid_emails", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"invalid@example.com"},
	}

	err := client.DeleteInvalidEmails(ctx, input)
	assert.NoError(t, err)
}

func TestDeleteInvalidEmails_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	input := &InputDeleteSuppressions{
		Emails: []string{"invalid@example.com"},
	}

	err := client.DeleteInvalidEmails(ctx, input)
	assert.Error(t, err)
}

func TestDeleteInvalidEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/suppression/invalid_emails/invalid@example.com", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteInvalidEmail(ctx, "invalid@example.com")
	assert.NoError(t, err)
}

func TestDeleteInvalidEmail_Failed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := New("test-api-key", OptionBaseURL(server.URL))
	ctx := context.Background()

	err := client.DeleteInvalidEmail(ctx, "invalid@example.com")
	assert.Error(t, err)
}