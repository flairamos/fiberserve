package respond

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

const (
	TokenExpire    = 1      // token过期
	TokenExpireMsg = "登录过期" // token过期
)

// ErrorData 定义自定义错误类型
type ErrorData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

// Error 错误返回值
func Error(c *fiber.Ctx, msg string) error {
	return respond(c, 0, msg, nil)
}

// ErrorCode 返回错误信息自定义错误码
func ErrorCode(c *fiber.Ctx, code int, msg string) error {
	return respond(c, code, msg, nil)
}

// 实现 error 接口的方法
func (e *ErrorData) Error() string {
	return e.Msg
}

// ErrorHandler 错误处理中间件
func ErrorHandler(c *fiber.Ctx, err error) error {

	var data *ErrorData

	if errors.As(err, &data) {
		return respond(c, data.Code, data.Msg, nil)
	}

	// 默认错误处理
	return respond(c, 500, err.Error(), nil)
}
