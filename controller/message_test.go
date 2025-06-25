package controller_test

import (
	"mini_douyin/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMessageAction(t *testing.T) {
	r := gin.Default()
	r.POST("/douyin/message/action/", controller.MessageAction)
	req, _ := http.NewRequest("POST", "/douyin/message/action/?to_user_id=2&action_type=1&content=hi", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestMessageChat(t *testing.T) {
	r := gin.Default()
	r.GET("/douyin/message/chat/", controller.MessageChat)
	req, _ := http.NewRequest("GET", "/douyin/message/chat/?to_user_id=2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
