// Package router 路由
package router

import (
	"strconv"
	"sunny/config"
	"sunny/pkg/respond"

	"github.com/gofiber/fiber/v2"
)

func Init() {

	app := fiber.New(fiber.Config{
		ErrorHandler: respond.ErrorHandler, // 设置全局错误处理函数
	})

	// JWT身份验证中间件
	// app.Use(func(c *fiber.Ctx) error {

	// 	// 全局捕获信息并返回，注意：在此后执行的方法才会捕获，如果有此之前执行的方法，想捕获时，把此方法放在要捕获的方法之前，或者复制一份过去
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			_ = c.JSON(r)
	// 		}
	// 	}()

	// 	url := c.OriginalURL()
	// 	avoidLogin := []string{"/admin/login"} // 免登录

	// 	// 排序数组
	// 	sort.Strings(avoidLogin)

	// 	// 使用二分查找
	// 	search := sort.SearchStrings(avoidLogin, url)

	// 	if search < len(avoidLogin) && avoidLogin[search] == url {
	// 		return c.Next()
	// 	}

	// 	tokenStr := c.Get("token")

	// 	if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
	// 		return respond.ErrorCode(respond.TokenExpire, tokenStr)
	// 	}

	// 	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	// 	mapClaims := Jwt.MapClaims{}

	// 	token, err := Jwt.ParseWithClaims(tokenStr, mapClaims, func(token *Jwt.Token) (interface{}, error) {
	// 		return config.JwtKeyVar, nil
	// 	})

	// 	if err != nil || !token.Valid {
	// 		return respond.ErrorCode(respond.TokenExpire, respond.TokenExpireMsg)
	// 	}

	// 	config.UserJwt = mapClaims

	// 	return c.Next()
	// })

	// 此处添加internal中的路由

	err := app.Listen(":" + strconv.Itoa(config.ServerVar.Port))

	if err != nil {
		panic(err)
	}
}
