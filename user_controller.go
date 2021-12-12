package main

import (
	"gocore/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
