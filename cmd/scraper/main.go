package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var urls = []string{
	"https://www.twilio.com/docs/sendgrid/api-reference/alerts/create-a-new-alert",
	"https://www.twilio.com/docs/sendgrid/api-reference/alerts/delete-an-alert",
	"https://www.twilio.com/docs/sendgrid/api-reference/alerts/retrieve-a-specific-alert",
	"https://www.twilio.com/docs/sendgrid/api-reference/alerts/retrieve-all-alerts",
	"https://www.twilio.com/docs/sendgrid/api-reference/alerts/update-an-alert",
	"https://www.twilio.com/docs/sendgrid/api-reference/api-keys/create-api-keys",
	"https://www.twilio.com/docs/sendgrid/api-reference/api-keys/delete-api-keys",
	"https://www.twilio.com/docs/sendgrid/api-reference/api-keys/retrieve-all-api-keys-belonging-to-the-authenticated-user",
	"https://www.twilio.com/docs/sendgrid/api-reference/api-keys/retrieve-an-existing-api-key",
	"https://www.twilio.com/docs/sendgrid/api-reference/bounces-api/delete-a-bounce",
	"https://www.twilio.com/docs/sendgrid/api-reference/bounces-api/delete-bounces",
	"https://www.twilio.com/docs/sendgrid/api-reference/bounces-api/retrieve-a-bounce",
	"https://www.twilio.com/docs/sendgrid/api-reference/bounces-api/retrieve-all-bounces",
	"https://www.twilio.com/docs/sendgrid/api-reference/bounces/remove-bounces",
	"https://www.twilio.com/docs/sendgrid/api-reference/blocks/retrieve-all-blocks",
	"https://www.twilio.com/docs/sendgrid/api-reference/blocks/retrieve-a-specific-block",
	"https://www.twilio.com/docs/sendgrid/api-reference/blocks/delete-blocks",
	"https://www.twilio.com/docs/sendgrid/api-reference/blocks/delete-a-specific-block",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/create-a-batch-id",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/cancel-or-pause-a-scheduled-send",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/delete-a-cancellation-or-pause-from-a-scheduled-send",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/retrieve-all-scheduled-sends",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/retrieve-scheduled-send",
	"https://www.twilio.com/docs/sendgrid/api-reference/cancel-scheduled-sends/update-a-scheduled-send",
	"https://www.twilio.com/docs/sendgrid/api-reference/certificates/delete-an-sso-certificate",
	"https://www.twilio.com/docs/sendgrid/api-reference/certificates/get-an-sso-certificate",
	"https://www.twilio.com/docs/sendgrid/api-reference/certificates/get-all-sso-certificates-by-integration",
	"https://www.twilio.com/docs/sendgrid/api-reference/certificates/update-sso-certificate",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/create-design",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/delete-design",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/duplicate-design",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/get-design",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/list-designs",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/list-sendgrid-pre-built-designs",
	"https://www.twilio.com/docs/sendgrid/api-reference/designs-api/update-design",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/add-an-ip-to-an-authenticated-domain",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/authenticate-a-domain",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/delete-an-authenticated-domain",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/list-all-authenticated-domains",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/remove-an-ip-from-an-authenticated-domain",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/update-an-authenticated-domain",
	"https://www.twilio.com/docs/sendgrid/api-reference/domain-authentication/validate-a-domain-authentication",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/add-one-or-more-ips-to-the-allow-list",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/remove-one-or-more-ips-from-the-allow-list",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-a-list-of-currently-allowed-ips",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-a-specific-allowed-ip",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/remove-a-specific-ip-from-the-allowed-list",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-access-management/retrieve-all-recent-access-attempts",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-addresses/retrieve-all-ip-addresses",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-addresses/retrieve-all-assigned-ips",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-addresses/retrieve-remaining-ip-count",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-addresses/retrieve-an-ip-address",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/retrieve-all-ip-pools",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/create-an-ip-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/retrieve-an-ip-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/update-an-ip-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/delete-an-ip-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/add-an-ip-address-to-a-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-pools/remove-an-ip-address-from-a-pool",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/start-ip-warmup",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/stop-ip-warmup",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/retrieve-ip-warmup-status",
	"https://www.twilio.com/docs/sendgrid/api-reference/ip-warmup/retrieve-all-ip-warmup-statuses",
	"https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/retrieve-all-invalid-emails",
	"https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/retrieve-an-invalid-email",
	"https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/delete-invalid-emails",
	"https://www.twilio.com/docs/sendgrid/api-reference/invalid-emails/delete-a-specific-invalid-email",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/associate-a-branded-link-with-a-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/create-a-branded-link",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/delete-a-branded-link",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/disassociate-a-branded-link-from-a-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/retrieve-a-branded-link",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/retrieve-all-branded-links",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/update-a-branded-link",
	"https://www.twilio.com/docs/sendgrid/api-reference/link-branding/validate-a-branded-link",
	"https://www.twilio.com/docs/sendgrid/api-reference/mail-send/mail-send",
	"https://www.twilio.com/docs/sendgrid/api-reference/reverse-dns/delete-a-reverse-dns-record",
	"https://www.twilio.com/docs/sendgrid/api-reference/reverse-dns/retrieve-a-reverse-dns-record",
	"https://www.twilio.com/docs/sendgrid/api-reference/reverse-dns/retrieve-all-reverse-dns-records",
	"https://www.twilio.com/docs/sendgrid/api-reference/reverse-dns/set-up-reverse-dns",
	"https://www.twilio.com/docs/sendgrid/api-reference/reverse-dns/validate-a-reverse-dns-record",
	"https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/delete-verified-sender",
	"https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/domain-warn-list",
	"https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/get-all-verified-senders",
	"https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/resend-verified-sender-request",
	"https://www.twilio.com/docs/sendgrid/api-reference/sender-verification/verify-sender-request",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-enforced-tls/retrieve-current-enforced-tls-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-enforced-tls/update-enforced-tls-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-inbound-parse/create-a-parse-setting",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-inbound-parse/delete-a-parse-setting",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-inbound-parse/retrieve-a-specific-parse-setting",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-inbound-parse/retrieve-all-parse-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-inbound-parse/update-a-parse-setting",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-mail/update-bounce-purge-mail-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/retrieve-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/retrieve-click-track-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/get-open-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/retrieve-google-analytics-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/retrieve-subscription-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/update-click-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/update-google-analytics-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/update-open-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/settings-tracking/update-subscription-tracking-settings",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-settings/create-an-sso-integration",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-settings/delete-an-sso-integration",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-settings/get-all-sso-integrations",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-settings/get-an-sso-integration",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-settings/update-an-sso-integration",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-teammates/create-sso-teammate",
	"https://www.twilio.com/docs/sendgrid/api-reference/single-sign-on-teammates/edit-an-sso-teammate",
	"https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/retrieve-all-spam-reports",
	"https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/retrieve-a-specific-spam-report",
	"https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/delete-spam-reports",
	"https://www.twilio.com/docs/sendgrid/api-reference/spam-reports/delete-a-specific-spam-report",
	"https://www.twilio.com/docs/sendgrid/api-reference/stats/retrieve-global-email-statistics",
	"https://www.twilio.com/docs/sendgrid/api-reference/categories-statistics/retrieve-email-statistics-for-categories",
	"https://www.twilio.com/docs/sendgrid/api-reference/categories-statistics/retrieve-sums-of-email-stats-for-each-category",
	"https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-email-statistics-for-your-subusers",
	"https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-sums-of-email-stats-for-each-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/subuser-statistics/retrieve-monthly-stats-for-all-subusers",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/create-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/delete-a-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/list-all-subusers",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/retrieve-subuser-reputations",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/enabledisable-a-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/subusers-api/update-ips-assigned-to-a-subuser",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/add-recipient-addresses-to-the-global-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/delete-a-global-suppression",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/retrieve-a-global-suppression",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-global-suppressions/retrieve-all-global-suppressions",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/add-suppressions-to-a-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/delete-a-suppression-from-a-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppressions",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppression-groups-for-an-email-address",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/retrieve-all-suppressions-for-a-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-suppressions/search-for-suppressions-within-a-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/create-a-new-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/delete-a-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/get-information-on-a-single-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/retrieve-all-suppression-groups-associated-with-the-user",
	"https://www.twilio.com/docs/sendgrid/api-reference/suppressions-unsubscribe-groups/update-a-suppression-group",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/delete-teammate",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/invite-teammate",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/retrieve-all-teammates",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/retrieve-all-pending-teammates",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/resend-teammate-invite",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/delete-pending-teammate",
	"https://www.twilio.com/docs/sendgrid/api-reference/teammates/get-teammate-subuser-access",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/create-a-transactional-template",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/delete-a-template",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/retrieve-a-single-transactional-template",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/duplicate-a-transactional-template",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/edit-a-transactional-template",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates/retrieve-paged-transactional-templates",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates-versions/create-a-new-transactional-template-version",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates-versions/activate-a-transactional-template-version",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates-versions/delete-a-transactional-template-version",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates-versions/edit-a-transactional-template-version",
	"https://www.twilio.com/docs/sendgrid/api-reference/transactional-templates-versions/retrieve-a-specific-transactional-template-version",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/create-an-event-webhook",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/delete-an-event-webhook",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/get-all-event-webhooks",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/get-an-event-webhook",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/update-an-event-webhook",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/toggle-signature-verification-for-an-event-webhook",
	"https://www.twilio.com/docs/sendgrid/api-reference/webhooks/get-signed-event-webhooks-public-key",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-batched-contacts-by-ids",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/add-or-update-a-contact",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/import-contacts",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/import-contacts-status",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-a-contact-by-id",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-batched-contacts-by-ids",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-contacts-by-emails",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/search-contacts",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-sample-contacts",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-total-contact-count",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/export-contacts",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/export-contacts-status",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-all-existing-exports",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/delete-contacts",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/delete-a-contact-identifier",
	"https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-contacts-by-identifiers",
}

const apiRefPrefix = "/docs/sendgrid/api-reference/"

func urlToFilename(rawURL string) string {
	idx := strings.Index(rawURL, apiRefPrefix)
	if idx == -1 {
		return ""
	}
	path := rawURL[idx+len(apiRefPrefix):]
	name := strings.ReplaceAll(path, "/", "_")
	return name + ".md"
}

func download(client *http.Client, rawURL, outDir string, force bool) error {
	filename := urlToFilename(rawURL)
	if filename == "" {
		return fmt.Errorf("failed to extract filename from %s", rawURL)
	}

	outPath := filepath.Join(outDir, filename)
	if !force {
		if _, err := os.Stat(outPath); err == nil {
			// fmt.Printf("SKIP: %s (already exists)\n", outPath)
			return nil
		}
	}

	mdURL := rawURL + ".md"

	resp, err := client.Get(mdURL)
	if err != nil {
		return fmt.Errorf("GET %s: %w", mdURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: status %d", mdURL, resp.StatusCode)
	}

	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("create %s: %w", outPath, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("write %s: %w", outPath, err)
	}
	return nil
}

func main() {
	outDir := flag.String("out", "docs", "output directory")
	concurrency := flag.Int("c", 5, "max concurrent downloads")
	force := flag.Bool("force", false, "force download even if file already exists")
	flag.Parse()

	if err := os.MkdirAll(*outDir, 0o755); err != nil {
		fmt.Fprintf(os.Stderr, "mkdir %s: %v\n", *outDir, err)
		os.Exit(1)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	sem := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup
	var failed int64

	for _, u := range urls {
		wg.Add(1)
		go func(rawURL string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			if err := download(client, rawURL, *outDir, *force); err != nil {
				fmt.Fprintf(os.Stderr, "FAIL: %v\n", err)
				atomic.AddInt64(&failed, 1)
				return
			}
		}(u)
	}

	wg.Wait()

	total := len(urls)
	f := atomic.LoadInt64(&failed)
	fmt.Printf("\nDone: %d/%d succeeded, %d failed\n", int64(total)-f, total, f)

	if f > 0 {
		os.Exit(1)
	}
}
