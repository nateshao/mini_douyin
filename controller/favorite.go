package controller

import (
	"net/http"
	"strconv"

	"mini_douyin/service"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoListResponse struct {
	BaseResponse
	VideoList interface{} `json:"video_list"`
}

type FavoriteActionRequest struct {
	VideoID    int64  `form:"video_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required,oneof=1 2"`
}

// FavoriteAction 点赞/取消点赞
// @Summary 点赞/取消点赞
// @Description 用户对视频进行点赞或取消点赞
// @Tags favorite
// @Accept json
// @Produce json
// @Param token query string true "用户鉴权token"
// @Param video_id query int true "视频ID"
// @Param action_type query string true "1-点赞 2-取消点赞"
// @Success 200 {object} BaseResponse
// @Router /douyin/favorite/action/ [post]
func FavoriteAction(c *gin.Context) {
	var req FavoriteActionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, BaseResponse{StatusCode: 1, StatusMsg: "未登录"})
		return
	}
	if err := service.FavoriteAction(userID.(int64), req.VideoID, req.ActionType); err != nil {
		c.JSON(http.StatusOK, BaseResponse{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{StatusCode: 0})
}

// FavoriteList 喜欢列表
// @Summary 喜欢列表
// @Description 获取用户喜欢的视频列表
// @Tags favorite
// @Accept json
// @Produce json
// @Param user_id query int true "用户ID"
// @Success 200 {object} VideoListResponse
// @Router /douyin/favorite/list/ [get]
func FavoriteList(c *gin.Context) {
	type FavoriteListRequest struct {
		UserID int64 `form:"user_id" binding:"required"`
	}
	var req FavoriteListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	videos, err := service.FavoriteList(req.UserID)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{BaseResponse: BaseResponse{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{BaseResponse: BaseResponse{StatusCode: 0}, VideoList: videos})
}

func getVideoLikeKey(videoID int64) string {
	return "video:like:" + strconv.FormatInt(videoID, 10)
}
