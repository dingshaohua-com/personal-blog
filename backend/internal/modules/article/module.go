package article

import (
	"backend/internal/modules/article/api"
	"backend/internal/modules/article/application/command"
	"backend/internal/modules/article/application/query"
	"backend/internal/modules/article/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func RegisterModule(db *gorm.DB, serverApi huma.API) {

	// 1. 写路径依赖组装 (Command Side)
	articleRepo := infrastructure.NewArticleRepository(db)
	articleTypeRepo := infrastructure.NewArticleTypeRepository(db)

	articleCmdSvc := command.NewArticleService(articleRepo) // 可以叫 CmdSvc，强调写路径
	articleTypeCmdSvc := command.NewArticleTypeService(articleTypeRepo)

	// 2. 读路径依赖组装 (Query Side - 直连 db，没有任何其他依赖！)
	articleQuerySvc := query.NewArticleService(db) // 推荐把结构体/构造命名为 Service 或 QueryService
	articleTypeQuerySvc := query.NewArticleTypeService(db)

	// 3. API 层同时注入【写服务】与【读服务】
	articleHandler := api.NewArticleHandler(articleCmdSvc, articleQuerySvc)
	articleTypeHandler := api.NewArticleTypeHandler(articleTypeCmdSvc, articleTypeQuerySvc)

	api.RegisterRoutes(articleHandler, articleTypeHandler, serverApi)
}
