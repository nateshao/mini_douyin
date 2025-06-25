package controller

import (
	"net/http"
	"strconv"

	"mini_douyin/service"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	BaseResponse
	CommentList interface{} `json:"comment_list"`
}

type CommentActionRequest struct {
	VideoID     int64  `form:"video_id" binding:"required"`
	ActionType  string `form:"action_type" binding:"required,oneof=1 2"`
	CommentText string `form:"comment_text"`
	CommentID   int64  `form:"comment_id"`
}

// CommentAction 评论/删除评论
// @Summary 评论/删除评论
// @Description 用户对视频进行评论或删除评论
// @Tags comment
// @Accept json
// @Produce json
// @Param token query string true "用户鉴权token"
// @Param video_id query int true "视频ID"
// @Param action_type query string true "1-发布评论 2-删除评论"
// @Param comment_text query string false "评论内容"
// @Param comment_id query int false "评论ID"
// @Success 200 {object} BaseResponse
// @Router /douyin/comment/action/ [post]
func CommentAction(c *gin.Context) {
	var req CommentActionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, BaseResponse{StatusCode: 1, StatusMsg: "未登录"})
		return
	}
	comment, err := service.CommentAction(userID.(int64), req.VideoID, req.ActionType, req.CommentText, req.CommentID)
	if err != nil {
		c.JSON(http.StatusOK, BaseResponse{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if req.ActionType == "1" {
		c.JSON(http.StatusOK, gin.H{"status_code": 0, "comment": comment})
	} else {
		c.JSON(http.StatusOK, BaseResponse{StatusCode: 0})
	}
}

// CommentList 评论列表
// @Summary 评论列表
// @Description 获取视频的评论列表
// @Tags comment
// @Accept json
// @Produce json
// @Param video_id query int true "视频ID"
// @Success 200 {object} CommentListResponse
// @Router /douyin/comment/list/ [get]
func CommentList(c *gin.Context) {
	type CommentListRequest struct {
		VideoID int64 `form:"video_id" binding:"required"`
	}
	var req CommentListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, BaseResponse{StatusCode: 1, StatusMsg: "参数错误: " + err.Error()})
		return
	}
	comments, err := service.CommentList(req.VideoID)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{BaseResponse: BaseResponse{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{BaseResponse: BaseResponse{StatusCode: 0}, CommentList: comments})
}

func getVideoCommentKey(videoID int64) string {
	return "video:comment:" + strconv.FormatInt(videoID, 10)
}
