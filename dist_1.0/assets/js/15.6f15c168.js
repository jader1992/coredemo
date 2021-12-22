(window.webpackJsonp=window.webpackJsonp||[]).push([[15],{479:function(e,o,a){"use strict";a.r(o);var n=a(61),r=Object(n.a)({},(function(){var e=this,o=e.$createElement,a=e._self._c||o;return a("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[a("h1",{attrs:{id:"安装"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#安装"}},[e._v("#")]),e._v(" 安装")]),e._v(" "),a("hr"),e._v(" "),a("h2",{attrs:{id:"可执行文件"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#可执行文件"}},[e._v("#")]),e._v(" 可执行文件")]),e._v(" "),a("p",[e._v("我们有两种方式来获取可执行的gocore文件，第一种是直接下载对应操作系统的gocore文件，另外一种是下载源码自己编译")]),e._v(" "),a("h3",{attrs:{id:"直接下载"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#直接下载"}},[e._v("#")]),e._v(" 直接下载")]),e._v(" "),a("p",[e._v("下载地址：\nxxx")]),e._v(" "),a("p",[e._v("将生成的可执行文件 gocore 放到 $PATH 目录中：\n"),a("code",[e._v("cp gocore /usr/local/bin/")])]),e._v(" "),a("h3",{attrs:{id:"源码编译"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#源码编译"}},[e._v("#")]),e._v(" 源码编译")]),e._v(" "),a("p",[e._v("下载 git 地址："),a("code",[e._v("git@github.com/jianfengye/gocore:cloud/gocore.git")]),e._v(" 到目录 gocore")]),e._v(" "),a("p",[e._v("在 gocore 目录中运行命令 "),a("code",[e._v("go run main.go build self")])]),e._v(" "),a("p",[e._v("将生成的可执行文件 gocore 放到 $PATH 目录中：\n"),a("code",[e._v("cp gocore /usr/local/bin/")])]),e._v(" "),a("h2",{attrs:{id:"初始化项目"}},[a("a",{staticClass:"header-anchor",attrs:{href:"#初始化项目"}},[e._v("#")]),e._v(" 初始化项目")]),e._v(" "),a("p",[e._v("使用命令 "),a("code",[e._v("gocore new [app]")]),e._v(" 在当前目录创建子项目")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v("[~/Documents/workspace/gocore_workspace]$ gocore new --help\ncreate a new app\n\nUsage:\n  gocore new [app] [flags]\n\nAliases:\n  new, create, init\n\nFlags:\n  -f, --force        if app exist, overwrite app, default: false\n  -h, --help         help for new\n  -m, --mod string   go mod name, default: folder name\n")])])]),a("p",[e._v("你可以通过 --mod 来定义项目名字的模块名字，默认项目名，目录名，模块名是相同的")]),e._v(" "),a("p",[e._v("接下来，可以通过命令 "),a("code",[e._v("go run main.go")]),e._v(" 看到如下信息：")]),e._v(" "),a("div",{staticClass:"language- extra-class"},[a("pre",{pre:!0,attrs:{class:"language-text"}},[a("code",[e._v('[~/Documents/workspace/gocore_workspace]$ go run main.go\ngocore commands\n\nUsage:\n  gocore [command]\n\nAvailable Commands:\n  app         start app serve\n  build       build gocore command\n  command     all about commond\n  cron        about cron command\n  deploy      deploy app by ssh\n  dev         dev mode\n  env         get current environment\n  help        get help info\n  middleware  gocore middleware\n  new         create a new app\n  provider    about gocore service provider\n  swagger     swagger operator\n\nFlags:\n  -h, --help   help for gocore\n\nUse "gocore [command] --help" for more information about a command.\n')])])]),a("p",[e._v("至此，项目安装成功。")])])}),[],!1,null,null,null);o.default=r.exports}}]);