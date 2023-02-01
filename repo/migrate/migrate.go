package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nico_minidouyin/config"
	feedModel "nico_minidouyin/service/feed/model"
	userModel "nico_minidouyin/service/user/model"
)

func main() {
	var err error
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: config.DSN,
			}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		panic(fmt.Errorf("db connection failed: %v", err))
	}
	err = db.AutoMigrate(&userModel.UserToken{}, &userModel.User{}, &feedModel.Video{})
	if err != nil {
		panic(fmt.Errorf("db migrate failed: %v", err))
	}
}
