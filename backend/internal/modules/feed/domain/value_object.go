package domain

import (
	"backend/internal/shared/validation"
	"strings"
)

// FeedContent 值对象
type FeedContent struct {
	value string
}

const MaxFeedContentLength = 100

func NewFeedContent(value string) (FeedContent, error) {
	value = strings.TrimSpace(value)
	if err := validation.String("content", "Feed 内容", value).
		Required().Max(MaxFeedContentLength).Validate(); err != nil {
		return FeedContent{}, err
	}
	return FeedContent{value: value}, nil
}

func (c FeedContent) String() string {
	return c.value
}
