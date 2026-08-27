package infrastructure

import (
	"backend/internal/modules/article/domain"
	"time"
)

type ArticleModel struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title"`
	TypeID    int       `gorm:"column:type_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Content   string    `gorm:"column:content"`
}

func (ArticleModel) TableName() string { return "article" }

func (m ArticleModel) toDomain() *domain.Article {
	title, _ := domain.NewArticleTitle(m.Title)
	return domain.RestoreArticle(m.ID, title, m.Content, m.TypeID, m.CreatedAt)
}

func toArticleModel(article *domain.Article) *ArticleModel {
	return &ArticleModel{
		ID:      article.ID(),
		Title:   article.Title().String(),
		Content: article.Content(),
		TypeID:  article.TypeID(),
	}
}
