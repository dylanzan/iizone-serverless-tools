package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"iizone-serverless-tools/model/codex"
	"iizone-serverless-tools/model/requestx"
	"iizone-serverless-tools/model/responsex"
	"iizone-serverless-tools/utilx/base64x"
	"io"
	"net/http"
)

type IIZoneTools struct {
}

// PingHandler 接口连通性测试
func (i *IIZoneTools) PingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "pong")
}

// IIZoneToolsBase64Decode 解码
func (i *IIZoneTools) IIZoneToolsBase64Decode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var reqBody requestx.RequestBody

	reqBodyBytes, err := io.ReadAll(r.Body)
	//base64 解码
	if err != nil {
		responsex.ResponseHandler(w, codex.ServiceInsideError, err)
		return
	} else {
		err = json.Unmarshal(reqBodyBytes, &reqBody)
		if err != nil {
			responsex.ResponseHandler(w, codex.ServiceInsideError, err)
			return
		}
	}

	if !base64x.JudgeBase64(reqBody.Data) {
		responsex.ResponseHandler(w, codex.RequestParamError, err)
		return
	}

	res, err := base64.StdEncoding.DecodeString(reqBody.Data)
	if err != nil {
		responsex.ResponseHandler(w, codex.ServiceInsideError, err)
		return
	}

	responsex.ResponseHandler(w, codex.Success, string(res))
}

// IIZoneToolsBase64Encode 编码
func (i *IIZoneTools) IIZoneToolsBase64Encode(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var reqBody requestx.RequestBody
	if reqBodyBytes, err := io.ReadAll(r.Body); err != nil {
		responsex.ResponseHandler(w, codex.ServiceInsideError, err)
		return
	} else {
		err = json.Unmarshal(reqBodyBytes, &reqBody)
		if err != nil {
			responsex.ResponseHandler(w, codex.ServiceInsideError, err)
			return
		}
	}

	res := base64.StdEncoding.EncodeToString([]byte(reqBody.Data))
	responsex.ResponseHandler(w, codex.Success, res)
}
