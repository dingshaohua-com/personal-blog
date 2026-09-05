//go:build !production

package webui

import "net/http"

// 开发页面由 Vite 提供，首次启动无需生成 dist。
func Handler() http.Handler { return http.NotFoundHandler() }
