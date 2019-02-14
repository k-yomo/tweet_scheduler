package models

type User struct {
	BaseModel
	Email     string   `json:"email" gorm:"unique;not null"`
	Password  string   `json:"password" gorm:"-"`
	PasswordHash  []byte   `json:"-"`
	Token     string   `json:"token,omitempty" gorm:"-"`
}

