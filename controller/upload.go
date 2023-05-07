package controller

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

var saveDir string = os.Getenv("tmp")

// UploadVuln1 单文件  任意文件上传   目录穿越
func UploadVuln1(ctx *gin.Context) {
	file, err := ctx.FormFile("files")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "status.html", gin.H{
			"message": "can't find file!",
		})
		return
	}
	savePath := filepath.ToSlash(saveDir + "\\" + file.Filename)
	//dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
	// 上传文件到指定的目录
	ctx.SaveUploadedFile(file, savePath)
	ctx.HTML(http.StatusOK, "status.html", gin.H{
		"message": fmt.Sprintf("files save in %s", savePath),
	})

}

// UploadVuln2 多文件  任意文件上传   目录穿越
func UploadVuln2(ctx *gin.Context) {
	fom, _ := ctx.MultipartForm()
	files := fom.File["files"]
	for _, f := range files {
		savePath := filepath.ToSlash(saveDir + "\\" + f.Filename)
		err := ctx.SaveUploadedFile(f, savePath)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "status.html", gin.H{
				"message": "file upload fail!",
			})
			return
		}
		ctx.HTML(http.StatusOK, "status.html", gin.H{
			"message": fmt.Sprintf("files save in %s", savePath),
		})
	}

}

func UploadImg(ctx *gin.Context) {
	file, err := ctx.FormFile("files")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	whilteList := []string{".jpg", ".png", ".jpeg", ".gif", ".bmp", ".ico"}
	fg := false
	for _, e := range whilteList {
		if strings.HasSuffix(file.Filename, e) {
			fg = true
			break
		}
	}
	if !fg {
		ctx.HTML(http.StatusForbidden, "status.html", gin.H{
			"message": "only can upload img",
		})
		return
	}
	mimeTypeBlackList := []string{
		"text/html",
		"text/javascript",
		"application/javascript",
		"application/ecmascript",
		"text/xml",
		"application/xml",
	}
	for _, m := range mimeTypeBlackList {
		if file.Header.Get("Content-Type") == m {

			ctx.HTML(http.StatusForbidden, "status.html", gin.H{
				"message": "only can upload img",
			})
			return
		}
	}
	fmt.Println(file.Header.Get("Content-Type"))
	//内容判断
	f, _, err := ctx.Request.FormFile("files")
	if err != nil {
		ctx.HTML(http.StatusForbidden, "status.html", gin.H{
			"message": "only can upload img",
		})
		return
	}
	defer f.Close()
	cont := make([]byte, 512)
	f.Read(cont)
	contentType := http.DetectContentType(cont) //返回的是mime类型
	fmt.Println(contentType)
	for _, m := range mimeTypeBlackList {
		if contentType == m {

			ctx.HTML(http.StatusForbidden, "status.html", gin.H{
				"message": "only can upload img",
			})
			return
		}
	}

	savePath := filepath.ToSlash(saveDir + "\\" + file.Filename)
	//dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
	// 上传文件到指定的目录
	ctx.SaveUploadedFile(file, savePath)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("files save in %s", savePath),
	})

}

// UploadZipVuln   zip slip
func UploadZipVuln(ctx *gin.Context) {
	file, err := ctx.FormFile("files")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	zipList := []string{".zip", ".tar"}
	fg := false
	for _, e := range zipList {
		if strings.HasSuffix(file.Filename, e) {
			fg = true
			break
		}
	}
	if !fg {
		ctx.HTML(http.StatusForbidden, "status.html", gin.H{
			"message": "only can upload zip",
		})
		return
	}
	savePath := filepath.ToSlash(saveDir + "\\" + file.Filename)
	// 保存
	err = ctx.SaveUploadedFile(file, savePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	//解压
	r, err := zip.OpenReader(savePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}
		rd, _ := f.Open()
		com, _ := ioutil.ReadAll(rd)
		p, _ := filepath.Abs(f.Name) //mian code
		ioutil.WriteFile(p, com, 0666)
		ctx.Writer.WriteString("save file in \n" + p)
	}
}
