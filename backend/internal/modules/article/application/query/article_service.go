package query

import (
	"backend/internal/shared/pagination"
	"context"

	"gorm.io/gorm"
)

type ArticleService struct {
	db *gorm.DB
}

func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{db: db}
}

func articleFilters(query ListArticlesQuery) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query.Title != "" {
			db = db.Where(
				"article.title ILIKE ?",
				"%"+query.Title+"%",
			)
		}

		if query.TypeID > 0 {
			db = db.Where(
				"article.type_id = ?",
				query.TypeID,
			)
		}

		if query.Content != "" {
			db = db.Where(
				"article.content ILIKE ?",
				"%"+query.Content+"%",
			)
		}

		return db
	}
}

func (s *ArticleService) ListDao(ctx context.Context, queryParam ListArticlesQuery) ([]*ArticleListItemModel, int64, error) {
	var articles []*ArticleListItemModel
	var total int64

	dbQuery := s.db.WithContext(ctx).
		Table("article").
		Joins("LEFT JOIN article_type ON article.type_id = article_type.id").
		Scopes(articleFilters(queryParam))
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := dbQuery.
		Select("article.*, LEFT(SPLIT_PART(REPLACE(article.content, CHR(13), ''), CHR(10) || CHR(10), 1), 160) AS description, article_type.name AS type_name").
		Order("article.created_at DESC, article.id DESC").
		Limit(queryParam.Page.Limit()).
		Offset(queryParam.Page.Offset()).
		Find(&articles).Error; err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

func (s *ArticleService) List(
	ctx context.Context,
	queryParam ListArticlesQuery,
) (pagination.Result[*ArticleListItemModel], error) {
	// 标准分页
	queryParam.Page = queryParam.Page.Normalize()
	articles, total, err := s.ListDao(ctx, queryParam)
	if err != nil {
		return pagination.Result[*ArticleListItemModel]{}, err
	}
	if articles == nil {
		articles = make([]*ArticleListItemModel, 0)
	}
	return pagination.Result[*ArticleListItemModel]{
		Items: articles, Total: total, Params: queryParam.Page,
	}, nil
}

func (s *ArticleService) GetDao(ctx context.Context, id int) (*ArticleModel, error) {
	var article ArticleModel
	if err := s.db.WithContext(ctx).
		Table("article").
		Select("article.*, article_type.name AS type_name").
		Joins("LEFT JOIN article_type ON article.type_id = article_type.id").
		Where("article.id = ?", id).
		Take(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}
func (s *ArticleService) Get(ctx context.Context, id int) (*ArticleModel, error) {
	article, err := s.GetDao(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}
