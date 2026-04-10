package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.green-api.com"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) buildURL(idInstance, apiToken, method string) string {
	return fmt.Sprintf("%s/waInstance%s/%s/%s", baseURL, idInstance, method, apiToken)
}

func (c *Client) GetSettings(idInstance, apiToken string) (map[string]interface{}, error) {
	url := c.buildURL(idInstance, apiToken, "getSettings")
	return c.doGet(url)
}

func (c *Client) GetStateInstance(idInstance, apiToken string) (map[string]interface{}, error) {
	url := c.buildURL(idInstance, apiToken, "getStateInstance")
	return c.doGet(url)
}

func (c *Client) SendMessage(idInstance, apiToken, chatID, message string) (map[string]interface{}, error) {
	url := c.buildURL(idInstance, apiToken, "sendMessage")
	body := map[string]string{
		"chatId":  chatID,
		"message": message,
	}
	return c.doPost(url, body)
}

func (c *Client) SendFileByURL(idInstance, apiToken, chatID, fileURL, fileName, caption string) (map[string]interface{}, error) {
	url := c.buildURL(idInstance, apiToken, "sendFileByUrl")
	body := map[string]string{
		"chatId":   chatID,
		"urlFile":  fileURL,
		"fileName": fileName,
	}
	if caption != "" {
		body["caption"] = caption
	}
	return c.doPost(url, body)
}

func (c *Client) doGet(url string) (map[string]interface{}, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	return c.parseResponse(resp)
}

func (c *Client) doPost(url string, body interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal body: %w", err)
	}
	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	return c.parseResponse(resp)
}

func (c *Client) parseResponse(resp *http.Response) (map[string]interface{}, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return result, nil
}
