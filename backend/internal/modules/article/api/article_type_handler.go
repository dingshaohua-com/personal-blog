package api

import (
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"
	"backend/internal/modules/article/domain"
	"backend/internal/shared/api"
	"context"
	"net/http"

	"gorm.io/gorm"
)

var articleTypeErrorMappings = []api.ErrorMapping{
	{Target: domain.ArticleTypeErrNotFound, Status: http.StatusNotFound},
	{Target: gorm.ErrRecordNotFound, Status: http.StatusNotFound},
}

type ArticleTypeHandler struct {
	service *command.ArticleTypeService
	query   *query.ArticleTypeService
}

func NewArticleTypeHandler(service *command.ArticleTypeService, query *query.ArticleTypeService) *ArticleTypeHandler {
	return &ArticleTypeHandler{
		service: service,
		query:   query,
	}
}

func (h *ArticleTypeHandler) Create(ctx context.Context, req *CreateArticleTypeRequest) (*api.Body[CreateArticleTypeResponse], error) {
	id, err := h.service.Create(ctx, req.ToCommand())
	if err != nil {
		return nil, api.MapError("创建文章类型失败", err, articleTypeErrorMappings...)
	}
	return api.NewBody(CreateArticleTypeResponse{ID: id}), nil
}

func (h *ArticleTypeHandler) Update(ctx context.Context, req *UpdateArticleTypeRequest) (*struct{}, error) {
	articleType := req.ToCommand()
	err := h.service.Update(ctx, articleType)
	if err != nil {
		return nil, api.MapError("更新文章类型失败", err, articleTypeErrorMappings...)
	}
	return nil, nil
}

func (h *ArticleTypeHandler) List(ctx context.Context, _ *struct{}) (*api.Body[[]ArticleTypeResponse], error) {
	articleTypes, err := h.query.List(ctx)
	if err != nil {
		return nil, api.MapError("查询文章类型失败", err, articleTypeErrorMappings...)
	}
	articleTypesResp, err := ToArticleTypeResponseList(articleTypes)
	if err != nil {
		return nil, api.MapError("文章类型数据转换失败", err)
	}
	return api.NewBody(articleTypesResp), nil
}

func (h *ArticleTypeHandler) Delete(ctx context.Context, req *DeleteArticleTypeRequest) (*struct{}, error) {
	err := h.service.Delete(ctx, req.ID)
	if err != nil {
		return nil, api.MapError("删除文章类型失败", err, articleTypeErrorMappings...)
	}
	return nil, nil
}
