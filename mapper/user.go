package mapper

import (
	"github.com/shivaak/demo-grpc/entity"
	proto "github.com/shivaak/demo-grpc/gen"
	"github.com/shivaak/demo-grpc/model"
)

func UserToModel(u *entity.User) *model.User {
	return &model.User{
		Id:      u.Id,
		Name:    u.Name,
		Age:     u.Age,
		City:    u.City,
		Country: u.Country,
	}
}

func ModelToUser(u *model.User) *entity.User {
	return &entity.User{
		Id:      u.Id,
		Name:    u.Name,
		Age:     u.Age,
		City:    u.City,
		Country: u.Country,
	}
}

func ProtoCreateRequestToUser(request *proto.CreateUserRequest) *entity.User {
	return &entity.User{
		Name:    request.User.Name,
		Age:     request.User.Age,
		City:    request.User.City,
		Country: request.User.Country,
	}
}

func UserEntityToProtoResponse(user *entity.User) *proto.UserData {
	return &proto.UserData{
		Name:    user.Name,
		Age:     user.Age,
		City:    user.City,
		Country: user.Country,
		Id:      user.Id.String(),
	}
}
