package quickbooks

import "net/http"

type Environment string

const (
	Sandbox    Environment = "sandbox"
	Production Environment = "production"
)

type Client struct {
	authClient *http.Client
	apiUrl     string
}

func NewClient(
	environment Environment,
	authClient *http.Client,
) *Client {
	var apiUrl string

	switch environment {
	case Production:
		apiUrl = "https://quickbooks.api.intuit.com/v3"
	case Sandbox:
		apiUrl = "https://sandbox-quickbooks.api.intuit.com/v3"
	}

	client := &Client{
		authClient: authClient,
		apiUrl:     apiUrl,
	}

	return client
}

func (c *Client) get() {}

func (c *Client) post() {}

func (c *Client) createObject() {}

func (c *Client) updateObject() {}

func (c *Client) deleteObject() {}

func (c *Client) query() {}

func (c *Client) batchOperation() {}

func (c *Client) miscOperation() {}
