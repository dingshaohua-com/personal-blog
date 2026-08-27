package query

type ArticleTypeModel struct {
	ID   int
	Name string
	Slug string
}

func (ArticleTypeModel) TableName() string {
	return "article_type"
}
