package initial

import "github.com/spf13/cobra"

// 项目初始化配置子命令
var Cmd = &cobra.Command{
	Use:     "init",
	Short:   "mkube init service",
	Long:    "mkube init service",
	Example: "mkube-api init -f xxx",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
