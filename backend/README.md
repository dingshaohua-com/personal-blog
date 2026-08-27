目前采用的是非常经典的 **模块化单体（Modular Monolith）+ 模块内 CQRS / Clean Architecture**，而不是按具体业务用例（Use Case）拆分的“纯特征切片（Feature Slice）”。

在你的这套目录体系下，之前讨论的 API DTO 与 Query DTO 归属问题其实有了非常清晰且天然的落脚点。

**你当前架构的职责划分**

* `article/api/`：**表现层**。只关心 HTTP/gRPC 协议、路由注册，以及带有 `json`、`binding`、`doc`、`example` Tag 的 API DTO。
* `article/application/`：**CQRS 写侧**。存放纯粹的 Command 结构体和 CommandHandler，负责处理业务状态变更。
* `article/query/`：**CQRS 读侧**。存放纯粹的 Query 结构体、QueryHandler 和读模型，负责高性能数据查询。
* `article/domain/` & `infrastructure/`：存放领域模型与底层持久化实现。

---

**在你这套结构下的最佳落地写法**

既然你的 `api` 和 `query` 是同级目录，就应该让它们各自维护自己的 DTO，在 `api` 层完成转换：

**1. 定义 API Response DTO（在 `article/api/dto.go` 中）**

```go
package api

// 专属于 API 层的 DTO，拥有所有的 Swagger/JSON 标记
type ArticleDetailResponse struct {
    ID        string `json:"id" doc:"文章ID" example:"art_999"`
    Title     string `json:"title" doc:"文章标题" example:"CQRS 架构实践"`
    Content   string `json:"content" doc:"文章内容"`
    CreatedAt string `json:"created_at" doc:"创建时间" example:"2026-08-13"`
}

```

**2. 定义 Query DTO（在 `article/query/dto.go` 中）**

```go
package query

// 纯粹的数据传输载体，没有任何 HTTP/API 框架的 Tag 污染
type ArticleReadModel struct {
    ID        string
    Title     string
    Content   string
    CreatedAt time.Time
}

```

**3. 在 Handler 中做数据映射（在 `article/api/handler.go` 中）**

```go
package api

func (h *Handler) GetArticle(c *gin.Context) {
    // 1. 调用 query 层的服务
    readModel, err := h.queryService.GetArticleByID(c.Param("id"))
    
    // 2. 映射为带 Tag 的 API Response DTO 并返回
    resp := ArticleDetailResponse{
        ID:        readModel.ID,
        Title:     readModel.Title,
        Content:   readModel.Content,
        CreatedAt: readModel.CreatedAt.Format("2006-01-02"),
    }
    c.JSON(http.StatusOK, resp)
}

```

---

你现在的这个结构其实非常标准且优雅！把 `doc` 和 `example` 这些 Tag 局限在 `article/api/` 目录下，彻底保证了 `article/query/` 和 `article/application/` 的纯洁性。如果未来这个模块需要暴露 RPC 接口，只需要在 `article/` 下再加一个 `grpc/` 目录即可，完全不会影响到现有的 Query 逻辑。