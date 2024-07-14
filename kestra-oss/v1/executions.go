package v1

import (
	"context"
	"fmt"
	"net/http"
)

type ExecutionService service

// Issue represents a Jira issue.
type Execution struct {
	ID           string `json:"id,omitempty" structs:"id,omitempty"`
	Namespace    string `json:"namespace,omitempty" structs:"namespace,omitempty"`
	FlowID       string `json:"flow_id,omitempty" structs:"flow_id,omitempty"`
	FlowRevision string `json:"flow_revision,omitempty" structs:"flow_revision,omitempty"`
}

func (s *ExecutionService) Get(ctx context.Context, executionID string) (*Execution, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/executions/%s", executionID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil, "")
	if err != nil {
		return nil, nil, err
	}

	execution := new(Execution)
	resp, err := s.client.Do(req, execution)

	return execution, resp, nil
}

func (s *ExecutionService) Create(ctx context.Context, namespace string, flowId string) (*Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/executions/%s/%s", namespace, flowId)
	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, nil, "multipart/form-data")
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	return resp, nil
}
