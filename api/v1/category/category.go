package category

import (
	"easyweb/global"
	"easyweb/pkg/response"
	categorySer "easyweb/service/category"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetInfo(c *gin.Context) {
	idStr := c.Param("id")
	global.Log.Info(idStr)
	if idStr == "" {
		global.Log.Debug("category-getinfo:参数id为空")
		response.Fail(response.ErrCategory, c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Warn("category-getinfo:转换id失败", zap.Error(err))
		response.Fail(response.ErrCategory, c)
		return
	}
	info, err := categorySer.GetInfo(c, uint(id))
	if err != nil {
		global.Log.Warn("category-getinfo:获取分类数据失败", zap.Error(err), zap.Int("id", id))
		response.FailWithMessage(response.ErrCategory, err.Error(), c)
		return
	}
	response.OkWithData(info, c)
}

func GetList(c *gin.Context) {
	name := c.Query("name")
	var err error
	var list []categorySer.ResponseCategoryInfo
	if name != "" {
		list, err = categorySer.SearchByName(c, name)
	} else {
		list, err = categorySer.GetList(c)
	}
	if err != nil {
		response.Fail(response.ErrDBNotFound, c)
		return
	}
	response.OkWithData(list, c)
}
