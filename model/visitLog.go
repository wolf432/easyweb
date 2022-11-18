package model

import "time"

type visitLog struct {
	Id        uint      `gorm:"column:id";primarykey`
	Lid       uint      `gorm:"column:lid";not null`
	Remark    string    `gorm:"column:remark"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (l *visitLog) TableName() string {
	return "visit_log"
}
