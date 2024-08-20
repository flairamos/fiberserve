// 主程序入口
package main

import (
	"sunny/config"
	"sunny/pkg/db"
	"sunny/pkg/log"
	"sunny/router"
)

func main() {

	// 实例化项目配置
	config.Init()

	// 实例化 zap 日志
	log.Init()

	// 实例化 GORM
	db.Init()

	// 实例化 router，这个一定要放在最后
	router.Init()
}
