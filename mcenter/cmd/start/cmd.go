package start

import "github.com/spf13/cobra"

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "mcenter start service",
	Long:    "mcenter start service",
	Example: "mcenter-api start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
