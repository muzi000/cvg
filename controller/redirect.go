package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzi000/vuln-go/security"
)

// UrlRedirectVuln  302
// http://localhost:8080/redirect/url?url=http://www.baidu.com
func UrlRedirectVuln(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.String(http.StatusBadRequest, "give me a query `url`")
		return
	}
	ctx.Redirect(http.StatusFound, url)
}

// HeadRedirectVuln 301
// http://localhost:8080/redirect/head?url=http://www.baidu.com
func HeadRedirectVuln(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.String(http.StatusBadRequest, "give me a query `url`")
		return
	}
	ctx.Status(http.StatusMovedPermanently)

	ctx.Header("location", url)
}

// Forword 路由代理
func Forword(ctx *gin.Context) {

}

func SecRedirect(ctx *gin.Context) {
	url := ctx.Query("url")
	if url == "" {
		ctx.String(http.StatusBadRequest, "give me a query `url`")
		return
	}
	if security.CheckUrl(url) {
		ctx.Redirect(http.StatusFound, url)
		return
	}
	ctx.String(http.StatusForbidden, "unsupport url: "+url)
}
