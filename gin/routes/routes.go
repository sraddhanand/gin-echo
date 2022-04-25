package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/sraddhanand/echo/app/v1"
)

func Routes(router *gin.Engine) {
	hcRoutes := router.Group("/")
	hcRoutes.GET("/echo/:msg", v1.Echo)
}
