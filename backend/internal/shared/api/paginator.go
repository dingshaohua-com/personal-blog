package api

import (
	"backend/internal/shared/pagination"
	"math"
)

// Page 分页请求对象
type Page struct {
	Page     int `query:"page" default:"1" doc:"当前页码"`
	PageSize int `query:"pageSize" default:"10" doc:"每页条数"`
}

// NewPage 构造并初始化
func NewPage(page, pageSize int) *Page {
	normalized := pagination.New(page, pageSize)
	return &Page{Page: normalized.Page, PageSize: normalized.PageSize}
}

// Normalization 规范化清洗参数（必须指接收器）
func (p *Page) Normalization() {
	if p == nil {
		return
	}
	normalized := pagination.New(p.Page, p.PageSize)
	p.Page = normalized.Page
	p.PageSize = normalized.PageSize
}

// Limit 保持一致，改为指针接收器
func (p *Page) Limit() int {
	p.Normalization() // 安全保障：获取 Limit 时顺便自动清洗
	return p.PageSize
}

// Offset 保持一致，改为指针接收器
func (p *Page) Offset() int {
	p.Normalization()
	return (p.Page - 1) * p.PageSize
}

// PageResult 通用响应结果
type PageResult[T any] struct {
	List      []T   `json:"list" doc:"数据列表"`
	Total     int64 `json:"total" doc:"总记录数"`
	Page      int   `json:"page" doc:"当前页码"`
	PageSize  int   `json:"pageSize" doc:"每页条数"`
	TotalPage int   `json:"totalPage" doc:"总页数"`
}

// NewPageResult 构造返回结果，接收 *Page 参数
func NewPageResult[T any](list []T, total int64, p *Page) PageResult[T] {
	if p == nil {
		p = NewPage(1, 10)
	} else {
		p.Normalization()
	}

	totalPage := 0
	if total > 0 {
		totalPage = int(math.Ceil(float64(total) / float64(p.PageSize)))
	}

	if list == nil {
		list = make([]T, 0)
	}

	return PageResult[T]{
		List:      list,
		Total:     total,
		Page:      p.Page,
		PageSize:  p.PageSize,
		TotalPage: totalPage,
	}
}
