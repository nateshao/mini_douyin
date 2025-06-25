package controller

import (
	"mini_douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageListResponse struct {
	BaseResponse
	MessageList interface{} `json:"message_list"`
}

type MessageActionRequest struct {
	ToUserID   int64  `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required,oneof=1"`
	Content    string `form:"content"`
}

// MessageAction 发送消息
// @Summary 发送消息
// @Description 用户发送消息
// @Tags message
// @Accept json
// @Produce json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query int true "对方用户ID"
// @Param action_type query string true "1-发送消息"
// @Param content query string false "消息内容"
// @Success 200 {object} BaseResponse
// @Router /douyin/message/action/ [post]
func MessageAction(c *gin.Context) {
	var req MessageActionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, BaseResponse{StatusCode: 1, StatusMsg: "未登录"})
		return
	}
	if err := service.MessageAction(userID.(int64), req.ToUserID, req.ActionType, req.Content); err != nil {
		c.JSON(http.StatusOK, BaseResponse{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{StatusCode: 0})
}

// MessageChat 聊天记录
// @Summary 聊天记录
// @Description 获取与指定用户的聊天消息记录
// @Tags message
// @Accept json
// @Produce json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query int true "对方用户ID"
// @Success 200 {object} MessageListResponse
// @Router /douyin/message/chat/ [get]
func MessageChat(c *gin.Context) {
	type MessageChatRequest struct {
		ToUserID int64 `form:"to_user_id" binding:"required"`
	}
	var req MessageChatRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, BaseResponse{StatusCode: 1, StatusMsg: "未登录"})
		return
	}
	messages, err := service.MessageChat(userID.(int64), req.ToUserID)
	if err != nil {
		c.JSON(http.StatusOK, MessageListResponse{BaseResponse: BaseResponse{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, MessageListResponse{BaseResponse: BaseResponse{StatusCode: 0}, MessageList: messages})
}
