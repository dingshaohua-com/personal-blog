package infrastructure

import (
	"backend/internal/modules/article/domain"
	"context"

	"gorm.io/gorm"
)

type ArticleTypeRepository struct {
	db *gorm.DB
}

func (r *ArticleTypeRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&ArticleTypeModel{}, id).Error
}

func (r *ArticleTypeRepository) Update(ctx context.Context, articleType *domain.ArticleType) error {
	po := toArticleTypeModel(articleType)
	result := r.db.WithContext(ctx).Model(&ArticleTypeModel{}).Where("id=?", articleType.ID()).Updates(po)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func NewArticleTypeRepository(db *gorm.DB) *ArticleTypeRepository {
	return &ArticleTypeRepository{db: db}
}

var _ domain.ArticleTypeRepository = (*ArticleTypeRepository)(nil)

func (r *ArticleTypeRepository) Create(ctx context.Context, articleType *domain.ArticleType) (int, error) {
	po := toArticleTypeModel(articleType)
	if err := r.db.WithContext(ctx).Create(po).Error; err != nil {
		return 0, err
	}
	return po.ID, nil
}

func (r *ArticleTypeRepository) FindByID(ctx context.Context, id int) (*domain.ArticleType, error) {
	model := &ArticleTypeModel{}
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.toDomain(), nil
}
