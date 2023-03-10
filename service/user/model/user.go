package model

import (
	"nico_minidouyin/repo/model"
)

// User 用户表 /*
type User struct {
	model.Model           // 基础模型
	Username      string  `gorm:"not null;unique;size: 32"` // 用户名
	Password      *string `gorm:"not null;size: 32"`        // 密码
	FollowCount   uint32  `gorm:"default:0"`                // 关注总数
	FollowerCount uint32  `gorm:"default:0"`                // 粉丝总数
	Name          string  `gorm:"not null;unique"`          // 用户名称
	Role          string  `gorm:"default:0"`                // 用户角色
}
