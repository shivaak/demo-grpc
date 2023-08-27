package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/shivaak/demo-grpc/entity"
	"github.com/shivaak/demo-grpc/mapper"
	proto "github.com/shivaak/demo-grpc/proto"
	"github.com/shivaak/demo-grpc/repository/inmemory"
)

type Controller interface {
	Create(ctx context.Context, request *proto.CreateUserRequest) (uuid.UUID, error)
	Get(ctx context.Context, request *proto.GetUserRequest) (*entity.User, error)
	List(context.Context, *proto.ListUserRequest) ([]*entity.User, error)
}

type controller struct {
	repository inmemory.Repository
}

func New(repository inmemory.Repository) Controller {
	return &controller{
		repository: repository,
	}
}

func (c *controller) Create(ctx context.Context, request *proto.CreateUserRequest) (uuid.UUID, error) {
	id, err := c.repository.CreateUser(mapper.ProtoCreateRequestToUser(request))
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (c *controller) Get(ctx context.Context, request *proto.GetUserRequest) (*entity.User, error) {
	user, err := c.repository.GetUser(uuid.MustParse(request.UserId))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *controller) List(ctx context.Context, request *proto.ListUserRequest) ([]*entity.User, error) {
	list, err := c.repository.ListUser()
	if err != nil {
		return nil, err
	}
	return list, nil
}
