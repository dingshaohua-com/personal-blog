package infrastructure

import (
	"backend/internal/modules/feed/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type FeedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *FeedRepository {
	return &FeedRepository{db: db}
}

func (r *FeedRepository) Create(ctx context.Context, feed *domain.Feed) (*domain.Feed, error) {
	model := FeedModel{
		Content:   feed.Content().String(),
		CreatedAt: feed.CreatedAt(),
	}
	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return nil, err
	}
	return model.toDomain()
}

func (r *FeedRepository) Update(ctx context.Context, feed *domain.Feed) error {
	result := r.db.
		WithContext(ctx).
		Model(&FeedModel{}).
		Where("id = ?", feed.ID()).
		Update("content", feed.Content().String())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrFeedNotFound
	}
	return nil
}

func (r *FeedRepository) Delete(ctx context.Context, id int) error {
	result := r.db.
		WithContext(ctx).
		Delete(&FeedModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrFeedNotFound
	}
	return nil
}

func (r *FeedRepository) List(ctx context.Context) ([]*domain.Feed, error) {
	var models []FeedModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	return toDomainList(models)
}

func (r *FeedRepository) FindByID(ctx context.Context, id int) (*domain.Feed, error) {
	var model FeedModel
	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(&model).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrFeedNotFound
	}
	if err != nil {
		return nil, err
	}
	return model.toDomain()
}
