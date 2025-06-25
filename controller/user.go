package controller

import (
	"mini_douyin/dao"
	"mini_douyin/middleware"
	"mini_douyin/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": 400, "status_msg": "参数错误"})
		return
	}
	var user model.User
	if err := dao.DB.Where("username = ?", req.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "用户名已存在"})
		return
	}
	user = model.User{Username: req.Username, Password: req.Password, Name: req.Username}
	if err := dao.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 2, "status_msg": "注册失败"})
		return
	}
	token, _ := middleware.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "user_id": user.ID, "token": token})
}

func Login(c *gin.Context) {
	var req struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status_code": 400, "status_msg": "参数错误"})
		return
	}
	var user model.User
	if err := dao.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status_code": 1, "status_msg": "用户名或密码错误"})
		return
	}
	token, _ := middleware.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"status_code": 0, "user_id": user.ID, "token": token})
}
