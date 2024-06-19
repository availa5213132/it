package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb/gvb_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok(map[string]string{}, "hello", c)
}
