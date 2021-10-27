package types

type CreateEnvelopeWithTemplatesResponse struct {
	Http                 int    `json:"http"`
	Message              string `json:"message"`
	Href                 string `json:"href"`
	EnvelopeTitle        string `json:"envelope_title"`
	EnvelopeFingerprint  string `json:"envelope_fingerprint"`
	EnvelopeSigningEmbed string `json:"envelope_signing_embed"`
	EnvelopeQueued       string `json:"envelope_queued"`
}

type CreateEnvelopeWithDocumentsResponse struct {
	Http                int    `json:"http"`
	Message             string `json:"message"`
	Href                string `json:"href"`
	EnvelopeTitle       string `json:"envelope_title"`
	EnvelopeFingerprint string `json:"envelope_fingerprint"`
	EnvelopeQueued      string `json:"envelope_queued"`
}

type ReadEnvelopeResponse struct {
	Http                int    `json:"http"`
	EnvelopeFingerprint string `json:"envelope_fingerprint"`
	EnvelopeTitle       string `json:"envelope_title"`
	EnvelopeStatus      string `json:"envelope_status"`
	EnvelopeRedirectUrl string `json:"envelope_redirect_url"`
	EnvelopeCreated     string `json:"envelope_created"`
	EnvelopeSent        string `json:"envelope_sent"`
	EnvelopeDocuments   []struct {
		DocumentFingerprint string   `json:"document_fingerprint"`
		DocumentTitle       string   `json:"document_title"`
		DocumentPageTotal   string   `json:"document_page_total"`
		DocumentPdfUrl      string   `json:"document_pdf_url"`
		DocumentThumbnails  []string `json:"document_thumbnails"`
		DocumentPages       []string `json:"document_pages"`
		DocumentFields      []struct {
			FieldId    string      `json:"field_id"`
			FieldTitle string      `json:"field_title"`
			FieldType  string      `json:"field_type"`
			FieldValue string      `json:"field_value"`
			PartyId    string      `json:"party_id"`
			FieldMerge interface{} `json:"field_merge"`
		} `json:"document_fields"`
		EnvelopeParties []struct {
			PartyId      string `json:"party_id"`
			PartyTitle   string `json:"party_title"`
			PartyStatus  string `json:"party_status"`
			ContactId    string `json:"contact_id"`
			ContactEmail string `json:"contact_email"`
		} `json:"envelope_parties"`
		EnvelopeSignedPdf string `json:"envelope_signed_pdf"`
		EnvelopeHistory   []struct {
			HistoryDetail    string `json:"history_detail"`
			HistoryIp        string `json:"history_ip"`
			HistoryUserAgent string `json:"history_user_agent"`
			HistoryDate      string `json:"history_date"`
		} `json:"envelope_history"`
		EnvelopeMeta struct {
			MyValue string `json:"my_value"`
		} `json:"envelope_meta"`
	} `json:"envelope_documents"`
}

type ListEnvelopesResponse struct {
	Http           int    `json:"http"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	TotalEnvelopes string `json:"total_envelopes"`
	Next           string `json:"next"`
	Href           string `json:"href"`
	Envelopes      []struct {
		EnvelopeFingerprint string      `json:"envelope_fingerprint"`
		EnvelopeTitle       string      `json:"envelope_title"`
		EnvelopeStatus      string      `json:"envelope_status"`
		EnvelopeRedirectUrl string      `json:"envelope_redirect_url"`
		EnvelopeCreated     string      `json:"envelope_created"`
		EnvelopeSent        interface{} `json:"envelope_sent"`
	} `json:"envelopes"`
}

type DeleteEnvelopeResponse struct {
	Http                int    `json:"http"`
	Message             string `json:"message"`
	EnvelopeFingerprint string `json:"envelope_fingerprint"`
	EnvelopeTitle       string `json:"envelope_title"`
}
