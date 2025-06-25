package service

import (
	"mini_douyin/cache"
	"mini_douyin/dao"
	"mini_douyin/model"
	"strconv"
)

func RelationAction(userID, toUserID int64, actionType string) error {
	if actionType == "1" {
		if err := dao.CreateRelation(userID, toUserID); err != nil {
			return err
		}
		cache.RDB.Incr(cache.Ctx, getUserFollowKey(userID))
		cache.RDB.Incr(cache.Ctx, getUserFollowerKey(toUserID))
	} else if actionType == "2" {
		if err := dao.DeleteRelation(userID, toUserID); err != nil {
			return err
		}
		cache.RDB.Decr(cache.Ctx, getUserFollowKey(userID))
		cache.RDB.Decr(cache.Ctx, getUserFollowerKey(toUserID))
	}
	return nil
}

func FollowList(userID int64) ([]model.User, error) {
	return dao.GetFollowList(userID)
}

func FollowerList(userID int64) ([]model.User, error) {
	return dao.GetFollowerList(userID)
}

func getUserFollowKey(userID int64) string {
	return "user:follow:" + strconv.FormatInt(userID, 10)
}
func getUserFollowerKey(userID int64) string {
	return "user:follower:" + strconv.FormatInt(userID, 10)
}
