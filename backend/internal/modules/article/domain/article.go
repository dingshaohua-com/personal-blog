package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	ErrTitleNotFound = errors.New("文章不存在")
	ErrTitleEmpty    = errors.New("文章标题不能为空")
	ErrTitleTooLong  = errors.New("文章标题太长")
)

type Article struct {
	id        int
	title     ArticleTitle
	typeID    int
	createdAt time.Time
	content   string
}

func (d *Article) ID() int             { return d.id }
func (d *Article) TypeID() int         { return d.typeID }
func (d *Article) Title() ArticleTitle { return d.title }
func (d *Article) Content() string     { return d.content }

func NewArticle(title ArticleTitle, typeID int, content string) *Article {
	return &Article{
		title:   title,
		typeID:  typeID,
		content: content,
	}
}

func RestoreArticle(
	id int,
	title ArticleTitle,
	content string,
	typeID int,
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
func (d *Article) ChangeTypeID(value int) {
	d.typeID = value
}

// ArticleTitle 值对象，不变量
type ArticleTitle struct {
	value string
}

const MaxArticleTitleLength = 10

func NewArticleTitle(value string) (ArticleTitle, error) {
	value = strings.TrimSpace(value)
	switch {
	case value == "":
		return ArticleTitle{}, ErrTitleEmpty
	case utf8.RuneCountInString(value) > MaxArticleTitleLength:
		return ArticleTitle{}, ErrTitleTooLong
	}
	return ArticleTitle{value: value}, nil
}

func (c ArticleTitle) String() string {
	return c.value
}
