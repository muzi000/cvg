package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPVuln(ctx *gin.Context) {
	fmt.Println(ctx.ClientIP())
	if ctx.ClientIP() == "127.0.0.1" || ctx.ClientIP() == "::1" {
		ctx.String(http.StatusOK, "loacl call")
		return
	}
	ctx.String(http.StatusForbidden, "not loacl call")
}
