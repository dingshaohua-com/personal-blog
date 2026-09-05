package api

import (
	sharedApi "backend/internal/shared/api"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(handler *FeedHandler, api huma.API) {
	feedGroup := sharedApi.NewGroup(api, "/feed", "feed")
	feedGroup.UseSimpleModifier(func(op *huma.Operation) {
		op.Description = "说说"
	})
	huma.Get(feedGroup, "", handler.List, func(op *huma.Operation) {
		op.OperationID = "list"
		op.Summary = "列表"
	})
	huma.Get(feedGroup, "/{id}", handler.Get, func(op *huma.Operation) {
		op.OperationID = "get"
		op.Summary = "获取单条"
	})
	huma.Post(feedGroup, "", handler.Create, func(op *huma.Operation) {
		op.OperationID = "create"
		op.Summary = "新增"
	})
	huma.Put(feedGroup, "/{id}", handler.Update, func(op *huma.Operation) {
		op.OperationID = "update"
		op.Summary = "更新"
	})
	huma.Delete(feedGroup, "/{id}", handler.Delete, func(op *huma.Operation) {
		op.OperationID = "remove"
		op.Summary = "删除"
	})
}
