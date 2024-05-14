package routers

import (
	"github.com/Touch/cmd/api/controllers"
	"github.com/Touch/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(router *gin.Engine) {
	RestHandler := controllers.NewHandler()

	// live probe
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	v1 := router.Group(controllers.APIV1)
	v1.POST("/login", RestHandler.Login)
	//v1.POST("/create-account", RestHandler.CreateAccount)
	//v1.POST("/get-account", RestHandler.GetAccount)
	//v1.POST("/create-personal-info", RestHandler.CreatePersonalInfo)
	//v1.POST("/recommend-users", RestHandler.RecommendUsers)
	//v1.POST("/show-nft", RestHandler.ShowNft)
	//v1.POST("/show-user-upgrade-nft", RestHandler.ShowUserUpgradeNft)

	// 静态资源访问
	v1.StaticFS("static", http.Dir(config.GetConfig().Options.UploadDir))
}
