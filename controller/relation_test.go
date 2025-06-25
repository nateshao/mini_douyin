package controller_test

import (
	"mini_douyin/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRelationAction(t *testing.T) {
	r := gin.Default()
	r.POST("/douyin/relation/action/", controller.RelationAction)
	req, _ := http.NewRequest("POST", "/douyin/relation/action/?to_user_id=2&action_type=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestFollowList(t *testing.T) {
	r := gin.Default()
	r.GET("/douyin/relation/follow/list/", controller.FollowList)
	req, _ := http.NewRequest("GET", "/douyin/relation/follow/list/?user_id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestFollowerList(t *testing.T) {
	r := gin.Default()
	r.GET("/douyin/relation/follower/list/", controller.FollowerList)
	req, _ := http.NewRequest("GET", "/douyin/relation/follower/list/?user_id=1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
