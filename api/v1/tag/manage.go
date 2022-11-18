package tag

import (
	"easyweb/global"
	"easyweb/model"
	"easyweb/pkg/response"
	"easyweb/pkg/validator"
	tagSer "easyweb/service/tag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Create(c *gin.Context) {
	var newTag tagSer.RequestTag
	if err := c.ShouldBind(&newTag); err != nil {
		response.FailWithMessage(response.ErrAuthFail, validator.GetErrorMsg(newTag, err), c)
		return
	}
	if err := tagSer.Create(c, newTag.Name); err != nil {
		response.FailWithMessage(response.ErrTag, err.Error(), c)
		return
	}
	response.Ok(c)
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		response.FailWithMessage(response.ErrTag, "参数错误", c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Warn("标签删除方法,id转换失败", zap.Error(err))
		response.FailWithMessage(response.ErrTag, "删除失败", c)
		return
	}
	if err := tagSer.Delete(c, uint(id)); err != nil {
		global.Log.Warn("删除标签失败", zap.Error(err))
		response.FailWithMessage(response.ErrTag, "删除失败", c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var tagInfo tagSer.RequestTag
	idStr := c.Param("id")
	if err := c.ShouldBind(&tagInfo); err != nil && idStr != "" {
		global.Log.Debug("tag-Update 接收参数错误", zap.Error(err))
		response.FailWithMessage(response.ErrAuthFail, "参数不能为空", c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Warn("tag-Update,转换id参数类型失败", zap.Error(err))
		response.FailWithMessage(response.ErrTag, "修改失败", c)
		return
	}
	var tagUpdate model.Tag
	tagUpdate.Id = uint(id)
	tagUpdate.Tname = tagInfo.Name
	if err := tagSer.Update(c, &tagUpdate); err != nil {
		global.Log.Warn("tag-Update,修改失败", zap.Error(err))
		response.FailWithMessage(response.ErrTag, "修改失败", c)
		return
	}

	response.Ok(c)
}
