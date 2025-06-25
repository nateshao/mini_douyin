package dao

import (
	"mini_douyin/model"
)

func CreateComment(userID, videoID int64, content string) (*model.Comment, error) {
	comment := &model.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: content,
	}
	if err := DB.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func DeleteComment(commentID, userID int64) error {
	return DB.Where("id = ? AND user_id = ?", commentID, userID).Delete(&model.Comment{}).Error
}

func GetComments(videoID int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := DB.Where("video_id = ?", videoID).Order("created_at desc").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
