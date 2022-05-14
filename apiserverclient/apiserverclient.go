package apiserverclient

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetResource() Resource {
	test:=getConfig()
	return  test
}

type Resource struct {
	SshHost       string `json:"ssh_host"`
	Port          string `json:"ssh_port"`
	SshUser       string `json:"ssh_user"`
	SshPassword   string `json:"ssh_password"`
	MysqlUser     string `json:"mysql_userser"`
	MysqlHost     string `json:"mysql_host"`
	MysqlPassword string `json:"mysql_password"`
}



type Resources struct {
	Resources []Resource `json:"resources"`
}

func getConfig()  Resource{
	var tmpConfig Resource
	filePtr, err := os.Open("resources.json")
	if err != nil {
		panic(err)
	}

	var resources Resources
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		tmpConfig  = resources.Resources[0]
		fmt.Println(tmpConfig.MysqlHost)
	}

	return tmpConfig
}
