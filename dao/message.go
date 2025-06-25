package dao

import "mini_douyin/model"

func CreateMessage(fromUserID, toUserID int64, content string) error {
	msg := model.Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
	}
	return DB.Create(&msg).Error
}

func GetMessages(userID, toUserID int64) ([]model.Message, error) {
	var messages []model.Message
	if err := DB.Where(
		"(from_user_id = ? AND to_user_id = ?) OR (from_user_id = ? AND to_user_id = ?)",
		userID, toUserID, toUserID, userID,
	).Order("created_at asc").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
