package domain

import (
	"strings"
	"testing"
)

func TestArticleTitleLength(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value string
		valid bool
	}{
		{"news title", "IBM 推出新一代 Nighthawk 量子处理器", true},
		{"unicode boundary", strings.Repeat("文", 200), true},
		{"over limit", strings.Repeat("文", 201), false},
		{"empty", "  ", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewArticleTitle(tc.value)
			if (err == nil) != tc.valid {
				t.Fatalf("NewArticleTitle() error = %v, valid = %v", err, tc.valid)
			}
		})
	}
}
