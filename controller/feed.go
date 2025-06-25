package controller

import (
	"mini_douyin/dao"
	"mini_douyin/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	latestTimeStr := c.Query("latest_time")
	var latestTime time.Time
	if latestTimeStr == "" {
		latestTime = time.Now()
	} else {
		ts, _ := strconv.ParseInt(latestTimeStr, 10, 64)
		latestTime = time.Unix(ts, 0)
	}
	var videos []model.Video
	dao.DB.Where("created_at < ?", latestTime).Order("created_at desc").Limit(30).Find(&videos)
	// TODO: 填充作者信息、点赞数、评论数、is_favorite等
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"video_list":  videos,
		"next_time":   videos[len(videos)-1].CreatedAt.Unix(),
	})
}
