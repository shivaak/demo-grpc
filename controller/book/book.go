package book

import (
	"context"
	"github.com/google/uuid"
	"github.com/shivaak/demo-grpc/entity"
	proto "github.com/shivaak/demo-grpc/gen"
	"github.com/shivaak/demo-grpc/mapper"
	"github.com/shivaak/demo-grpc/repository/inmemory"
)

type BookController interface {
	Create(ctx context.Context, request *proto.CreateBookRequest) (uuid.UUID, error)
	List(ctx context.Context, request *proto.EmptyRequest) ([]*entity.Book, error)
}

type bookcontroller struct {
	repository inmemory.Repository
}

func New(r inmemory.Repository) BookController {
	return &bookcontroller{
		repository: r,
	}
}

func (c *bookcontroller) Create(ctx context.Context, request *proto.CreateBookRequest) (uuid.UUID, error) {
	id, err := c.repository.CreateBook(mapper.ProtoCreateRequestToBook(request))
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (c *bookcontroller) List(ctx context.Context, request *proto.EmptyRequest) ([]*entity.Book, error) {
	list, err := c.repository.ListBook()
	if err != nil {
		return nil, err
	}
	return list, nil
}
