package api

import (
	"backend/internal/modules/common/application/query"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
	"gorm.io/gorm"
)

type teacherReaderStub struct {
	teacher *query.TeacherModel
	err     error
}

func (s teacherReaderStub) Get(_ context.Context, _ int) (*query.TeacherModel, error) {
	return s.teacher, s.err
}

func TestHealth(t *testing.T) {
	_, testAPI := humatest.New(t, huma.DefaultConfig("Test API", "1.0.0"))
	RegisterRoutes(NewCommonHandler(nil), testAPI)

	response := testAPI.Get("/health")
	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body HealthResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Status != "ok" {
		t.Fatalf("expected status ok, got %q", body.Status)
	}
}

func TestGetTeacher(t *testing.T) {
	_, testAPI := humatest.New(t, huma.DefaultConfig("Test API", "1.0.0"))
	RegisterRoutes(NewCommonHandler(teacherReaderStub{
		teacher: &query.TeacherModel{ID: 1, Name: "张老师", Avatar: "/avatar.png", Introduction: "个人简介"},
	}), testAPI)

	response := testAPI.Get("/teacher/1")
	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body TeacherResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.ID != 1 || body.Name != "张老师" {
		t.Fatalf("unexpected teacher: %+v", body)
	}
}

func TestGetTeacherNotFound(t *testing.T) {
	_, testAPI := humatest.New(t, huma.DefaultConfig("Test API", "1.0.0"))
	RegisterRoutes(NewCommonHandler(teacherReaderStub{err: gorm.ErrRecordNotFound}), testAPI)

	response := testAPI.Get("/teacher/1")
	if response.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", response.Code)
	}

	var body huma.ErrorModel
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Status != http.StatusNotFound {
		t.Fatalf("unexpected error response: %+v", body)
	}
}
