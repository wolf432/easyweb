package visitLog

import "easyweb/pkg/validator"

type RequestVisitLogAdd struct {
	Lid    uint   `json:"lid" form:"lid" binding:"required"`
	Remark string `json:"remark" form:"remark" binding:"min=2"`
}

func (v RequestVisitLogAdd) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"Lid.required": "访问的链接不能为空",
		"Remark.min":   "标注长度不能小于2个",
	}
}
