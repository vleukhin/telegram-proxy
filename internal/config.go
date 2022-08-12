package internal

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

type Config struct {
	Addr     string `env:"ADDRESS" envDefault:"localhost:8080"`
	BotToken string `env:"BOT_TOKEN"`
}

func (cfg *Config) Parse() error {
	err := env.Parse(cfg)
	if err != nil {
		return err
	}

	addr := pflag.StringP("addr", "a", cfg.Addr, "App address")
	token := pflag.StringP("token", "t", cfg.BotToken, "Telegram bot token")
	pflag.Parse()

	cfg.Addr = *addr
	cfg.BotToken = *token
	return nil
}
