package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"strings"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print the version number of cf",
	Long:                  `cf version  长的部分`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strings.Join(args, " "))
		fmt.Println("cf v2.0")
	},
}
