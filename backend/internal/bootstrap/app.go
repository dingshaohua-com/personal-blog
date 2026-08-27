package bootstrap

import (
	"backend/internal/modules/article"
	"backend/internal/modules/feed"
	sharedApi "backend/internal/shared/api"
	"fmt"
	"log"
	"net/http"

	"backend/internal/infrastructure"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	Router   *http.ServeMux
	Config   *infrastructure.Config
	Database *gorm.DB
	Server   *http.Server
	Redis    *redis.Client
}

func (app *App) Run() {
	addr := ":" + app.Config.HTTPPort
	server := &http.Server{
		Addr:    addr,
		Handler: sharedApi.Cors(app.Router),
	}
	log.Printf(
		"HTTP已服务启动: http://localhost%s",
		addr,
	)
	log.Printf(
		"OpenAPI文档: http://localhost%s/docs",
		addr,
	)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("HTTP 服务启动失败: %v", err)
	}
}

// NewApp 这里负责组装整个系统
func NewApp() (*App, error) {
	// 加载环境变量
	cfg := infrastructure.LoadConfig()

	// 初始化数据库和缓存
	db, err := infrastructure.NewPostgres(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("初始化 PostgreSQL: %w", err)
	}
	redisClient, redisErr := infrastructure.NewRedis(cfg.Redis)
	if redisErr != nil {
		sqlDB, dbErr := db.DB()
		if dbErr == nil {
			_ = sqlDB.Close()
		}
		return nil, fmt.Errorf("初始化 Redis: %w", redisErr)
	}

	// 初始化服务
	router := http.NewServeMux()
	config := huma.DefaultConfig("My API", "1.0.0")
	config.CreateHooks = nil
	//config.DocsRenderer = huma.DocsRendererScalar
	//config.DocsRenderer = huma.DocsRendererSwaggerUI
	api := humago.New(router, config)

	// 初始化服务模块
	article.RegisterModule(db, api)
	feed.RegisterModule(db, api)

	return &App{
		Router:   router,
		Config:   cfg,
		Database: db,
		Redis:    redisClient,
	}, nil
}
