package rpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"nico_minidouyin/gen/proto/gen/feed"
	pb "nico_minidouyin/gen/proto/gen/feed"
	feedApp "nico_minidouyin/service/feed/biz/app"
)

type FeedServer struct {
	pb.UnimplementedFeedServiceServer
}

func NewServer() *FeedServer {
	s := &FeedServer{}
	return s
}

func (f FeedServer) GetFeed(ctx context.Context, request *feed.FeedRequest) (*feed.FeedResponse, error) {

	uid, err := strconv.ParseUint(*request.Token, 10, 32)
	uidConverted := uint32(uid)
	if err != nil {
		return nil, errors.New("user doesn't exist")
	}

	resp := new(feed.FeedResponse)
	queryFeed, err := feedApp.QueryFeed(ctx, &uidConverted)
	if err != nil {
		return nil, fmt.Errorf("error happend: %v", err)
	}
	for i, video := range queryFeed.VideoList {
		resp.VideoList[i] = &feed.Video{
			Id: video.Id,
			Author: &feed.User{
				Id:            video.Author.Id,
				Name:          video.Author.Name,
				FollowCount:   video.Author.FollowerCount,
				FollowerCount: video.Author.FollowCount,
				IsFollow:      video.Author.IsFollow,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
	}
	msg := "success"
	resp.StatusMsg = &msg
	resp.StatusCode = 0

	return resp, nil
}
