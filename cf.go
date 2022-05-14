package main

import (
	"fmt"
	"github.com/yimtun/cf/cmd"
	"os"
)







func main() {
	// first get userToken from thoreau-api-server os.Getenv("thoreau-token")
	// and  get all resource  and privileges list for current user
	//
	execute()
}

func execute() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
