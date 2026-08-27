package dto

type GetFeedRequest struct {
	ID int `path:"id" minimum:"1" doc:"说说 ID"`
}

// CreateFeedRequest 对应 POST /feed 的 JSON 请求体。
type CreateFeedRequest struct {
	Body struct {
		Content string `json:"content" minLength:"1" maxLength:"100" doc:"说说内容"`
	}
}

// UpdateFeedRequest 对应 PUT /feed/{id}。
type UpdateFeedRequest struct {
	ID   int `path:"id" minimum:"1" doc:"文章 ID"`
	Body struct {
		Content string `json:"content" minLength:"1" maxLength:"100" doc:"文章内容"`
	}
}

// DeleteFeedRequest 对应 DELETE /feed/{id}。
type DeleteFeedRequest struct {
	ID int `path:"id" minimum:"1" doc:"文章 ID"`
}
