package gin

import "github.com/jader1992/gocore/framework/contract"

// MustMakeApp 从容器中获取App服务
func (c *Context) MustMakeApp() contract.App {
	return c.MustMake(contract.APP_KEY).(contract.App)
}

// MustMakeKernel 从容器中获取Kernel服务
func (c *Context) MustKernel() contract.Kernel {
	return c.MustMake(contract.KERNEL_KEY).(contract.Kernel)
}
