package handlers

import (
	"net/http"
	"fmt"
	"../utils"
	"../logger"
	"strings"
	"os/exec"
	"encoding/json"
)

type Result struct {
	FilePath string
}

// 导出图片
func ExportHandler(responseWriter http.ResponseWriter, request *http.Request)  {
	responseWriter.Header().Set("Content-Type", "application/json")
	logger.Info.Printf("%s %s", request.Method, request.URL)
	request.ParseForm()
	fileName := request.Form.Get("name")
	filePath := fmt.Sprintf("%s/%s", utils.GetCurrentDirectory(), fileName)
	contentLength := request.ContentLength
	requestContent := make([]byte, contentLength)
	request.Body.Read(requestContent)
	utils.WriteToFile(filePath, requestContent)
	command := fmt.Sprintf("./phantomjs_linux  highcharts-convert.js -infile %s.json -outfile %s.png -scale 2.5 -width 800 -constr Chart -resources highcharts.js,jquery.js", filePath, filePath)
	parts := strings.Fields(command)
	head := parts[0]
	parts = parts[1:]
	_, cmdErr := exec.Command(head,parts...).Output()
	if cmdErr != nil {
		fmt.Printf("%s", cmdErr)
	}
	result := Result{
		FilePath:fmt.Sprintf("%s.png", filePath),
	}
	jsonResult, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		fmt.Fprintf(responseWriter, jsonErr.Error())
	}
	fmt.Fprintf(responseWriter, string(jsonResult))
}
