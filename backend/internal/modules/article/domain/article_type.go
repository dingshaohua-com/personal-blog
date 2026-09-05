package domain

import (
	"backend/internal/shared/validation"
	"errors"
	"strings"
)

var ArticleTypeErrNotFound = errors.New("文章类型不存在")

type ArticleType struct {
	id   int
	name ArticleTypeName
	slug string
}

func (d *ArticleType) ID() int               { return d.id }
func (d *ArticleType) Name() ArticleTypeName { return d.name }
func (d *ArticleType) Slug() string          { return d.slug }

func (*ArticleType) TableName() string { return "article_type" }

func RestoreArticleType(id int, name ArticleTypeName, slug string) *ArticleType {
	return &ArticleType{
		id:   id,
		name: name,
		slug: slug,
	}
}

func NewArticleType(name ArticleTypeName, slug string) *ArticleType {
	return &ArticleType{
		name: name,
		slug: slug,
	}
}

func (d *ArticleType) ChangeName(name string) error {
	articleTypeName, err := NewArticleTypeName(name)
	if err != nil {
		return err
	}
	d.name = *articleTypeName
	return nil
}

func (d *ArticleType) ChangeSlug(slug string) {
	d.slug = slug
}

// ArticleTypeName 值对象，不变量
type ArticleTypeName struct {
	value string
}

const MaxArticleTypeNameLength = 10

func NewArticleTypeName(value string) (*ArticleTypeName, error) {
	value = strings.TrimSpace(value)
	if err := validation.String("name", "文章类型名称", value).
		Required().Max(MaxArticleTypeNameLength).Validate(); err != nil {
		return nil, err
	}
	return &ArticleTypeName{
		value: value,
	}, nil
}

func (c ArticleTypeName) String() string {
	return c.value
}
