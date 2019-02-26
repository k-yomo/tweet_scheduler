package models

type TweetLog struct {
	BaseModel
	TweetBody string `json:"tweet_body" gorm:"not null"`
}
