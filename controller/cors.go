package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsVuln1(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Add("Access-Control-Allow-Origin", origin)      // set origin from header
	ctx.Writer.Header().Add("Access-Control-Allow-Credentials", "true") // allow cookie
}

// CorsVuln2 设置为任意 *
func CorsVuln2(ctx *gin.Context) {
	ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")

}

// CorsSec 只允许白名单
func CorsSec(ctx *gin.Context) {
	allowOrigin := []string{"baidu.com", "google.com"}
	origin := ctx.Request.Header.Get("origin")
	// 如果origin不为空并且origin不在白名单内，认定为不安全。
	// 如果origin为空，表示是同域过来的请求或者浏览器直接发起的请求。
	if origin == "" || !listCont(allowOrigin, origin) {
		ctx.Status(http.StatusForbidden)
		return
	}
	ctx.String(http.StatusOK, "safety")
}

func listCont(l []string, c string) bool {

	for _, n := range l {
		if n == c {
			return true
		}
	}
	return false
}
