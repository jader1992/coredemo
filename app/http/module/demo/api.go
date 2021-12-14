package demo

import (
	demoService "github.com/jader1992/gocore/app/provider/demo"
	"github.com/jader1992/gocore/framework/contract"
	"github.com/jader1992/gocore/framework/gin"
)

func Register(r *gin.Engine) error {
	// 绑定demoPrivider提供者
	// r.Bind(&demoService.DemoProvider{})

	// 注册路由
	api := NewDemoApi()
	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo_post", api.DemoPost)
	return nil
}

// DemoApi 测试api的提供者
type DemoApi struct {
	service *Service  // 嵌套了与user方法的service
}

// NewDemoApi 初始化DemoApi
func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo godoc
// @Summary 获取所有用户
// @Description 获取所有用户
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo [get]
func (api *DemoApi) Demo(c *gin.Context)  {
	// 获取app服务提供者
	appService := c.MustMake(contract.APP_KEY).(contract.App)
	// 获取项目基础目录
	baseFolder := appService.BaseFolder()
	c.JSON(200, baseFolder)
}

// Demo godoc
// @Summary 获取所有学生
// @Description 获取所有学生
// @Produce  json
// @Tags demo
// @Success 200 array []UserDTO
// @Router /demo/demo2 [get]
func (api *DemoApi) Demo2(c *gin.Context)  {
	// 获取demo服务的提供者
	demoProvide := c.MustMake(demoService.DEMO_KEY).(demoService.IService)
	//
	students := demoProvide.GetAllStudent()
	usersDto := StudentsToUsersDTOs(students)
	c.JSON(200, usersDto)
}

func (api *DemoApi) DemoPost(c *gin.Context)  {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

