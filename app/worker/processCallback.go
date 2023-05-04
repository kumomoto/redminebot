package worker

import (
	"log"

	Flow "github.com/kumomoto/redminebot/app/flow"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (w *Worker) ProcessCallback(BotAPI *tgbotapi.BotAPI, updateCollback tgbotapi.Update, UpdatesChannel chan tgbotapi.Update) error {

	FlowOfUser := Flow.NewFlow(w.users)
	chatID := int64(updateCollback.CallbackQuery.From.ID)

	FlowOfUser.SetApiAndUpdates(BotAPI)

	exist, err := w.userExists(chatID)
	if err != nil {
		FlowOfUser.SendMessage("Internal error, try lather", chatID)
		return err
	}

	if !exist {
		err := w.tryRegister(BotAPI, UpdatesChannel, chatID)
		return err
	}

	switch updateCollback.CallbackQuery.Data {
	case "newiss":
		FlowOfUser.FlowProcessCreateIssue(chatID, BotAPI, UpdatesChannel)
		FlowOfUser.CreateButtonsNewIssMenu(chatID)
	case "getstatus":
		ok := FlowOfUser.FlowProcessStatusIssue(BotAPI, chatID, UpdatesChannel)
		if !ok {
			FlowOfUser.CreateButtonsMainMenu(chatID)
		} else {
			FlowOfUser.CreateButtonsStatIssMenu(chatID)
		}
	case "gotomain":
		FlowOfUser.SetStepUser(chatID, "main")
		FlowOfUser.SetChannelClosedState(chatID)
		FlowOfUser.CreateButtonsMainMenu(chatID)
		FlowOfUser.SetChannelClosedState(chatID)
	default:
		FlowOfUser.SendMessage("Данной команды нет в представленном функционале бота. Просим выбрать другое действие.", chatID)
		FlowOfUser.CreateButtonsMainMenu(chatID)
		FlowOfUser.SetChannelClosedState(chatID)
	}

	FlowOfUser.Close()

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered: %v", panicValue)
		}
	}()
	return nil
}
