package category

import (
	"easyweb/global"
	"easyweb/model"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ExistCategoryByName(c *gin.Context, name string) bool {
	if err := global.DB.First(&model.Category{}, "cname=?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func ExistCategoryById(c *gin.Context, id uint) bool {
	if err := global.DB.First(&model.Category{}, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func Create(c *gin.Context, name string) error {
	var newCategory model.Category
	if ExistCategoryByName(c, name) {
		return errors.New("分类已存在,不能重复添加")
	}
	newCategory.Cname = name
	if err := global.DB.Create(&newCategory).Error; err != nil {
		global.Log.Error("service-category-create", zap.Error(err))
		return errors.New("添加失败")
	}
	return nil
}

func Update(c *gin.Context, category RequestCategoryUpdate) error {
	if !ExistCategoryById(c, category.Id) {
		return errors.New("分类不存在")
	}
	if err := global.DB.Model(&model.Category{}).Where("id=?", category.Id).Update("cname", category.Name).Error; err != nil {
		global.Log.Error("service-category-update", zap.Error(err))
		return errors.New("修改失败")
	}
	return nil
}

func Delete(c *gin.Context, id uint) error {
	if !ExistCategoryById(c, id) {
		return errors.New("分类不存在")
	}
	if err := global.DB.Where("id=?", id).Delete(&model.Category{}).Error; err != nil {
		return errors.New("删除失败")
	}
	return nil
}

// Increment 数量自增
func Increment(c *gin.Context, id uint) error {
	return global.DB.Model(&model.Category{}).Where("id=?", id).Update("count", gorm.Expr("count+1")).Error
}

// DeCrement 数量自减
func DeCrement(c *gin.Context, id uint) error {
	return global.DB.Model(&model.Category{}).Where("id=?", id).Update("count", gorm.Expr("count-1")).Error
}

func GetList(c *gin.Context) ([]ResponseCategoryInfo, error) {
	var list []ResponseCategoryInfo
	err := global.DB.Model(&model.Category{}).Order("count desc,id desc").Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func GetInfo(c *gin.Context, id uint) (ResponseCategoryInfo, error) {
	var info ResponseCategoryInfo
	if err := global.DB.Model(&model.Category{}).First(&info, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return info, errors.New("分类不存在")
		}
		global.Log.Warn("serviceCategory-GetInfo,获取数据失败", zap.Error(err))
		return info, errors.New("获取失败")
	}
	return info, nil
}

func SearchByName(c *gin.Context, name string) ([]ResponseCategoryInfo, error) {
	var list []ResponseCategoryInfo
	err := global.DB.Model(&model.Category{}).Where("cname like ?", "%"+name+"%").Order("count desc,id desc").Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}
