package app

import (
	"context"
	"fmt"
	"nico_minidouyin/config"
	"nico_minidouyin/service/feed/biz/model/douyin/db"
	"nico_minidouyin/service/feed/biz/model/douyin/domain"
	"nico_minidouyin/service/feed/biz/repo"
)

func QueryFeed(ctx context.Context, uid *uint32) (feed *domain.Feed, err error) {
	feed = &domain.Feed{}

	/**
	 * TODO: These authentication code should be placed in auth service.
	 */
	var dbUsr db.User
	authenticated := uid != nil
	if authenticated {
		if err := repo.DB.First(&dbUsr, *uid).Error; err != nil {
			return nil, fmt.Errorf("user not found: %p", uid)
		}
		feed.User = domain.DomainizeUser(&dbUsr)
	}

	dbVidList := make([]*db.Video, config.FeedVideoCount)
	repo.DB.Limit(config.FeedVideoCount).Order("created_at desc").Find(&dbVidList)

	domainVidList := make([]*domain.Video, len(dbVidList))
	feed.VideoList = domainVidList
	for i, video := range dbVidList {
		var dbAuthor db.User
		if err := repo.DB.First(&dbAuthor, video.AuthorId).Error; err != nil {
			return nil, fmt.Errorf("user not found: %p", uid)
		}
		domainVidList[i] = domain.DomainizeVideo(video, &dbAuthor)
	}

	return
}
