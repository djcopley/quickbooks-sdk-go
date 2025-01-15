package quickbooks

import (
	"bytes"
	"encoding/json"
	"github.com/djcopley/quickbooks-sdk-go/model"
	"io"
	"net/http"
	neturl "net/url"
)

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

func (c *Client) constructUrl(endpoint string) string {
	return c.apiUrl + endpoint
}

func (c *Client) get(url string) ([]byte, error) {
	resp, err := c.authClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (c *Client) post(url string, object model.QuickbooksEntity) ([]byte, error) {
	reqBody, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	resp, err := c.authClient.Post(url, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func (c *Client) CreateObject(object model.QuickbooksEntity, realmId string) (model.QuickbooksEntity, error) {
	entityName := object.GetObjectInfo().EntityName
	url, err := neturl.JoinPath(c.apiUrl, entityName, "company", realmId, entityName)
	if err != nil {
		return nil, err
	}
	resp, err := c.post(url, object)
	if err != nil {
		return nil, err
	}
	var result model.QuickbooksEntity
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) UpdateObject() {}

func (c *Client) DeleteObject() {}

func (c *Client) Query() {}

func (c *Client) BatchOperation() {}

func (c *Client) MiscOperation() {}
