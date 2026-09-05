package bootstrap

import (
	sharedApi "backend/internal/shared/api"
	"backend/internal/webui"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"backend/internal/infrastructure"

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
		Handler: sharedApi.Cors(NewHTTPHandler(app.Router, webui.Handler())),
	}
	serviceURLs := []string{"http://localhost" + addr}
	docURLs := []string{"http://localhost" + addr + "/docs"}
	for _, ip := range localLANIPv4Addresses() {
		serviceURLs = append(serviceURLs, "http://"+ip+addr)
		docURLs = append(docURLs, "http://"+ip+addr+"/docs")
	}
	log.Printf("HTTP已服务启动: %s", strings.Join(serviceURLs, ", "))
	log.Printf("OpenAPI文档: %s", strings.Join(docURLs, ", "))
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("HTTP 服务启动失败: %v", err)
	}
}

func localLANIPv4Addresses() []string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var result []string
	seen := make(map[string]struct{})
	for _, networkInterface := range interfaces {
		if networkInterface.Flags&net.FlagUp == 0 || networkInterface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addresses, err := networkInterface.Addrs()
		if err != nil {
			continue
		}
		for _, address := range addresses {
			ip, _, err := net.ParseCIDR(address.String())
			if err != nil || ip.To4() == nil || !ip.IsPrivate() {
				continue
			}
			value := ip.String()
			if _, exists := seen[value]; exists {
				continue
			}
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
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
	NewAPI(router, db)

	return &App{
		Router:   router,
		Config:   cfg,
		Database: db,
		Redis:    redisClient,
	}, nil
}
