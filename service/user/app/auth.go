package app

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"nico_minidouyin/mw"
	"strconv"

	genDB "nico_minidouyin/gen/db"
	"nico_minidouyin/gen/douyin/auth"
	pb "nico_minidouyin/gen/douyin/auth"
	"nico_minidouyin/service/user/model"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
		return nil, mw.NewBizError("數據庫查詢失敗", err)
	}
	if len(dbUser) > 0 {
		return nil, mw.NewBizError("你嘅用戶名已被占用")
	}
	hashedPwd, err := HashPassword(request.Password)
	if err != nil {
		return nil, mw.NewBizError("密碼唔好使", err)
	}
	newUser := model.User{
		Username:      request.Username,
		Password:      &hashedPwd,
		FollowCount:   0,
		FollowerCount: 0,
		Name:          request.Username,
		Role:          "0",
	}
	err = user.WithContext(ctx).Save(
		&newUser)
	if err != nil {
		return nil, mw.NewBizError("數據庫保存失敗", err)
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
		return nil, mw.NewBizError("數據庫查詢失敗", err)
	}
	if len(dbUser) == 0 {
		return nil, mw.NewBizError("冇搵到用戶")
	}
	hashedPwd, err := HashPassword(request.Password)
	if err != nil {
		return nil, mw.NewBizError("密碼唔好使", err)
	}
	if CheckPasswordHash(*dbUser[0].Password, hashedPwd) {
		return nil, mw.NewBizError("密碼唔對")
	}
	token := strconv.Itoa(int(dbUser[0].ID)) + dbUser[0].Username
	resp.Token = token
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return resp, nil
}
