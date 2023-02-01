package app

import (
	"context"
	"strconv"

	genDB "nico_minidouyin/gen/db"
	"nico_minidouyin/gen/douyin/auth"
	pb "nico_minidouyin/gen/douyin/auth"
	"nico_minidouyin/service/user/model"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func (a AuthService) Register(ctx context.Context, request *auth.RegisterRequest) (
	resp *auth.RegisterResponse,
	err error) {
	resp = &auth.RegisterResponse{}
	user := genDB.Q.User
	dbUser, err := user.WithContext(ctx).Where(genDB.User.Username.Eq(request.Username)).Select().Find()
	if err != nil {
		return nil, err
	}
	if len(dbUser) > 0 {
		resp.StatusCode = 114514
		resp.StatusMsg = "user already exists"
		return
	}
	newUser := model.User{
		Username:      request.Username,
		Password:      &request.Password,
		FollowCount:   0,
		FollowerCount: 0,
		Name:          request.Username,
		Role:          "0",
	}
	err = user.WithContext(ctx).Save(
		&newUser)
	if err != nil {
		return nil, err
	}
	resp.Token = strconv.Itoa(int(newUser.ID)) + newUser.Username
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return
}

func (a AuthService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp := &auth.LoginResponse{}
	user := genDB.Q.User
	dbUser, err := user.WithContext(ctx).Where(genDB.User.Username.Eq(request.Username)).Select().Find()
	if err != nil {
		return nil, err
	}
	if len(dbUser) == 0 {
		errorMessage := "user not found"
		resp.StatusCode = 114514
		resp.StatusMsg = errorMessage
		return resp, nil
	}
	if *dbUser[0].Password != request.Password {
		errorMessage := "password incorrect"
		resp.StatusCode = 114514
		resp.StatusMsg = errorMessage
		return resp, nil
	}
	token := strconv.Itoa(int(dbUser[0].ID)) + dbUser[0].Username
	resp.Token = token
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return resp, nil
}
