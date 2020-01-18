package controller

import (
	"fastIM/app/util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	os.MkdirAll("./resource", os.ModePerm)
}

func FileUpload(writer http.ResponseWriter, request *http.Request) {
	UploadLocal(writer, request)
}

//将文件存储在本地/im_resource目录下
func UploadLocal(writer http.ResponseWriter, request *http.Request) {
	//获得上传源文件
	srcFile, head, err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	//创建一个新的文件
	suffix := ".png"
	srcFilename := head.Filename
	splitMsg := strings.Split(srcFilename, ".")
	if len(splitMsg) > 1 {
		suffix = "." + splitMsg[len(splitMsg)-1]
	}
	filetype := request.FormValue("filetype")
	if len(filetype) > 0 {
		suffix = filetype
	}
	filename := fmt.Sprintf("%d%s%s", time.Now().Unix(), util.GenRandomStr(32), suffix)
	//创建文件
	filepath := "./resource/" + filename
	dstfile, err := os.Create(filepath)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}
	//将源文件拷贝到新文件
	_, err = io.Copy(dstfile, srcFile)
	if err != nil {
		util.RespFail(writer, err.Error())
		return
	}

	util.RespOk(writer, filepath, "")
}




