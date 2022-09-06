package main

import (
	"iizone-serverless-tools/function"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	var iizoneTools *function.IIZoneTools

	router := httprouter.New()
	router.GET("/ping", iizoneTools.PingHandler)
	router.POST("/api/iizone-tools/base64", iizoneTools.IIZoneToolsBase64Decode)

	http.ListenAndServe(":9000", router)

}