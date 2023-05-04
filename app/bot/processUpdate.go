package api

import (
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kumomoto/redminebot/app/mongodb"
)

func (telegramBot Bot) Start() {

	for update := range telegramBot.Updates {

		go checkOpenChannels()

		if update.CallbackQuery != nil {
			telegramBot.processCallback(update)
		} else if update.Message != nil {
			telegramBot.processMessage(update)
		}
		time.Sleep(time.Second)
	}
}

func (telegramBot Bot) processCallback(update tgbotapi.Update) {

	if strings.Compare(getStatusUser(int64(update.CallbackQuery.Message.Chat.ID)), "waitSubject") == 0 {
		Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.CallbackQuery.Message.Chat.ID)), "waitDesc") == 0 {
		Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.CallbackQuery.Message.Chat.ID)), "waitAttach") == 0 {
		Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.CallbackQuery.Message.Chat.ID)), "search") == 0 {
		Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.CallbackQuery.Message.Chat.ID)), "edit") == 0 {
		Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
		return
	}
	makeChannelForUser(int64(update.CallbackQuery.Message.Chat.ID))
	go telegramBot.distributionCallbackBetweenUsersChannels(telegramBot.API, Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel)
	Channels[int64(update.CallbackQuery.Message.Chat.ID)].channel <- update
}

func (telegramBot Bot) processMessage(update tgbotapi.Update) {

	if strings.Compare(getStatusUser(int64(update.Message.Chat.ID)), "reg") == 0 {
		Channels[int64(update.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.Message.Chat.ID)), "waitSubject") == 0 {
		Channels[int64(update.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.Message.Chat.ID)), "waitDesc") == 0 {
		Channels[int64(update.Message.Chat.ID)].channel <- update
		return
	} else if strings.Compare(getStatusUser(int64(update.Message.Chat.ID)), "waitAttach") == 0 {
		Channels[int64(update.Message.Chat.ID)].channel <- update
		return
	}
	makeChannelForUser(int64(update.Message.Chat.ID))
	go telegramBot.distributionMessagesBetweenUsersChannels(telegramBot.API, Channels[int64(update.Message.Chat.ID)].channel)
	Channels[int64(update.Message.Chat.ID)].channel <- update
}

func getStatusUser(chatID int64) string {
	user, _ := mongodb.Сonnection.GetUser(chatID)
	return user.Step
}
