package types

type ReadTemplateResponse struct {
	Http                 int    `json:"http"`
	TemplateId           string `json:"template_id"`
	TemplateFingerprint  string `json:"template_fingerprint"`
	TemplateTitle        string `json:"template_title"`
	TemplatePageTotal    string `json:"template_page_total"`
	TemplateInProgress   string `json:"template_in_progress"`
	TemplatePartiesTotal string `json:"template_parties_total"`
	TemplateParties      []struct {
		PartyId          string `json:"party_id"`
		PartyName        string `json:"party_name"`
		PartyMergeFields []struct {
			FieldId    string `json:"field_id"`
			FieldMerge string `json:"field_merge"`
			FieldType  string `json:"field_type"`
		} `json:"party_merge_fields"`
	} `json:"template_parties"`
	TemplateWidgetUrl   string   `json:"template_widget_url"`
	TemplateWidgetEmbed string   `json:"template_widget_embed"`
	TemplatePdfUrl      string   `json:"template_pdf_url"`
	TemplateThumbnails  []string `json:"template_thumbnails"`
	TemplatePages       []string `json:"template_pages"`
	TemplateUploaded    string   `json:"template_uploaded"`
}

type ListTemplatesResponse struct {
	Http           int    `json:"http"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	TotalTemplates string `json:"total_templates"`
	Next           string `json:"next"`
	Templates      []struct {
		TemplateId           string   `json:"template_id"`
		TemplateFingerprint  string   `json:"template_fingerprint"`
		TemplateTitle        string   `json:"template_title"`
		TemplatePageTotal    string   `json:"template_page_total"`
		TemplateInProgress   string   `json:"template_in_progress"`
		TemplatePartiesTotal string   `json:"template_parties_total"`
		TemplateWidgetUrl    string   `json:"template_widget_url"`
		TemplateWidgetEmbed  string   `json:"template_widget_embed"`
		TemplatePdfUrl       string   `json:"template_pdf_url"`
		TemplateThumbnails   []string `json:"template_thumbnails"`
		TemplatePages        []string `json:"template_pages"`
		TemplateParties      []struct {
			PartyId          string `json:"party_id"`
			PartyName        string `json:"party_name"`
			PartyMergeFields []struct {
				FieldId    string `json:"field_id"`
				FieldMerge string `json:"field_merge"`
				FieldType  string `json:"field_type"`
			} `json:"party_merge_fields"`
		} `json:"template_parties"`
		TemplateUploaded string `json:"template_uploaded"`
	} `json:"templates"`
}

type DeleteTemplateResponse struct {
	Http               int    `json:"http"`
	Message            string `json:"message"`
	TemplateId         string `json:"template_id"`
	TempateFingerprint string `json:"tempate_fingerprint"`
	TempateTitle       string `json:"tempate_title"`
}
