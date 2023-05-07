package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"
)

func GenZip() {

	//创建压缩包
	z, _ := os.Create("a.zip")
	zipw := zip.NewWriter(z)
	//创建文件夹
	_, err := zipw.Create("a\\")
	if err != nil {
		fmt.Println(err)
	}
	//创建文件
	w1, _ := zipw.Create("a\\..\\..\\a.txt")
	io.Copy(w1, strings.NewReader("sssss"))
	zipw.Close()
	z.Close()

}
