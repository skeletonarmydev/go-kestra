package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"testing"
)

const (
	testKestraInstanceURL = "http://localhost:8080/"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the Jira client being tested.
	testClient *Client

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a jira.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	// kestra client configured to use test server
	testClient, _ = NewClient(testServer.URL, nil)
}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}

func testRequestParams(t *testing.T, r *http.Request, want map[string]string) {
	params := r.URL.Query()

	if len(params) != len(want) {
		t.Errorf("Request params: %d, want %d", len(params), len(want))
	}

	for key, val := range want {
		if got := params.Get(key); val != got {
			t.Errorf("Request params: %s, want %s", got, val)
		}

	}

}

func TestCheckResponse(t *testing.T) {
	type args struct {
		r *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("CheckResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Client(t *testing.T) {
	type fields struct {
		clientMu  sync.Mutex
		client    *http.Client
		UserAgent string
		common    service
		BaseURL   *url.URL
		Flow      *FlowService
		Execution *ExecutionService
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:  tt.fields.clientMu,
				client:    tt.fields.client,
				UserAgent: tt.fields.UserAgent,
				common:    tt.fields.common,
				BaseURL:   tt.fields.BaseURL,
				Flow:      tt.fields.Flow,
				Execution: tt.fields.Execution,
			}
			if got := c.Client(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Do(t *testing.T) {
	type fields struct {
		clientMu  sync.Mutex
		client    *http.Client
		UserAgent string
		common    service
		BaseURL   *url.URL
		Flow      *FlowService
		Execution *ExecutionService
	}
	type args struct {
		req *http.Request
		v   interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:  tt.fields.clientMu,
				client:    tt.fields.client,
				UserAgent: tt.fields.UserAgent,
				common:    tt.fields.common,
				BaseURL:   tt.fields.BaseURL,
				Flow:      tt.fields.Flow,
				Execution: tt.fields.Execution,
			}
			got, err := c.Do(tt.args.req, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	type fields struct {
		clientMu  sync.Mutex
		client    *http.Client
		UserAgent string
		common    service
		BaseURL   *url.URL
		Flow      *FlowService
		Execution *ExecutionService
	}
	type args struct {
		ctx          context.Context
		method       string
		urlStr       string
		body         *string
		content_type string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				clientMu:  tt.fields.clientMu,
				client:    tt.fields.client,
				UserAgent: tt.fields.UserAgent,
				common:    tt.fields.common,
				BaseURL:   tt.fields.BaseURL,
				Flow:      tt.fields.Flow,
				Execution: tt.fields.Execution,
			}
			got, err := c.NewRequest(tt.args.ctx, tt.args.method, tt.args.urlStr, tt.args.body, tt.args.content_type)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		baseURL    string
		httpClient *http.Client
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.baseURL, tt.args.httpClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_populatePageValues(t *testing.T) {
	type fields struct {
		Response   *http.Response
		StartAt    int
		MaxResults int
		Total      int
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Response:   tt.fields.Response,
				StartAt:    tt.fields.StartAt,
				MaxResults: tt.fields.MaxResults,
				Total:      tt.fields.Total,
			}
			r.populatePageValues(tt.args.v)
		})
	}
}

func Test_newResponse(t *testing.T) {
	type args struct {
		r *http.Response
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newResponse(tt.args.r, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
