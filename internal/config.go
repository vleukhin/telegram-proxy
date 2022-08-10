package internal

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

type Config struct {
	Addr string `env:"ADDRESS" envDefault:"localhost:8080"`
}

func (cfg *Config) Parse() error {
	err := env.Parse(cfg)
	if err != nil {
		return err
	}

	addr := pflag.StringP("addr", "a", cfg.Addr, "Server address")
	pflag.Parse()

	cfg.Addr = *addr
	return nil
}
