package domain

import (
	"backend/internal/shared/validation"
	"errors"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var ErrArticleNotFound = errors.New("文章不存在")

var articleValidator = validator.New()

func validateArticleTypeID(value int) error {
	return validation.Wrap("typeId", "文章类型 ID", articleValidator.Var(value, "gt=0"))
}

type Article struct {
	id        int
	title     ArticleTitle
	typeID    *int
	createdAt time.Time
	content   string
}

func (d *Article) ID() int             { return d.id }
func (d *Article) TypeID() *int        { return d.typeID }
func (d *Article) Title() ArticleTitle { return d.title }
func (d *Article) Content() string     { return d.content }

func NewArticle(title ArticleTitle, typeID *int, content string) (*Article, error) {
	if typeID != nil {
		if err := validateArticleTypeID(*typeID); err != nil {
			return nil, err
		}
	}
	return &Article{
		title:   title,
		typeID:  typeID,
		content: content,
	}, nil
}

func RestoreArticle(
	id int,
	title ArticleTitle,
	content string,
	typeID *int,
	createdAt time.Time,
) *Article {
	return &Article{
		id:        id,
		title:     title,
		content:   content,
		typeID:    typeID,
		createdAt: createdAt,
	}
}

func (d *Article) ChangeTitle(value string) error {
	articleTitle, err := NewArticleTitle(value)
	if err != nil {
		return err
	}
	d.title = articleTitle
	return nil
}

func (d *Article) ChangeContent(value string) {
	d.content = value
}
func (d *Article) ChangeTypeID(value int) error {
	if err := validateArticleTypeID(value); err != nil {
		return err
	}
	d.typeID = &value
	return nil
}

// ArticleTitle 值对象，不变量
type ArticleTitle struct {
	value string
}

const MaxArticleTitleLength = 10

func NewArticleTitle(value string) (ArticleTitle, error) {
	value = strings.TrimSpace(value)
	err := validation.String("title", "文章标题", value).
		Required().
		Max(MaxArticleTitleLength).
		Validate()
	if err != nil {
		return ArticleTitle{}, err
	}
	return ArticleTitle{value: value}, nil
}

func (c ArticleTitle) String() string {
	return c.value
}
