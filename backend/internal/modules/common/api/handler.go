package api

import (
	"backend/internal/modules/common/application/query"
	sharedAPI "backend/internal/shared/api"
	"context"
	"net/http"

	"gorm.io/gorm"
)

type teacherReader interface {
	Get(ctx context.Context, id int) (*query.TeacherModel, error)
}

type CommonHandler struct {
	teacherReader teacherReader
}

type HealthResponse struct {
	Status string `json:"status" doc:"服务状态" example:"ok"`
}

func NewCommonHandler(teacherReader teacherReader) *CommonHandler {
	return &CommonHandler{teacherReader: teacherReader}
}

// Health 返回服务进程的存活状态。
func (h *CommonHandler) Health(_ context.Context, _ *struct{}) (*sharedAPI.Body[*HealthResponse], error) {
	return sharedAPI.NewBody(&HealthResponse{Status: "ok"}), nil
}

func (h *CommonHandler) GetTeacher(ctx context.Context, req *GetTeacherRequest) (*sharedAPI.Body[*TeacherResponse], error) {
	teacher, err := h.teacherReader.Get(ctx, req.ID)
	if err != nil {
		return nil, sharedAPI.MapError(
			"获取老师信息失败：",
			err,
			sharedAPI.ErrorMapping{Target: gorm.ErrRecordNotFound, Status: http.StatusNotFound},
		)
	}
	return sharedAPI.NewBody(ToTeacherResponse(teacher)), nil
}
