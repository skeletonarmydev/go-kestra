package v1

import (
	"context"
	"reflect"
	"testing"
)

func TestLogService_Get(t *testing.T) {
	type args struct {
		ctx         context.Context
		executionID string
	}
	tests := []struct {
		name    string
		s       LogService
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
