package tag

import (
	"easyweb/global"
	"easyweb/model"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UpdateRealtionTag 批量更新链接与标签关联数据
func UpdateRealtionTag(c *gin.Context, lid uint, tid []uint) error {
	//删除所有关联lid的标签数据
	if err := DeleteRealtionTagByLinkId(c, lid); err != nil {
		return errors.New("删除失败")
	}
	//重新添加关联
	var list []model.TagRelation
	for _, id := range tid {
		list = append(list, model.TagRelation{
			Lid: lid,
			Tid: id,
		})
	}
	if err := global.DB.Create(&list).Error; err != nil {
		return err
	}

	return nil
}

func GetAllByLinkId(c *gin.Context, linkId uint) ([]ResponseTagInfo, error) {
	var list []ResponseTagInfo
	err := global.DB.Table("tag_relation as tr").Select("t.id,t.tname").Joins("inner join tag as t on t.id=tr.tid").Where("tr.lid=?", linkId).Scan(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

func GetLinkIdByTagName(c *gin.Context, name string) ([]uint, error) {
	var ids, tids []uint
	var err error
	//搜索对应的标签id
	tids, err = SearchMultiIdByName(c, name)
	if err != nil || len(tids) == 0 {
		global.Log.Debug("GetLinkIdByTagName调用SearchMultiIdByName返回为空", zap.Error(err), zap.Any("tids", tids))
		return []uint{}, gorm.ErrRecordNotFound
	}
	err = global.DB.Model(model.TagRelation{}).Select("lid").Where("tid in(?)", tids).Scan(&ids).Error
	if err != nil || len(ids) == 0 {
		global.Log.Debug("GetLinkIdByTagName查询tid结果为空", zap.Any("ids", tids))
		return []uint{}, gorm.ErrRecordNotFound
	}
	return ids, nil
}

func DeleteRealtionTagByLinkId(c *gin.Context, linkId uint) error {
	if err := global.DB.Where("lid=?", linkId).Delete(&model.TagRelation{}).Error; err != nil {
		global.Log.Warn("TagRelation-DeleteRealtionTagByLinkId,删除链接管理标签失败", zap.Error(err))
		return errors.New("删除失败")
	}
	return nil
}
