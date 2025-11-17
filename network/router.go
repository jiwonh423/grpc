package network

import (
	"grpc/config"
	"grpc/service"
)

type Network struct{
	cfg *config.Config

	service *service.Service
}

// 네트워크는 전송한 요청에 대해서 서비스한테만 보낼 예정
func NewNetwork(cfg *config.Config, service *service.Service) (*Network,error){
	r:= &Network{cfg:cfg, service:service}
	return r,nil 
}