package api

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kumomoto/redminebot/app/mongodb"
)

type ChannelOfUser struct {
	chatID  int64
	channel chan tgbotapi.Update
}

func NewChannelOfUser(chatID int64) *ChannelOfUser {
	return &ChannelOfUser{
		chatID:  chatID,
		channel: make(chan tgbotapi.Update),
	}
}

func addChannelsToMap(chatID int64, chann *ChannelOfUser) {

	user, err := mongodb.Сonnection.GetUser(chatID)

	if err != nil {
		fmt.Errorf("Cant find user in DB")
	}

	user.StatusOfChannel = true

	mongodb.Сonnection.UpdateUser(user, chatID)

	Channels[chatID] = chann
}

func makeChannelForUser(chatID int64) *ChannelOfUser {

	chann := NewChannelOfUser(chatID)

	addChannelsToMap(chatID, chann)

	return chann
}

func checkOpenChannels() {
	for ID, channel := range Channels {
		user, _ := mongodb.Сonnection.GetUser(ID)
		if user.StatusOfChannel == false {
			close(channel.channel)
			delete(Channels, ID)
		}
	}
}

func (telegramBot Bot) getStatusChannels(chatID int64) bool {

	user, err := mongodb.Сonnection.GetUser(chatID)

	var ok bool = false

	if err != nil {
		fmt.Errorf("Error to get user - %d", chatID)
		telegramBot.API.Send(tgbotapi.NewMessage(chatID, "Error to get user from DB, try later"))
		return ok
	}

	if user.StatusOfChannel == true {
		ok = true
	}

	return ok
}

func (telegramBot Bot) distributionCallbackBetweenUsersChannels(api *tgbotapi.BotAPI, updatesChan chan tgbotapi.Update) {
	update := <-updatesChan
	er := telegramBot.Worker.ProcessCallback(api, update, updatesChan)
	if er != nil {
		fmt.Println(er)
	}
}

func (telegramBot Bot) distributionMessagesBetweenUsersChannels(api *tgbotapi.BotAPI, updatesChan chan tgbotapi.Update) {
	update := <-updatesChan
	er := telegramBot.Worker.ProcessStartMessage(api, update, updatesChan)
	if er != nil {
		fmt.Println(er)
	}
}
