package main

import (
	"fmt"

	"auth-service/internal/config"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/pkg/handler"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	cfg, err := config.LoadConfig("app/config.yaml")
	if err != nil {
		sugar.Fatal("Failed to load configuration:", err)
	}

	db := initializeDatabase(cfg, sugar)

	memcachedClient := initializeMemcached(cfg, sugar)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, memcachedClient, sugar, []byte(cfg.JWT.Secret))

	authController := handler.NewAuthController(authService, sugar)

	router := gin.Default()
	authController.RegisterRoutes(router)

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := router.Run(address); err != nil {
		sugar.Fatal("Failed to start server:", err)
	}
}

func initializeDatabase(cfg *config.Config, logger *zap.SugaredLogger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}

	logger.Info("Connected to the database")

	return db
}

func initializeMemcached(cfg *config.Config, logger *zap.SugaredLogger) *memcache.Client {
	memcachedAddr := fmt.Sprintf("%s:%d", cfg.Memcached.Host, cfg.Memcached.Port)
	memcachedClient := memcache.New(memcachedAddr)
	logger.Info("Connected to Memcached at ", memcachedAddr)
	return memcachedClient
}
