package api

import "github.com/danielgtaylor/huma/v2"

func NewError(status int, msg string) huma.StatusError {
	return huma.NewError(status, msg)
}

func InternalError(messages ...string) huma.StatusError {
	msg := "服务器内部错误"
	if len(messages) > 0 && messages[0] != "" {
		msg += "：" + messages[0]
	}
	return huma.Error500InternalServerError(msg)
}
