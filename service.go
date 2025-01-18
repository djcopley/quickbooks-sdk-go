package quickbooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/djcopley/quickbooks-sdk-go/batch"
	"io"
	"net/http"
	neturl "net/url"
)

type Environment string

const (
	Sandbox    Environment = "sandbox"
	Production Environment = "production"
)

type Service struct {
	baseURL      string
	realmID      string
	authClient   *http.Client
	minorVersion string
}

func NewService(environment Environment, authClient *http.Client, realmID string) *Service {
	var baseURL string

	switch environment {
	case Production:
		baseURL = "https://quickbooks.api.intuit.com/v3"
	case Sandbox:
		baseURL = "https://sandbox-quickbooks.api.intuit.com/v3"
	}

	return &Service{
		baseURL:      baseURL,
		realmID:      realmID,
		authClient:   authClient,
		minorVersion: "73",
	}
}

func (c *Service) WithMinorVersion(minorVersion string) *Service {
	return &Service{
		baseURL:      c.baseURL,
		realmID:      c.realmID,
		authClient:   c.authClient,
		minorVersion: minorVersion,
	}
}

func (c *Service) request(method, path string, params map[string]string, contentType string, body io.Reader) ([]byte, error) {
	url, err := neturl.JoinPath(c.baseURL, path)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	query := req.URL.Query()
	query.Set("minorversion", c.minorVersion)
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := c.authClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute %s request to '%s': %w", method, url, err)
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

func (c *Service) CreateEntity(entity Entity, response any) error {
	path := fmt.Sprintf("company/%s/%s", c.realmID, entity.GetEntityInfo().EntityName)
	body, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("failed to encode entity: %w", err)
	}
	resp, err := c.request("POST", path, nil, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("could not create entity: %w", err)
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}
	return nil
}

func (c *Service) UpdateEntity(entity Entity, response any) error {
	path := fmt.Sprintf("company/%s/%s", c.realmID, entity.GetEntityInfo().EntityName)
	body, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("failed to encode entity: %w", err)
	}
	resp, err := c.request("POST", path, nil, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("could not update entity: %w", err)
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}
	return nil
}

func (c *Service) DeleteEntity(entity Entity, response any) error {
	path := fmt.Sprintf("company/%s/%s", c.realmID, entity.GetEntityInfo().EntityName)
	params := map[string]string{"operation": "delete"}
	body, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("failed to encode entity: %w", err)
	}
	resp, err := c.request("POST", path, params, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("could not delete entity: %w", err)
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}
	return nil
}

func (c *Service) Query(query string, response any) error {
	path := fmt.Sprintf("company/%s/query", c.realmID)
	resp, err := c.request("POST", path, nil, "application/text", bytes.NewBufferString(query))
	if err != nil {
		return fmt.Errorf("could not query entity: %w", err)
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}
	return nil
}

func (c *Service) BatchOperation(batchRequest batch.Request, response any) error {
	path := fmt.Sprintf("company/%s/batch", c.realmID)
	body, err := json.Marshal(batchRequest)
	if err != nil {
		return fmt.Errorf("failed to encode batch request: %w", err)
	}
	resp, err := c.request("POST", path, nil, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("could not batch operation: %w", err)
	}
	err = json.Unmarshal(resp, &response)
	return nil
}
