package network

import (
	"grpc/config"
	"grpc/service"

	"github.com/gin-gonic/gin"
)

type Network struct{
	cfg *config.Config

	service *service.Service

	engin *gin.Engine
}

// 네트워크는 전송한 요청에 대해서 서비스한테만 보낼 예정
func NewNetwork(cfg *config.Config, service *service.Service) (*Network,error){
	r:= &Network{cfg:cfg, service:service, engin:gin.New()}
	return r,nil 
}

func (n *Network) StartServer(){
	n.engin.Run(":8080")

}