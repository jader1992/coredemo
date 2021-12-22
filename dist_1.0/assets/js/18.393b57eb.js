(window.webpackJsonp=window.webpackJsonp||[]).push([[18],{484:function(e,r,t){"use strict";t.r(r);var n=t(61),a=Object(n.a)({},(function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"服务提供者"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#服务提供者"}},[e._v("#")]),e._v(" 服务提供者")]),e._v(" "),t("h2",{attrs:{id:"指南"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#指南"}},[e._v("#")]),e._v(" 指南")]),e._v(" "),t("p",[e._v("gocore框架使用ServiceProvider机制来满足协议，通过service Provder提供某个协议服务的具体实现。这样如果开发者对具体的实现协议的服务类的具体实现不满意，则可以很方便的通过切换具体协议的Service Provider来进行具体服务的切换。")]),e._v(" "),t("p",[e._v("一个ServiceProvider是一个单独的文件夹，它包含服务提供和服务实现。具体可以参考framework/provider/demo")]),e._v(" "),t("p",[e._v("一个SerivceProvider就是一个独立的包，这个包可以作为插件独立地发布和分享。")]),e._v(" "),t("p",[e._v("你也可以定义一个无contract的ServiceProvider，其中的Name()需要保证唯一。")]),e._v(" "),t("h2",{attrs:{id:"创建"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#创建"}},[e._v("#")]),e._v(" 创建")]),e._v(" "),t("p",[e._v("我们可以使用命令 "),t("code",[e._v("./gocore provider new")]),e._v(" 来创建一个新的service provider")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v("[~/Documents/workspace/gocore_workspace]$ ./gocore provider new\ncreate a provider\n? please input provider name test\n? please input provider folder(default: provider name):\ncreate provider success, folder path: /Users/Documents/workspace/gocore_workspace/app/provider/test\nplease remember add provider to kernel\n")])])]),t("p",[e._v("该命令会在"),t("code",[e._v("app/provider/")]),e._v(" 目录下创建一个对应的服务提供者。并且初始化好三个文件： "),t("code",[e._v("contract.go")]),e._v(", "),t("code",[e._v("provider.go")]),e._v(", "),t("code",[e._v("service.go")])]),e._v(" "),t("h2",{attrs:{id:"自定义"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#自定义"}},[e._v("#")]),e._v(" 自定义")]),e._v(" "),t("p",[e._v("我们需要编写这三个文件：")]),e._v(" "),t("h3",{attrs:{id:"contract-go"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#contract-go"}},[e._v("#")]),e._v(" contract.go")]),e._v(" "),t("p",[e._v("contract.go 定义了这个服务提供方提供的协议接口。gocore 框架任务，作为一个业务的服务提供者，定义一个好的协议是最重要的事情。")]),e._v(" "),t("p",[e._v("所以 contract.go 中定义了一个 Service 接口，在其中定义各种方法，包含输入参数和返回参数。")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('package demo\n\nconst DemoKey = "demo"\n\ntype IService interface {\n\tGetAllStudent() []Student\n}\n\ntype Student struct {\n\tID   int\n\tName string\n}\n\n')])])]),t("p",[e._v("其中还定义了一个Key， 这个 Key 是全应用唯一的，服务提供者将服务以 Key 关键字注入到容器中。服务使用者使用 Key 关键字获取服务。")]),e._v(" "),t("h3",{attrs:{id:"provider"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#provider"}},[e._v("#")]),e._v(" provider")]),e._v(" "),t("p",[e._v("provider.go 提供服务适配的实现，实现一个Provider必须实现对应的五个方法")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('package demo\n\nimport (\n\t"github.com/gogocore/gocore/framework"\n)\n\ntype DemoProvider struct {\n\tframework.ServiceProvider\n\n\tc framework.Container\n}\n\nfunc (sp *DemoProvider) Name() string {\n\treturn DemoKey\n}\n\nfunc (sp *DemoProvider) Register(c framework.Container) framework.NewInstance {\n\treturn NewService\n}\n\nfunc (sp *DemoProvider) IsDefer() bool {\n\treturn false\n}\n\nfunc (sp *DemoProvider) Params() []interface{} {\n\treturn []interface{}{sp.c}\n}\n\nfunc (sp *DemoProvider) Boot(c framework.Container) error {\n\tsp.c = c\n\treturn nil\n}\n')])])]),t("ul",[t("li",[e._v("Name() // 指定这个服务提供者提供的服务对应的接口的关键字")]),e._v(" "),t("li",[e._v("Register() // 这个服务提供者注册的时候调用的方法，一般是指定初始化服务的函数名")]),e._v(" "),t("li",[e._v("IsDefer() // 这个服务是否是使用时候再进行初始化，false为注册的时候直接进行初始化服务")]),e._v(" "),t("li",[e._v("Params() // 初始化服务的时候对服务注入什么参数，一般把 container 注入到服务中")]),e._v(" "),t("li",[e._v("Boot() // 初始化之前调用的函数，一般设置一些全局的Provider")])]),e._v(" "),t("h3",{attrs:{id:"service-go"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#service-go"}},[e._v("#")]),e._v(" service.go")]),e._v(" "),t("p",[e._v("service.go提供具体的实现，它至少需要提供一个实例化的方法 "),t("code",[e._v("NewService(params ...interface{}) (interface{}, error)")]),e._v("。")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('package demo\n\nimport "github.com/gogocore/gocore/framework"\n\ntype Service struct {\n\tcontainer framework.Container\n}\n\nfunc NewService(params ...interface{}) (interface{}, error) {\n\tcontainer := params[0].(framework.Container)\n\treturn &Service{container: container}, nil\n}\n\nfunc (s *Service) GetAllStudent() []Student {\n\treturn []Student{\n\t\t{\n\t\t\tID:   1,\n\t\t\tName: "foo",\n\t\t},\n\t\t{\n\t\t\tID:   2,\n\t\t\tName: "bar",\n\t\t},\n\t}\n}\n\n')])])]),t("h2",{attrs:{id:"注入"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#注入"}},[e._v("#")]),e._v(" 注入")]),e._v(" "),t("p",[e._v("gocore的路由，controller的定义是选择基于gin框架进行扩展的。所有的gin框架的路由，参数获取，验证，context都和gin框架是相同的。唯一不同的是gin的全局路由gin.Engine实现了gocore的容器结构，可以对gin.Engine进行服务提供的注入，且可以从context中获取具体的服务。")]),e._v(" "),t("p",[e._v("gocore提供两种服务注入的方法：")]),e._v(" "),t("ul",[t("li",[e._v("Bind: 将一个ServiceProvider绑定到容器中，可以控制其是否是单例")]),e._v(" "),t("li",[e._v("Singleton: 将一个单例ServiceProvider绑定到容器中")])]),e._v(" "),t("p",[e._v("建议在文件夹 "),t("code",[e._v("app/provider/kernel.go")]),e._v(" 中进行服务注入")]),e._v(" "),t("div",{staticClass:"language-golang extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v("func RegisterCustomProvider(c framework.Container) {\n\tc.Bind(&demo.DemoProvider{}, true)\n}\n")])])]),t("p",[e._v("当然你也可以在某个业务模块路由注册的时候进行服务注入")]),e._v(" "),t("div",{staticClass:"language-golang extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('func Register(r *gin.Engine) error {\n\tapi := NewDemoApi()\n\tr.Container().Singleton(&demoService.DemoProvider{})\n\n\tr.GET("/demo/demo", api.Demo)\n\tr.GET("/demo/demo2", api.Demo2)\n\treturn nil\n}\n')])])]),t("h2",{attrs:{id:"获取"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#获取"}},[e._v("#")]),e._v(" 获取")]),e._v(" "),t("p",[e._v("gocore提供了三种服务获取的方法：")]),e._v(" "),t("ul",[t("li",[e._v("Make: 根据一个Key获取服务，获取不到获取报错")]),e._v(" "),t("li",[e._v("MustMake: 根据一个Key获取服务，获取不到返回空")]),e._v(" "),t("li",[e._v("MakeNew: 根据一个Key获取服务，每次获取都实例化，对应的ServiceProvider必须是以非单例形式注入")])]),e._v(" "),t("p",[e._v("你可以在任意一个可以获取到 container 的地方进行服务的获取。")]),e._v(" "),t("p",[e._v("业务模块中:")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v("func (api *DemoApi) Demo2(c *gin.Context) {\n\tdemoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)\n\tstudents := demoProvider.GetAllStudent()\n\tusersDTO := StudentsToUserDTOs(students)\n\tc.JSON(200, usersDTO)\n}\n")])])]),t("p",[e._v("命令行中：")]),e._v(" "),t("div",{staticClass:"language-golang extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('var CenterCommand = &cobra.Command{\n\tUse:   "direct_center",\n\tShort: "计算区域中心点",\n\tRunE: func(c *cobra.Command, args []string) error {\n\t\tcontainer := util.GetContainer(c.Root())\n\t\tapp := container.MustMake(contract.AppKey).(contract.App)\n        return nil\n    }\n')])])]),t("p",[e._v("甚至于另外一个服务提供者中：")]),e._v(" "),t("div",{staticClass:"language-golang extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v('type Service struct {\n\tc framework.Container\n\n\tbaseURL string\n\tuserID  string\n\ttoken   string\n\tlogger  contract.Log\n}\n\nfunc NewService(params ...interface{}) (interface{}, error) {\n\tc := params[0].(framework.Container)\n\tconfig := c.MustMake(contract.ConfigKey).(contract.Config)\n\tbaseURL := config.GetString("app.stsmap.url")\n\tuserID := config.GetString("app.stsmap.user_id")\n\ttoken := config.GetString("app.stsmap.token")\n\n\tlogger := c.MustMake(contract.LogKey).(contract.Log)\n\treturn &Service{baseURL: baseURL, logger: logger, userID: userID, token: token}, nil\n}\n\n')])])]),t("h2",{attrs:{id:"gocore-provider"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#gocore-provider"}},[e._v("#")]),e._v(" gocore provider")]),e._v(" "),t("p",[e._v("gocore 框架默认自带了一些服务提供者，提供基础的服务接口协议，可以通过 "),t("code",[e._v("./gocore provider list")]),e._v(" 来获取已经安装的服务提供者。")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v("[~/Documents/workspace/gocore_workspace]$ ./gocore provider list\ngocore:app\ngocore:env\ngocore:config\ngocore:log\ngocore:ssh\ngocore:kernel\n")])])]),t("p",[e._v("gocore 框架自带的服务提供者的 key 是以 "),t("code",[e._v("gocore:")]),e._v(" 开头。目的为的是与自定义服务提供者的 key 区别开。")]),e._v(" "),t("p",[e._v("gocore 框架自带的服务提供者具体定义的协议可以参考："),t("RouterLink",{attrs:{to:"/provider/"}},[e._v("provider")])],1)])}),[],!1,null,null,null);r.default=a.exports}}]);