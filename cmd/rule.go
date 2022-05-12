package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	ruleCmd.Flags().StringVarP(&name, "name", "n", "", "rule name")
	RootCmd.AddCommand(ruleCmd)
}

// 添加变量 name
var name string

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "rule",
	Long:  "Rule Command.",
	Run: func(cmd *cobra.Command, args []string) {
		// 如果没有输入 name
		if len(name) == 0 {
			cmd.Help()
			return
		}
		fmt.Printf("Create rule %s success.\n", name)
	},
}
