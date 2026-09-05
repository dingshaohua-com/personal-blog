package bootstrap

import (
	"backend/internal/modules/article"
	"backend/internal/modules/feed"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"gorm.io/gorm"
)

// NewAPI registers the HTTP routes and their OpenAPI documentation.
// Registration does not query the database; tests can pass nil.
func NewAPI(router *http.ServeMux, db *gorm.DB) huma.API {
	config := huma.DefaultConfig("My API", "1.0.0")
	config.CreateHooks = nil
	api := humago.New(router, config)
	article.RegisterModule(db, api)
	feed.RegisterModule(db, api)
	return api
}
