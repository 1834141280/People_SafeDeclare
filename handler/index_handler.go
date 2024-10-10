package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software/data"
	"software/database"
)

var (
	userRegister data.UserRegister
	userSign     data.UserSign
)

func Getindex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

func GetSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&userSign); err != nil {
			fmt.Println("绑定失败")
			c.JSON(http.StatusOK, gin.H{
				"msg": "error",
			})
			panic(err)
		} else {
			fmt.Println("登录成功")
			c.SetCookie("index", userSign.Username, 3600, "/", "localhost", false, false)
			a := database.SearchUser(userSign)
			c.JSON(http.StatusOK, gin.H{
				"msg": a,
			})
			//缺身份比对，看是管理员还是用户

		}
		//val := c.PostForm("json")
		data.Initpeople(userSign)
		//fmt.Println("hello")
		fmt.Println(userSign)
	}

}

func GetRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetRegister")
		if err := c.BindJSON(&userRegister); err != nil {
			fmt.Println("绑定失败")
			c.JSON(http.StatusOK, gin.H{
				"msg": "error",
			})
			panic(err)
		} else {
			fmt.Println("注册成功")
			c.SetCookie("index", userRegister.Username, 3600, "/", "localhost", false, false)
			database.Init(userRegister)
			a := database.CreateUser(userRegister)
			c.JSON(http.StatusOK, gin.H{
				"msg": a,
			})
			//缺少数据库判断，错误要返回json：错误
		}
		fmt.Println(userRegister)
	}
}
