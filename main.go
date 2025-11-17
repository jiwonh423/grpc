package main

import (
	"flag"
	"fmt"
	"grpc/config"
)

var configFlag = flag.String("config","./config.toml","config path")
func main(){
	flag.Parse()

	fmt.Println(*configFlag)
	
	config.NewConfig(*configFlag)
}
