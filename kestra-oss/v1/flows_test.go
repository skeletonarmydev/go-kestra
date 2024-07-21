package v1

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFlowService_Get(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/flows/tutorial/hello_world", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/api/v1/flows/tutorial/hello_world")

		fmt.Fprint(w, `{"id":"hello_world","namespace":"tutorial","revision":21,"disabled":false,"deleted":false,"description":"Hello World","tasks":[{"id":"log","type":"io.kestra.plugin.core.log.Log","message":"Hello World"}]}`)
	})

	type args struct {
		ctx       context.Context
		flowID    string
		namespace string
	}
	tests := []struct {
		name    string
		s       FlowService
		args    args
		want    *Flow
		code    int
		wantErr bool
	}{
		{"should find Hello World", *testClient.Flow,
			args{context.Background(), "hello_world", "tutorial"},
			&Flow{"hello_world", "tutorial", "21", "Hello World", []FlowTask{{"log", "io.kestra.plugin.core.log.Log", "", nil}}, ""},
			200,
			false,
		},
		{"should not find anything", *testClient.Flow,
			args{context.Background(), "hello_world", "badnamespace"},
			nil,
			404,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, resp, err := tt.s.Get(tt.args.ctx, tt.args.namespace, tt.args.flowID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(resp.StatusCode, tt.code) {
				t.Errorf("StatusCode got = %v, want %v", resp.StatusCode, tt.code)
			}
		})
	}

}

func TestFlowService_Search(t *testing.T) {

	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/flows/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		params := r.URL.Query()
		if params.Get("q") == "hello" {

			m := make(map[string]string)
			m["q"] = "hello"
			testRequestParams(t, r, m)

			fmt.Fprint(w, `{"results":[{"id":"hello_world","namespace":"tutorial","revision":21,"disabled":false,"deleted":false,"description":"Hello World","tasks":[{"id":"log","type":"io.kestra.plugin.core.log.Log","message":"Hello World"}]}],"total":1}`)
		} else {
			w.WriteHeader(404)
		}
	})

	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		s       FlowService
		args    args
		want    *SearchResult
		code    int
		wantErr bool
	}{
		{"should find Hello World", *testClient.Flow,
			args{context.Background(), "hello"},
			&SearchResult{[]Flow{{"hello_world", "tutorial", "21", "Hello World", []FlowTask{{"log", "io.kestra.plugin.core.log.Log", "", nil}}, ""}}, 1},
			200,
			false,
		},
		{"should not find Hello World", *testClient.Flow,
			args{context.Background(), "another"},
			nil,
			404,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, resp, err := tt.s.Search(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(resp.StatusCode, tt.code) {
				t.Errorf("StatusCode got = %v, want %v", resp.StatusCode, tt.code)
			}
		})
	}
}

func TestFlowService_Create(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/flows", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testRequestURL(t, r, "/api/v1/flows")

		fmt.Fprint(w, `{"id":"hello_world","namespace":"tutorial","revision":21,"disabled":false,"deleted":false,"description":"Hello World","tasks":[{"id":"log","type":"io.kestra.plugin.core.log.Log","message":"Hello World"}]}`)
	})

	type args struct {
		ctx     context.Context
		content string
	}
	tests := []struct {
		name    string
		s       FlowService
		args    args
		want    *Flow
		code    int
		wantErr bool
	}{
		{"should create Hello World", *testClient.Flow,
			args{context.Background(), "hello_world"},
			&Flow{"hello_world", "tutorial", "21", "Hello World", []FlowTask{{"log", "io.kestra.plugin.core.log.Log", "", nil}}, ""},
			200,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, resp, err := tt.s.Create(tt.args.ctx, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(resp.StatusCode, tt.code) {
				t.Errorf("StatusCode got = %v, want %v", resp.StatusCode, tt.code)
			}
		})
	}
}

func TestFlowService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/api/v1/flows/tutorial", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testRequestURL(t, r, "/api/v1/flows/tutorial")

		fmt.Fprint(w, `[{"id":"hello_world","namespace":"tutorial","revision":21,"disabled":false,"deleted":false,"description":"Hello World","tasks":[{"id":"log","type":"io.kestra.plugin.core.log.Log","message":"Hello World"}]}]`)
	})

	type args struct {
		ctx       context.Context
		namespace string
	}
	tests := []struct {
		name    string
		s       FlowService
		args    args
		want    *[]Flow
		code    int
		wantErr bool
	}{
		{"should get Hello World", *testClient.Flow,
			args{context.Background(), "tutorial"},
			&[]Flow{{"hello_world", "tutorial", "21", "Hello World", []FlowTask{{"log", "io.kestra.plugin.core.log.Log", "", nil}}, ""}},
			200,
			false,
		},
		{"should not get Hello World", *testClient.Flow,
			args{context.Background(), "another"},
			nil,
			404,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, resp, err := tt.s.GetAll(tt.args.ctx, tt.args.namespace)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(resp.StatusCode, tt.code) {
				t.Errorf("StatusCode got1 = %v, want %v", resp.StatusCode, tt.code)
			}
		})
	}
}
