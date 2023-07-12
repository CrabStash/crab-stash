package config

import (
	"log"

	"github.com/caarlos0/env/v9"
)

type config struct {
	addr         string `env:"ADDR"`
	db_addr      string `env:"DB_ADDR"`
	jwt_duration uint16 `env:"JWT_DURATION"`
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

func (cfg *config) GetDbAddr() string {
	return cfg.db_addr
}

func (cfg *config) GetJWTDuration() uint16 {
	return cfg.jwt_duration
}
