package api

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestUnifiedResponses(t *testing.T) {
	_, testAPI := humatest.New(t, huma.DefaultConfig("Test API", "1.0.0"))

	type input struct {
		Value int `query:"value" required:"true"`
	}

	huma.Get(testAPI, "/test", func(_ context.Context, in *input) (*Body[int], error) {
		return NewBody(in.Value), nil
	})

	t.Run("success", func(t *testing.T) {
		response := testAPI.Get("/test?value=42")
		if response.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", response.Code)
		}

		var body int
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if body != 42 {
			t.Fatalf("unexpected success response: %+v", body)
		}
	})

	t.Run("validation error", func(t *testing.T) {
		response := testAPI.Get("/test?value=invalid")
		if response.Code != http.StatusUnprocessableEntity {
			t.Fatalf("expected status 422, got %d", response.Code)
		}

		var body huma.ErrorModel
		if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
			t.Fatalf("decode response: %v", err)
		}
		if body.Status != http.StatusUnprocessableEntity || body.Detail == "" {
			t.Fatalf("unexpected error response: %+v", body)
		}
	})
}

func TestNoContent(t *testing.T) {
	if response := NoContent(); response == nil {
		t.Fatal("expected non-nil empty response")
	}
}

func TestInternalErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  []string
		expected string
	}{
		{name: "default", expected: "服务器内部错误"},
		{name: "with context", message: []string{"文章分类加载失败"}, expected: "服务器内部错误：文章分类加载失败"},
		{name: "empty context", message: []string{""}, expected: "服务器内部错误"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := InternalError(test.message...)
			if err.GetStatus() != http.StatusInternalServerError || err.Error() != test.expected {
				t.Fatalf("unexpected internal error: %+v", err)
			}
		})
	}
}
