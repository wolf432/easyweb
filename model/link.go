package model

type Link struct {
	BaseModel
	Title  string `gorm:"column:title;not null;"` //标题
	Cid    uint   `gorm:"column:cid;not null;"`   //分类表id
	Url    string `gorm:"column:url;not null;"`   //
	Remark string `gorm:"column:remark;"`
	Click  uint   `gorm:"column:click;default:0;"`
	Zank   uint   `gorm:"column:zank;default:1"`
}

func (l *Link) TableName() string {
	return "link"
}
