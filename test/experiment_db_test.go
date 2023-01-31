package test

import (
	"fmt"
	"log"
	"nico_minidouyin/config"
	feedModel "nico_minidouyin/service/feed/model"
	userModel "nico_minidouyin/service/user/model"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {
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
	err = db.AutoMigrate(&userModel.User{}, &feedModel.Video{})
	assert.Nil(t, err)
}

func TestDB(t *testing.T) {
	/**
	1. create db connection
	*/
	var err error
	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  config.DSN,
				PreferSimpleProtocol: true, // disables implicit prepared statement usage
			}), &gorm.Config{
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(fmt.Errorf("db connection failed: %v", err))
	}

	/**
	2. create default user as author
	*/
	password := "password"
	author := &userModel.User{
		Username:      "nyanki",
		Password:      &password,
		FollowCount:   0,
		FollowerCount: 0,
		Name:          "Nyanki",
	}
	if err := db.Create(&author).Error; err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			log.Printf("dup user, already created: %s(%d)", author.Name, author.ID)
		} else {
			panic(fmt.Errorf("author user creation failed: %v", err))
		}
	}

	/**
	3. create video with author
	*/
	video := &feedModel.Video{
		AuthorId:      author.ID,
		PlayUrl:       "https://internal-api-drive-stream.feishu.cn/space/api/box/stream/download/all/boxcneWtGYeCxvnZq4ZuSyZMxit",
		CoverUrl:      "https://internal-api-drive-stream.feishu.cn/space/api/box/stream/download/v2/cover/boxcnHd0AC6LDN4A4QBVp16q9q1?width=480&height=270&policy=near&fall_back=1",
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         "Real World Reinforcement Learning by Microsoft",
	}
	if err := db.Create(&video).Error; err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == pgerrcode.UniqueViolation {
			log.Printf("dup video, already created: %s(%d)", video.Title, video.ID)
		} else {
			panic(fmt.Errorf("video creation failed: %v", err))
		}
	}

	/**
	4. validate the data
	*/
	func() {
		var authorFound userModel.User
		if firstResult := db.Where("id = 1").First(&authorFound); firstResult.Error == gorm.ErrRecordNotFound {
			panic(fmt.Errorf("default UID0 user not found: %v", firstResult.Error))
		}
		assert.DeepEqual(t, authorFound.Name, author.Name)

		videoFound := make([]*feedModel.Video, 0)
		firstResult := db.Where("author_id = 1").Find(&videoFound)
		log.Printf("found %d rows: %v", firstResult.RowsAffected, videoFound)

		assert.True(t, firstResult.RowsAffected > 0)
	}()
}
