package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/k-yomo/tweet_scheduler/handler"
	"github.com/k-yomo/tweet_scheduler/helper"
	"github.com/k-yomo/tweet_scheduler/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main()  {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwt_generator.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	db, err := gorm.Open("postgres", "host=localhost dbname=tweet_scheduler_development user=postgres sslmode=disable")
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Tweet{})

	h := &handler.Handler{DB: db}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	// Tweet CRUD
	e.GET("/tweets", h.GetTweets)
	e.POST("/tweets", h.CreateTweet)
	e.PUT("/tweets/:id", h.UpdateTweet)
	e.DELETE("/tweets/:id", h.DeleteTweet)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}