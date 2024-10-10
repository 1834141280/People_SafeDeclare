package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software/data"
	"software/database"
)

func Getcontrol() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "control.html", nil)
	}
}

func Getwrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		//x := people.Write
		x := 0
		if x == 1 {
			c.JSON(http.StatusOK, gin.H{
				"msg": "true",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "false",
			})
		}
	}
}

func GetSearch() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("id")
		fmt.Println(username)
		information := database.GetInformation(username)
		c.JSON(http.StatusOK, gin.H{
			"msg": information,
		})
	}
}

func GetIncrease() gin.HandlerFunc {
	return func(c *gin.Context) {
		var information data.UserInformation
		if err := c.BindJSON(&information); err != nil {
			fmt.Println("绑定失败")
			c.JSON(http.StatusOK, gin.H{
				"msg": "error",
			})
			panic(err)
		} else {
			fmt.Println("提交成功")
			//c.SetCookie("", userSign.Username, 3600, "/", "localhost", false, false)
			a := database.IncreaseInformation(information)
			c.JSON(http.StatusOK, gin.H{
				"msg": a,
			})
		}
	}
}

func GetEdit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var information data.UserInformation
		if err := c.BindJSON(&information); err != nil {
			fmt.Println("绑定失败")
			c.JSON(http.StatusOK, gin.H{
				"msg": "error",
			})
			panic(err)
		} else {
			fmt.Println("提交成功")
			//c.SetCookie("", userSign.Username, 3600, "/", "localhost", false, false)
			a := database.IncreaseInformation(information)
			c.JSON(http.StatusOK, gin.H{
				"msg": a,
			})
		}
	}
}

func GetDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("id")
		fmt.Println("提交成功")
		fmt.Printf(username)
		userinformation := database.GetInformation(username)
		//c.SetCookie("", userSign.Username, 3600, "/", "localhost", false, false)

		userinformation.High_region = ""
		userinformation.Isolation = ""
		userinformation.Typical_symptoms = ""
		userinformation.Fever = ""
		userinformation.Other = ""
		userinformation.Temperature_morning = ""
		userinformation.Temperature_moon = ""
		userinformation.Temperature_aftermoon = ""
		userinformation.T = ""
		userinformation.Write = ""
		a := database.IncreaseInformation(userinformation)
		c.JSON(http.StatusOK, gin.H{
			"msg": a,
		})
	}
}
