package handler

import (
	"context"
	"github.com/shivaak/demo-grpc/controller/user"
	proto "github.com/shivaak/demo-grpc/gen"
	"github.com/shivaak/demo-grpc/mapper"
)

type UserHandler proto.UserManagementServer

type userhandler struct {
	controller user.Controller
	proto.UnimplementedUserManagementServer
}

func NewUserHandler(c user.Controller) UserHandler {
	return &userhandler{
		controller: c,
	}
}

func (c *userhandler) Create(ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	id, err := c.controller.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResponse{
		UserId: id.String(),
	}, nil
}

func (c *userhandler) Get(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	user, err := c.controller.Get(ctx, request)
	if err != nil {
		return nil, err
	}
	return &proto.GetUserResponse{
		User: mapper.UserEntityToProtoResponse(user),
	}, nil
}

func (c *userhandler) List(ctx context.Context, request *proto.ListUserRequest) (*proto.ListUserResponse, error) {
	list, err := c.controller.List(ctx, request)
	if err != nil {
		return nil, err
	}

	var finalList []*proto.UserData
	for _, l := range list {
		finalList = append(finalList, mapper.UserEntityToProtoResponse(l))
	}

	return &proto.ListUserResponse{
		UserList: finalList,
	}, nil
}
