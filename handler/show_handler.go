package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"software/data"
	"software/database"
)

var showinformation []data.UserInformation

func GetShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "show.html", nil)
	}
}

func GetInformation() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.ConnectMySQL()
		defer db.Close()
		db.Table(data.TableUserInformation).Find(&showinformation)
		showtabel := make([]data.UserInformation, 0)
		for i := 0; i < len(showinformation); i++ {
			showtabel = append(showtabel, showinformation[i])
		}
		fmt.Print("showtable:%+v\n", showtabel)
		c.JSON(http.StatusOK, gin.H{
			"rows": showtabel,
		})
	}
}
