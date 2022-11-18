package link

import (
	"easyweb/global"
	"easyweb/pkg/response"
	"easyweb/pkg/validator"
	linkSer "easyweb/service/link"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Create(c *gin.Context) {
	var newLink linkSer.RequestLinkAdd
	if err := c.ShouldBind(&newLink); err != nil {
		response.FailWithMessage(response.ErrCategory, validator.GetErrorMsg(newLink, err), c)
		return
	}
	if err := linkSer.Create(c, newLink); err != nil {
		response.Fail(response.ErrLink, c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.Fail(response.ErrLink, c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Warn("link-Update,转换id参数失败", zap.Error(err))
		response.Fail(response.ErrLink, c)
		return
	}
	if err := linkSer.Delete(c, uint(id)); err != nil {
		response.FailWithMessage(response.ErrLink, err.Error(), c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var linkUpdate linkSer.RequestLinkUpdate
	if err := c.ShouldBind(&linkUpdate); err != nil {
		response.FailWithMessage(response.ErrLink, validator.GetErrorMsg(linkUpdate, err), c)
		return
	}
	if err := linkSer.Update(c, linkUpdate); err != nil {
		response.Fail(response.ErrLink, c)
		return
	}
	response.Ok(c)
}
