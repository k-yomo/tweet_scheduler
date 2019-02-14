package db

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/k-yomo/tweet_scheduler/models"
)

var config = struct {
	DBName   string `default:"tweet_scheduler_development"`
	User     string `default:"postgres"`
	Host     string `default:"localhost"`
	Password string `default:"password" env:"DBPassword"`
	Port     string `default:"5433"`
}{}

func New() (db *gorm.DB) {

	_ = configor.Load(&config)

	args := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Host,
		config.Port,
		config.User,
		config.DBName,
		config.Password)

	db, err := gorm.Open("postgres", args)

	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	autoMigrate(db)

	return
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Tweet{})
}
