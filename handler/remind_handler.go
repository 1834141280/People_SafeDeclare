package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software/data"
	"software/database"
)

func GetRemind() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "remind.html", nil)
	}
}

func GetUn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var a []string
		var u []data.UserInformation
		db := database.ConnectMySQL()
		defer db.Close()
		if err := db.Table(data.TableUserInformation).Find(&u).Error; err == nil {
			fmt.Printf("old:%+v\n", u)
		}
		for i := 0; i < len(u); i++ {
			if u[i].Write == "未填报" {
				a = append(a, u[i].Name)
			}

		}
		fmt.Printf("old:%+v\n", a)
		c.JSON(http.StatusOK, gin.H{
			"msg": a,
		})
	}
}

func GetDangerUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var a []string
		var u []data.UserInformation
		db := database.ConnectMySQL()
		defer db.Close()
		if err := db.Table(data.TableUserInformation).Find(&u).Error; err == nil {
			fmt.Printf("old:%+v\n", u)
		}
		for i := 0; i < len(u); i++ {
			if u[i].High_region == "是" || u[i].Isolation == "是" || u[i].Typical_symptoms == "是" || u[i].Fever == "是" {
				a = append(a, u[i].Name)
			}

		}
		fmt.Printf("old:%+v\n", a)
		c.JSON(http.StatusOK, gin.H{
			"msg": a,
		})
	}
}
