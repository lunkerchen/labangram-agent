package labangram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const DefaultBaseURL = "https://labangram.kamera-ichi.com"

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func New(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{BaseURL: DefaultBaseURL, HTTPClient: httpClient}
}

type Project struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Price struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}

type projectResponse struct {
	Projects []Project `json:"projects"`
}

type servicesResponse struct {
	Services []Service `json:"services"`
}

type pricingResponse struct {
	Pricing []Price `json:"pricing"`
}

type MCPResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      any             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *MCPError       `json:"error,omitempty"`
}

type MCPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) get(ctx context.Context, path string, target any) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+path, nil)
	if err != nil {
		return err
	}
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(response.Body, 4096))
		return fmt.Errorf("labangram: GET %s returned %s: %s", path, response.Status, bytes.TrimSpace(body))
	}
	return json.NewDecoder(response.Body).Decode(target)
}

func (c *Client) GetProjects(ctx context.Context, query, category string) ([]Project, error) {
	params := url.Values{}
	if query != "" {
		params.Set("q", query)
	}
	if category != "" {
		params.Set("category", category)
	}
	path := "/api/v1/projects"
	if encoded := params.Encode(); encoded != "" {
		path += "?" + encoded
	}
	var response projectResponse
	if err := c.get(ctx, path, &response); err != nil {
		return nil, err
	}
	return response.Projects, nil
}

func (c *Client) GetProject(ctx context.Context, slug string) (Project, error) {
	var project Project
	err := c.get(ctx, "/api/v1/projects/"+url.PathEscape(slug), &project)
	return project, err
}

func (c *Client) GetServices(ctx context.Context) ([]Service, error) {
	var response servicesResponse
	if err := c.get(ctx, "/api/v1/services", &response); err != nil {
		return nil, err
	}
	return response.Services, nil
}

func (c *Client) GetPricing(ctx context.Context) ([]Price, error) {
	var response pricingResponse
	if err := c.get(ctx, "/api/v1/pricing", &response); err != nil {
		return nil, err
	}
	return response.Pricing, nil
}

func (c *Client) CallMCP(ctx context.Context, method string, params any) (MCPResponse, error) {
	payload := map[string]any{"jsonrpc": "2.0", "id": 1, "method": method}
	if params != nil {
		payload["params"] = params
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return MCPResponse{}, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/api/product-mcp", bytes.NewReader(body))
	if err != nil {
		return MCPResponse{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json, text/event-stream")
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return MCPResponse{}, err
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return MCPResponse{}, fmt.Errorf("labangram: MCP returned %s", response.Status)
	}
	var result MCPResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return MCPResponse{}, err
	}
	return result, nil
}
