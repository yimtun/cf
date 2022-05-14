package cmd

import (
	//	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cf",
	Short: "not Cloud Foundy cf",
	Long: `cf long output
	                      love by spf13 and friends in Go.
			      `,
	Args: cobra.NoArgs,
	//Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()
	},
}
