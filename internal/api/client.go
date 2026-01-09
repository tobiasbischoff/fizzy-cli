package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	BaseURL      string
	Token        string
	SessionToken string
	Agent        string
	HTTP         *http.Client
}

type Response struct {
	Status  int
	Headers http.Header
	Body    []byte
}

type APIError struct {
	Status int
	Body   []byte
}

func (e *APIError) Error() string {
	msg := strings.TrimSpace(string(e.Body))
	if msg == "" {
		return fmt.Sprintf("api error: status %d", e.Status)
	}
	return fmt.Sprintf("api error: status %d: %s", e.Status, msg)
}

func NewClient(baseURL, token, sessionToken, agent string) *Client {
	return &Client{
		BaseURL:      strings.TrimRight(baseURL, "/"),
		Token:        token,
		SessionToken: sessionToken,
		Agent:        agent,
		HTTP:         &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) Do(ctx context.Context, method, path string, query url.Values, body io.Reader, contentType string, headers map[string]string) (*Response, error) {
	urlStr, err := buildURL(c.BaseURL, path, query)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	if body != nil {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, body); err != nil {
			return nil, err
		}
		bodyReader = buf
	}

	req, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.Agent != "" {
		req.Header.Set("User-Agent", c.Agent)
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	} else if c.SessionToken != "" {
		if headers == nil {
			req.Header.Set("Cookie", "session_token="+c.SessionToken)
		} else if _, ok := headers["Cookie"]; !ok {
			req.Header.Set("Cookie", "session_token="+c.SessionToken)
		}
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, &APIError{Status: resp.StatusCode, Body: respBody}
	}

	return &Response{
		Status:  resp.StatusCode,
		Headers: resp.Header,
		Body:    respBody,
	}, nil
}

func buildURL(baseURL, path string, query url.Values) (string, error) {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		u, err := url.Parse(path)
		if err != nil {
			return "", err
		}
		if query != nil {
			u.RawQuery = query.Encode()
		}
		return u.String(), nil
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	u := base.ResolveReference(rel)
	if query != nil {
		u.RawQuery = query.Encode()
	}
	return u.String(), nil
}
