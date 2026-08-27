package command

type CreateArticleCommand struct {
	Title   string
	TypeID  int
	Content string
}

type UpdateArticleCommand struct {
	ID      int
	Title   *string
	TypeID  *int
	Content *string
}
