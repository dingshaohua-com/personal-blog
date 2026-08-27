package query

import (
	"backend/internal/shared/pagination"
	"time"
)

// ArticleListItemModel 读模型 / 视图模型（Output Read Model）
type ArticleListItemModel struct {
	ID          int
	Title       string
	Description string
	TypeID      int
	TypeName    string
	CreatedAt   time.Time
}

type ArticleModel struct {
	ArticleListItemModel
	Content string
}

// ListArticlesQuery 查询条件模型（Input Query）
type ListArticlesQuery struct {
	Title   string
	TypeID  int
	Content string
	Page    pagination.Params
}
