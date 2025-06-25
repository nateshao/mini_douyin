package dao

import "mini_douyin/model"

func CreateRelation(userID, toUserID int64) error {
	rel := model.Relation{UserID: userID, ToUserID: toUserID}
	return DB.Where(rel).FirstOrCreate(&rel).Error
}

func DeleteRelation(userID, toUserID int64) error {
	return DB.Where("user_id = ? AND to_user_id = ?", userID, toUserID).Delete(&model.Relation{}).Error
}

func GetFollowList(userID int64) ([]model.User, error) {
	var rels []model.Relation
	if err := DB.Where("user_id = ?", userID).Find(&rels).Error; err != nil {
		return nil, err
	}
	var users []model.User
	var ids []int64
	for _, rel := range rels {
		ids = append(ids, rel.ToUserID)
	}
	if len(ids) > 0 {
		if err := DB.Where("id IN ?", ids).Find(&users).Error; err != nil {
			return nil, err
		}
	}
	return users, nil
}

func GetFollowerList(userID int64) ([]model.User, error) {
	var rels []model.Relation
	if err := DB.Where("to_user_id = ?", userID).Find(&rels).Error; err != nil {
		return nil, err
	}
	var users []model.User
	var ids []int64
	for _, rel := range rels {
		ids = append(ids, rel.UserID)
	}
	if len(ids) > 0 {
		if err := DB.Where("id IN ?", ids).Find(&users).Error; err != nil {
			return nil, err
		}
	}
	return users, nil
}
