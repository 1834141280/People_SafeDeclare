package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software/data"
	"software/database"
)

var people data.People

func Getdeclare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "declare.html", nil)
	}
}

func GetSubmit() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GetSubmit")
		if err := c.BindJSON(&people); err != nil {
			fmt.Println("绑定失败")
			c.JSON(http.StatusOK, gin.H{
				"msg": "error",
			})
			panic(err)
		} else {
			fmt.Println("提交成功")
			//c.SetCookie("", userSign.Username, 3600, "/", "localhost", false, false)
			a := database.GetUserInformation(people)
			c.JSON(http.StatusOK, gin.H{
				"msg": a,
			})
			//确实判断今天是否提交
		}
	}

}
