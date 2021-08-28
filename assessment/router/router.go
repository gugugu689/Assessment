package router

import (
	"assessment/Controller"
	"assessment/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	//日志中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/blog")

	//注册
	v1.POST("/signup", Controller.SignUp)
	//登陆
	v1.POST("/login", Controller.Login)
	//查看帖子
	v1.GET("/post/:id", Controller.GetPostDetail)
	//按照 浏览数/时间 排序 查看帖子列表
	v1.GET("posts_by", Controller.GetPostsBy)
	//查看评论
	v1.GET("/post/:id/get_comment", Controller.GetComment)
	//查看所有社区
	v1.GET("/communities", Controller.GetCommunities)

	//刷新token
	v1.GET("/refresh_token", Controller.RefreshToken)

	//需要token鉴权的
	v1.Use(Controller.JWTAuthMiddleware())
	{
		//发表帖子
		v1.POST("/post", Controller.CreatePost)
		//更新帖子
		v1.PUT("/post/:id/update", Controller.UpdatePost)
		//删除帖子 以及其评论
		v1.DELETE("/post/:id/delete", Controller.DeletePost)
		//发表评论
		v1.POST("/post/:id/post_comment", Controller.CreateComment)
		//创建社区
		v1.POST("/create_community", Controller.CreateCommunity)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
