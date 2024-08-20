package serr

// code为业务错误码，用于隐藏错误信息 开发者自定义，如 10000001--Not Found
// msg为错误信息，用于给用户展示，如 "用户不存在"
// 该类为error的子类使用断言解析具体的类型
type Serr struct {
	Code string
	Msg  string
}

func (s Serr) Error() string {
	return s.Msg
}
