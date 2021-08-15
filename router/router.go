package router

import (
	"fast-gin/library/util"
	"github.com/gin-gonic/gin"
	"net/http"

	"fast-gin/config"
	"fast-gin/router/api/v1"
	"fast-gin/service"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", GetAuth)

	r.POST("/image/upload", UploadImage)
	r.StaticFS("/upload/images", http.Dir(util.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	gin.SetMode(config.Cfg.Server.Mode)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(util.JWT())
	{
		// user
		apiv1.POST("user/add", v1.AddUser)
		apiv1.POST("user/edit", v1.UpdateUser)
		apiv1.GET("user/info", v1.GetUserInfo)
		apiv1.GET("user/list", v1.GetAllUsers)
		apiv1.DELETE("user/del", v1.DelUser)
	}

	wx := r.Group("/wx")
	{
		wx.GET("/check", service.WXCheckSignature)
		wx.POST("/check", service.WXMsgReceive)
	}

	return r
}
