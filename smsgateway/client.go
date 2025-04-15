package smsgateway

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/android-sms-gateway/client-go/rest"
)

//nolint:revive // backward compatibility
const BASE_URL = "https://api.sms-gate.app/3rdparty/v1"

type Config struct {
	Client   *http.Client // Optional HTTP Client, defaults to `http.DefaultClient`
	BaseURL  string       // Optional base URL, defaults to `https://api.sms-gate.app/3rdparty/v1`
	User     string       // Required username
	Password string       // Required password
}

type Client struct {
	*rest.Client

	headers map[string]string
}

// Sends an SMS message.
func (c *Client) Send(ctx context.Context, message Message) (MessageState, error) {
	path := "/message"
	resp := new(MessageState)

	if err := c.Do(ctx, http.MethodPost, path, c.headers, &message, resp); err != nil {
		return *resp, fmt.Errorf("failed to send message: %w", err)
	}

	return *resp, nil
}

// Gets the state of an SMS message by ID.
func (c *Client) GetState(ctx context.Context, messageID string) (MessageState, error) {
	path := fmt.Sprintf("/message/%s", messageID)
	resp := new(MessageState)

	if err := c.Do(ctx, http.MethodGet, path, c.headers, nil, resp); err != nil {
		return *resp, fmt.Errorf("failed to get message state: %w", err)
	}

	return *resp, nil
}

// ListWebhooks retrieves all registered webhooks.
// Returns a slice of Webhook objects or an error if the request fails.
func (c *Client) ListWebhooks(ctx context.Context) ([]Webhook, error) {
	path := "/webhooks"
	resp := []Webhook{}

	if err := c.Do(ctx, http.MethodGet, path, c.headers, nil, &resp); err != nil {
		return resp, fmt.Errorf("failed to list webhooks: %w", err)
	}

	return resp, nil
}

// RegisterWebhook registers a new webhook.
// Returns the registered webhook with server-assigned fields or an error if the request fails.
func (c *Client) RegisterWebhook(ctx context.Context, webhook Webhook) (Webhook, error) {
	path := "/webhooks"
	resp := new(Webhook)

	if err := c.Do(ctx, http.MethodPost, path, c.headers, &webhook, resp); err != nil {
		return *resp, fmt.Errorf("failed to register webhook: %w", err)
	}

	return *resp, nil
}

// DeleteWebhook removes a webhook with the specified ID.
// Returns an error if the deletion fails.
func (c *Client) DeleteWebhook(ctx context.Context, webhookID string) error {
	path := fmt.Sprintf("/webhooks/%s", url.PathEscape(webhookID))

	if err := c.Do(ctx, http.MethodDelete, path, c.headers, nil, nil); err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}

	return nil
}

// ListDevices retrieves all registered devices in the account.
// Returns a slice of Device objects or an error if the request fails.
func (c *Client) ListDevices(ctx context.Context) ([]Device, error) {
	path := "/device"
	resp := []Device{}

	if err := c.Do(ctx, http.MethodGet, path, c.headers, nil, &resp); err != nil {
		return resp, fmt.Errorf("failed to list devices: %w", err)
	}

	return resp, nil
}

// NewClient creates a new instance of the API Client.
func NewClient(config Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = BASE_URL
	}

	return &Client{
		Client: rest.NewClient(rest.Config{
			Client:  config.Client,
			BaseURL: config.BaseURL,
		}),
		headers: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(config.User+":"+config.Password)),
		},
	}
}
