package bootstrap

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2"
)

func TestOpenAPIEndpointMatchesRegisteredSchema(t *testing.T) {
	// No config, database, Redis, or listening socket is needed for registration.
	router := http.NewServeMux()
	api := NewAPI(router, nil)
	registered, err := json.Marshal(api.OpenAPI())
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/openapi.json", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("schema endpoint returned %d", response.Code)
	}
	var want, got any
	if err := json.Unmarshal(registered, &want); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(response.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatal("registered schema differs from HTTP schema")
	}

	validKey := regexp.MustCompile(`^[a-zA-Z0-9.\-_]+$`)
	for name := range api.OpenAPI().Components.Schemas.Map() {
		if !validKey.MatchString(name) {
			t.Errorf("invalid OpenAPI schema name: %q", name)
		}
	}
	validAction := regexp.MustCompile(`^[a-z][a-zA-Z0-9]*$`)
	seen := map[string]bool{}
	for path, item := range api.OpenAPI().Paths {
		for _, op := range []*huma.Operation{item.Get, item.Post, item.Put, item.Patch, item.Delete, item.Head, item.Options, item.Trace} {
			if op == nil {
				continue
			}
			if len(op.Tags) != 1 {
				t.Errorf("%s %s needs exactly one tag", op.Method, path)
				continue
			}
			suffix := "-" + op.Tags[0]
			action := strings.TrimSuffix(op.OperationID, suffix)
			if !strings.HasSuffix(op.OperationID, suffix) || !validAction.MatchString(action) || action == "delete" {
				t.Errorf("invalid operation ID: %s", op.OperationID)
			}
			if seen[op.OperationID] {
				t.Errorf("duplicate operation ID: %s", op.OperationID)
			}
			seen[op.OperationID] = true
		}
	}
	for _, path := range []string{"/article/{id}", "/article-type/{id}", "/feed/{id}"} {
		item := api.OpenAPI().Paths[path]
		for _, op := range []*huma.Operation{item.Put, item.Delete} {
			if op.Method == http.MethodPut || op.Method == http.MethodDelete {
				response := op.Responses["204"]
				if response == nil || len(response.Content) != 0 {
					t.Errorf("%s %s should document an empty 204 response", op.Method, path)
				}
			}
		}
	}
}
