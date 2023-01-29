package main

import (
	"context"
	"fmt"
	"nico_minidouyin/config"
	"nico_minidouyin/kitex_gen/douyin"
	"nico_minidouyin/service/feed/repo"
	dbModel "nico_minidouyin/service/feed/repo/model"
	"strconv"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *douyin.FeedRequest) (resp *douyin.FeedResponse, err error) {
	resp = new(douyin.FeedResponse)
	uid := req.GetToken()
	var msg string
	parseInt, err := strconv.ParseInt(uid, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("unknown token: %s", req.GetToken())
	}
	uint32Value := uint32(parseInt)

	//feed, err := app.QueryFeed(ctx, &uint32Value)

	/**
	 * TODO: These authentication code should be placed in auth service.
	 */
	var dbUsr dbModel.User
	if err := repo.DB.First(&dbUsr, uint32Value).Error; err != nil {
		return nil, fmt.Errorf("user not found: %p", uid)
	}
	dbVidList := make([]*dbModel.Video, config.FeedVideoCount)
	repo.DB.Limit(config.FeedVideoCount).Order("created_at desc").Find(&dbVidList)

	douyinVidList := make([]*douyin.Video, len(dbVidList))
	resp.VideoList = douyinVidList
	for i, video := range dbVidList {
		var dbAuthor dbModel.User
		if err := repo.DB.First(&dbAuthor, video.AuthorId).Error; err != nil {
			return nil, fmt.Errorf("user not found: %p", uid)
		}
		douyinVidList[i] = &douyin.Video{
			Id: video.ID,
			Author: &douyin.User{
				Id:            dbAuthor.ID,
				Name:          dbAuthor.Name,
				FollowCount:   dbAuthor.FollowCount,
				FollowerCount: dbAuthor.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		}
	}

	if err != nil {
		return nil, err
	}
	resp.StatusMsg = &msg

	return
}
