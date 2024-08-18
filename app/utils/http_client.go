// httpclient/httpclient.go
package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client:  &http.Client{Timeout: 100 * time.Second},
		headers: make(map[string]string),
	}
}

func (c *HTTPClient) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *HTTPClient) SetHeader(key, value string) {
	c.headers[key] = value
}

// GET method to perform HTTP GET requests
func (c *HTTPClient) Get(endpoint string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return "", err
	}

	// Set request headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body and decode it into the provided response object
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Received Details: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		return string(body), fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(body))
	}

	return string(body), nil
}

// Post method to perform HTTP POST requests
func (c *HTTPClient) Post(endpoint string, body *bytes.Buffer) (string, error) {
	req, err := http.NewRequest(http.MethodPost, c.baseURL+endpoint, body)
	if err != nil {
		return "", err
	}

	// Set request headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Received Response: %s\n", string(responseBody))

	// Check for success based on status code
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return string(responseBody), nil
	}

	return string(responseBody), fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(responseBody))
}

// PostForm handles multipart form data, including optional file uploads
func (c *HTTPClient) PostForm(endpoint string, formData map[string]string, files map[string]string) (string, error) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Add form fields
	for key, value := range formData {
		writer.WriteField(key, value)
	}

	// Add file fields if any
	for key, filepath := range files {
		if filepath == "" {
			continue // Skip empty file paths
		}

		file, err := os.Open(filepath)
		if err != nil {
			return "", fmt.Errorf("failed to open file %s: %v", filepath, err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile(key, filepath)
		if err != nil {
			return "", fmt.Errorf("failed to create form file %s: %v", filepath, err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return "", fmt.Errorf("failed to copy file content: %v", err)
		}
	}

	// Close the writer to set the terminating boundary
	writer.Close()

	// Create the request
	req, err := http.NewRequest(http.MethodPost, c.baseURL+endpoint, &b)
	if err != nil {
		return "", err
	}

	// Set request headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Check for success based on status code
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return string(responseBody), nil
	}

	return "", fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(responseBody))
}

// DELETE method to perform HTTP DELETE requests
func (c *HTTPClient) Delete(endpoint string) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, c.baseURL+endpoint, nil)
	if err != nil {
		return "", err
	}

	// Set request headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body and decode it into the provided response object
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Received UserDetails: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		return string(body), fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(body))
	}

	return string(body), nil
}

// Put method to perform HTTP PUT requests
func (c *HTTPClient) Put(endpoint string, body *bytes.Buffer) (string, error) {
	req, err := http.NewRequest(http.MethodPut, c.baseURL+endpoint, body)
	if err != nil {
		return "", err
	}

	// Set request headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("Received Response: %s\n", string(responseBody))

	// Check for success based on status code
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return string(responseBody), nil
	}

	return string(responseBody), fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(responseBody))
}
