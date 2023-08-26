package app

import (
	"github.com/shivaak/demo-grpc/controller"
	"github.com/shivaak/demo-grpc/handler"
	"github.com/shivaak/demo-grpc/loggerfx"
	"github.com/shivaak/demo-grpc/repository"
	"go.uber.org/fx"
)

var Module = fx.Options(
	loggerfx.Module,
	controller.Module,
	handler.Module,
	repository.Module,
)
