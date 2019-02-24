package lib

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/k-yomo/tweet_scheduler/db"
	"github.com/k-yomo/tweet_scheduler/models"
	"github.com/labstack/gommon/log"
	"math/rand"
	"os"
	"time"
)

func NewTwitterApi() *anaconda.TwitterApi {
	apiKey := os.Getenv("API_KEY")
	apiSecretKey := os.Getenv("API_SECRET_KEY")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, apiKey, apiSecretKey)
	return api
}

// Pick one tweet from registered tweets randomly
func TweetRandomly() {
	database, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	api := NewTwitterApi()
	tweets := models.PrepareWeightedTweets(database)
	if len(tweets) == 0 {
		return
	}
	if len(tweets) == 1 {

	}
	rand.Seed(time.Now().UnixNano())
	tweetIndex := rand.Intn(len(tweets))
	_, err = api.PostTweet(tweets[tweetIndex].Body, nil)
	if err != nil {
		log.Error(err)
	}
}

