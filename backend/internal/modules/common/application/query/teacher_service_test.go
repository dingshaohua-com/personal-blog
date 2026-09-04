package query

import (
	"context"
	"errors"
	"testing"

	"gorm.io/gorm"
)

func TestTeacherServiceGet(t *testing.T) {
	service := NewTeacherService()

	for _, id := range []int{1, 2} {
		teacher, err := service.Get(context.Background(), id)
		if err != nil {
			t.Fatalf("get teacher %d: %v", id, err)
		}
		if teacher.ID != id || teacher.Name == "" {
			t.Fatalf("unexpected teacher: %+v", teacher)
		}
	}

	if _, err := service.Get(context.Background(), 3); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("expected record not found, got %v", err)
	}
}
