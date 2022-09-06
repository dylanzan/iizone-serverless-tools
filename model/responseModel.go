package model

import (
	"encoding/json"
	"iizone-serverless-tools/model/codex"
	"net/http"
)

// ResponseBody： 响应body
type ResponseBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseHandler(w http.ResponseWriter, respStatus *codex.Code, data interface{}) {
	respBody := ResponseBody{
		Code: respStatus.Status,
		Msg:  respStatus.Message,
		Data: data,
	}
	respBodyJson, _ := json.Marshal(respBody)
	w.Write(respBodyJson)
}
