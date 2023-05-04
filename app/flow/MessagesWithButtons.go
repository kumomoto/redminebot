package Flow

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (us *FlowUserStruct) SendMessage(message string, chatID int64) {
	responce := tgbotapi.NewMessage(chatID, message)
	API.Send(responce)
}

func (us *FlowUserStruct) SendMessageWithButtons(message string, chatID int64, buttonsMarkup tgbotapi.InlineKeyboardMarkup) error {
	responce := tgbotapi.NewMessage(chatID, message)
	responce.ReplyMarkup = buttonsMarkup
	_, err := API.Send(responce)
	return err
}
