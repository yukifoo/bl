package api

import (
	"github.com/kenzo0107/backlog"
)

// Client represents the Backlog API client
type Client struct {
	backlog *backlog.Client
	baseURL string
	apiKey  string
}

// NewClient creates a new Backlog API client
func NewClient(baseURL, apiKey string) *Client {
	client := backlog.New(baseURL, apiKey)
	return &Client{
		backlog: client,
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}