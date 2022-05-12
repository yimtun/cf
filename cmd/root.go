package cmd

import (
	//	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cf",
	Short: "仅仅是借用的cf 命令的名字 和 大家所知道的 cf 不是一样的东西",
	Long: `cf long 输出的内容
	                      love by spf13 and friends in Go.
			      这个是cf Long 参数所输出的内容`,
	Args: cobra.NoArgs,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}
