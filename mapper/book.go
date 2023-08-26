package mapper

import (
	"github.com/shivaak/demo-grpc/entity"
	proto "github.com/shivaak/demo-grpc/gen"
	"github.com/shivaak/demo-grpc/model"
)

func BookToModel(b *entity.Book) *model.Book {
	return &model.Book{
		Id:    b.Id,
		Title: b.Title,
		Count: b.Count,
	}
}

func ModelToBook(b *model.Book) *entity.Book {
	return &entity.Book{
		Id:    b.Id,
		Title: b.Title,
		Count: b.Count,
	}
}

func BookEntityToProtoResponse(b *entity.Book) *proto.Book {
	return &proto.Book{
		Id:    b.Id.String(),
		Title: b.Title,
		Count: b.Count,
	}
}

func ProtoCreateRequestToBook(r *proto.CreateBookRequest) *entity.Book {
	return &entity.Book{
		Title: r.Book.Title,
		Count: r.Book.Count,
	}
}
