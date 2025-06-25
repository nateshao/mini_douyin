package controller_test

import (
	"mini_douyin/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFavoriteAction(t *testing.T) {
	r := gin.Default()
	r.POST("/douyin/favorite/action/", controller.FavoriteAction)
	req, _ := http.NewRequest("POST", "/douyin/favorite/action/?video_id=1&action_type=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestFavoriteList(t *testing.T) {
	r := gin.Default()
	r.GET("/douyin/favorite/list/", controller.FavoriteList)
	req, _ := http.NewRequest("GET", "/douyin/favorite/list/?user_id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
