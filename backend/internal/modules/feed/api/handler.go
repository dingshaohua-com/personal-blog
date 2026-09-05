package api

import (
	"backend/internal/modules/feed/api/dto"
	"backend/internal/modules/feed/application"
	"backend/internal/modules/feed/domain"
	sharedApi "backend/internal/shared/api"
	"context"
	"net/http"
)

type FeedHandler struct {
	service *application.FeedService
}

var errorMappings = []sharedApi.ErrorMapping{
	{
		Target: domain.ErrFeedNotFound,
		Status: http.StatusNotFound,
	},
}

func NewFeedHandler(service *application.FeedService) *FeedHandler {
	return &FeedHandler{service: service}
}

func (h *FeedHandler) List(ctx context.Context, _ *struct{}) (*sharedApi.Body[[]*dto.FeedDTO], error) {
	result, err := h.service.List(ctx)
	if err != nil {
		return nil, sharedApi.MapError("查询 Feed 列表失败", err, errorMappings...)
	}
	return sharedApi.NewBody(dto.ToFeedDTOList(result)), nil
}

func (h *FeedHandler) Get(ctx context.Context, req *dto.GetFeedRequest) (*sharedApi.Body[*dto.FeedDTO], error) {
	feed, err := h.service.Get(ctx, req.ID)
	if err != nil {
		return nil, sharedApi.MapError("获取失败", err, errorMappings...)
	}
	return sharedApi.NewBody(dto.ToFeedDTO(feed)), nil
}

func (h *FeedHandler) Create(ctx context.Context, req *dto.CreateFeedRequest) (*sharedApi.Body[*dto.FeedDTO], error) {
	feed, err := h.service.Create(ctx, req.Body.Content)
	if err != nil {
		return nil, sharedApi.MapError("插入失败", err, errorMappings...)
	}
	return sharedApi.NewBody(dto.ToFeedDTO(feed)), nil
}

func (h *FeedHandler) Update(ctx context.Context, req *dto.UpdateFeedRequest) (*struct{}, error) {
	err := h.service.Update(ctx, req.ID, req.Body.Content)
	if err != nil {
		return nil, sharedApi.MapError("更新失败", err, errorMappings...)
	}
	return sharedApi.NoContent(), nil
}

func (h *FeedHandler) Delete(ctx context.Context, req *dto.DeleteFeedRequest) (*struct{}, error) {
	err := h.service.Delete(ctx, req.ID)
	if err != nil {
		return nil, sharedApi.MapError("删除失败", err, errorMappings...)
	}
	return sharedApi.NoContent(), nil
}
