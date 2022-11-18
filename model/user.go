package model

import (
	"strconv"
)

type User struct {
	BaseModel
	UserName string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
	Email    string `gorm:"column:email"`
}

func (u *User) TableName() string {
	return "user"
}

func (u User) GetUid() string {
	return strconv.Itoa(int(u.Id))
}
