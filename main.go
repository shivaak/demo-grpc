package main

import (
	"context"
	"github.com/shivaak/demo-grpc/app"
	proto "github.com/shivaak/demo-grpc/gen"
	"github.com/shivaak/demo-grpc/handler"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	app := fx.New(
		app.Module,
		fx.Invoke(registerHooks),
	)
	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	userHandler handler.UserHandler,
	bookHandler handler.BookHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				listener, err := net.Listen("tcp", ":8089")
				if err != nil {
					logger.Fatalf("Unable to open listener %s", err)
				}
				logger.Info("Server starting at 8089...")

				// rpc server
				var opts []grpc.ServerOption
				grpcServer := grpc.NewServer(opts...)
				proto.RegisterUserManagementServer(grpcServer, userHandler)
				proto.RegisterBookManagementServer(grpcServer, bookHandler)
				go grpcServer.Serve(listener)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
