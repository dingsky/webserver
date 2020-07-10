package ctrl

import (
	"github.com/gin-gonic/gin"
	"webserver/pkg/service/v0Service"
)
func HttpRouter() *gin.Engine {
	router := gin.Default()
	router.POST("api/v0/webserver/register/account"， v0Service.DefService.RegisterAccount)
	return router
}