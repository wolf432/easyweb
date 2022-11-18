package tag

import (
	"easyweb/global"
	"easyweb/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

//Create 添加标签
func Create(c *gin.Context, tName string) error {
	var newTag model.Tag

	if ExistTagByName(c, tName) {
		return errors.New("标签已存在")
	}
	newTag.Tname = tName
	newTag.Count = 0
	if err := global.DB.Create(&newTag).Error; err != nil {
		return errors.New("添加失败")
	}
	return nil
}

// ExistTagById 根据标签id判断数据是否存在
func ExistTagById(c *gin.Context, id uint) bool {
	if err := global.DB.First(&model.Tag{}, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// ExistTagByName 根据标签id判断数据是否存在
func ExistTagByName(c *gin.Context, name string) bool {
	if err := global.DB.First(&model.Tag{}, "tname=?", name).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func Delete(c *gin.Context, tid uint) error {
	if err := global.DB.Where("id=?", tid).Delete(&model.Tag{}).Error; err != nil {
		return errors.New("删除失败")
	}
	return nil
}

func Update(c *gin.Context, tagData *model.Tag) error {
	if err := ExistTagById(c, tagData.Id); !err {
		return errors.New("标签不存在")
	}
	if err := global.DB.Model(&model.Tag{}).Where("id", tagData.Id).Update("tname", tagData.Tname).Error; err != nil {
		return errors.New("修改失败")
	}
	return nil
}

func GetTagInfo(c *gin.Context, tid uint) (ResponseTagInfo, error) {
	var info ResponseTagInfo
	var tag model.Tag
	err := global.DB.Table(tag.TableName()).First(&info, tid).Error
	return info, err
}

func GetTagInfoByName(c *gin.Context, name string) (ResponseTagInfo, error) {
	var info ResponseTagInfo
	var tag model.Tag
	err := global.DB.Table(tag.TableName()).Where("tname=?", name).First(&info).Error
	return info, err
}

// GetList 获取全部的标签列表
func GetList(c *gin.Context) (error, []ResponseTagInfo) {
	var list []ResponseTagInfo
	var tag model.Tag
	err := global.DB.Table(tag.TableName()).Order("count desc").Find(&list).Error
	if err != nil {
		return err, list
	}
	return nil, list
}

func SearchByName(c *gin.Context, name string) ([]ResponseTagInfo, error) {
	var list []ResponseTagInfo
	var tag model.Tag
	err := global.DB.Table(tag.TableName()).Where("tname like ?", "%"+name+"%").Order("count desc").Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func SearchMultiIdByName(c *gin.Context, name string) ([]uint, error) {
	nameArr := strings.Split(name, ",")

	//如果搜索的是单个标签直接查询后返回
	if len(nameArr) == 0 {
		if info, err := GetTagInfoByName(c, name); err == nil {
			return []uint{info.Id}, nil
		}
		return []uint{}, gorm.ErrRecordNotFound
	}

	ids := make([]uint, 0, len(nameArr))
	err := global.DB.Model(model.Tag{}).Select("id").Where("tname in(?)", nameArr).Find(&ids).Error
	if err != nil {
		return ids, err
	}
	return ids, nil
}
