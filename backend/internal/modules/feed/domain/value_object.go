package domain

import (
	"strings"
	"unicode/utf8"
)

// FeedContent 值对象
type FeedContent struct {
	value string
}

const MaxFeedContentLength = 100

func NewFeedContent(value string) (FeedContent, error) {
	value = strings.TrimSpace(value)
	switch {
	case value == "":
		return FeedContent{}, ErrFeedContentEmpty
	case utf8.RuneCountInString(value) > MaxFeedContentLength:
		return FeedContent{}, ErrFeedContentTooLong
	}
	return FeedContent{value: value}, nil
}

func (c FeedContent) String() string {
	return c.value
}
