package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	ClientVersion = "1.0.0"

	defaultUserAgent = "go-kestra" + "/" + ClientVersion
)

type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify it.
	client   *http.Client // HTTP client used to communicate with the API.

	// User agent used when communicating with the Kestra API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Base URL for API requests.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	Flow      *FlowService
	Execution *ExecutionService
	Log       *LogService
}

// service is the base structure to bundle API services
// under a sub-struct.
type service struct {
	client *Client
}

// Client returns the http.Client used by this Kestra client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

// NewClient returns a new Kestra API client with provided base URL
// If a nil httpClient is provided, a new http.Client will be used.
// To use API methods which require authentication, provide an http.Client that will perform the authentication for you (such as that provided by the golang.org/x/oauth2 library).
// baseURL is the HTTP endpoint of your Kestra instance and should always be specified with a trailing slash.
func NewClient(baseURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	// ensure the baseURL contains a trailing slash so that all paths are preserved in later calls
	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}

	c := &Client{
		client:    httpClient,
		BaseURL:   baseEndpoint,
		UserAgent: defaultUserAgent,
	}
	c.common.client = c

	c.Flow = (*FlowService)(&c.common)
	c.Execution = (*ExecutionService)(&c.common)
	c.Log = (*LogService)(&c.common)

	return c, nil
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body *string, content_type string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Relative URLs should be specified without a preceding slash since BaseURL will have the trailing slash
	rel.Path = strings.TrimLeft(rel.Path, "/")

	u := c.BaseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		buf = strings.NewReader(*body)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if content_type == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", content_type)
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(httpResp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return newResponse(httpResp, nil), err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer httpResp.Body.Close()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}

	resp := newResponse(httpResp, v)
	return resp, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}

// Response represents Kestra API response. It wraps http.Response returned from
// API and provides information about paging.
type Response struct {
	*http.Response

	StartAt    int
	MaxResults int
	Total      int
}

func newResponse(r *http.Response, v interface{}) *Response {
	resp := &Response{Response: r}
	resp.populatePageValues(v)
	return resp
}

func (r *Response) populatePageValues(v interface{}) {

}
