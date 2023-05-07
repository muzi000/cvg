package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SSRF(ctx *gin.Context) {
	u := ctx.Query("url")
	if u == "" {
		ctx.String(http.StatusBadRequest, "give me a query `url`. ")
		return
	}
	resp, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}
	ctx.Data(http.StatusOK, "text/html", body)
	//fmt.Fprintf(w, "%s", body)

}
