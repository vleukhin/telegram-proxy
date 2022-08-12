package handlers

import (
	"github.com/vleukhin/telegram-proxy/internal/proxy"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"
)

type SMSHandler struct {
	proxy *proxy.Proxy
}

func NewSMSHandler(proxy *proxy.Proxy) *SMSHandler {
	return &SMSHandler{proxy: proxy}
}

func (h *SMSHandler) ProxyMessage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userIDraw := query.Get("user_id")
	if userIDraw == "" {
		log.Warn().Msg("Empty user ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDraw)
	if err != nil {
		log.Warn().Msgf("Bad user ID: ", userIDraw)
		w.WriteHeader(http.StatusBadRequest)
	}

	text := query.Get("text")
	if text == "" {
		log.Warn().Msg("Empty text")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.proxy.SendMessage(userID, text)
}
