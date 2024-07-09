// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"lin-cms-go/internal/biz"
	"lin-cms-go/internal/conf"
	"lin-cms-go/internal/data"
	"lin-cms-go/internal/server"
	"lin-cms-go/internal/service"
)

// Injectors from wire.go:

// NewWire init kratos application.
func NewWire(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, nil, err
	}
	linUserRepo := data.NewLinUserRepo(dataData, logger)
	linUserusecase := biz.NewLinUserUsecase(linUserRepo, logger)
	cmsService := service.NewCmsService(linUserusecase)
	bookRepo := data.NewBookRepo(dataData, logger)
	bookUsecase := biz.NewBookUsecase(bookRepo, logger)
	lessonRepo := data.NewLessonRepo(dataData, logger)
	lessonUsecase := biz.NewLessonUsecase(lessonRepo, logger)
	teacherRepo := data.NewTeacherRepo(dataData, logger)
	teacherUsecase := biz.NewTeacherUsecase(teacherRepo, logger)
	appService := service.NewAppService(bookUsecase, lessonUsecase, teacherUsecase)
	httpServer := server.NewHTTPServer(confServer, cmsService, appService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}