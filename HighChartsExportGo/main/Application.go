package main

import (
	"net/http"
	"../logger"
	"../utils"
	"../handlers"
)

func main() {
	logger.Info.Printf("Server is running ...\n")
	// url配置
	http.HandleFunc("/export", handlers.ExportHandler)
	// 端口监听
	httpServerErr := http.ListenAndServe(utils.ServerIPPort, nil)
	if httpServerErr != nil {
		logger.Error.Printf("error is: %s", httpServerErr)
	}
}
