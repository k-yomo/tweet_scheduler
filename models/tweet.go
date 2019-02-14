package models

type Tweet struct {
	BaseModel
	Body      string `json:"body" gorm:"not null"`
	Weight    int `json:"weight" gorm:"default:1"`
}
