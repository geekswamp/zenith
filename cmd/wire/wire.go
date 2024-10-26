//go:build wireinject

package wire

import (
	"github.com/geekswap/zenith/cmd/wire/handler"
	"github.com/geekswap/zenith/cmd/wire/middleware"
	"github.com/geekswap/zenith/config"
	"github.com/geekswap/zenith/pkg/logger"
	"github.com/geekswap/zenith/pkg/server/http"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func InitializeRouter(db *gorm.DB, redis *redis.Client, cfg *config.Config, log logger.Logger) *gin.Engine {
	wire.Build(
		handler.ProvideAccountHandler,
		handler.ProvideNotificationHandler,
		middleware.WireMiddlewareSet,
		http.ProvideGinEngine,
	)
	return &gin.Engine{}
}
