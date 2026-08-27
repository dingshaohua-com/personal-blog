package api

import (
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"

	"github.com/jinzhu/copier"
)

type ArticleTypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func ToArticleTypeResponseList(pos []*query.ArticleTypeModel) []ArticleTypeResponse {
	var dst []ArticleTypeResponse
	err := copier.Copy(&dst, pos)
	if err != nil {
		return nil
	}
	return dst
}

type CreateArticleTypeRequest struct {
	Body struct {
		Name string `json:"name" minLength:"1"`
		Slug string `json:"slug" minLength:"1"`
	}
}

type UpdateArticleTypeRequest struct {
	ID   int `path:"id" minimum:"1" doc:"类型 ID"`
	Body struct {
		Name string `json:"name,omitempty" minLength:"1"`
		Slug string `json:"slug,omitempty" minLength:"1"`
	}
}

func (r *CreateArticleTypeRequest) ToCommand() command.CreateArticleTypeCommand {
	return command.CreateArticleTypeCommand{
		Name: r.Body.Name,
		Slug: r.Body.Slug,
	}
}

type DeleteArticleTypeRequest struct {
	ID int `path:"id" minimum:"1" doc:"类型 ID"`
}

type CreateArticleTypeResponse struct {
	ID int `json:"id"`
}

func (r *UpdateArticleTypeRequest) ToCommand() command.UpdateArticleTypeCommand {
	return command.UpdateArticleTypeCommand{
		ID:   r.ID,
		Name: &r.Body.Name,
		Slug: &r.Body.Slug,
	}
}
