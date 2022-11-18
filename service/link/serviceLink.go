package link

import (
	"easyweb/global"
	"easyweb/model"
	"easyweb/pkg/str"
	categorySer "easyweb/service/category"
	tagSer "easyweb/service/tag"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math"
)

func ExistLinkByName(c *gin.Context, title string) bool {
	if err := global.DB.First(&model.Link{}, "title=?", title).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func ExistLinkById(c *gin.Context, id uint) bool {
	if err := global.DB.First(&model.Link{}, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func Create(c *gin.Context, data RequestLinkAdd) error {
	newLink := &model.Link{
		Title:  data.Title,
		Cid:    data.Cid,
		Url:    data.Url,
		Remark: data.Remark,
		Zank:   data.Zank,
	}
	if err := global.DB.Create(newLink).Error; err != nil {
		global.Log.Warn("serviceLink-Create,添加失败", zap.Error(err))
		return err
	}
	if len(data.TagIds) != 0 {
		tids := str.SplitToUintList(data.TagIds, ",")
		if err := tagSer.UpdateRealtionTag(c, newLink.Id, tids); err != nil {
			return err
		}
	}

	//更新关联分类的统计字段
	if err := categorySer.Increment(c, data.Cid); err != nil {
		return err
	}
	return nil
}

func Update(c *gin.Context, data RequestLinkUpdate) error {
	if err := global.DB.Where("id=?", data.Id).Updates(&model.Link{
		Title:  data.Title,
		Cid:    data.Cid,
		Url:    data.Url,
		Remark: data.Remark,
		Zank:   data.Zank,
	}).Error; err != nil {
		return err
	}
	if len(data.TagIds) != 0 {
		tids := str.SplitToUintList(data.TagIds, ",")
		tagSer.UpdateRealtionTag(c, data.Id, tids)
	} else {
		if err := tagSer.DeleteRealtionTagByLinkId(c, data.Id); err != nil {
			return errors.New("更新成功")
		}
	}
	return nil
}

func Delete(c *gin.Context, id uint) error {
	if !ExistLinkById(c, id) {
		return errors.New("链接不存在")
	}
	if err := global.DB.Where("id=?", id).Delete(&model.Link{}).Error; err != nil {
		global.Log.Warn("serviceLink-Delete,删除失败", zap.Error(err))
		return errors.New("删除失败")
	}
	return nil
}

func GetInfo(c *gin.Context, id uint) (ResponseInfo, error) {
	var info ResponseInfo
	if err := global.DB.Model(&model.Link{}).First(&info, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return info, errors.New("链接不存在")
		}
		global.Log.Warn("serviceCategory-GetInfo,获取数据失败", zap.Error(err))
		return info, errors.New("获取失败")
	}
	//获取分类数据
	if category, err := categorySer.GetInfo(c, info.Cid); !errors.Is(err, gorm.ErrRecordNotFound) {
		info.Category = category
	}
	//获取标签数据
	if tags, err := tagSer.GetAllByLinkId(c, info.Id); !errors.Is(err, gorm.ErrRecordNotFound) {
		info.Tag = tags
	}
	//获取标签数据
	return info, nil
}

func GetConditionByPage(c *gin.Context, page, limit int, condition ConditionLink) (PageLink, error) {
	var list []ResponseInfo
	var count int64
	query := global.DB.Model(&model.Link{})
	//按照标题搜索
	if condition.Title != "" {
		query = query.Where("title like ?", "%"+condition.Title+"%")
	}
	//按照分类id搜索
	if condition.Cid != 0 {
		query = query.Where("cid = ?", condition.Cid)
	}
	//按照分类名搜索
	if condition.Cname != "" {
		cateList, err := categorySer.SearchByName(c, condition.Cname)
		lenCateList := len(cateList)
		if err != nil || lenCateList == 0 {
			return PageLink{}, nil
		}
		ids := make([]uint, 0, lenCateList)
		for _, cate := range cateList {
			ids = append(ids, cate.Id)
		}
		query = query.Where("cid in(?)", ids)
	}
	//按照标签搜索
	if condition.TagName != "" {
		linkIds, err := tagSer.GetLinkIdByTagName(c, condition.TagName)

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			query = query.Where("id", linkIds)
		} else {
			return PageLink{}, nil
		}
	}

	query.Count(&count)
	pageAmount := int(math.Ceil(float64(count) / float64(limit)))
	if count == 0 || page > pageAmount {
		return PageLink{}, nil
	}
	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&list).Error; err != nil {
		return PageLink{}, err
	}

	for index, link := range list {
		if category, err := categorySer.GetInfo(c, link.Cid); !errors.Is(err, gorm.ErrRecordNotFound) {
			list[index].Category = category
		}
		if tag, err := tagSer.GetAllByLinkId(c, link.Id); !errors.Is(err, gorm.ErrRecordNotFound) {
			list[index].Tag = tag
		}
	}

	return PageLink{
		Amount: pageAmount,
		Data:   list,
	}, nil
}
