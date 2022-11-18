package tag

import (
	"easyweb/global"
	"easyweb/pkg/response"
	tagSer "easyweb/service/tag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func GetInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	info, err := tagSer.GetTagInfo(c, uint(id))
	if err != nil {
		global.Log.Debug("获取单个标签数据失败:", zap.Error(err))
		response.FailWithMessage(response.ErrDBNotFound, "获取数据失败", c)
		return
	}

	response.OkWithData(info, c)
}

func GetList(c *gin.Context) {
	tname := c.Query("name")
	var err error
	var list []tagSer.ResponseTagInfo
	if tname != "" {
		list, err = tagSer.SearchByName(c, tname)
	} else {
		err, list = tagSer.GetList(c)
	}
	if err != nil {
		response.Fail(response.ErrDBNotFound, c)
		return
	}
	response.OkWithData(list, c)
}
