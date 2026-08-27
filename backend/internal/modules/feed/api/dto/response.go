package dto

import (
	"backend/internal/modules/feed/domain"
	"time"
)

// FeedDTO 1. 封装结构体
type FeedDTO struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToFeedDTO 2. 封装单个对象的转换：PO -> DTO
func ToFeedDTO(feed *domain.Feed) *FeedDTO {
	if feed == nil {
		return nil
	}
	return &FeedDTO{
		ID:        feed.ID(),
		Content:   feed.Content().String(),
		CreatedAt: feed.CreatedAt(),
		UpdatedAt: feed.UpdatedAt(),
	}
}

// ToFeedDTOList 3. 封装切片/列表的批量转换：[]PO -> []DTO
func ToFeedDTOList(feeds []*domain.Feed) []*FeedDTO {
	list := make([]*FeedDTO, 0, len(feeds))
	for _, feed := range feeds {
		list = append(list, ToFeedDTO(feed))
	}
	return list
}
