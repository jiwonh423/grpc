package main

import (
	"flag"
	// "fmt"
	"grpc/cmd"
	"grpc/config"
)

var configFlag = flag.String("config","./config.toml","config path")
func main(){
	flag.Parse()

	// fmt.Println(*configFlag)
	
	cfg := config.NewConfig(*configFlag)
	cmd.NewApp(cfg)
}
