package category

type ResponseCategoryInfo struct {
	Id    uint   `json:"id" gorm:"column:id"`
	Cname string `json:"cname" gorm:"column:cname"`
}
