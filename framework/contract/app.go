package contract

const APP_KEY = "gocore:app"

// App 定义接口
type App interface {
	// AppId 表示当前这个app的唯一id，可以用于分布式锁
	AppId() string
	// Version 定义当前版本
	Version() string
	// BaseFolder 定义项目基础地址
	BaseFolder() string
	// ConfigFolder 定义配置文件的地址
	ConfigFolder() string
	// LogFolder 定义了日志所在路径
	LogFolder() string
	// ProviderFolder 定义业务自己定义的服务提供者位置
	ProviderFolder() string
	// MiddlewareFolder 定义业务自己定义的中间件
	MiddlewareFolder() string
	// CommandFolder 定义业务定义的命令
	CommandFolder() string
	// RuntimeFolder 定义业务运行中间件信息
	RuntimeFolder() string
	// TestFolder 存放测试所需要的信息
	TestFolder() string

	// LoadAppConfig 加载新的AppConfig，key为对应的函数转为小写下划线，比如ConfigFolder => config_folder
	LoadAppConfig(kv map[string]string)
}
