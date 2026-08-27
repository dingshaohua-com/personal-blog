package command

import (
	"backend/internal/modules/article/domain"
	"context"
)

type ArticleTypeService struct {
	repo domain.ArticleTypeRepository
}

func NewArticleTypeService(repo domain.ArticleTypeRepository) *ArticleTypeService {
	return &ArticleTypeService{
		repo: repo,
	}
}

func (s *ArticleTypeService) Create(ctx context.Context, cmd CreateArticleTypeCommand) (int, error) {
	articleTypeName, err := domain.NewArticleTypeName(cmd.Name)
	if err != nil {
		return 0, err
	}
	articleType := domain.NewArticleType(*articleTypeName, cmd.Slug)
	return s.repo.Create(ctx, articleType)
}

func (s *ArticleTypeService) Update(ctx context.Context, cmd UpdateArticleTypeCommand) error {
	articleTypeDomain, err := s.repo.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	if cmd.Name != nil {
		err := articleTypeDomain.ChangeName(*cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.Slug != nil {
		articleTypeDomain.ChangeSlug(*cmd.Slug)
	}
	updateErr := s.repo.Update(ctx, articleTypeDomain)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (s *ArticleTypeService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
