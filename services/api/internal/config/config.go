package config

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type config struct {
	addr       string `env:"ADDR"`
	grpcServer string `env:"GRPC_SERVER"`
}

var Cfg config

func InitializeConfig() {
	Cfg = config{}

	if err := env.Parse(&Cfg); err != nil {
		log.Fatalf("error while loading config:%v", err)
	}
}

func (cfg *config) GetAddr() string {
	return cfg.addr
}

func (cfg *config) GetGrpcServer() string {
	return cfg.grpcServer
}
