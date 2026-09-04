package api

import "github.com/danielgtaylor/huma/v2"

func RegisterRoutes(handler *CommonHandler, serverAPI huma.API) {
	commonGroup := huma.NewGroup(serverAPI, "")
	commonGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Tags = []string{"common"}
	})

	huma.Get(commonGroup, "/health", handler.Health, func(op *huma.Operation) {
		op.OperationID = "health"
		op.Summary = "健康检查"
		op.Description = "检查 HTTP 服务是否正常运行"
	})

	huma.Get(commonGroup, "/teacher/{id}", handler.GetTeacher, func(op *huma.Operation) {
		op.OperationID = "get-teacher"
		op.Summary = "获取老师信息"
		op.Description = "根据 ID 获取单个老师的资料"
	})
}
