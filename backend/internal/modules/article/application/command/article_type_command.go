package command

type CreateArticleTypeCommand struct {
	Name string
	Slug string
}

type UpdateArticleTypeCommand struct {
	ID   int
	Name *string
	Slug *string
}
