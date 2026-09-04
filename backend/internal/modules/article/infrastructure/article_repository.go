package infrastructure

import (
	"backend/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

var _ domain.ArticleRepository = (*ArticleRepository)(nil)

func (r *ArticleRepository) Create(ctx context.Context, article *domain.Article) (int, error) {
	po := toArticleModel(article)
	if err := r.db.WithContext(ctx).Create(po).Error; err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *ArticleRepository) Update(ctx context.Context, article *domain.Article) error {
	po := toArticleModel(article)
	result := r.db.WithContext(ctx).
		Model(&ArticleModel{}).
		Where("id = ?", article.ID()).
		Updates(po)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrTitleEmpty
	}
	return nil
}

func (r *ArticleRepository) FindByID(ctx context.Context, id int) (*domain.Article, error) {
	var model ArticleModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.toDomain()
}

func (r *ArticleRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&ArticleModel{}, id).Error
}
