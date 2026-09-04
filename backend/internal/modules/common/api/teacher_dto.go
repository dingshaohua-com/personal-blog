package api

import (
	"backend/internal/modules/common/application/query"
	"time"
)

type GetTeacherRequest struct {
	ID int `path:"id" minimum:"1" doc:"老师 ID"`
}

type TeacherResponse struct {
	ID           int       `json:"id" doc:"老师 ID"`
	Name         string    `json:"name" doc:"老师姓名"`
	Avatar       string    `json:"avatar" doc:"头像地址"`
	Introduction string    `json:"introduction" doc:"个人简介"`
	CreatedAt    time.Time `json:"createdAt" doc:"创建时间"`
}

func ToTeacherResponse(teacher *query.TeacherModel) *TeacherResponse {
	return &TeacherResponse{
		ID:           teacher.ID,
		Name:         teacher.Name,
		Avatar:       teacher.Avatar,
		Introduction: teacher.Introduction,
		CreatedAt:    teacher.CreatedAt,
	}
}
