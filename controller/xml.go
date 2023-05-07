package controller

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func XML(ctx *gin.Context) {
	xmlData, _ := ioutil.ReadAll(ctx.Request.Body)

	data := mydata{}

	err := xml.Unmarshal(xmlData, &data)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusInternalServerError, "parase xml error")
		return
	}
	ctx.String(http.StatusOK, data.Name)

}

func XML2(ctx *gin.Context) {
	r := xml.NewDecoder(ctx.Request.Body)
	data := mydata{}
	err := r.Decode(&data)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusInternalServerError, "parase xml error")
		return
	}
	ctx.Data(http.StatusOK, "text/html", []byte(data.Name))
}

type mydata struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
}
