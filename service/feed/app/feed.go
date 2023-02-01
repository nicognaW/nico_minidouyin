package app

import (
	"context"
	"fmt"
	"log"

	"nico_minidouyin/config"
	pb "nico_minidouyin/gen/douyin/feed"
	"nico_minidouyin/repo"
	feedModel "nico_minidouyin/service/feed/model"
	userModel "nico_minidouyin/service/user/model"
)

type FeedService struct {
	pb.UnimplementedFeedServiceServer
}

func (f FeedService) GetFeed(ctx context.Context, request *pb.FeedRequest) (feed *pb.FeedResponse, err error) {
	uid := request.GetToken()
	feed = &pb.FeedResponse{}

	/**
	 * TODO: These authentication code should be placed in auth service.
	 */
	var dbUsr userModel.User
	if err := repo.DB.First(&dbUsr, uid).Error; err != nil {
		return nil, fmt.Errorf("user not found: %s", uid)
	}

	log.Printf("getting feed for user %s", uid)

	dbVidList := make([]*feedModel.Video, config.FeedVideoCount)
	repo.DB.Limit(config.FeedVideoCount).Order("created_at desc").Find(&dbVidList)

	domainVidList := make([]*pb.Video, len(dbVidList))
	feed.VideoList = domainVidList
	for i, video := range dbVidList {
		var dbAuthor userModel.User
		if err := repo.DB.First(&dbAuthor, video.AuthorId).Error; err != nil {
			return nil, fmt.Errorf("user not found: %s", uid)
		}
		domainVidList[i] = &pb.Video{
			Id: video.ID,
			Author: &pb.User{
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

	return

}
