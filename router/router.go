package router

import (
	"abc/controllers"
	"abc/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {

	//创建一个默认的 Gin 引擎
	r := gin.Default()

	//中间件调用日志
	r.Use(gin.LoggerWithConfig(utils.LoggerToFile()))
	r.Use(utils.Recover)

	//路由
	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetUserInfoList)
		user.POST("/listJson", controllers.UserController{}.GetUserInfoJson)
		user.POST("/lists", controllers.UserController{}.GetUserInfoStruct)

		//gorm进行数据库添加操作
		user.POST("/add", controllers.UserController{}.AddUser)
		//gorm修改名字
		user.POST("/updateName", controllers.UserController{}.UpdateUser)
		//gorm删除
		user.POST("/deleteById", controllers.UserController{}.DeleteUser)
		//gorm查找
		user.POST("/findById", controllers.UserController{}.FindUsers)
		//测试日志生成
		user.POST("/logger", controllers.UserController{}.LoggerTest)

		user.PUT("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "add")
		})
		user.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "delete")
		})
	}

	return r
}
