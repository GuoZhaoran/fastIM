package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//定义一个结构体
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

//失败的返回结果
func RespFail(writer http.ResponseWriter, msg string) {
	Resp(writer, -1, nil, msg)
}

//返回成功
func RespOk(writer http.ResponseWriter, data interface{}, msg string) {
	Resp(writer, 0, data, msg)
}

//返回列表数据
func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	RespList(w,0,lists,total)
}

//返回列表
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {
	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	//设置header 为JSON 默认是test/html,所以特别指出返回的数据类型为application/json
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	rep := ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	//将结构体转化为json字符串
	ret, err := json.Marshal(rep)
	if err != nil {
		log.Panicln(err.Error())
	}

	//返回json ok
	writer.Write(ret)
}

