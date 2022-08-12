package proxy

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog/log"
)

type (
	Proxy struct {
		api *tgbotapi.BotAPI
	}
	Account struct {
		ID int
	}
)

func NewProxy(botToken string) (*Proxy, error) {
	api, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}
	return &Proxy{
		api: api,
	}, nil
}

func (p *Proxy) SendMessage(userID int, text string) error {
	log.Debug().Msgf("Sending \"%s\" to %d", text, userID)
	msg := tgbotapi.NewMessage(int64(userID), text)
	_, err := p.api.Send(msg)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to send message")
		return err
	}

	return nil
}
