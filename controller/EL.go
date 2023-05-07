package controller

import (
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
	"github.com/gin-gonic/gin"
)

// 没完成
func EL(ctx *gin.Context) {
	expression, err := govaluate.NewEvaluableExpression("system(command)")
	if err != nil {
		fmt.Println(err)
		return
	}
	parameters := make(map[string]interface{})
	parameters["command"] = "ls"
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result) // 输出 ls 的结果
	ctx.Status(http.StatusOK)
}
