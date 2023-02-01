package app

import (
	"context"
	genDB "nico_minidouyin/gen/db"
	pb "nico_minidouyin/gen/douyin/user"
	_ "nico_minidouyin/repo"
	"strconv"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (u UserService) GetUser(ctx context.Context, request *pb.UserRequest) (resp *pb.UserResponse, err error) {
	uidInt, err := strconv.Atoi(request.UserId)
	if err != nil {
		return nil, err
	}
	uidUint32 := uint32(uidInt)
	resp = &pb.UserResponse{}
	dbUser, err := genDB.Q.User.WithContext(ctx).Where(genDB.Q.User.ID.Eq(uidUint32)).First()
	if err != nil {
		return nil, err
	}
	resp.User = &pb.User{
		Id:            uidUint32,
		Name:          dbUser.Name,
		FollowCount:   dbUser.FollowCount,
		FollowerCount: dbUser.FollowerCount,
		IsFollow:      false,
	}
	return
}
