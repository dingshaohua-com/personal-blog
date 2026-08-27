package query

import (
	"context"

	"gorm.io/gorm"
)

type ArticleTypeService struct {
	db *gorm.DB
}

func NewArticleTypeService(db *gorm.DB) *ArticleTypeService {
	return &ArticleTypeService{
		db: db,
	}
}

func (q *ArticleTypeService) List(ctx context.Context) ([]*ArticleTypeModel, error) {
	var articleTypes []*ArticleTypeModel
	dbQuery := q.db.WithContext(ctx)
	if err := dbQuery.
		Model(&ArticleTypeModel{}).
		Find(&articleTypes).Error; err != nil {
		return nil, err
	}
	return articleTypes, nil
}
