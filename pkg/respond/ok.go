package respond

import (
	"github.com/gofiber/fiber/v2"
)

// Response 定义一个标准响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// Ok 函数返回一个成功响应
// 返回书只能是一个值，多多数封装到一个结构体
// 入参为多个数据时只会返回第一个
func Ok(c *fiber.Ctx, data ...any) error {
	if len(data) > 0 {
		return respond(c, 0, "操作成功!", data[0])
	}
	return respond(c, 0, "操作成功!", nil)
}

// OkData Ok 函数返回一个成功响应
//func OkData(c *fiber.Ctx, data interface{}) error {
//	return respond(c, 0, "操作成功！", data)
//}

// 创建一个辅助函数来生成响应
func respond(c *fiber.Ctx, code int, msg string, data interface{}) error {
	return c.JSON(Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
