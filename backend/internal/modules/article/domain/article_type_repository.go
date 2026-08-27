package domain

import "context"

type ArticleTypeRepository interface {
	FindByID(ctx context.Context, id int) (*ArticleType, error)
	Create(ctx context.Context, articleType *ArticleType) (int, error)
	Update(ctx context.Context, articleType *ArticleType) error
	Delete(ctx context.Context, id int) error
}
