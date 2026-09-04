package query

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type TeacherModel struct {
	ID           int       `gorm:"column:id"`
	Name         string    `gorm:"column:name"`
	Avatar       string    `gorm:"column:avatar"`
	Introduction string    `gorm:"column:introduction"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type TeacherService struct{}

func NewTeacherService() *TeacherService {
	return &TeacherService{}
}

var mockTeachers = []TeacherModel{
	{
		ID:           1,
		Name:         "王老师",
		Avatar:       "https://api.dicebear.com/9.x/avataaars/svg?seed=wang",
		Introduction: "专注于前端工程化与用户体验设计。",
		CreatedAt:    time.Date(2026, time.January, 10, 9, 0, 0, 0, time.Local),
	},
	{
		ID:           2,
		Name:         "李老师",
		Avatar:       "https://api.dicebear.com/9.x/avataaars/svg?seed=li",
		Introduction: "专注于 Go 后端开发与分布式系统。",
		CreatedAt:    time.Date(2026, time.February, 18, 14, 30, 0, 0, time.Local),
	},
}

func (s *TeacherService) Get(_ context.Context, id int) (*TeacherModel, error) {
	for i := range mockTeachers {
		if mockTeachers[i].ID == id {
			teacher := mockTeachers[i]
			return &teacher, nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}
