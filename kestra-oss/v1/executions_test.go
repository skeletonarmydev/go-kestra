package v1

import (
	"context"
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
	type args struct {
		ctx         context.Context
		executionID string
	}
	tests := []struct {
		name    string
		s       ExecutionService
		args    args
		want    *Execution
		want1   *Response
		wantErr bool
	}{

		// TODO: Add test cases.
	}
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
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
