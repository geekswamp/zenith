// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package service

import (
	"github.com/geekswap/zenith/config"
	"github.com/geekswap/zenith/internal/repository"
	"github.com/geekswap/zenith/internal/service"
	"github.com/geekswap/zenith/pkg/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func ProvideService(cfg *config.Config, log logger.Logger) *service.Service {
	serviceService := service.New(cfg, log)
	return serviceService
}

func ProvideAccountService(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log logger.Logger) service.AccountService {
	serviceService := service.New(cfg, log)
	repositoryRepository := repository.New(db, rdb)
	accountRepository := repository.NewAccountRepository(repositoryRepository)
	accountService := service.NewAccountService(serviceService, accountRepository)
	return accountService
}

func ProvideNotificationService(db *gorm.DB, rdb *redis.Client, cfg *config.Config, log logger.Logger) service.NotificationService {
	serviceService := service.New(cfg, log)
	repositoryRepository := repository.New(db, rdb)
	notificationRepository := repository.NewNotificationRepository(repositoryRepository)
	notificationService := service.NewNotificationService(serviceService, notificationRepository)
	return notificationService
}
