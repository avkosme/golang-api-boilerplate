package telegram

import (
	"fmt"

	"github.com/avkosme/golang-api-boilerplate/internal/config"
	"github.com/avkosme/golang-api-boilerplate/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Teleram struct {
	username  string
	chatGroup int64
}

// Send set WebHook request to teleram api server
func (t *Teleram) WebHook() {

	var teleram Teleram

	bot, err := teleram.NewBot()

	if err != nil {
		logger.ForError(err)
		panic(err)
	}
	bot.SetWebhook(
		tgbotapi.NewWebhookWithCert(
			fmt.Sprintf("https://%s:%s", config.BotAddress, config.BotPort), config.CertPath))
}

func (t *Teleram) NewBot() (bot *tgbotapi.BotAPI, err error) {

	bot, err = tgbotapi.NewBotAPI(fmt.Sprintf("%s:%s", config.BotId, config.BotKey))

	if err != nil {
		logger.ForError(err)
		panic(err)
	}

	return bot, err
}

func (t *Teleram) Send() {

	var teleram Teleram
	var replyKeyboardMarkup tgbotapi.ReplyKeyboardMarkup

	replyKeyboardMarkup.OneTimeKeyboard = true

	messageConfig := tgbotapi.NewMessage(t.chatGroup, "Список страниц")

	messageConfig.ReplyMarkup = replyKeyboardMarkup

	var keyboardButton tgbotapi.KeyboardButton
	keyboardButton.Text = "-=Добавить страницу=-"
	keyboard := []tgbotapi.KeyboardButton{keyboardButton}

	replyKeyboardMarkup.Keyboard = append(replyKeyboardMarkup.Keyboard, keyboard)

	bot, err := teleram.NewBot()

	if err != nil {
		logger.ForError(err)
		panic(err)
	}

	bot.Send(messageConfig)
}
