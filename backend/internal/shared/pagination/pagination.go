package pagination

const (
	DefaultPage     = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

type Params struct {
	Page     int
	PageSize int
}

func New(page, pageSize int) Params {
	return Params{Page: page, PageSize: pageSize}.Normalize()
}

func (p Params) Normalize() Params {
	if p.Page <= 0 {
		p.Page = DefaultPage
	}
	if p.PageSize <= 0 {
		p.PageSize = DefaultPageSize
	} else if p.PageSize > MaxPageSize {
		p.PageSize = MaxPageSize
	}
	return p
}

func (p Params) Limit() int {
	return p.Normalize().PageSize
}

func (p Params) Offset() int {
	p = p.Normalize()
	return (p.Page - 1) * p.PageSize
}

type Result[T any] struct {
	Items  []T
	Total  int64
	Params Params
}
