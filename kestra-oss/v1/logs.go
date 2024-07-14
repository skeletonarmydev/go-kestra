package v1

import (
	"context"
	"fmt"
	"net/http"
)

type LogService service

// Issue represents a Ketra Log.
type Log struct {
	TaskId      string `json:"taskId,omitempty" structs:"taskId,omitempty"`
	Namespace   string `json:"namespace,omitempty" structs:"namespace,omitempty"`
	FlowID      string `json:"flowId,omitempty" structs:"flowId,omitempty"`
	ExecutionId string `json:"executionId,omitempty" structs:"executionId,omitempty"`
	TaskRunId   string `json:"taskRunId,omitempty" structs:"taskRunId,omitempty"`
	Timestamp   string `json:"timestamp,omitempty" structs:"timestamp,omitempty"`
	Level       string `json:"level,omitempty" structs:"level,omitempty"`
	Message     string `json:"message,omitempty" structs:"message,omitempty"`
}

func (s *LogService) Get(ctx context.Context, executionID string) (*Log, *Response, error) {
	apiEndpoint := fmt.Sprintf("/api/v1/logs/%s", executionID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, apiEndpoint, nil, "")
	if err != nil {
		return nil, nil, err
	}

	log := new(Log)
	resp, err := s.client.Do(req, log)

	return log, resp, nil
}
