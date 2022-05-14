package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"strings"
)

func init() {
	RootCmd.AddCommand(&podRootCmd) //registry cmd to RootCmd

	podRootCmd.AddCommand(podCreateCmd) //

	// 创建 Pod 的选项参数
	podCreateCmd.Flags().StringVar(&podCreateName, "name", "", "pod name")
	podCreateCmd.Flags().StringVar(&podCreateImage, "image", "", "image name")

	//podCreateCmd.Flags().StringVarP(&age, "age", "a", "", "person's age")

	podCreateCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")

}

var podRootCmd = cobra.Command{
	Use:   "pod",
	Short: "pod is used to manage kubernetes Pods",
}

var age int
var podCreateName string
var podCreateImage string

var podCreateCmd = &cobra.Command{
	Use:                   "create",
	Short:                 "xx",
	Long:                  ``,
	DisableFlagsInUseLine: false,
	//Args:                  cobra.NoArgs,
	//	Args: cobra.ExactArgs(1), // must one arg
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}
