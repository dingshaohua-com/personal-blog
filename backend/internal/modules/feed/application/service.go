package application

import (
	"backend/internal/modules/feed/domain"
	"context"
)

type FeedService struct {
	repo domain.FeedRepository
}

func NewFeedService(repo domain.FeedRepository) *FeedService {
	return &FeedService{
		repo: repo,
	}
}

func (s *FeedService) Create(ctx context.Context, rawContent string) (*domain.Feed, error) {
	content, err := domain.NewFeedContent(rawContent)
	if err != nil {
		return nil, err
	}
	feed, err := domain.NewFeed(content)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(ctx, feed)
}

func (s *FeedService) Update(ctx context.Context, id int, content string) error {
	feed, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if err := feed.ChangeContent(content); err != nil {
		return err
	}
	return s.repo.Update(ctx, feed)
}

func (s *FeedService) List(c context.Context) ([]*domain.Feed, error) {
	return s.repo.List(c)
}

func (s *FeedService) Get(ctx context.Context, id int) (*domain.Feed, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *FeedService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
