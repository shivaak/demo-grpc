package repository

import (
	"github.com/shivaak/demo-grpc/repository/inmemory"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	inmemory.NewInMemoryRepo,
)
