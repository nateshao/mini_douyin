package service

import (
	"mini_douyin/cache"
	"mini_douyin/dao"
	"mini_douyin/model"
	"strconv"
)

func CommentAction(userID, videoID int64, actionType, text string, commentID int64) (*model.Comment, error) {
	if actionType == "1" {
		comment, err := dao.CreateComment(userID, videoID, text)
		if err != nil {
			return nil, err
		}
		cache.RDB.Incr(cache.Ctx, getVideoCommentKey(videoID))
		return comment, nil
	} else if actionType == "2" {
		if err := dao.DeleteComment(commentID, userID); err != nil {
			return nil, err
		}
		cache.RDB.Decr(cache.Ctx, getVideoCommentKey(videoID))
		return nil, nil
	}
	return nil, nil
}

func CommentList(videoID int64) ([]model.Comment, error) {
	return dao.GetComments(videoID)
}

func getVideoCommentKey(videoID int64) string {
	return "video:comment:" + strconv.FormatInt(videoID, 10)
}
