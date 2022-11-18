package category

import (
	"easyweb/global"
	"easyweb/pkg/response"
	"easyweb/pkg/validator"
	categorySer "easyweb/service/category"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func Create(c *gin.Context) {
	var newCategory categorySer.RequestCategoryAdd
	if err := c.ShouldBind(&newCategory); err != nil {
		response.FailWithMessage(response.ErrCategory, validator.GetErrorMsg(newCategory, err), c)
		return
	}
	if err := categorySer.Create(c, newCategory.Name); err != nil {
		response.FailWithMessage(response.ErrCategory, err.Error(), c)
		return
	}

	response.Ok(c)

}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		global.Log.Debug("category-api-delete,id参数为空")
		response.FailWithMessage(response.ErrCategory, "参数不能为空", c)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		global.Log.Debug("category-api-delete,id参数转换格式失败", zap.Error(err))
		response.Fail(response.ErrCategory, c)
		return
	}
	if err := categorySer.Delete(c, uint(id)); err != nil {
		response.FailWithMessage(response.ErrCategory, "分类不存在", c)
		return
	}
	response.Ok(c)
}

func Update(c *gin.Context) {
	var newCategory categorySer.RequestCategoryUpdate
	if err := c.ShouldBind(&newCategory); err != nil {
		response.FailWithMessage(response.ErrCategory, validator.GetErrorMsg(newCategory, err), c)
		return
	}
	if err := categorySer.Update(c, newCategory); err != nil {
		response.FailWithMessage(response.ErrCategory, err.Error(), c)
		return
	}

	response.Ok(c)
}
