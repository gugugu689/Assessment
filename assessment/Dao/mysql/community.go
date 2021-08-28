package mysql

import (
	"assessment/Models"

	"go.uber.org/zap"
)

// CreateCommunity 创建社区
func CreateCommunity(com *Models.Community) (err error) {
	if err = db.Create(com).Error; err != nil {
		zap.L().Error("create community failed", zap.Error(err))
		return
	}
	return
}

//GetCommunityByID 通过community id 获取社区信息
func GetCommunityByID(id int64) (com *Models.Community, err error) {
	com = new(Models.Community)
	if err = db.Where("id=?", id).Find(com).Error; err != nil {
		zap.L().Error("GetCommunityByID failed", zap.Error(err))
	}
	return
}

//GetCommunities 获取所有社区
func GetCommunities() (coms []*Models.Community, err error) {
	coms = make([]*Models.Community, 0, 5)
	if err = db.Find(coms).Error; err != nil {
		return
	}
	return
}
