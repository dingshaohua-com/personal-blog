package infrastructure

import "backend/internal/modules/article/domain"

type ArticleTypeModel struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
	Slug string `gorm:"column:slug"`
}

func (ArticleTypeModel) TableName() string { return "article_type" }

func (m ArticleTypeModel) toDomain() *domain.ArticleType {
	articleTypeName, _ := domain.NewArticleTypeName(m.Name)
	return domain.RestoreArticleType(m.ID, *articleTypeName, m.Slug)
}

func toArticleTypeModel(article *domain.ArticleType) *ArticleTypeModel {
	return &ArticleTypeModel{
		ID:   article.ID(),
		Name: article.Name().String(),
		Slug: article.Slug(),
	}
}
