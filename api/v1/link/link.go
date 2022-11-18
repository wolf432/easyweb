package link

import (
	"easyweb/global"
	"easyweb/pkg/response"
	linkSer "easyweb/service/link"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetInfo(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(response.ErrLink, c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Error("link-GetInfo,转换id失败", zap.Error(err))
		response.Fail(response.ErrLink, c)
		return
	}
	info, err := linkSer.GetInfo(c, uint(id))
	if err != nil {
		global.Log.Error("link-GetInfo", zap.Error(err))
		response.Fail(response.ErrLink, c)
		return
	}
	response.OkWithData(info, c)
}

func GetList(c *gin.Context) {
	title := c.Query("title")
	cidStr := c.Query("cid")
	cname := c.Query("cname")
	tag := c.Query("tag")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	var condition linkSer.ConditionLink
	if title != "" {
		condition.Title = title
	}
	if cidStr != "" {
		cid, err := strconv.Atoi(cidStr)
		if err != nil {
			response.Fail(response.ErrLink, c)
			return
		}
		condition.Cid = uint(cid)
	}
	if tag != "" {
		condition.TagName = tag
	}
	if cname != "" {
		condition.Cname = cname
	}

	list, err := linkSer.GetConditionByPage(c, page, limit, condition)

	if err != nil {
		response.Fail(response.ErrDBNotFound, c)
		return
	}
	response.OkWithData(list, c)
}
