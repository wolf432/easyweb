package link

import "easyweb/pkg/validator"

type RequestLinkAdd struct {
	Id     uint   `json:"id" form:"id"`
	Title  string `json:"title" form:"title" binding:"required,min=2"`
	Cid    uint   `json:"cid" form:"cid" binding:"required"`
	Url    string `json:"url" form:"url" binding:"required,url"`
	Remark string `json:"remark" form:"remark" binding:"max=500"`
	Zank   uint   `json:"zank" form:"zank"`
	TagIds string `json:"tag_ids" form:"tag_ids"`
}

func (l RequestLinkAdd) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"title.required": "标题不能为空",
		"title.min":      "标题长度不能小于2",
		"cid.required":   "请选择一个分类",
		"url.required":   "链接不能空",
		"url.url":        "请填写正确格式的链接",
		"remark.Max":     "备注长度不能超过500",
	}
}

type RequestLinkUpdate struct {
	Id     uint   `json:"id" form:"id" binding:"required"`
	Title  string `json:"title" form:"title" binding:"required,min=2"`
	Cid    uint   `json:"cid" form:"cid" binding:"required"`
	Url    string `json:"url" form:"url" binding:"required,url"`
	Remark string `json:"remark" form:"remark" binding:"max=500"`
	Zank   uint   `json:"zank" form:"zank"`
	TagIds string `json:"tag_ids" form:"tag_ids"`
}

func (l RequestLinkUpdate) GetMessages() validator.ValidatorMessageMap {
	return validator.ValidatorMessageMap{
		"id.required":    "链接id不能为空 ",
		"title.required": "标题不能为空",
		"title.min":      "标题长度不能小于2",
		"cid.required":   "请选择一个分类",
		"url.required":   "链接不能空",
		"url.url":        "请填写正确格式的链接",
		"remark.Max":     "备注长度不能超过500",
	}
}
