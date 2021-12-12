package middleware

import (
	"fmt"
	"gocore/framework"
)

func Test1() framework.ControllerHandler  {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware after test1")
		return nil
	}
}

func Test2() framework.ControllerHandler  {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware after test2")
		return nil
	}
}

func Test3() framework.ControllerHandler  {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		c.Next()
		fmt.Println("middleware after test3")
		return nil
	}
}
