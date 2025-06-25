package controller

import (
	"net/http"
	"strconv"

	"mini_douyin/service"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	BaseResponse
	UserList interface{} `json:"user_list"`
}

type RelationActionRequest struct {
	ToUserID   int64  `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required,oneof=1 2"`
}

// RelationAction 关注/取消关注
// @Summary 关注/取消关注
// @Description 用户关注或取消关注其他用户
// @Tags relation
// @Accept json
// @Produce json
// @Param token query string true "用户鉴权token"
// @Param to_user_id query int true "对方用户ID"
// @Param action_type query string true "1-关注 2-取消关注"
// @Success 200 {object} BaseResponse
// @Router /douyin/relation/action/ [post]
func RelationAction(c *gin.Context) {
	var req RelationActionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, BaseResponse{StatusCode: 1, StatusMsg: "未登录"})
		return
	}
	if err := service.RelationAction(userID.(int64), req.ToUserID, req.ActionType); err != nil {
		c.JSON(http.StatusOK, BaseResponse{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{StatusCode: 0})
}

// FollowList 关注列表
// @Summary 关注列表
// @Description 获取用户关注的用户列表
// @Tags relation
// @Accept json
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} UserListResponse
// @Router /douyin/relation/follow/list/ [get]
func FollowList(c *gin.Context) {
	type FollowListRequest struct {
		UserID int64 `form:"user_id" binding:"required"`
	}
	var req FollowListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	users, err := service.FollowList(req.UserID)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{BaseResponse: BaseResponse{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{BaseResponse: BaseResponse{StatusCode: 0}, UserList: users})
}

// FollowerList 粉丝列表
// @Summary 粉丝列表
// @Description 获取用户的粉丝列表
// @Tags relation
// @Accept json
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} UserListResponse
// @Router /douyin/relation/follower/list/ [get]
func FollowerList(c *gin.Context) {
	type FollowerListRequest struct {
		UserID int64 `form:"user_id" binding:"required"`
	}
	var req FollowerListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	users, err := service.FollowerList(req.UserID)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{BaseResponse: BaseResponse{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{BaseResponse: BaseResponse{StatusCode: 0}, UserList: users})
}

func getUserFollowKey(userID int64) string {
	return "user:follow:" + strconv.FormatInt(userID, 10)
}
func getUserFollowerKey(userID int64) string {
	return "user:follower:" + strconv.FormatInt(userID, 10)
}
