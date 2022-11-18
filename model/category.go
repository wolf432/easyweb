package model

type Category struct {
	BaseModel
	Cname string `gorm:"column:cname;not null"`
	Count uint   `gorm:"column:count"`
}

func (c *Category) TableName() string {
	return "category"
}
