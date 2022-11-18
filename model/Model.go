package model

import "time"

type BaseModel struct {
	Id        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
