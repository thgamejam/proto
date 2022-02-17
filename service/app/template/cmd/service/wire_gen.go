// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"service/app/template/internal/biz"
	"service/app/template/internal/data"
	"service/app/template/internal/server"
	"service/app/template/internal/service"
	"service/pkg/conf"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confService *conf.Service, registrar registry.Registrar, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confService, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	app := newApp(logger, registrar, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
