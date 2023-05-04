package api

import (
	"log"

	"github.com/kumomoto/redminebot/app/conf"
	"github.com/kumomoto/redminebot/app/redmine"
	"github.com/kumomoto/redminebot/app/worker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var task redmine.Task
var Channels map[int64]*ChannelOfUser = make(map[int64]*ChannelOfUser)

type Bot struct {
	API                   *tgbotapi.BotAPI        //API TG
	Updates               tgbotapi.UpdatesChannel //Updates users
	ActiveContactRequests []int64                 //ID chats where waiting phone numbers
	Worker                *worker.Worker
}

func (telegramBot *Bot) Init() {
	botAPI, err := tgbotapi.NewBotAPI(conf.TELEGRAM_BOT_API_KEY) //Init connection to Bot
	if err != nil {
		log.Panic(err) //Print err and exit
	}

	telegramBot.API = botAPI
	botUpdate := tgbotapi.NewUpdate(conf.TLEGRAM_BOT_UPDATE_OFFSET) //Init update channel
	botUpdate.Timeout = conf.TELEGRAM_BOT_UPDATE_TIMEOUT
	botUpdates, err := telegramBot.API.GetUpdatesChan(botUpdate)

	if err != nil {
		log.Fatal(err) //Print err and exit
	}
	Channels = make(map[int64]*ChannelOfUser)
	telegramBot.Updates = botUpdates
}
