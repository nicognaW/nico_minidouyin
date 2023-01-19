package test

import (
	"fmt"
	"log"
	"nico_minidouyin/config"
	dbModels "nico_minidouyin/service/feed/biz/model/douyin/db"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
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
	err = db.AutoMigrate(&dbModels.User{}, &dbModels.Video{})

	if err != nil {
		panic(fmt.Errorf("db migration failed: %v", err))
	}

	author := &dbModels.User{
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

	video := &dbModels.Video{
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

	func() {
		var authorFound dbModels.User
		if firstResult := db.Where("id = 1").First(&authorFound); firstResult.Error == gorm.ErrRecordNotFound {
			panic(fmt.Errorf("default UID0 user not found: %v", firstResult.Error))
		}
		assert.DeepEqual(t, authorFound.Name, author.Name)

		videoFound := make([]*dbModels.Video, 0)
		firstResult := db.Where("author_id = 1").Find(&videoFound)
		log.Printf("found %d rows: %v", firstResult.RowsAffected, videoFound)

		assert.True(t, firstResult.RowsAffected > 0)
	}()
}
