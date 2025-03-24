package controllers

import (
	"abc/src/utils"
	"abc/userTest"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserController 避免同包方法名重复错误
type UserController struct{}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (u UserController) GetUserInfo(c *gin.Context) {
	//Get获取参数
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	user := userTest.GetUserInfoTest(id)
	ReturnSuccess(c, 0, "success", user, 2)
}
func (u UserController) GetUserInfoList(c *gin.Context) {
	//Post获取参数
	postForm := c.PostForm("id")
	ReturnErr(c, 404, postForm)
}
func (u UserController) GetUserInfoJson(c *gin.Context) {
	//Post获取参数转成JSON
	param := make(map[string]interface{})
	err := c.BindJSON(&param)
	if err == nil {
		ReturnSuccess(c, 0, "success", param, 11)
		return
	}
	ReturnErr(c, 400, err)
}
func (u UserController) GetUserInfoStruct(c *gin.Context) {
	//Post获取参数转成结构体
	user := &User{}
	err := c.BindJSON(&user)
	if err == nil {
		ReturnSuccess(c, 0, "success", user, 11)
		return
	}
	ReturnErr(c, 400, "param err")
}

// AddUser 单个添加
func (u UserController) AddUser(c *gin.Context) {
	user := &userTest.Contact01{}
	c.BindJSON(&user)
	userTest.AddUserById(user)
	ReturnSuccess(c, 0, "success", user, 11)
}

// UpdateUser 修改名字
func (u UserController) UpdateUser(c *gin.Context) {
	user := &userTest.Contact01{}
	c.BindJSON(&user)
	userTest.UpdateUserById(user)
	ReturnSuccess(c, 0, "success", user, 11)
}

// DeleteUser 删除用户
func (u UserController) DeleteUser(c *gin.Context) {
	user := userTest.Contact01{}
	c.BindJSON(&user)
	userTest.DeleteUserById(user)
	ReturnSuccess(c, 0, "success", user, 11)
}

// FindUsers 查找用户
func (u UserController) FindUsers(c *gin.Context) {
	user := userTest.Contact01{}
	c.BindJSON(&user)
	users := userTest.FindUsers(user)
	ReturnSuccess(c, 0, "success", users, 11)

}

// 测试自定义日志生成
func (u UserController) LoggerTest(c *gin.Context) {
	utils.Write("日志信息", "user")
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("异常捕获", err)
	//	}
	//}()
	num1 := 100
	num2 := 0
	num3 := num1 / num2
	ReturnErr(c, 4004, num3)
}
