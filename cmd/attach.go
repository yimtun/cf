package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yimtun/cf/apiserverclient"
	"github.com/yimtun/cf/host"
	"github.com/yimtun/cf/mysql"
	"os"
	"strings"
)

func init() {
	RootCmd.AddCommand(&attachRootCmd)    //将父命令注册到根命令
	attachRootCmd.AddCommand(attachDbCmd) // 将子命令 create  注册到 父命令 pod
	attachRootCmd.AddCommand(attachHostCmd) // 将子命令 create  注册到 父命令 pod

}

var attachRootCmd = cobra.Command{
	Use:   "attach",
	Short: "login db or host",
}

var attachDbCmd = &cobra.Command{
	Use:                   "db",
	Short:                 "login db",
	Long:                  ``,
	DisableFlagsInUseLine: false,
	Args:                  cobra.ExactArgs(1), //
	Run: func(cmd *cobra.Command, args []string) {
		arg := strings.Join(args, " ")
		if arg == "mysql01" {
			mysql.ConnectMysqlDemo(apiserverclient.GetResource())

		} else {
			fmt.Println("db：", arg, "not exists")
		}

	},
}


var attachHostCmd = &cobra.Command{
	Use:                   "host",
	Short:                 "login server",
	Long:                  ``,
	DisableFlagsInUseLine: false,
	Args:                  cobra.ExactArgs(1), //指定只接收一个参数  作为 db id
	Run: func(cmd *cobra.Command, args []string) {
		// use thoreau user  and thoreau token  request  api-server
		arg := strings.Join(args, " ")
		if arg == "host01" {
			host.Connect(os.Getenv("thoreauUser"),apiserverclient.GetResource().SshHost,apiserverclient.GetResource().Port,apiserverclient.GetResource().SshUser,apiserverclient.GetResource().SshPassword)
		} else {
			fmt.Println("host：", arg, "not exists")
		}

	},
}


