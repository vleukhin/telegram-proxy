package main

import (
	"fmt"
	"log"

	"github.com/vleukhin/telegram-proxy/internal"
)

func main() {
	cfg := internal.Config{}
	if err := cfg.Parse(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Addr)
}
