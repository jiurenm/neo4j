package di

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"neo4j/internal/dao"
	mygrpc "neo4j/internal/server/grpc"
	"neo4j/internal/service"
	"neo4j/pkg/conf"
)

func InitApp1() (*grpc.Server, func(), error) {
	panic(wire.Build(conf.ConfigFile, conf.New, dao.New, service.New, mygrpc.New))
}
