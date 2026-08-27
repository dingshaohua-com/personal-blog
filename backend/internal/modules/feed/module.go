package feed

import (
	"backend/internal/modules/feed/api"
	"backend/internal/modules/feed/application"
	"backend/internal/modules/feed/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, serverApi huma.API) {
	repo := infrastructure.NewFeedRepository(db)
	serv := application.NewFeedService(repo)
	hand := api.NewFeedHandler(serv)
	api.RegisterRoutes(hand, serverApi)
}
