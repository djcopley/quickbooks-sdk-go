package quickbooks

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	apiUrl       string
	authClient   *http.Client
	realmId      string
	minorVersion string
}

func NewClient(
	environment Environment,
	authClient *http.Client,
	realmId string,
) *Client {
	var apiUrl string

	switch environment {
	case Production:
		apiUrl = "https://quickbooks.api.intuit.com/v3"
	case Sandbox:
		apiUrl = "https://sandbox-quickbooks.api.intuit.com/v3"
	}

	client := &Client{
		apiUrl:       apiUrl,
		authClient:   authClient,
		realmId:      realmId,
		minorVersion: "73",
	}

	return client
}

func (c *Client) WithMinorVersion(minorVersion string) *Client {
	c.minorVersion = minorVersion
	return c
}

func addParams(req *http.Request, params map[string]string) {
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
}

func (c *Client) get(url string, params map[string]string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	addParams(req, map[string]string{"minorversion": c.minorVersion})
	addParams(req, params)
	resp, err := c.authClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to GET from '%s': %w", url, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return respBody, nil
}

func (c *Client) post(url string, params map[string]string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	addParams(req, map[string]string{"minorversion": c.minorVersion})
	addParams(req, params)
	resp, err := c.authClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to POST to '%s': %w", url, err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}
	return respBody, nil
}

func (c *Client) CreateObject(object model.QuickbooksEntity) ([]byte, error) {
	entityInfo := object.GetEntityInfo()
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, entityInfo.EntityName)
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(object)
	if err != nil {
		return nil, fmt.Errorf("failed to encode quickbooks object: %w", err)
	}
	resp, err := c.post(url, nil, &body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) UpdateObject(object model.QuickbooksEntity) ([]byte, error) {
	entityInfo := object.GetEntityInfo()
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, entityInfo.EntityName)
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(object)
	if err != nil {
		return nil, fmt.Errorf("failed to encode quickbooks object: %w", err)
	}
	resp, err := c.post(url, nil, &body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) DeleteObject(object model.QuickbooksEntity) ([]byte, error) {
	entityInfo := object.GetEntityInfo()
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, entityInfo.EntityName)
	if err != nil {
		return nil, err
	}
	params := map[string]string{
		"operation": "delete",
	}
	return c.post(url, params, nil)
}

func (c *Client) Query(query string) ([]byte, error) {
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, "query")
	if err != nil {
		return nil, err
	}
	return c.post(url, nil, bytes.NewBufferString(query))
}

func (c *Client) BatchOperation() ([]byte, error) {
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, "batch")
	if err != nil {
		return nil, err
	}
	return c.post(url, nil, nil)
}
