package handler

import (
	"github.com/shivaak/demo-grpc/handler/book"
	"github.com/shivaak/demo-grpc/handler/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.NewUserHandler,
	book.NewBookHandler,
)
