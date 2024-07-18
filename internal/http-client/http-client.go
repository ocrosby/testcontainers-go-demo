package http_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

type Client struct {
	baseUrl    string
	httpClient *http.Client
}

// NewClient creates a new Client instance
func NewClient(baseUrl string) (*Client, error) {
	if baseUrl == "" {
		return nil, errors.New("invalid baseUrl provided")
	}

	httpClient := http.Client{Timeout: time.Duration(5) * time.Second}

	return &Client{
		baseUrl:    baseUrl,
		httpClient: &httpClient,
	}, nil
}

func (c *Client) GetAge(name string) (*AgeResponse, error) {
	url := fmt.Sprintf("%s?name=%s", c.baseUrl, name)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		err = errors.New(fmt.Sprintf("error: %v", err))
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			err = errors.New(fmt.Sprintf("error: %v", err))
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("error: %v", resp.Status))
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	fmt.Printf("Body: %s\n", body)

	ageResponse := &AgeResponse{}

	// Unmarshal the response body into the AgeResponse struct
	if err = json.Unmarshal(body, &ageResponse); err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return ageResponse, nil
}
