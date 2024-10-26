package middleware

import (
	"github.com/geekswap/zenith/internal/middleware"
	"github.com/google/wire"
)

var WireMiddlewareSet = wire.NewSet(
	middleware.New,
	middleware.NewStrictAuthMiddleware,
)
