package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlowService service

// Issue represents a Jira issue.
type Flow struct {
	ID        string      `json:"id,omitempty" structs:"id,omitempty"`
	Namespace string      `json:"namespace,omitempty" structs:"namespace,omitempty"`
	Revision  json.Number `json:"revision,omitempty" structs:"revision,omitempty"`
}

type SearchResult struct {
	Results []Flow `json:"results,omitempty" structs:"results,omitempty"`
	Total   int32  `json:"total,omitempty" structs:"total,omitempty"`
}

func (s *FlowService) GetAll(ctx context.Context, namespace string) (*[]Flow, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/flows/%s", namespace)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil, "")
	if err != nil {
		return nil, nil, err
	}

	flow := new([]Flow)
	resp, err := s.client.Do(req, flow)
	if resp.StatusCode == 404 {
		return nil, resp, nil
	}

	return flow, resp, nil
}

func (s *FlowService) Get(ctx context.Context, namespace string, flowID string) (*Flow, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/flows/%s/%s", namespace, flowID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil, "")
	if err != nil {
		return nil, nil, err
	}

	flow := new(Flow)
	resp, err := s.client.Do(req, flow)
	if resp.StatusCode == 404 {
		return nil, resp, nil
	}

	return flow, resp, nil
}

func (s *FlowService) Search(ctx context.Context, query string) (*SearchResult, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/flows/search?q=%s", query)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil, "")
	if err != nil {
		return nil, nil, err
	}

	searchResult := new(SearchResult)
	resp, err := s.client.Do(req, searchResult)
	if resp.StatusCode == 404 {
		return nil, resp, nil
	}

	return searchResult, resp, nil
}

func (s *FlowService) Create(ctx context.Context, content string) (*Flow, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/flows")
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, &content, "application/json")
	if err != nil {
		return nil, nil, err
	}

	flow := new(Flow)
	resp, err := s.client.Do(req, flow)

	return flow, resp, nil
}
