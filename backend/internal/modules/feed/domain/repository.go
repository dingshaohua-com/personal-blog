package domain

import "context"

type FeedRepository interface {
	Create(ctx context.Context, feed *Feed) (*Feed, error)
	FindByID(ctx context.Context, id int) (*Feed, error)
	List(ctx context.Context) ([]*Feed, error)
	Update(ctx context.Context, feed *Feed) error
	Delete(ctx context.Context, id int) error
}
