package api

import (
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"
	"backend/internal/shared/api"
	"context"
)

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
		return nil, err
	}
	return api.NewBody(CreateArticleTypeResponse{ID: id}), nil
}

func (h *ArticleTypeHandler) Update(ctx context.Context, req *UpdateArticleTypeRequest) (*api.Body[*struct{}], error) {
	articleType := req.ToCommand()
	err := h.service.Update(ctx, articleType)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *ArticleTypeHandler) List(ctx context.Context, _ *struct{}) (*api.Body[[]ArticleTypeResponse], error) {
	articleTypes, err := h.query.List(ctx)
	if err != nil {
		return nil, err
	}
	articleTypesResp := ToArticleTypeResponseList(articleTypes)
	return api.NewBody(articleTypesResp), nil
}

func (h *ArticleTypeHandler) Delete(ctx context.Context, req *DeleteArticleTypeRequest) (*api.Body[*struct{}], error) {
	err := h.service.Delete(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
