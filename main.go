package main

import (
	"context"
	"github.com/shivaak/demo-grpc/app"
	"github.com/shivaak/demo-grpc/handler/book"
	"github.com/shivaak/demo-grpc/handler/user"
	"github.com/shivaak/demo-grpc/proto"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func opts() fx.Option {
	return fx.Options(
		app.Module,
		fx.Invoke(registerHooks),
	)
}

func main() {
	app := fx.New(opts())
	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle,
	logger *zap.SugaredLogger,
	userHandler user.UserHandler,
	bookHandler book.BookHandler) {
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
