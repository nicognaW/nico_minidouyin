package model

import (
	"nico_minidouyin/repo/model"
)

type Video struct {
	model.Model
	AuthorId      uint32
	PlayUrl       string // 视频播放地址
	CoverUrl      string // 视频封面地址
	FavoriteCount uint32 // 视频的点赞总数
	CommentCount  uint32 // 视频的评论总数
	Title         string // 视频标题
}
