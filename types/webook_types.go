package types

type CreateWebhookResponse struct {
	Http           int    `json:"http"`
	Message        string `json:"message"`
	Href           string `json:"href"`
	WebhookId      string `json:"webhook_id"`
	WebhookType    string `json:"webhook_type"`
	WebhookUrl     string `json:"webhook_url"`
	WebhookCreated string `json:"webhook_created"`
}

type ReadWebhookResponse struct {
	Http           int    `json:"http"`
	WebhookId      int    `json:"webhook_id"`
	WebhookType    string `json:"webhook_type"`
	WebhookUrl     string `json:"webhook_url"`
	WebhookCreated string `json:"webhook_created"`
}

type ListWebhooksResponse struct {
	Http          int    `json:"http"`
	Offset        int    `json:"offset"`
	Limit         int    `json:"limit"`
	TotalWebhooks string `json:"total_webhooks"`
	Next          string `json:"next"`
	Webhooks      []struct {
		WebhookId      string `json:"webhook_id"`
		WebhookType    string `json:"webhook_type"`
		WebhookUrl     string `json:"webhook_url"`
		WebhookCreated string `json:"webhook_created"`
	} `json:"webhooks"`
}

type UpdateWebhookResponse struct {
	Http           int    `json:"http"`
	Message        string `json:"message"`
	Href           string `json:"href"`
	WebhookId      int    `json:"webhook_id"`
	WebhookUrl     string `json:"webhook_url"`
	WebhookType    string `json:"webhook_type"`
	WebhookUpdated string `json:"webhook_updated"`
}

type DeleteWebhookResponse struct {
	Http      int    `json:"http"`
	Message   string `json:"message"`
	WebhookId int    `json:"webhook_id"`
}
