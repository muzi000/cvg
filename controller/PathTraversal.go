package controller

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func PathTraVuln(ctx *gin.Context) {
	fileName := ctx.Query("file")
	// _,err:=os.Open(fileName)
	// if err!=nil{
	// 	ctx.String(http.StatusNotFound,"open err")
	// 	return
	// }
	ctx.File(fileName)

}

func PathSec(ctx *gin.Context) {
	fileName := ctx.Query("file")

	//fileName,err=url.QueryUnescape(fileName)
	for strings.Contains(fileName, "%") {
		var err error
		fileName, err = url.QueryUnescape(fileName)
		if err != nil {
			ctx.String(http.StatusBadRequest, "decode fail")
			return
		}
	}
	if strings.Contains(fileName, "..") || fileName[0] == '/' || strings.Contains(fileName, ":") {
		ctx.String(http.StatusForbidden, "can't include ..")
		return
	}
	ctx.File(fileName)

}
