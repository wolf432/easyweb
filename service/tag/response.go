package tag

type ResponseTagInfo struct {
	Id    uint   `json:"id" gorm:"column:id"`
	Tname string `json:"tname" gorm:"column:tname"`
}
