package routers

import (
	"github.com/gin-gonic/gin"
	"gvb/gvb_server/api"
)

func SettingsRouter(router *gin.Engine) {
	SettingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", SettingsApi.SettingsInfoView)
}
