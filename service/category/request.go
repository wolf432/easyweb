package category

import "easyweb/pkg/validator"

// RequestCategoryAdd 添加分类
type RequestCategoryAdd struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (c RequestCategoryAdd) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"name.required": "分类名不能为空",
	}
}

// RequestCategoryUpdate 修改分类
type RequestCategoryUpdate struct {
	Id   uint   `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
}

func (c RequestCategoryUpdate) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"name.required": "分类名不能为空",
		"id.required":   "分类id不能为空",
	}
}
