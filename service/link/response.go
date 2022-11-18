package link

import (
	categorySer "easyweb/service/category"
	tagSer "easyweb/service/tag"
	"time"
)

//ResponseInfo 链接列表
type ResponseInfo struct {
	Id        uint                             `json:"lid" gorm:"column:id"`
	Title     string                           `json:"title" gorm:"column:title"`
	Url       string                           `json:"url" gorm:"column:url"`
	Remark    string                           `json:"remark" gorm:"column:remark"`
	Click     uint                             `json:"click" gorm:"column:click"`
	Zank      uint                             `json:"zand" gorm:"column:zank"`
	CreatedAt time.Time                        `json:"created_at" gorm:"column:created_at"`
	Cid       uint                             `json:"cid" gorm:"column:cid"`
	Tag       []tagSer.ResponseTagInfo         `json:"tag" gorm:"-"`
	Category  categorySer.ResponseCategoryInfo `json:"category" gorm:"-"`
}

type PageLink struct {
	Data   []ResponseInfo `json:"data"`
	Amount int            `json:"amount"`
}

// ConditionLink 搜索条件
type ConditionLink struct {
	Title   string
	Cid     uint
	Cname   string
	TagName string
}
