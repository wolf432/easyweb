package model

type TagRelation struct {
	Id  uint `gorm:"column:id;primarykey"`
	Tid uint `gorm:"column:tid"`
	Lid uint `gorm:"column:lid"`
}

func (r *TagRelation) TableName() string {
	return "tag_relation"
}
