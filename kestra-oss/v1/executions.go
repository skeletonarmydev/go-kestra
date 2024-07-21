package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ExecutionService service

type ExecutionTaskState struct {
	Current   string    `json:"current,omitempty" structs:"current,omitempty"`
	Duration  string    `json:"duration,omitempty" structs:"duration,omitempty"`
	StartDate time.Time `json:"startDate,omitempty" structs:"startDate,omitempty"`
	EndDate   time.Time `json:"endDate,omitempty" structs:"endDate,omitempty"`
}

type ExecutionTaskRun struct {
	ID          string             `json:"id,omitempty" structs:"id,omitempty"`
	TaskId      string             `json:"taskId,omitempty" structs:"taskId,omitempty"`
	Description string             `json:"description,omitempty" structs:"description,omitempty"`
	State       ExecutionTaskState `json:"state,omitempty" structs:"state,omitempty"`
}

type ExecutionHistory struct {
	State string    `json:"state,omitempty" structs:"state,omitempty"`
	Date  time.Time `json:"date,omitempty" structs:"date,omitempty"`
}

type ExecutionState struct {
	Current   string             `json:"current,omitempty" structs:"current,omitempty"`
	History   []ExecutionHistory `json:"histories,omitempty" structs:"histories,omitempty"`
	Duration  string             `json:"duration,omitempty" structs:"duration,omitempty"`
	StartDate time.Time          `json:"startDate,omitempty" structs:"startDate,omitempty"`
	EndDate   time.Time          `json:"endDate,omitempty" structs:"endDate,omitempty"`
}

// Kestra Execution.
type Execution struct {
	ID           string             `json:"id,omitempty" structs:"id,omitempty"`
	Namespace    string             `json:"namespace,omitempty" structs:"namespace,omitempty"`
	FlowID       string             `json:"flowId,omitempty" structs:"flowId,omitempty"`
	FlowRevision json.Number        `json:"flowRevision,omitempty" structs:"flowRevision,omitempty"`
	State        ExecutionState     `json:"state,omitempty" structs:"state,omitempty"`
	TaskRunList  []ExecutionTaskRun `json:"taskRunList,omitempty" structs:"taskRunList,omitempty"`
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

func (s *ExecutionService) Create(ctx context.Context, namespace string, flowId string, input map[string]string) (*Execution, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/executions/%s/%s", namespace, flowId)

	comb := []string{}

	for key, val := range input {
		comb = append(comb, url.QueryEscape(key)+"="+url.QueryEscape(val))
	}

	body := strings.Join(comb, "&")

	req, err := s.client.NewRequest(ctx, http.MethodPost, apiEndpoint, &body, "multipart/form-data")
	if err != nil {
		return nil, nil, err
	}

	execution := new(Execution)
	resp, err := s.client.Do(req, execution)

	return execution, resp, nil
}
