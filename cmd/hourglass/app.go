package main

import (
	"github.com/ku1bu/hourglass/internal/conf"
	grpcv1 "github.com/ku1bu/hourglass/internal/grpc/v1"
	"github.com/ku1bu/hourglass/internal/server"
	"github.com/ku1bu/hourglass/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// initApp init kratos application.
func initApp(confServer *conf.Server, logger log.Logger) (*kratos.App, func(), error) {
	coo := service.NewCoordinator()
	provier := grpcv1.NewProvider(coo, logger)
	customer := grpcv1.NewCustomer(coo, logger)

	grpcServer := server.NewGRPCServer(confServer, provier, customer, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
	}, nil
}
