package demo

import (
	"fmt"
	"github.com/jader1992/gocore/framework"
)

// DemoServiceProvider 服务提供方
type DemoServiceProvider struct {

}

// Name方法直接将服务对应的字符串凭证返回，在这个例子中就是“jade.demo"
func (sp *DemoServiceProvider) Name() string {
	return KEY
}

// Register方法是注册初始化服务实例的方法，我们这里先暂定为NewDemoService
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance  {
	return NewDemoService
}

// IsDefer方法表示是否延迟实例化，我们这里设置为true，将这个服务的实例化延迟到第一次make的时候
func (sp *DemoServiceProvider) IsDefer() bool{
	return true
}

func (sp *DemoServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (sp *DemoServiceProvider) Boot(c framework.Container) error {
	fmt.Println("demo service boot")
	return nil
}



