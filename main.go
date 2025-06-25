package main

import (
	"mini_douyin/cache"
	"mini_douyin/config"
	"mini_douyin/controller"
	"mini_douyin/dao"
	"mini_douyin/middleware"

	"github.com/gin-gonic/gin"
	// "mini_douyin/controller"
	// swagger 文档相关
	// _ "mini_douyin/docs"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/files"
)

func main() {
	conf := config.LoadConfig()
	if err := dao.InitMySQL(conf.MySQL.DSN); err != nil {
		panic(err)
	}
	cache.InitRedis(conf.Redis.Addr, conf.Redis.Password, conf.Redis.DB)

	r := gin.Default()
	r.Use(middleware.RateLimitMiddleware())

	// 未登录接口
	r.POST("/douyin/user/register/", controller.Register)
	r.POST("/douyin/user/login/", controller.Login)
	r.GET("/douyin/feed/", controller.Feed)
	r.GET("/douyin/comment/list/", controller.CommentList)

	// 需要登录的接口分组
	auth := r.Group("/", middleware.JWTAuth())
	{
		auth.POST("/douyin/favorite/action/", controller.FavoriteAction)
		auth.GET("/douyin/favorite/list/", controller.FavoriteList)
		auth.POST("/douyin/comment/action/", controller.CommentAction)
		auth.POST("/douyin/relation/action/", controller.RelationAction)
		auth.GET("/douyin/relation/follow/list/", controller.FollowList)
		auth.GET("/douyin/relation/follower/list/", controller.FollowerList)
		auth.POST("/douyin/message/action/", controller.MessageAction)
		auth.GET("/douyin/message/chat/", controller.MessageChat)
	}

	// swagger 文档路由（如需开启，取消注释）
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
