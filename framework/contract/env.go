package contract

const (
	// ENV_PRODUCTION 代表生产环境
	ENV_PRODUCTION = "production"
	// ENV_TESTING 代表测试环境
	ENV_TESTING = "testing"
	// ENV_DEVELOPMENT 代表开发环境
	ENV_DEVELOPMENT = "development"

	// ENV_KEY 是环境变量服务字符串凭证
	ENV_KEY = "gocore:env"
)

// Env 定义环境变量的获取服务
type Env interface {
	// AppEnv 获取当前的环境，建议分为developent/testing/production
	AppEnv() string

	// IsExist 判断一个环境变量是否被设置
	IsExist(string) bool

	// Get 获取环境变量, 如果没有设置，返回""
	Get(string) string

	// All 获取所有的的环境变量
	All() map[string]string
}
