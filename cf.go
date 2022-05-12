package main

import (
	"fmt"
	"github.com/yimtun/cf/cmd"
	"os"
)

func main() {

	execute()
}

func execute() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
