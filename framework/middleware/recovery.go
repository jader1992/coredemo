package middleware

import "gocore/framework"

// Recovery 将函数异常进行捕获
func Recovery() framework.ControllerHandler  {
	// 使用函数回调
	return func(c *framework.Context) error {
		defer func() {
			if err := recover(); err != nil {
				_ = c.SetStatus(500).Json(err)
			}
		}()

		// 使用next执行具体的业务逻辑
		_ = c.Next()

		return nil
	}
}