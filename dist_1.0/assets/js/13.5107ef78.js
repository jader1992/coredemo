(window.webpackJsonp=window.webpackJsonp||[]).push([[13],{478:function(e,t,n){"use strict";n.r(t);var r=n(61),o=Object(r.a)({},(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[n("h1",{attrs:{id:"调试模式"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#调试模式"}},[e._v("#")]),e._v(" 调试模式")]),e._v(" "),n("h2",{attrs:{id:"命令"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#命令"}},[e._v("#")]),e._v(" 命令")]),e._v(" "),n("p",[e._v("gocore 框架自带调试模式，不管是前端还是后端，都可以启动调试模式，边修改代码，边编译运行服务。")]),e._v(" "),n("p",[e._v("对应的命令为 "),n("code",[e._v("./gocore dev")])]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v('[~/Documents/workspace/gocore_workspace]$ ./gocore dev\ndev mode\n\nUsage:\n  gocore dev [flags]\n  gocore dev [command]\n\nAvailable Commands:\n  all         dev mode from both frontend and backend\n  backend     dev mode for backend, hot reload\n  frontend    dev mode for frontend\n\nFlags:\n  -h, --help   help for dev\n\nUse "gocore dev [command] --help" for more information about a command.\n')])])]),n("ul",[n("li",[e._v("调试前端")]),e._v(" "),n("li",[e._v("调试后端")]),e._v(" "),n("li",[e._v("同时调试")])]),e._v(" "),n("h2",{attrs:{id:"调试前端"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#调试前端"}},[e._v("#")]),e._v(" 调试前端")]),e._v(" "),n("p",[e._v("使用命令 "),n("code",[e._v("./gocore dev frontend")])]),e._v(" "),n("p",[e._v("要求当前编译机器安装 npm 软件，并且当前项目已经运行了 npm install，安装完成前端依赖。")]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v("[~/Documents/workspace/gocore_workspace]$ ./gocore dev frontend\n\n> gocore@0.1.0 serve /Users/Documents/workspace/gocore_workspace\n> vue-cli-service serve\n\n INFO  Starting development server...\n98% after emitting\n\n DONE  Compiled successfully in 2589ms                                                                                                     下午6:07:06\n\n\n  App running at:\n  - Local:   http://localhost:8080\n  - Network: http://172.24.34.34:8080\n\n  Note that the development build is not optimized.\n  To create a production build, run npm run build.\n")])])]),n("p",[e._v("实际上是调用 "),n("code",[e._v("npm run dev")]),e._v(" 来调试前端")]),e._v(" "),n("h2",{attrs:{id:"调试后端"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#调试后端"}},[e._v("#")]),e._v(" 调试后端")]),e._v(" "),n("p",[e._v("使用命令 "),n("code",[e._v("./gocore dev backend")])]),e._v(" "),n("p",[e._v("要求当前编译机器安装 go 软件，版本 > 1.3。")]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v("[~/Documents/workspace/gocore_workspace]$  ./gocore dev backend\n./gocore dev backend\nbuild success please run ./gocore direct\nbackend server: http://127.0.0.1:15060\nproxy backend server: http://0.0.0.0:8073\n[PID] 29034\napp serve url: http://127.0.0.1:15060\n")])])]),n("p",[e._v("可以通过 proxy backend server 地址进行访问：")]),e._v(" "),n("p",[n("code",[e._v("http://0.0.0.0:8073/demo/demo")])]),e._v(" "),n("div",{staticClass:"custom-block tip"},[n("p",{staticClass:"custom-block-title"},[e._v("TIP")]),e._v(" "),n("p",[e._v("后端调试默认是最后一次操作后3秒启动后端编译启动命令。")]),e._v(" "),n("p",[e._v("gocore 也允许通过配置修改这个等待时间。")]),e._v(" "),n("p",[e._v("可以配置 "),n("code",[e._v("development/app.yaml")]),e._v(" 里面的 "),n("code",[e._v("dev_fresh")]),e._v(" 参数修改这个等待时间。")])]),e._v(" "),n("h1",{attrs:{id:"同时调试"}},[n("a",{staticClass:"header-anchor",attrs:{href:"#同时调试"}},[e._v("#")]),e._v(" 同时调试")]),e._v(" "),n("p",[e._v("也可以选择同时调试，这个时候会同时运行调试前端和调试后端的程序")]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v("[~/Documents/workspace/gocore_workspace]$ ./gocore dev all\n\n> gocore@0.1.0 serve /Users/Documents/workspace/gocore_workspace\n> vue-cli-service serve\n\n INFO  Starting development server...\nbuild success please run ./gocore direct\nbackend server: http://127.0.0.1:19866\nproxy backend server: http://0.0.0.0:8073\nproxy frontend server: http://0.0.0.0:8073/dist/#/\n[PID] 29761\napp serve url: http://127.0.0.1:19866\n98% after emitting\n\n DONE  Compiled successfully in 1421ms                                                                                                     下午6:19:51\n\n\n  App running at:\n  - Local:   http://localhost:19073\n  - Network: http://172.24.34.34:19073\n\n  Note that the development build is not optimized.\n  To create a production build, run npm run build.\n\n[GIN] 2020/09/16 - 18:20:26 | 200 |     134.079µs |       127.0.0.1 | GET      /demo/demo\n\n")])])]),n("p",[e._v("前端和后端的访问地址分别为：")]),e._v(" "),n("div",{staticClass:"language- extra-class"},[n("pre",{pre:!0,attrs:{class:"language-text"}},[n("code",[e._v("proxy backend server: http://0.0.0.0:8073\nproxy frontend server: http://0.0.0.0:8073/dist/#/\n")])])])])}),[],!1,null,null,null);t.default=o.exports}}]);