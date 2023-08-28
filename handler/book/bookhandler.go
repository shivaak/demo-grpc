package book

import (
	"context"
	"github.com/shivaak/demo-grpc/controller/book"
	"github.com/shivaak/demo-grpc/mapper"
	proto "github.com/shivaak/demo-grpc/proto"
)

type BookHandler proto.BookManagementServer

type bookhandler struct {
	controller book.BookController
	proto.UnimplementedBookManagementServer
}

func NewBookHandler(c book.BookController) BookHandler {
	return &bookhandler{
		controller: c,
	}
}

func (h *bookhandler) ListBooks(c context.Context, r *proto.EmptyRequest) (*proto.ListBookResponse, error) {
	books, err := h.controller.List(c, r)
	if err != nil {
		return nil, err
	}

	var finalList []*proto.Book
	for _, l := range books {
		finalList = append(finalList, mapper.BookEntityToProtoResponse(l))
	}

	return &proto.ListBookResponse{
		BookList: finalList,
	}, nil
}

func (h *bookhandler) CreateBook(ctx context.Context, request *proto.CreateBookRequest) (*proto.CreateBookResponse, error) {
	id, err := h.controller.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return &proto.CreateBookResponse{
		Id: id.String(),
	}, nil
}
