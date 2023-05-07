package controller

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Template(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.String(http.StatusBadRequest, "give me a query `name` ")
		return
	}
	ctx.HTML(http.StatusOK, "xxs.html", gin.H{
		"name": name,
	})

}

func TemplateVuln(ctx *gin.Context) {
	t := `
	<html lang="zh-CN">
    <head>
        <title>vuln-go</title>
    </head>
    <body>
        holle {{.}}
    </body>
	</html>
	`
	name := ctx.Query("name")
	tmpl, err := template.New("test").Parse(t) //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		ctx.String(http.StatusBadRequest, "")
		return
	}
	b := new(strings.Builder)
	err = tmpl.Execute(b, name) //将string与模板合成，变量name的内容会替换掉{{.}}
	if err != nil {
		ctx.String(http.StatusInternalServerError, "")
		return
	}
	ctx.Data(http.StatusOK, "text/html", []byte(b.String()))
}
