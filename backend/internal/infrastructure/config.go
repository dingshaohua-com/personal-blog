package infrastructure

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv   string
	HTTPPort string
	Database string
	Redis    RedisConfig
}

func LoadConfig() *Config {
	// 如果没有 .env（例如线上环境），忽略错误即可
	_ = godotenv.Load()
	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	return &Config{
		AppEnv:   getEnv("APP_ENV", "dev"),
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		Database: getEnv("PGSQL_PARAMS", ""),
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "127.0.0.1:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       redisDB,
		},
	}
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
