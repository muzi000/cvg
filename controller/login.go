package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muzi000/vuln-go/auth"
	"github.com/muzi000/vuln-go/security"
)

// 登录处理
func Login(ctx *gin.Context) {
	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	if name == "" || passwd == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "账号密码不能为空",
			"data":   nil,
		})
		return
	}
	var id int
	var hassPwd string
	var role int
	err := DB.QueryRow("select id,passwd,role from users where name=?", name).Scan(&id, &hassPwd, &role)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -2,
			"msg":    "账号或密码错误",
			"data":   nil,
		})
		return
	}
	if !security.ComparePasswd([]byte(passwd), hassPwd) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": -2,
			"msg":    "账号或密码错误",
			"data":   nil,
		})
		return
	}
	token, err := auth.GenerateToken(name, role)
	if err != nil {
		fmt.Println(err)
	}
	ctx.SetCookie("jwt", token, 3600, "/", "", false, true)
	returnPath := ctx.Query("return")
	if returnPath == "" {
		returnPath = "/"
	}
	if !security.CheckUrl(returnPath) {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": -2,
			"msg":    "非法跳转地址",
			"data":   nil,
		})
		return
	}
	ctx.HTML(http.StatusFound, "redirect.html", gin.H{
		"redirectUrl": returnPath,
	})
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"status": 1,
	// 	"msg":    "成功登录",
	// 	"data":   nil,
	// })
}
