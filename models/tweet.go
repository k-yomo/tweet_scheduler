package models

import "github.com/jinzhu/gorm"

type Tweet struct {
	BaseModel
	Body      string `json:"body" gorm:"not null"`
	// Weight is to calculate how often this tweet should be tweeted
	// if weight is set to 3, possibility will be 3/(number of all weighted tweets)
	Weight    int `json:"weight" gorm:"default:1"`
}

func PrepareWeightedTweets(db *gorm.DB) []Tweet {
	tweets := make([]Tweet, 0)
	db.Find(&tweets)
	if len(tweets) == 0 {
		return tweets
	}

	weightedTweets := make([]Tweet, 0)
	for _, v := range tweets {
		for i := 0; i < v.Weight; i++ {
			weightedTweets = append(weightedTweets, v)
		}
	}
	return weightedTweets
}
