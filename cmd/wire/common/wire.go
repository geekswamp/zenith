//go:build wireinject

package common

import (
	"github.com/geekswap/zenith/pkg/common"
	"github.com/google/wire"
)

func ProvideResponse() *common.Response {
	wire.Build(common.NewResponse)
	return &common.Response{}
}
