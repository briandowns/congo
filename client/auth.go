package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// OAUTH2 callback endpoint
func (c *Client) CallbackAuth(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// OAUTH2 login endpoint
func (c *Client) OauthAuth(path string) (*http.Response, error) {
	var body io.Reader
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// RefreshAuthPayload is the data structure used to initialize the auth refresh request body.
type RefreshAuthPayload struct {
	// UUID of requesting application
	Application string `json:"application,omitempty"`
	// email
	Email string `json:"email,omitempty"`
	// password
	Password string `json:"password,omitempty"`
}

// Obtain a refreshed access token
func (c *Client) RefreshAuth(path string, payload *RefreshAuthPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}

// TokenAuthPayload is the data structure used to initialize the auth token request body.
type TokenAuthPayload struct {
	// UUID of requesting application
	Application string `json:"application,omitempty"`
	// email
	Email string `json:"email,omitempty"`
	// password
	Password string `json:"password,omitempty"`
}

// Obtain an access token
func (c *Client) TokenAuth(path string, payload *TokenAuthPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	u := url.URL{Host: c.Host, Scheme: c.Scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(req)
}
