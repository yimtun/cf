package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"strings"
)

func init() {
	RootCmd.AddCommand(&listRootCmd)
	//list cmd registry root cmd

	listRootCmd.AddCommand(listAppCmd)
	listRootCmd.AddCommand(listEnvCmd)
	listAppCmd.Flags().StringVarP(&envName, "envname", "e", "", "the environment name")
}

var listRootCmd = cobra.Command{
	Use:   "list",
	Short: "list current user's app",
}

var envName string

var listAppCmd = &cobra.Command{
	Use:                   "app",
	Example:               "cf  list  app | cf list app  -e test",
	DisableFlagsInUseLine: false,
	Args:                  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args", args)
		fmt.Println("op val", envName)
	},
}

var listEnvCmd = &cobra.Command{
	Use:                   "env",
	Example:               "cf  list  env",
	DisableFlagsInUseLine: false,
	Args:                  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args", args)
		fmt.Println("op val", envName)
	},
}
