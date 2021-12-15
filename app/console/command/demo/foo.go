package demo

import (
	"github.com/jader1992/gocore/framework/cobra"
	"log"
)

func InitFoo() *cobra.Command {
	FooCommand.AddCommand(Foo1Command)
	return FooCommand
}

// FooCommand 代表Foo命令
var FooCommand = &cobra.Command{
	Use:     "foo",
	Short:   "foo的简要说明",
	Long:    "food的长说明",
	Aliases: []string{"fo", "f"},
	Example: "foo命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		//container := cmd.GetContainer()
		//log.Println(container)
		log.Println("execute foo command")
		return nil
	},
}

// Foo1Command 代表Foo命令的子命令Foo1
var Foo1Command = &cobra.Command{
	Use:     "foo1",
	Short:   "foo1的简要说明",
	Long:    "foo1的长说明",
	Aliases: []string{"fo1", "f"},
	Example: "foo1命令的例子",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		log.Println(container)
		return nil
	},
}
