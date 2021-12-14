package cobra

import "github.com/jader1992/gocore/framework"

// SetContainer 设置服务容器的方法是为了在创建根 Command 之后，能将服务容器设置到里面去
func (c *Command) SetContainer(container framework.Container)  {
	c.container = container
}

// GetContainer 获取容器是为了在执行命令的 RunE 函数的时候，能从参数 Command 中获取到服务容器
func (c *Command) GetContainer() framework.Container  {
	return c.Root().container
}
