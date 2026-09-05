package domain

import (
	"errors"
	"time"
)

// 领域错误
var ErrFeedNotFound = errors.New("feed 不存在")

// Feed 领域实体（也是业务的对象）
type Feed struct {
	id        int
	content   FeedContent
	createdAt time.Time
	updatedAt time.Time
}

func (f *Feed) ID() int {
	return f.id
}
func (f *Feed) Content() FeedContent {
	return f.content
}
func (f *Feed) CreatedAt() time.Time {
	return f.createdAt
}
func (f *Feed) UpdatedAt() time.Time {
	return f.updatedAt
}

// NewFeed 领域工厂方法
func NewFeed(content FeedContent) (*Feed, error) {
	feed := &Feed{
		content: content,
	}
	return feed, nil
}

// ChangeContent 领域行为，用于修改文章内容
func (f *Feed) ChangeContent(content string) error {
	feedContent, err := NewFeedContent(content)
	if err != nil {
		return err
	}
	f.content = feedContent
	return nil
}

// RestoreFeed 业务的（领域）对象重建方法
func RestoreFeed(id int, content FeedContent, createdAt time.Time, updatedAt time.Time) *Feed {
	return &Feed{
		id:        id,
		content:   content,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
