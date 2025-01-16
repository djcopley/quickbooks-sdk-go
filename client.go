package quickbooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/djcopley/quickbooks-sdk-go/model"
	"github.com/djcopley/quickbooks-sdk-go/model/batch"
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

func (c *Client) WithMinorVersion(minorVersion string) Client {
	client := *c
	client.minorVersion = minorVersion
	return client
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
	req.Header.Add("Accept", "application/json")
	addParams(req, map[string]string{"minorversion": c.minorVersion})
	addParams(req, params)
	resp, err := c.authClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to GET from '%s': %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, parseFailure(resp)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return respBody, nil
}

func (c *Client) post(url string, params map[string]string, contentType string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", "application/json")
	addParams(req, map[string]string{"minorversion": c.minorVersion})
	addParams(req, params)
	resp, err := c.authClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to POST to '%s': %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, parseFailure(resp)
	}
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
	resp, err := c.post(url, nil, "application/json", &body)
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
	resp, err := c.post(url, nil, "applications/json", &body)
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
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(object)
	if err != nil {
		return nil, fmt.Errorf("failed to encode quickbooks object: %w", err)
	}
	return c.post(url, params, "application/json", &body)
}

func (c *Client) Query(query string) ([]byte, error) {
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, "query")
	if err != nil {
		return nil, err
	}
	return c.post(url, nil, "application/text", bytes.NewBufferString(query))
}

func (c *Client) BatchOperation(batchRequest batch.Request) ([]byte, error) {
	url, err := neturl.JoinPath(c.apiUrl, "company", c.realmId, "batch")
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(batchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to encode quickbooks object: %w", err)
	}
	return c.post(url, nil, "application/json", &body)
}
