package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewMiniBlogCommand 创建一个*cobra.Command对象，之后就可以使用Command对象的Execute方法来启动应用程序
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.
		Find more information about miniblog at:
		https://github.com/tuxnotes/miniblog#README.MD`,
		// 设置为true时，命令出错时，不打印帮助信息，这样可以直接看到错误信息
		SilenceUsage: true,
		// 指定调用cmd.Execute()时，执行的Run函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	return cmd
}

// run函数是实际的业务代码入口
func run() error {
	fmt.Println("Hello MiniBlog")
	return nil
}
