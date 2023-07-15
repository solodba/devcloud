package initial

import "github.com/spf13/cobra"

// 项目初始化配置子命令
var Cmd = &cobra.Command{
	Use:     "init",
	Short:   "mpaas init service",
	Long:    "mpaas init service",
	Example: "mpaas-api init -f xxx",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
