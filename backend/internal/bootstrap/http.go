package bootstrap

import (
	"net/http"
	"strings"
)

// NewHTTPHandler 保留原 API 地址，并提供与 Vite 代理一致的 /api 前缀。
func NewHTTPHandler(api *http.ServeMux, frontend http.Handler) http.Handler {
	prefixed := http.StripPrefix("/api", api)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" || strings.HasPrefix(r.URL.Path, "/api/") {
			prefixed.ServeHTTP(w, r)
			return
		}
		if _, pattern := api.Handler(r); pattern != "" {
			api.ServeHTTP(w, r)
			return
		}
		frontend.ServeHTTP(w, r)
	})
}
