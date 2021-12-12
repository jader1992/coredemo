package main

import (
	"context"
	"fmt"
	"gocore/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	// finish 负责通知结束
	finish := make(chan struct{}, 1)

	// panicChan 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1 * time.Second))
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		// Do read action
		time.Sleep(10 * time.Millisecond)
		_ = c.SetOkStatus().Json("ok")

		finish <- struct{}{}
	}()

	// 3: 监听三个事件：异常事件、结束事件、超时事件
	select {
	case p := <-panicChan:
		c.WriteMux().Lock() // 对c.ResponseWrite进行保护
		defer c.WriteMux().Unlock()

		log.Println(p)
		_ = c.SetStatus(500).Json("panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationCtx.Done():
		c.WriteMux().Lock()
		defer c.WriteMux().Unlock()

		_ = c.SetStatus(500).Json("time out")
		c.SetHasTimeout()
	}
	return nil
}

// func Foo(request *http.Request, response http.ResponseWriter) {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errMsg": "inner error",
// 		"data":   nil,
// 	}

// 	response.Header().Set("Content-Type", "application/json")

// 	foo := request.PostFormValue("foo")
// 	if foo == "" {
// 		foo = "10"
// 	}
// 	fooInt, err := strconv.Atoi(foo)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	obj["data"] = fooInt
// 	byt, err := json.Marshal(obj)
// 	if err != nil {
// 		response.WriteHeader(500)
// 		return
// 	}
// 	response.WriteHeader(200)
// 	response.Write(byt)
// 	return
// }

// func Foo2(ctx *framework.Context) error {
// 	obj := map[string]interface{}{
// 		"errno":  50001,
// 		"errMsg": "inner error",
// 		"data":   nil,
// 	}

// 	fooInt := ctx.FormInt("foo", 10)
// 	obj["data"] = fooInt
// 	return ctx.Json(http.StatusOK, obj)
// }

// func Foo3(ctx *framework.Context) error {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	return rdb.Set(ctx, "key", "value", 0).Err()
// }


