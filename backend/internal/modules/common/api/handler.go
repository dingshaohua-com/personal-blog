package api

import (
	"backend/internal/modules/common/application/query"
	sharedAPI "backend/internal/shared/api"
	"context"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type CommonHandler struct {
	teacherQuerySvc *query.TeacherService
}

type HealthResponse struct {
	Status string `json:"status" doc:"服务状态" example:"ok"`
}

func NewCommonHandler(teacherQuerySvc *query.TeacherService) *CommonHandler {
	return &CommonHandler{teacherQuerySvc: teacherQuerySvc}
}

// Health 返回服务进程的存活状态。
func (h *CommonHandler) Health(_ context.Context, _ *struct{}) (*sharedAPI.Body[*HealthResponse], error) {
	return sharedAPI.NewBody(&HealthResponse{Status: "ok"}), nil
}

func (h *CommonHandler) GetTeacher(ctx context.Context, req *GetTeacherRequest) (*sharedAPI.Body[*TeacherResponse], error) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
	}

	teacher, err := h.teacherQuerySvc.Get(ctx, req.ID)
	if err != nil {
		return nil, sharedAPI.MapError(
			"获取老师信息失败：",
			err,
			sharedAPI.ErrorMapping{Target: gorm.ErrRecordNotFound, Status: http.StatusNotFound},
		)
	}
	return sharedAPI.NewBody(ToTeacherResponse(teacher)), nil
}
