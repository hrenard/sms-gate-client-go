package smsgateway

// The type of event a webhook can be triggered for.
type WebhookEvent string

const (
	// Triggered when an SMS is received.
	WebhookEventSmsReceived WebhookEvent = "sms:received"
)

// A webhook configuration.
type Webhook struct {
	// The unique identifier of the webhook.
	ID string `json:"id" validate:"max=36" example:"123e4567-e89b-12d3-a456-426614174000"`

	// The URL the webhook will be sent to.
	URL string `json:"url" validate:"required,http_url" example:"https://example.com/webhook"`

	// The type of event the webhook is triggered for.
	Event WebhookEvent `json:"event" validate:"required" example:"sms:received"`
}
