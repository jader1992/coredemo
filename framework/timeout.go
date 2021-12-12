package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

// TimeoutHandler 超时的中间件
func TimeoutHandler(fun ControllerHandler, d time.Duration) ControllerHandler  {
	return func(c *Context) error{
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()

		c.request.WithContext(durationCtx)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			// 执行具体的逻辑
			_ = fun(c)

			finish <- struct{}{}
		}()

		// 执行业务逻辑后操作
		select {
		case p := <-panicChan:
			log.Println(p)
			c.responseWriter.WriteHeader(500)
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.SetHasTimeout()
			_, _ = c.responseWriter.Write([]byte("time out"))
		}
		return nil
	}
}
