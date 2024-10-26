//go:build wireinject

package logger

import (
	"github.com/geekswap/zenith/pkg/logger"
	"github.com/google/wire"
)

func ProvideLogger() logger.Logger {
	wire.Build(logger.New)
	return logger.Logger{}
}
