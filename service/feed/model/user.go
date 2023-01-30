package model

import (
	"nico_minidouyin/repo"
)

type User struct {
	repo.Model
	FollowCount   uint32 `gorm:"default:0"`       // 关注总数
	FollowerCount uint32 `gorm:"default:0"`       // 粉丝总数
	Name          string `gorm:"not null;unique"` // 用户名称
}
