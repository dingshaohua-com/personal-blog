package infrastructure

import (
	"backend/internal/modules/article/domain"
	"time"
)

type ArticleModel struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title"`
	TypeID    *int      `gorm:"column:type_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Content   string    `gorm:"column:content"`
}

func (ArticleModel) TableName() string { return "article" }

func (m ArticleModel) toDomain() (*domain.Article, error) {
	title, err := domain.NewArticleTitle(m.Title)
	if err != nil {
		return nil, err
	}
	return domain.RestoreArticle(m.ID, title, m.Content, m.TypeID, m.CreatedAt), nil
}

func toArticleModel(article *domain.Article) *ArticleModel {
	return &ArticleModel{
		ID:      article.ID(),
		Title:   article.Title().String(),
		Content: article.Content(),
		TypeID:  article.TypeID(),
	}
}
