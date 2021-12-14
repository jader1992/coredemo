package app

import (
	"errors"
	"flag"
	"github.com/jader1992/gocore/framework"
	"github.com/jader1992/gocore/framework/util"
	"path/filepath"
)

// GocoreApp 代表gocore框架的app实现
type GocoreApp struct {
	container framework.Container // 服务容器
	baseFolder string             // 基础路径
}

// Version 实现版本
func (g GocoreApp) Version()  string{
	return "0.0.3"
}

func (g GocoreApp) BaseFolder() string {
	if g.baseFolder != "" {
		return g.baseFolder
	}

	// 如果没有设置，则使用参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
	flag.Parse()

	if baseFolder != "" {
		return baseFolder
	}

	// 如果参数也没有。表示默认的当前路径
	return util.GetExecDirectory()
}

// ConfigFolder 表示配置文件地址
func (g GocoreApp) ConfigFolder() string {
	return filepath.Join(g.BaseFolder(), "config")
}

func (g GocoreApp) StorageFolder() string {
	return filepath.Join(g.BaseFolder(), "storage")
}

func (g GocoreApp) LogFolder() string {
	return filepath.Join(g.StorageFolder(), "log")
}

func (g GocoreApp) RuntimeFolder() string {
	return filepath.Join(g.StorageFolder(), "runtime")
}

func (g GocoreApp) HttpFolder() string {
	return filepath.Join(g.BaseFolder(), "http")
}

func (g GocoreApp) ConsoleFolder() string {
	return filepath.Join(g.BaseFolder(), "console")
}

func (g GocoreApp) CommandFolder() string {
	return filepath.Join(g.ConsoleFolder(), "command")
}

func (g GocoreApp) MiddlewareFolder() string {
	return filepath.Join(g.BaseFolder(), "middleware")
}

func (g GocoreApp) TestFolder() string {
	return filepath.Join(g.BaseFolder(), "test")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (h GocoreApp) ProviderFolder() string {
	return filepath.Join(h.BaseFolder(), "provider")
}

func NewGocoreApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	// 有两个参数
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &GocoreApp{baseFolder: baseFolder, container: container}, nil
}












