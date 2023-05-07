package route

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzi000/vuln-go/auth"
	"github.com/muzi000/vuln-go/controller"
)

func InitRoute() *gin.Engine {
	r := gin.Default()
	//gin.DisableConsoleColor()

	// 记录到文件。
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//全局模板
	r.LoadHTMLGlob("./view/*")

	//全局静态文件
	r.Static("/css", "/static/css")

	//认证
	r.Use(auth.Auth)

	//ssrf-token
	r.GET("/err", func(ctx *gin.Context) {
		var err error = errors.New("error test")
		ctx.Error(err)
		return
	})

	//index
	r.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"name": "admin",
		})
	})

	r.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", "")
	})
	r.POST("/login", controller.Login)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	//mysql
	mysqlInject := r.Group("/mysql")
	mysqlInject.GET("/1", controller.MysqlVuln1)
	mysqlInject.GET("/2", controller.MysqlVuln2)

	//upload
	upload := r.Group("/upload")
	upload.GET("/com", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", gin.H{
			"url": "/upload/com",
		})
	})
	upload.POST("/com", controller.UploadVuln1)
	upload.GET("/img", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", gin.H{
			"url": "/upload/img",
		})
	})
	upload.POST("/img", controller.UploadImg)
	upload.GET("/zip", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", gin.H{
			"url": "/upload/zip",
		})
	})
	upload.POST("/zip", controller.UploadZipVuln)

	//cors
	r.OPTIONS("/cors", controller.CorsVuln1)
	r.GET("/cors", func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("origin")
		ctx.String(http.StatusOK, "you can from "+origin)
	})

	//ip
	r.GET("/ip", controller.IPVuln)

	//jsonp
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		// /JSONP?callback=x
		// 将输出：x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	//path tral
	r.GET("path", controller.PathTraVuln)
	r.GET("path/sec", controller.PathSec)

	//el
	r.GET("/el", controller.EL)

	//ssrf
	r.GET("ssrf", controller.SSRF)

	//xxs
	r.GET("/xxs", controller.Template)

	//template
	r.GET("/template", controller.TemplateVuln)

	//xml
	r.GET("/xml", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "upload.html", gin.H{
			"url": "/xml",
		})
	})
	r.POST("/xml", controller.XML2)

	//redirect
	redirect := r.Group("/redirect")
	redirect.GET("/url", controller.UrlRedirectVuln)
	redirect.GET("/head", controller.HeadRedirectVuln)
	redirect.GET("/sec", controller.SecRedirect)

	//xss
	xss := r.Group("/xss")
	xss.GET("/reflect", controller.XSSReflectVuln)
	xss.GET("/store", controller.XSSStoreVuln)
	xss.GET("/store/show", controller.XSSShow)
	xss.GET("/dom", controller.XSSDomVuln)

	return r
}
