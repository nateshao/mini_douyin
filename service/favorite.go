package service

import (
	"mini_douyin/cache"
	"mini_douyin/dao"
	"mini_douyin/model"
	"strconv"
)

func FavoriteAction(userID, videoID int64, actionType string) error {
	if actionType == "1" {
		if err := dao.CreateFavorite(userID, videoID); err != nil {
			return err
		}
		cache.RDB.Incr(cache.Ctx, getVideoLikeKey(videoID))
	} else if actionType == "2" {
		if err := dao.DeleteFavorite(userID, videoID); err != nil {
			return err
		}
		cache.RDB.Decr(cache.Ctx, getVideoLikeKey(videoID))
	}
	return nil
}

func FavoriteList(userID int64) ([]model.Video, error) {
	return dao.GetFavoriteVideos(userID)
}

func getVideoLikeKey(videoID int64) string {
	return "video:like:" + strconv.FormatInt(videoID, 10)
}
