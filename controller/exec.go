package controller

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func ExecVuln1(ctx *gin.Context) {
	path := ctx.Query("path")
	if path == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "give me a param `path` ",
		})
		return
	}
	cmd := exec.Command("bash", "-c", "ls"+" -alh "+path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "system wrong",
		})
		return
	}
	ctx.Writer.WriteString(string(output))
}
