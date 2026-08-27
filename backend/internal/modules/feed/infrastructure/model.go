package infrastructure

import (
	"backend/internal/modules/feed/domain"
	"time"
)

type FeedModel struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (FeedModel) TableName() string { return "feed" }

func (m FeedModel) toDomain() (*domain.Feed, error) {
	content, err := domain.NewFeedContent(m.Content)
	if err != nil {
		return nil, err
	}
	return domain.RestoreFeed(m.ID, content, m.CreatedAt, m.UpdatedAt), nil
}

func toDomainList(models []FeedModel) ([]*domain.Feed, error) {
	res := make([]*domain.Feed, 0, len(models))
	for _, m := range models {
		feed, err := m.toDomain()
		if err != nil {
			return nil, err
		}
		res = append(res, feed) // 调用单个转换
	}
	return res, nil
}
