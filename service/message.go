package service

import (
	"mini_douyin/dao"
	"mini_douyin/model"
)

func MessageAction(userID, toUserID int64, actionType, content string) error {
	if actionType == "1" {
		return dao.CreateMessage(userID, toUserID, content)
	}
	return nil
}

func MessageChat(userID, toUserID int64) ([]model.Message, error) {
	return dao.GetMessages(userID, toUserID)
}
