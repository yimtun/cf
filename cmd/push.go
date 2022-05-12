package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"strings"
)

func init() {
	RootCmd.AddCommand(&pushRootCmd) // 命令 push 注册到 根命令

	pushRootCmd.AddCommand(pushAppCmd) // 将 子命令 app   注册到 父命令  push

	// 创建 Pod 的选项参数
	//podCreateCmd.Flags().StringVar(&podCreateName, "name", "", "pod name")
	//podCreateCmd.Flags().StringVar(&podCreateImage, "image", "", "image name")

	//podCreateCmd.Flags().StringVarP(&age, "age", "a", "", "person's age")

	//	podCreateCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	pushAppCmd.Flags().StringVarP(&appname, "appname", "n", "", "the app name")
	pushAppCmd.MarkFlagRequired("appname")

}

var pushRootCmd = cobra.Command{
	Use:   "push",
	Short: "发布代码到指定的测试环境",
}

//var age int
var appname string

var pushAppCmd = &cobra.Command{
	Use: "app appname",
	//Short:                 "发布版本到指定环境",
	//Long:                  ``,
	Example:               "push app lz-eoms-hr -e test",
	DisableFlagsInUseLine: false,
	//Args:                  cobra.NoArgs,
	Args: cobra.ExactArgs(1), //指定只接收一个参数  作为app name
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(strings.Join(args, " "))
		fmt.Println("参数列表", args)
		fmt.Println("选项后的值", appname)
	},
}
