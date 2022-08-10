package internal

import (
	proxy2 "github.com/vleukhin/telegram-proxy/internal/proxy"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/vleukhin/telegram-proxy/internal/handlers"
)

type (
	App struct {
		cfg   Config
		proxy *proxy2.Proxy
	}
)

func NewApp(cfg Config) *App {
	proxy := proxy2.NewProxy()

	return &App{
		proxy: proxy,
		cfg:   cfg,
	}
}

func (a *App) Run(err chan<- error) {
	log.Info().Msgf("Telegram proxy server listen at: %s", a.cfg.Addr)
	err <- http.ListenAndServe(a.cfg.Addr, NewRouter(a))
}

func NewRouter(app *App) *mux.Router {
	smsHandler := handlers.NewSMSHandler(app.proxy)
	r := mux.NewRouter()
	r.Handle("/send", http.HandlerFunc(smsHandler.ProxyMessage)).Methods(http.MethodGet)

	return r
}
