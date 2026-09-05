package api

import (
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"
	"backend/internal/modules/article/domain"
	"backend/internal/shared/api"
	"backend/internal/shared/pagination"
	"context"
	"net/http"

	"gorm.io/gorm"
)

var articleErrorMappings = []api.ErrorMapping{
	{Target: domain.ErrArticleNotFound, Status: http.StatusNotFound},
	{Target: gorm.ErrRecordNotFound, Status: http.StatusNotFound},
}

type ArticleHandler struct {
	cmdSvc   *command.ArticleService
	querySvc *query.ArticleService
}

func NewArticleHandler(cmdSvc *command.ArticleService, querySvc *query.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		cmdSvc:   cmdSvc,
		querySvc: querySvc,
	}
}

func (h *ArticleHandler) Create(ctx context.Context, req *CreateArticleRequest) (*api.Body[CreateArticleResponse], error) {
	createArticleCommand := req.ToCommand()
	createId, err := h.cmdSvc.Create(ctx, createArticleCommand)
	if err != nil {
		return nil, api.MapError("创建文章失败", err, articleErrorMappings...)
	}
	return api.NewBody(CreateArticleResponse{ID: createId}), nil
}

func (h *ArticleHandler) Update(ctx context.Context, req *UpdateArticleRequest) (*struct{}, error) {
	updateArticleCommand := req.ToCommand()
	if err := h.cmdSvc.Update(ctx, updateArticleCommand); err != nil {
		return nil, api.MapError("更新文章失败", err, articleErrorMappings...)
	}
	return nil, nil
}

func (h *ArticleHandler) Delete(ctx context.Context, req *DeleteArticleRequest) (*struct{}, error) {
	if err := h.cmdSvc.Delete(ctx, req.ID); err != nil {
		return nil, api.MapError("删除文章失败", err, articleErrorMappings...)
	}
	return nil, nil
}

func (h *ArticleHandler) List(ctx context.Context, req *ListArticleRequest) (*api.Body[api.PageResult[ArticleResponse]], error) {
	queryParam := query.ListArticlesQuery{
		Title:   req.Title,
		TypeID:  req.TypeID,
		Content: req.Content,
		Page: pagination.New(
			req.Page.Page,
			req.Page.PageSize,
		),
	}
	result, err := h.querySvc.List(ctx, queryParam)
	if err != nil {
		return nil, api.MapError("查询列表失败", err)
	}
	items, err := ToArticleResponseList(result.Items)
	if err != nil {
		return nil, api.MapError("文章数据转换失败", err)
	}
	page := api.Page{Page: result.Params.Page, PageSize: result.Params.PageSize}
	return api.NewBody(api.NewPageResult(items, result.Total, &page)), nil
}

func (h *ArticleHandler) Get(ctx context.Context, req *GetArticleRequest) (*api.Body[ArticleDetailResponse], error) {
	article, err := h.querySvc.Get(ctx, req.ID)
	if err != nil {
		return nil, api.MapError("获取失败", err, articleErrorMappings...)
	}
	response, err := ToArticleDetailResponse(article)
	if err != nil {
		return nil, api.MapError("文章数据转换失败", err)
	}
	return api.NewBody(response), nil
}
