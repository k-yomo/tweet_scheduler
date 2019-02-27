package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/k-yomo/tweet_scheduler/db"
	"github.com/k-yomo/tweet_scheduler/handler"
	"github.com/k-yomo/tweet_scheduler/lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
)

const (
	AuthApiRoot = "/auth/api/v1"
	ApiRoot = "/api/v1"
)

func main()  {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error while loading .env file")
		}
	}

	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(lib.JwtSigningKey),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for a request to root or login path
			if c.Path() == "/" || c.Path() == AuthApiRoot + "/login" {
				return true
			}
			return false
		},
	}))

	database, err := db.New()
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler.Handler{DB: database}

	// Routes
	// Not to let heroku instance sleep, we use monitor service(http://www.uptimerobot.com/)
	// And it require the app to return 200, so we use root path for it.
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "")
	})
	// Auth
	e.POST(AuthApiRoot + "/sign_up", h.Signup)
	e.POST(AuthApiRoot + "/login", h.Login)
	// Tweet CRUD
	e.GET(ApiRoot + "/tweets", h.GetTweets)
	e.POST(ApiRoot + "/tweets", h.CreateTweet)
	e.PUT(ApiRoot + "/tweets/:id", h.UpdateTweet)
	e.DELETE(ApiRoot + "/tweets/:id", h.DeleteTweet)
	// Tweet Logs
	e.GET(ApiRoot + "/tweet_logs", h.GetTweetLogs)

	// Tweet once every 3 hours
	gocron.Start()
	gocron.Every(3).Hours().Do(lib.TweetRandomly)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}