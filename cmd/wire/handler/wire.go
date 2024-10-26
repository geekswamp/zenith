//go:build wireinject

package handler

import (
	cmn "github.com/geekswap/zenith/cmd/wire/common"
	repo "github.com/geekswap/zenith/cmd/wire/repository"
	"github.com/geekswap/zenith/config"
	"github.com/geekswap/zenith/internal/handler"
	"github.com/geekswap/zenith/internal/repository"
	"github.com/geekswap/zenith/internal/service"
	"github.com/geekswap/zenith/pkg/logger"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ProvideHandler() *handler.Handler {
	wire.Build(cmn.ProvideResponse, handler.New)
	return &handler.Handler{}
}

func ProvideAccountHandler(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log logger.Logger) *handler.AccountHandler {
	wire.Build(cmn.ProvideResponse, repo.ProvideRepository, handler.New, service.New, repository.NewAccountRepository, service.NewAccountService, handler.NewAccountHandler)
	return &handler.AccountHandler{}
}

func ProvideNotificationHandler(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log logger.Logger) *handler.NotificationHandler {
	wire.Build(cmn.ProvideResponse, repo.ProvideRepository, handler.New, service.New, repository.NewNotificationRepository, service.NewNotificationService, handler.NewNotificationHandler)
	return &handler.NotificationHandler{}
}
