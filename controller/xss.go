package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// http://localhost:8080/xss/reflect?xss=%3Cscript%3Ealert(1)%3C/script%3E
func XSSReflectVuln(ctx *gin.Context) {
	reflect := ctx.Query("xss")
	ctx.Data(http.StatusOK, "text/html", []byte(reflect))
}

// http://localhost:8080/xss/dom?xss=%3Cscript%3Ealert(1)%3C/script%3E
func XSSDomVuln(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "xssdom.html", "")
}

// http://localhost:8080/xss/store?xss=%3Cscript%3Ealert(1)%3C/script%3E
func XSSStoreVuln(ctx *gin.Context) {
	xss := ctx.Query("xss")
	ctx.SetCookie("xss", xss, 3600, "/", "", false, true)
	ctx.String(http.StatusOK, "Set param into cookie")
}

func XSSShow(ctx *gin.Context) {
	xss, _ := ctx.Cookie("xss")
	ctx.Data(http.StatusOK, "text/html", []byte(xss))
}
