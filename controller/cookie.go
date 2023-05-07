package controller

import "github.com/gin-gonic/gin"

func CookieVuln1(ctx *gin.Context) {
	user, err := ctx.Cookie("user")
	if err != nil {
		ctx.SetCookie("user", "test", 3600, "/", "localhost", false, true)
		return
	}
	ctx.Writer.WriteString("hello " + user)
}
