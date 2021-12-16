package command

import "github.com/jader1992/gocore/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// demo
	root.AddCommand(DemoCommand)
    // app 命令
	root.AddCommand(initAppCommand())
	// env 命令
	root.AddCommand(InitEnvCommand())
	// cron 命令
	root.AddCommand(initCronCommand())
	// config 命令
	root.AddCommand(initConfigCommand())

//	root.AddCommand(deployCommand)
//
	// cron
//	// cmd
//	cmdCommand.AddCommand(cmdListCommand)
//	cmdCommand.AddCommand(cmdCreateCommand)
//	root.AddCommand(cmdCommand)
//
//	// build
//	buildCommand.AddCommand(buildSelfCommand)
//	buildCommand.AddCommand(buildBackendCommand)
//	buildCommand.AddCommand(buildFrontendCommand)
//	buildCommand.AddCommand(buildAllCommand)
//	root.AddCommand(buildCommand)
//

//
//	// dev
//	root.AddCommand(initDevCommand())
//
//	// middleware
//	middlewareCommand.AddCommand(middlewareAllCommand)
//	middlewareCommand.AddCommand(middlewareAddCommand)
//	middlewareCommand.AddCommand(middlewareRemoveCommand)
//	root.AddCommand(middlewareCommand)
//
//	// swagger
//	swagger.IndexCommand.AddCommand(swagger.InitServeCommand())
//	swagger.IndexCommand.AddCommand(swagger.GenCommand)
//	root.AddCommand(swagger.IndexCommand)
//
//	// provider
//	providerCommand.AddCommand(providerListCommand)
//	providerCommand.AddCommand(providerCreateCommand)
//	root.AddCommand(providerCommand)
//
//	// new
//	root.AddCommand(initNewCommand())
}
