package utils

import (
	"path/filepath"
	"os"
	"../logger"
	"strings"
	"io/ioutil"
	"fmt"
)

//获取项目当前路径
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Error.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

//将字节数组写入文件
func WriteToFile(filePath string, data []byte) (error) {
	jsonPath := fmt.Sprintf("%s.json", filePath)
	err := ioutil.WriteFile(jsonPath, data, 0666)
	return err
}
