package types

type ListCompanySettingsType struct {
	Http                                 int    `json:"http"`
	SettingSignatureMoreInfo             bool   `json:"setting_signature_more_info"`
	SettingSignatureFormatAccepted       string `json:"setting_signature_format_accepted"`
	SettingSignatureFormatDefault        bool   `json:"setting_signature_format_default"`
	SettingEmailPdfAttach                bool   `json:"setting_email_pdf_attach"`
	SettingEmailEnvelopeOpen             bool   `json:"setting_email_envelope_open"`
	SettingEmailMainEmail                bool   `json:"setting_email_main_email"`
	SettingEmailReturnAddressUser        bool   `json:"setting_email_return_address_user"`
	SettingSigningConsumerRegulationsAct bool   `json:"setting_signing_consumer_regulations_act"`
	SettingSigningTooltips               bool   `json:"setting_signing_tooltips"`
	SettingSigningQuestions              bool   `json:"setting_signing_questions"`
	SettingPdfAuditTrailHide             bool   `json:"setting_pdf_audit_trail_hide"`
}

type UpdateCompanySettingsType struct {
	Http                                 int    `json:"http"`
	Message                              string `json:"message"`
	Href                                 string `json:"href"`
	CompanyId                            string `json:"company_id"`
	CompanyName                          string `json:"company_name"`
	SettingSignatureMoreInfo             bool   `json:"setting_signature_more_info"`
	SettingSignatureFormatAccepted       string `json:"setting_signature_format_accepted"`
	SettingSignatureFormatDefault        string `json:"setting_signature_format_default"`
	SettingEmailPdfAttach                bool   `json:"setting_email_pdf_attach"`
	SettingEmailEnvelopeOpen             bool   `json:"setting_email_envelope_open"`
	SettingEmailMainEmail                bool   `json:"setting_email_main_email"`
	SettingEmailReturnAddressUser        bool   `json:"setting_email_return_address_user"`
	SettingSigningConsumerRegulationsAct bool   `json:"setting_signing_consumer_regulations_act"`
	SettingSigningTooltips               bool   `json:"setting_signing_tooltips"`
	SettingSigningQuestions              bool   `json:"setting_signing_questions"`
	SettingPdfAuditTrailHide             bool   `json:"setting_pdf_audit_trail_hide"`
}
