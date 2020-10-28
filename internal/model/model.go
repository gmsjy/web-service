package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gmsjy/web-service/global"
	"github.com/gmsjy/web-service/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	*gorm.Model
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	logLevel := logger.Silent

	if global.ServerSetting.RunMode == "debug" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.DBName,
		databaseSetting.Host,
	)
	//dsn = "user=postgres password=xlwebsql1234 dbname=xldata port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: newLogger})

	if err != nil {
		return nil, err
	}

	if !db.Migrator().HasTable(&Article{}) {
		db.Migrator().CreateTable(&Article{})
	} else {
		db.AutoMigrate(&Article{})
	}

	return db, nil

}
