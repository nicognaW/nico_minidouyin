package repo

import (
	genDB "nico_minidouyin/gen/db"
	"time"

	"nico_minidouyin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

func init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(
		postgres.Open(config.DSN),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	genDB.SetDefault(DB)
	if err != nil {
		panic(err)
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}
