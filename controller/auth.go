package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzi000/vuln-go/module"
)

// AuthVuln1
func AuthVuln1(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.String(http.StatusBadRequest, "who are you")
		return
	}
	if name == "admin" {
		ctx.String(http.StatusOK, "welcome, admin")
		return
	}
	ctx.String(http.StatusUnauthorized, "only admin can access")
}

// AuthVuln2  http改包越权
func AuthVuln2(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	u, ok := (user).(module.User)
	if !ok {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"role": u.Role,
	})
}
