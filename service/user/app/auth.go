package app

import (
	"context"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	genDB "nico_minidouyin/gen/db"
	"nico_minidouyin/gen/douyin/auth"
	pb "nico_minidouyin/gen/douyin/auth"
	"nico_minidouyin/mw"
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

func (a AuthService) Authenticate(ctx context.Context, request *auth.AuthenticateRequest) (resp *auth.AuthenticateResponse, err error) {
	userToken := genDB.Q.UserToken
	first, err := userToken.WithContext(ctx).Where(userToken.Token.Eq(request.Token)).First()
	if err != nil {
		return nil, mw.NewBizError("你系邊個啊", err)
	}
	resp = &auth.AuthenticateResponse{
		StatusCode: 6,
		StatusMsg:  "6",
		UserId:     first.UserID,
	}
	return
}

func (a AuthService) Register(ctx context.Context, request *auth.RegisterRequest) (resp *auth.RegisterResponse, err error) {
	resp = &auth.RegisterResponse{}
	user := genDB.Q.User
	userToken := genDB.Q.UserToken
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
	token := ksuid.New().String()
	err = userToken.WithContext(ctx).Save(&model.UserToken{
		Token:    token,
		Username: newUser.Username,
		UserID:   newUser.ID,
		Role:     newUser.Role,
	})

	if err != nil {
		return nil, mw.NewBizError("數據庫保存失敗", err)
	}
	resp.Token = token
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return
}

func (a AuthService) Login(ctx context.Context, request *auth.LoginRequest) (resp *auth.LoginResponse, err error) {
	resp = &auth.LoginResponse{}
	user := genDB.Q.User
	userToken := genDB.Q.UserToken
	dbUser, err := user.WithContext(ctx).Where(genDB.User.Username.Eq(request.Username)).Select().Find()
	if err != nil {
		return nil, mw.NewBizError("數據庫查詢失敗", err)
	}
	if len(dbUser) != 1 {
		return nil, mw.NewBizError("冇搵到用戶")
	}
	if err != nil {
		return nil, mw.NewBizError("密碼唔好使", err)
	}
	if !CheckPasswordHash(request.Password, *dbUser[0].Password) {
		return nil, mw.NewBizError("密碼唔對")
	}

	tokens, err := userToken.WithContext(ctx).Where(userToken.UserID.Eq(dbUser[0].ID)).Find()
	if len(tokens) == 1 {
		resp.Token = tokens[0].Token
		resp.StatusCode = 0
		resp.StatusMsg = "success"
		return
	}
	if err != nil {
		return nil, err
	}
	token := ksuid.New().String()
	err = userToken.WithContext(ctx).Save(&model.UserToken{
		Token:    token,
		Username: dbUser[0].Username,
		UserID:   dbUser[0].ID,
		Role:     dbUser[0].Role,
	})

	if err != nil {
		return nil, mw.NewBizError("數據庫保存失敗", err)
	}
	resp.Token = token
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	return
}
