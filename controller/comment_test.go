package controller_test

import (
	"mini_douyin/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCommentAction(t *testing.T) {
	r := gin.Default()
	r.POST("/douyin/comment/action/", controller.CommentAction)
	req, _ := http.NewRequest("POST", "/douyin/comment/action/?video_id=1&action_type=1&comment_text=hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCommentList(t *testing.T) {
	r := gin.Default()
	r.GET("/douyin/comment/list/", controller.CommentList)
	req, _ := http.NewRequest("GET", "/douyin/comment/list/?video_id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
