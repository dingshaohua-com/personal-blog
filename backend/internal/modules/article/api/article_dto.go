package api

import (
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"
	"backend/internal/shared/api"
	"time"

	"github.com/jinzhu/copier"
)

type GetArticleRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`
}

type DeleteArticleRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`
}

type ListArticleRequest struct {
	api.Page
	Title   string `query:"title" minLength:"1" doc:"标题"`
	TypeID  int    `query:"typeId" minimum:"1" doc:"类型ID"`
	Content string `query:"content" minLength:"1" doc:"内容"`
}

type CreateArticleResponse struct {
	ID int `json:"id" doc:"新创建的文章 ID"`
}

func (r *CreateArticleRequest) ToCommand() command.CreateArticleCommand {
	content := ""
	if r.Body.Content != nil {
		content = *r.Body.Content
	}
	return command.CreateArticleCommand{
		Title:   r.Body.Title,
		TypeID:  r.Body.TypeID,
		Content: content,
	}
}

type ArticleResponse struct {
	ID          int       `json:"id" doc:"ID"`
	Title       string    `json:"title" doc:"标题"`
	Description string    `json:"description" doc:"简介描述"`
	TypeID      int       `json:"typeId" doc:"文章类型ID"`
	TypeName    string    `json:"typeName" doc:"文章类型名称"` // 新增
	CreatedAt   time.Time `json:"createdAt" doc:"创建时间"`
}

type ArticleDetailResponse struct {
	ArticleResponse
	Content string `json:"content" doc:"文章内容"` // 新增
}

// ToArticleDetailResponse 封装单个对象的转换：MODEL -> Response
func ToArticleDetailResponse(article *query.ArticleModel) (ArticleDetailResponse, error) {
	var result ArticleDetailResponse
	err := copier.Copy(&result, article)
	if err != nil {
		return ArticleDetailResponse{}, err
	}
	return result, nil
}

// ToArticleResponseList 封装切片/列表的批量转换：[]MODEL -> []Response
func ToArticleResponseList(pos []*query.ArticleListItemModel) ([]ArticleResponse, error) {
	var dst []ArticleResponse
	err := copier.Copy(&dst, pos)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

type CreateFeedRequest struct {
	Content string `json:"content"`
}

type CreateArticleRequest struct {
	Body struct {
		Title   string  `json:"title" minLength:"1"`
		TypeID  *int    `json:"typeId,omitempty" minimum:"1"`
		Content *string `json:"content,omitempty"`
	}
}

type UpdateArticleRequest struct {
	ID   int `path:"id" minimum:"1" doc:"文章 ID"`
	Body struct {
		Title   *string `json:"title,omitempty" minLength:"1"`
		TypeID  *int    `json:"typeId,omitempty" minimum:"1"`
		Content *string `json:"content,omitempty"`
	}
}

func (r *UpdateArticleRequest) ToCommand() command.UpdateArticleCommand {
	return command.UpdateArticleCommand{
		ID:      r.ID,
		TypeID:  r.Body.TypeID,
		Title:   r.Body.Title,
		Content: r.Body.Content,
	}
}
