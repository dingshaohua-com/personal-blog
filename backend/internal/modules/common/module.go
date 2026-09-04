package common

import (
	"backend/internal/modules/common/api"
	"backend/internal/modules/common/application/query"

	"github.com/danielgtaylor/huma/v2"
)

// RegisterModule 组装并注册不属于具体业务领域的通用接口。
func RegisterModule(serverAPI huma.API) {
	teacherService := query.NewTeacherService()
	handler := api.NewCommonHandler(teacherService)
	api.RegisterRoutes(handler, serverAPI)
}
