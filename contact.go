package sendgrid

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type Contact struct {
	ID                  string                 `json:"id,omitempty"`
	FirstName           string                 `json:"first_name,omitempty"`
	LastName            string                 `json:"last_name,omitempty"`
	UniqueName          string                 `json:"unique_name,omitempty"`
	Email               string                 `json:"email,omitempty"`
	PhoneNumberID       string                 `json:"phone_number_id,omitempty"`
	ExternalID          string                 `json:"external_id,omitempty"`
	AnonymousID         string                 `json:"anonymous_id,omitempty"`
	AlternateEmails     []string               `json:"alternate_emails,omitempty"`
	AddressLine1        string                 `json:"address_line_1,omitempty"`
	AddressLine2        string                 `json:"address_line_2,omitempty"`
	City                string                 `json:"city,omitempty"`
	StateProvinceRegion string                 `json:"state_province_region,omitempty"`
	Country             string                 `json:"country,omitempty"`
	PostalCode          string                 `json:"postal_code,omitempty"`
	PhoneNumber         string                 `json:"phone_number,omitempty"`
	Whatsapp            string                 `json:"whatsapp,omitempty"`
	Line                string                 `json:"line,omitempty"`
	Facebook            string                 `json:"facebook,omitempty"`
	ListIDs             []string               `json:"list_ids,omitempty"`
	SegmentIDs          []string               `json:"segment_ids,omitempty"`
	CustomFields        map[string]interface{} `json:"custom_fields,omitempty"`
	CreatedAt           string                 `json:"created_at,omitempty"`
	UpdatedAt           string                 `json:"updated_at,omitempty"`
	Metadata            *_Metadata             `json:"_metadata,omitempty"`
}

// ---

type ContactRequest struct {
	AddressLine1        string                 `json:"address_line_1,omitempty"`
	AddressLine2        string                 `json:"address_line_2,omitempty"`
	AlternateEmails     []string               `json:"alternate_emails,omitempty"`
	City                string                 `json:"city,omitempty"`
	Country             string                 `json:"country,omitempty"`
	Email               string                 `json:"email,omitempty"`
	PhoneNumberID       string                 `json:"phone_number_id,omitempty"`
	ExternalID          string                 `json:"external_id,omitempty"`
	AnonymousID         string                 `json:"anonymous_id,omitempty"`
	FirstName           string                 `json:"first_name,omitempty"`
	LastName            string                 `json:"last_name,omitempty"`
	PostalCode          string                 `json:"postal_code,omitempty"`
	StateProvinceRegion string                 `json:"state_province_region,omitempty"`
	CustomFields        map[string]interface{} `json:"custom_fields,omitempty"`
}

type InputAddOrUpdateContacts struct {
	ListIDs  []string          `json:"list_ids,omitempty"`
	Contacts []*ContactRequest `json:"contacts"`
}

type OutputAddOrUpdateContacts struct {
	JobID string `json:"job_id,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/add-or-update-a-contact
func (c *Client) AddOrUpdateContacts(ctx context.Context, input *InputAddOrUpdateContacts) (*OutputAddOrUpdateContacts, error) {
	req, err := c.NewRequest("PUT", "/marketing/contacts", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputAddOrUpdateContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputImportContacts struct {
	ListIDs       []string `json:"list_ids,omitempty"`
	FileType      string   `json:"file_type"`
	FieldMappings []string `json:"field_mappings"`
}

type ImportContactsUploadHeader struct {
	Header string `json:"header,omitempty"`
	Value  string `json:"value,omitempty"`
}

type OutputImportContacts struct {
	JobID         string                        `json:"job_id,omitempty"`
	UploadURI     string                        `json:"upload_uri,omitempty"`
	UploadHeaders []*ImportContactsUploadHeader `json:"upload_headers,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/import-contacts
func (c *Client) ImportContacts(ctx context.Context, input *InputImportContacts) (*OutputImportContacts, error) {
	req, err := c.NewRequest("PUT", "/marketing/contacts/imports", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputImportContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ImportContactsResults struct {
	RequestedCount int64  `json:"requested_count,omitempty"`
	CreatedCount   int64  `json:"created_count,omitempty"`
	UpdatedCount   int64  `json:"updated_count,omitempty"`
	DeletedCount   int64  `json:"deleted_count,omitempty"`
	ErroredCount   int64  `json:"errored_count,omitempty"`
	ErrorsURL      string `json:"errors_url,omitempty"`
}

type OutputGetImportContactsStatus struct {
	ID         string                 `json:"id,omitempty"`
	Status     string                 `json:"status,omitempty"`
	JobType    string                 `json:"job_type,omitempty"`
	Results    *ImportContactsResults `json:"results,omitempty"`
	StartedAt  string                 `json:"started_at,omitempty"`
	FinishedAt string                 `json:"finished_at,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/import-contacts-status
func (c *Client) GetImportContactsStatus(ctx context.Context, id string) (*OutputGetImportContactsStatus, error) {
	path := fmt.Sprintf("/marketing/contacts/imports/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetImportContactsStatus)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetContact struct {
	Contact *Contact `json:"contact,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-a-contact-by-id
func (c *Client) GetContact(ctx context.Context, id string) (*OutputGetContact, error) {
	path := fmt.Sprintf("/marketing/contacts/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetContact)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputGetBatchedContactsByIDs struct {
	IDs []string `json:"ids"`
}

type OutputGetBatchedContactsByIDs struct {
	Result []*Contact `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-batched-contacts-by-ids
func (c *Client) GetBatchedContactsByIDs(ctx context.Context, input *InputGetBatchedContactsByIDs) (*OutputGetBatchedContactsByIDs, error) {
	req, err := c.NewRequest("POST", "/marketing/contacts/batch", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetBatchedContactsByIDs)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ContactByEmailResult struct {
	Contact *Contact `json:"contact,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type InputGetContactsByEmails struct {
	Emails        []string `json:"emails"`
	PhoneNumberID string   `json:"phone_number_id,omitempty"`
	ExternalID    string   `json:"external_id,omitempty"`
	AnonymousID   string   `json:"anonymous_id,omitempty"`
}

type OutputGetContactsByEmails struct {
	Result map[string]*ContactByEmailResult `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-contacts-by-emails
func (c *Client) GetContactsByEmails(ctx context.Context, input *InputGetContactsByEmails) (*OutputGetContactsByEmails, error) {
	req, err := c.NewRequest("POST", "/marketing/contacts/search/emails", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetContactsByEmails)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputSearchContacts struct {
	Query string `json:"query,omitempty"`
}

type OutputSearchContacts struct {
	Result       []*Contact `json:"result,omitempty"`
	ContactCount int64      `json:"contact_count,omitempty"`
	Metadata     _Metadata  `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/search-contacts
func (c *Client) SearchContacts(ctx context.Context, input *InputSearchContacts) (*OutputSearchContacts, error) {
	req, err := c.NewRequest("POST", "/marketing/contacts/search", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputSearchContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetSampleContacts struct {
	Result       []*Contact `json:"result,omitempty"`
	ContactCount int64      `json:"contact_count,omitempty"`
	Metadata     _Metadata  `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-sample-contacts
func (c *Client) GetSampleContacts(ctx context.Context) (*OutputGetSampleContacts, error) {
	req, err := c.NewRequest("GET", "/marketing/contacts", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetSampleContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ContactBillableBreakdown struct {
	Total     int64            `json:"total,omitempty"`
	Breakdown map[string]int64 `json:"breakdown,omitempty"`
}

type OutputGetTotalContactCount struct {
	ContactCount      int64                     `json:"contact_count,omitempty"`
	BillableCount     int64                     `json:"billable_count,omitempty"`
	BillableBreakdown *ContactBillableBreakdown `json:"billable_breakdown,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-total-contact-count
func (c *Client) GetTotalContactCount(ctx context.Context) (*OutputGetTotalContactCount, error) {
	req, err := c.NewRequest("GET", "/marketing/contacts/count", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetTotalContactCount)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ExportContactsNotification struct {
	Email bool `json:"email,omitempty"`
}

type InputExportContacts struct {
	ListIDs       []string                    `json:"list_ids,omitempty"`
	SegmentIDs    []string                    `json:"segment_ids,omitempty"`
	Notifications *ExportContactsNotification `json:"notifications,omitempty"`
	FileType      string                      `json:"file_type,omitempty"`
	MaxFileSize   int64                       `json:"max_file_size,omitempty"`
}

type OutputExportContacts struct {
	ID       string    `json:"id,omitempty"`
	Metadata _Metadata `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/export-contacts
func (c *Client) ExportContacts(ctx context.Context, input *InputExportContacts) (*OutputExportContacts, error) {
	req, err := c.NewRequest("POST", "/marketing/contacts/exports", input)
	if err != nil {
		return nil, err
	}

	r := new(OutputExportContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type OutputGetExportContactsStatus struct {
	ID           string    `json:"id,omitempty"`
	Status       string    `json:"status,omitempty"`
	CreatedAt    string    `json:"created_at,omitempty"`
	UpdatedAt    string    `json:"updated_at,omitempty"`
	CompletedAt  string    `json:"completed_at,omitempty"`
	ExpiresAt    string    `json:"expires_at,omitempty"`
	URLs         []string  `json:"urls,omitempty"`
	Message      string    `json:"message,omitempty"`
	Metadata     _Metadata `json:"_metadata,omitempty"`
	ContactCount int64     `json:"contact_count,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/export-contacts-status
func (c *Client) GetExportContactsStatus(ctx context.Context, id string) (*OutputGetExportContactsStatus, error) {
	path := fmt.Sprintf("/marketing/contacts/exports/%s", id)

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetExportContactsStatus)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ExportJobSegment struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type ExportJobList struct {
	ID   string `json:"ID,omitempty"`
	Name string `json:"Name,omitempty"`
}

type ContactExportJob struct {
	ID          string              `json:"id,omitempty"`
	Status      string              `json:"status,omitempty"`
	CreatedAt   string              `json:"created_at,omitempty"`
	CompletedAt string              `json:"completed_at,omitempty"`
	ExpiresAt   string              `json:"expires_at,omitempty"`
	URLs        []string            `json:"urls,omitempty"`
	UserID      string              `json:"user_id,omitempty"`
	ExportType  string              `json:"export_type,omitempty"`
	Segments    []*ExportJobSegment `json:"segments,omitempty"`
	Lists       []*ExportJobList    `json:"lists,omitempty"`
	Metadata    *_Metadata          `json:"_metadata,omitempty"`
}

type OutputGetAllExistingExports struct {
	Result   []*ContactExportJob `json:"result,omitempty"`
	Metadata *_Metadata          `json:"_metadata,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-all-existing-exports
func (c *Client) GetAllExistingExports(ctx context.Context) (*OutputGetAllExistingExports, error) {
	req, err := c.NewRequest("GET", "/marketing/contacts/exports", nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetAllExistingExports)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputDeleteContacts struct {
	IDs               []string `url:"ids,omitempty"`
	DeleteAllContacts string   `url:"delete_all_contacts,omitempty"`
}

type OutputDeleteContacts struct {
	JobID interface{} `json:"job_id,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/delete-contacts
func (c *Client) DeleteContacts(ctx context.Context, input *InputDeleteContacts) (*OutputDeleteContacts, error) {
	u, _ := url.Parse("/marketing/contacts")

	q := u.Query()
	if len(input.IDs) > 0 {
		q.Set("ids", strings.Join(input.IDs, ","))
	} else if input.DeleteAllContacts != "" {
		q.Set("delete_all_contacts", input.DeleteAllContacts)
	}
	u.RawQuery = q.Encode()

	req, err := c.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}

	r := new(OutputDeleteContacts)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type InputDeleteContactIdentifier struct {
	IdentifierType  string `json:"identifier_type"`
	IdentifierValue string `json:"identifier_value"`
}

type OutputDeleteContactIdentifier struct {
	JobID interface{} `json:"job_id,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/delete-a-contact-identifier
func (c *Client) DeleteContactIdentifier(ctx context.Context, contactID string, input *InputDeleteContactIdentifier) (*OutputDeleteContactIdentifier, error) {
	path := fmt.Sprintf("/marketing/contacts/%s/identifiers", contactID)

	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputDeleteContactIdentifier)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}

type ContactByIdentifierResult struct {
	Contact *Contact `json:"contact,omitempty"`
	Error   string   `json:"error,omitempty"`
}

type InputGetContactsByIdentifiers struct {
	Identifiers []string `json:"identifiers"`
}

type OutputGetContactsByIdentifiers struct {
	Result map[string]*ContactByIdentifierResult `json:"result,omitempty"`
}

// see: https://www.twilio.com/docs/sendgrid/api-reference/contacts/get-contacts-by-identifiers
func (c *Client) GetContactsByIdentifiers(ctx context.Context, identifierType string, input *InputGetContactsByIdentifiers) (*OutputGetContactsByIdentifiers, error) {
	path := fmt.Sprintf("/marketing/contacts/search/identifiers/%s", identifierType)

	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	r := new(OutputGetContactsByIdentifiers)
	if err := c.Do(ctx, req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
