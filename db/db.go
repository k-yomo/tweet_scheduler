package db

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/k-yomo/tweet_scheduler/models"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var config = struct {
	DBName   string `default:"tweet_scheduler_development" env:"DB_NAME"`
	User     string `default:"postgres" env:"DB_USER"`
	Host     string `default:"db" env:"DB_HOST"`
	Password string `default:"postgres" env:"DB_PASSWORD"`
	Port     string `default:"5432" env:"DB_PORT"`
}{}

func New() (*gorm.DB, error) {

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
		return nil, err
	}

	autoMigrate(db)

	env := os.Getenv("ENV")
	if env != "production" {
		seedData(db)
		db.LogMode(true)
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Tweet{})
}

func seedData(db *gorm.DB) {
	u := &models.User{
		Email: "test@example.com",
		Password: "password",
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.PasswordHash = passwordHash
	db.Create(u)
}