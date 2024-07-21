package v1

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestExecutionService_Create(t *testing.T) {
	type args struct {
		ctx       context.Context
		namespace string
		flowId    string
		body      map[string]string
	}
	tests := []struct {
		name    string
		s       ExecutionService
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, err := tt.s.Create(tt.args.ctx, tt.args.namespace, tt.args.flowId, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutionService_Get(t *testing.T) {

	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/executions/1CcnlV1DwvXXZauauyirIO", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/api/v1/executions/1CcnlV1DwvXXZauauyirIO")

		_, err := fmt.Fprint(w, `{"id":"1CcnlV1DwvXXZauauyirIO","namespace":"tutorial",
			"flowId":"hello_world",
			"flowRevision":26,
			"taskRunList":[{
				"id":"321HwiEUBACDQzkJcP8J4r",
				"executionId":"1CcnlV1DwvXXZauauyirIO",
				"namespace":"tutorial",
				"flowId":"hello_world",
				"taskId":"log",
				"attempts":[{"state":{"current":"SUCCESS",
					"histories":[{"state":"CREATED","date":"2024-07-15T09:27:24.172Z"},
						{"state":"RUNNING","date":"2024-07-15T09:27:24.172Z"},
						{"state":"SUCCESS","date":"2024-07-15T09:27:24.176Z"}],
					"duration":"PT0.004S","endDate":"2024-07-15T09:27:24.176Z",
					"startDate":"2024-07-15T09:27:24.172Z"}}],
			"outputs":{},
			"state":{"current":"SUCCESS",
				"histories":[{"state":"CREATED","date":"2024-07-15T09:27:24.048Z"},
					{"state":"RUNNING","date":"2024-07-15T09:27:24.170Z"},
					{"state":"SUCCESS","date":"2024-07-15T09:27:24.177Z"}],
				"duration":"PT0.129S"
				}}],
			"inputs":{"name":"go.app"},
			"state":{
				"current":"SUCCESS",
				"histories":[
					{"state":"CREATED"},
					{"state":"RUNNING"},
					{"state":"SUCCESS"}],
				"duration":"PT1.203S"},
			"originalId":"1CcnlV1DwvXXZauauyirIO",
			"deleted":false,
			"metadata":{"attemptNumber":1,"originalCreatedDate":"2024-07-15T09:27:23.801Z"}}`)
		if err != nil {
			return
		}
	})

	type args struct {
		ctx         context.Context
		executionID string
	}
	tests := []struct {
		name    string
		s       ExecutionService
		args    args
		want    *Execution
		code    int
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.s.Get(tt.args.ctx, tt.args.executionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.code) {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.code)
			}
		})
	}
}
