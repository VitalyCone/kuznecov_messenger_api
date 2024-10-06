package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/VitalyCone/kuznecov_messenger_api/internal/app/apiserver"
)

var (
	configPath string
	debugConfigPath string
)

func init(){
	configPath = "config/apiserver.toml"
	debugConfigPath = "D:/projects/gitrep/kuznecov_hack/kuznecov_messenger_api/config/apiserver.toml"
}


func main() {
	config := apiserver.NewConfig()

	_,err:= toml.DecodeFile(configPath,config)

	if err != nil{
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err:=s.Start(); err!= nil{
		log.Fatal(err)
	}
}