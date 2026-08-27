package command

import (
	"backend/internal/modules/article/domain"
	"context"
)

type ArticleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(repo domain.ArticleRepository) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) Create(ctx context.Context, cmd CreateArticleCommand) (int, error) {
	articleTitle, err := domain.NewArticleTitle(cmd.Title)
	if err != nil {
		return 0, err
	}
	articleDomain := domain.NewArticle(articleTitle, cmd.TypeID, cmd.Content)
	return s.repo.Create(ctx, articleDomain)
}

func (s *ArticleService) Update(ctx context.Context, cmd UpdateArticleCommand) error {
	articleDomain, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if cmd.Title != nil {
		if err := articleDomain.ChangeTitle(*cmd.Title); err != nil {
			return err
		}
	}
	if cmd.Content != nil {
		articleDomain.ChangeContent(*cmd.Content)
	}
	if cmd.TypeID != nil {
		articleDomain.ChangeTypeID(*cmd.TypeID)
	}
	return s.repo.Update(ctx, articleDomain)
}

func (s *ArticleService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *ArticleService) FindByID(ctx context.Context, id int) (*domain.Article, error) {
	return s.repo.FindByID(ctx, id)
}
