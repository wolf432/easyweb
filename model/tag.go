package model

type Tag struct {
	BaseModel
	Tname string `gorm:"column:tname;not null"`
	Count uint   `gorm:"column:count;"`
}

func (t *Tag) TableName() string {
	return "tag"
}
