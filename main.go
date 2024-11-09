package main

import (
	"fmt"
	"os"
	"shoollove/config"
	minioutil "shoollove/pkg/minio_util"
)

func main() {
	// 获取文件地址
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("----获取文件当前目录失败---", err.Error())
		return
	}
	// 初始化配置
	config.Init(path)
	// 初始化minio
	minioutil.Init()
}
