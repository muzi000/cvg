package controller

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	if err := InitDB(); err != nil {
		fmt.Println(err)
		return
	}
	rows, err := DB.Query("select id,passwd,role from users where user=?", "admin")
	if err != nil {
		fmt.Println(err)
		return
	}
	var id int
	var hassPwd string
	var role int
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &hassPwd, &role)
	}

}
