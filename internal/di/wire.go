// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/znyh/account/internal/dao"
	"github.com/znyh/account/internal/server/grpc"
	"github.com/znyh/account/internal/server/http"
	"github.com/znyh/account/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
