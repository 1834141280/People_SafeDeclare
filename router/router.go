package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"software/handler"
)

func InitRouter(r *gin.Engine) {
	r.LoadHTMLGlob("views/html/*")
	r.StaticFS("/css", http.Dir("./views/css"))
	r.StaticFS("/image", http.Dir("./views/image"))
	r.StaticFS("/js", http.Dir("./views/js"))
	////////////////////////////////////////////////////
	r.StaticFS("/bootstrap/css", http.Dir("./views/bootstrap/css"))
	r.StaticFS("/bootstrap/js", http.Dir("./views/bootstrap/js"))
	r.StaticFS("/bootstrap-table-master/dist", http.Dir("./views/bootstrap-table-master/dist"))
	////////////////////////////////////////////////////
	indexgroup := r.Group("/index")
	{
		indexgroup.GET("", handler.Getindex())
		indexgroup.POST("/sign", handler.GetSign())
		indexgroup.POST("/register", handler.GetRegister())
	}
	declaregroup := r.Group("/declare")
	{
		declaregroup.GET("", handler.Getdeclare())
		declaregroup.POST("/submit", handler.GetSubmit())
	}
	remindgroup := r.Group("/remind")
	{
		remindgroup.GET("", handler.GetRemind())
		remindgroup.GET("/getun", handler.GetUn())
		remindgroup.GET("/getdanger", handler.GetDangerUser())
	}
	controlgroup := r.Group("/control")
	{
		controlgroup.GET("", handler.Getcontrol())
		controlgroup.GET("/write", handler.Getwrite())
		controlgroup.POST("/search", handler.GetSearch())
		controlgroup.POST("/increase", handler.GetIncrease())
		controlgroup.POST("/edit", handler.GetEdit())
		controlgroup.POST("/delete", handler.GetDelete())
	}
	showgroup := r.Group("/show")
	{
		showgroup.GET("", handler.GetShow())
		showgroup.GET("/information", handler.GetInformation())
	}

}
