package v1

import (
	"net/http"

	"github.com/sraddhanand/echo/utils/app"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	appG := app.Gin{C: c}
	msg := c.Param("msg")
	appG.Response(http.StatusOK, msg)

}
