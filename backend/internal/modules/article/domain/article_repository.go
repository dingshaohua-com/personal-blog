package domain

import "context"

type ArticleRepository interface {
	Create(ctx context.Context, article *Article) (int, error)
	Update(ctx context.Context, article *Article) error
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id int) (*Article, error)
}
