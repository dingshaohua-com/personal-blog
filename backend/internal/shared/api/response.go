package api

// Body 是不额外包装业务信封的 Huma 响应结构。
type Body[T any] struct {
	Body T
}

func NewBody[T any](data T) *Body[T] {
	return &Body[T]{Body: data}
}

func NoContent() *struct{} {
	return &struct{}{}
}
