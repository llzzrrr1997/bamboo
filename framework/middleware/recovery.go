package middleware

import (
	"github.com/llzzrrr1997/bamboo/framework"
	"net/http"
)

// recovery机制，将协程中的函数异常进行捕获
func Recovery() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				_ = c.SetStatus(http.StatusInternalServerError).Json(err)
			}
		}()
		// 使用next执行具体的业务逻辑
		_ = c.Next()

		return nil
	}
}