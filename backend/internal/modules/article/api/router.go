package api

import (
	"github.com/danielgtaylor/huma/v2"
)

func operation(id, summary string) func(*huma.Operation) {
	return func(op *huma.Operation) {
		op.OperationID = id
		op.Summary = summary
	}
}

func useSimpleModifierCb(id string) func(*huma.Operation) {
	return func(op *huma.Operation) {
		op.OperationID = op.OperationID + "-" + id
		op.Tags = []string{id}
	}
}

func RegisterRoutes(articleHandler *ArticleHandler, articleTypeHandler *ArticleTypeHandler, api huma.API) {
	// 1. 文章相关的路由组
	articleGroup := huma.NewGroup(api, "/article")
	articleGroup.UseSimpleModifier(useSimpleModifierCb("article"))
	huma.Get(articleGroup, "", articleHandler.List, operation("list", "列表"))
	huma.Get(articleGroup, "/{id}", articleHandler.Get, operation("get", "获取单条"))
	huma.Post(articleGroup, "", articleHandler.Create, operation("create", "新增"))
	huma.Put(articleGroup, "/{id}", articleHandler.Update, operation("update", "更新"))
	huma.Delete(articleGroup, "/{id}", articleHandler.Delete, operation("delete", "删除"))

	// 2. 文章分类 相关的路由组
	articleTypeGroup := huma.NewGroup(api, "/article-type") // 或者 /categories
	articleTypeGroup.UseSimpleModifier(useSimpleModifierCb("article-type"))
	huma.Get(articleTypeGroup, "", articleTypeHandler.List, operation("list", "列表"))
	huma.Post(articleTypeGroup, "", articleTypeHandler.Create, operation("create", "新增"))
	huma.Put(articleTypeGroup, "/{id}", articleTypeHandler.Update, operation("update", "更新"))
	huma.Delete(articleTypeGroup, "/{id}", articleTypeHandler.Delete, operation("delete", "删除"))
}
