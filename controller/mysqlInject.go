package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MysqlVuln1 Query简单拼接
// http://localhost:8080/mysql/1?name=admin' or '1'='1
func MysqlVuln1(ctx *gin.Context) {
	name := ctx.Query("name")
	query := fmt.Sprintf("select * from users where name = '%s'", name)
	//query:="select * from users where name = '" + name + "'"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	defer rows.Close()
	var (
		id       int
		username string
		passwd   string
		role     int
	)
	for rows.Next() {
		err := rows.Scan(&id, &username, &passwd, &role)
		if err != nil {
			ctx.JSON(http.StatusBadGateway, "")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"id":     id,
			"name":   username,
			"passwd": passwd,
			"role":   role,
		})
	}

}

// MysqlVuln2 QueryRow简单拼接
func MysqlVuln2(ctx *gin.Context) {
	name := ctx.Query("name")
	query := fmt.Sprintf("select * from users where name = '%s'", name)
	//query:="select * from users where name = '" + name + "'"
	var (
		id       int
		username string
		passwd   string
		role     int
	)
	err := DB.QueryRow(query).Scan(&id, &username, &passwd, &role)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   username,
		"passwd": passwd,
		"role":   role,
	})

}

// MysqlSec1 参数化查询
func MysqlSec1(ctx *gin.Context) {
	name := ctx.Query("name")
	var (
		id       int
		username string
		passwd   string
		role     int
	)
	err := DB.QueryRow("select * from users where name = ?", name).Scan(&id, &username, &passwd, &role)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   username,
		"passwd": passwd,
		"role":   role,
	})
}

// MysqlSec2 预编译，只适用与没有返回值的  如插入、更新
func MysqlSec2(ctx *gin.Context) {
	name := ctx.Query("name")
	dbpre, err := DB.Prepare("select * from users where name = ?")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	result, err := dbpre.Exec(name)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	if row, err := result.RowsAffected(); row != 1 || err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, "")
		return
	}
	ctx.JSON(http.StatusCreated, "")

}
