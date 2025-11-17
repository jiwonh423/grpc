package service

import (
	"grpc/config"
	"grpc/repository"
)


type Service struct{
	cfg *config.Config

	repository *repository.Repository
}

// 서비스는 레포지토리와만 연결
func NewService(cfg *config.Config, repository *repository.Repository) (*Service,error){
	r:= &Service{cfg:cfg, repository:repository}

	return r,nil
}