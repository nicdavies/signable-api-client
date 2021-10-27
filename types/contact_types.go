package types

type ListContactsOptions struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListContactsResponse struct {
	Http          int    `json:"http"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	TotalContacts string `json:"total_contacts"`
	Next          string `json:"next"`
	Contacts      []struct {
		ContactId                   string `json:"contact_id"`
		ContactName                 string `json:"contact_name"`
		ContactEmail                string `json:"contact_email"`
		ContactOutstandingDocuments string `json:"contact_outstanding_documents"`
		ContactCreated              string `json:"contact_created"`
		ContactApps                 struct {
			CapsuleId  string `json:"capsule_id"`
			KashflowId string `json:"kashflow_id"`
		} `json:"contact_apps"`
	} `json:"contacts"`
}

type CreateContactsOptions struct {
	Name  string `json:"contact_name"`
	Email string `json:"contact_email"`
}

type CreateContactsResponse struct {
	Http           int    `json:"http"`
	Message        string `json:"message"`
	Href           string `json:"href"`
	ContactId      string `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	ContactEmail   string `json:"contact_email"`
	ContactCreated string `json:"contact_created"`
}

type ReadContactOptions struct {
	Id int
}

type ReadContactResponse struct {
	Http                        int    `json:"http"`
	ContactId                   int    `json:"contact_id"`
	ContactName                 string `json:"contact_name"`
	ContactEmail                string `json:"contact_email"`
	ContactOutstandingDocuments string `json:"contact_outstanding_documents"`
	ContactCreated              string `json:"contact_created"`
	ContactApps                 struct {
		CapsuleId  string `json:"capsule_id"`
		KashflowId string `json:"kashflow_id"`
	} `json:"contact_apps"`
}

type UpdateContactOptions struct {
	Name  string `json:"contact_name"`
	Email string `json:"contact_email"`
}

type UpdateContactResponse struct {
	Http           int    `json:"http"`
	Message        string `json:"message"`
	Href           string `json:"href"`
	ContactId      int    `json:"contact_id"`
	ContactName    string `json:"contact_name"`
	ContactEmail   string `json:"contact_email"`
	ContactUpdated string `json:"contact_updated"`
}

type DeleteContactOptions struct {
	Id int
}

type DeleteContactResponse struct {
	Http         int    `json:"http"`
	Message      string `json:"message"`
	ContactId    int    `json:"contact_id"`
	ContactName  string `json:"contact_name"`
	ContactEmail string `json:"contact_email"`
}

type ReadContactEnvelopesOptions struct {
	Id int
}

type ReadContactEnvelopesResponse struct {
	Http           int    `json:"http"`
	TotalEnvelopes string `json:"total_envelopes"`
	Envelopes      []struct {
		EnvelopeFingerprint string  `json:"envelope_fingerprint"`
		EnvelopeTitle       string  `json:"envelope_title"`
		EnvelopeStatus      string  `json:"envelope_status"`
		EnvelopeCreated     string  `json:"envelope_created"`
		EnvelopeSent        string  `json:"envelope_sent"`
		EnvelopeProcessed   *string `json:"envelope_processed"`
		EnvelopeSignedPdf   string  `json:"envelope_signed_pdf,omitempty"`
	} `json:"envelopes"`
}
