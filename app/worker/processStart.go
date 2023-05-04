package worker

import (
	"log"

	Flow "github.com/kumomoto/redminebot/app/flow"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (w *Worker) ProcessStartMessage(BotAPI *tgbotapi.BotAPI, updateCollback tgbotapi.Update, UpdatesChannel chan tgbotapi.Update) error {

	FlowOfUser := Flow.NewFlow(w.users)

	chatID := updateCollback.Message.Chat.ID

	exist, _ := w.userExists(chatID)

	if !exist {
		err := w.tryRegister(BotAPI, UpdatesChannel, chatID)
		FlowOfUser.SetApiAndUpdates(BotAPI)
		FlowOfUser.CreateButtonsMainMenu(chatID)
		return err
	}

	if updateCollback.Message.Command() == "start" {
		FlowOfUser.SetApiAndUpdates(BotAPI)
		FlowOfUser.CreateButtonsMainMenu(chatID)
	} else {
		message := tgbotapi.NewMessage(chatID, "Данной команды нет в функционале бота")
		BotAPI.Send(message)
	}

	FlowOfUser.SetChannelClosedState(chatID)

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered: %v", panicValue)
		}
	}()

	return nil
}
