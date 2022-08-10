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
	userIDraw, ok := query["user_id"]
	if !ok || len(userIDraw) == 0 || userIDraw[0] == "" {
		log.Warn().Msg("Empty user ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDraw[0])
	if err != nil {
		log.Warn().Msgf("Bad user ID: ", userIDraw[0])
		w.WriteHeader(http.StatusBadRequest)
	}

	text, ok := query["text"]
	if !ok || len(text) == 0 || text[0] == "" {
		log.Warn().Msg("Empty text")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.proxy.SendMessage(userID, text[0])
}
