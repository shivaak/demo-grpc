package controller

import (
	"github.com/shivaak/demo-grpc/controller/book"
	"github.com/shivaak/demo-grpc/controller/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	user.New,
	book.New,
)
