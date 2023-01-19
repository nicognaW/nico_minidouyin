package domain

import "nico_minidouyin/service/feed/biz/model/douyin/db"

type User struct {
	Id            uint32
	Name          string
	FollowCount   uint32
	FollowerCount uint32
	IsFollow      bool
}

type Author User

type Video struct {
	Id            uint32
	Author        *Author
	PlayUrl       string
	CoverUrl      string
	FavoriteCount uint32
	CommentCount  uint32
	IsFavorite    bool
	Title         string
}

type Feed struct {
	VideoList []*Video
	User      User
}

func DomainizeUser(dbUsr *db.User) User {
	return User{
		Id:            dbUsr.ID,
		Name:          dbUsr.Name,
		FollowCount:   dbUsr.FollowCount,
		FollowerCount: dbUsr.FollowerCount,
		IsFollow:      false,
	}
}

func DomainizeAuthor(dbAuthor *db.User) *Author {
	return (*Author)(&User{
		Id:            dbAuthor.ID,
		Name:          dbAuthor.Name,
		FollowCount:   dbAuthor.FollowCount,
		FollowerCount: dbAuthor.FollowerCount,
		IsFollow:      false,
	})
}

func DomainizeVideo(video *db.Video, dbAuthor *db.User) *Video {
	return &Video{
		Id:            video.ID,
		Author:        DomainizeAuthor(dbAuthor),
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    false,
		Title:         video.Title,
	}
}
