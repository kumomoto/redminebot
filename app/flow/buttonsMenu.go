package Flow

import (
	redmineAPI "github.com/kumomoto/redminebot/app/redmine"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (us *FlowUserStruct) CreateButtonsMainMenu(chatID int64) {

	newIssueButton := tgbotapi.NewInlineKeyboardButtonData("Новое обращение", "newiss")
	getStatusIssueButton := tgbotapi.NewInlineKeyboardButtonData("Узнать статус обращения", "getstatus")
	ReplyKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{newIssueButton, getStatusIssueButton})

	us.SendMessageWithButtons("Выберите необходимую функцию", chatID, ReplyKeyboardMarkup)

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)
}

func (us *FlowUserStruct) CreateButtonsStatIssMenu(chatID int64) {

	newIssueButton := tgbotapi.NewInlineKeyboardButtonData("Новое обращение", "newiss")

	getStatusIssueButton := tgbotapi.NewInlineKeyboardButtonData("Узнать статус другого обращения", "getstatus")

	backToMain := tgbotapi.NewInlineKeyboardButtonData("Вернуться в основное меню", "gotomain")

	ReplyKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{newIssueButton, getStatusIssueButton, backToMain})

	us.SendMessageWithButtons("Выберите необходимую функцию", chatID, ReplyKeyboardMarkup)

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)

}

func (us *FlowUserStruct) CreateButtonsNewIssMenu(chatID int64) {

	newIssueButton := tgbotapi.NewInlineKeyboardButtonData("Ещё одно обращение", "newiss")

	getStatusIssueButton := tgbotapi.NewInlineKeyboardButtonData("Узнать статус обращения", "getstatus")

	backToMain := tgbotapi.NewInlineKeyboardButtonData("Вернуться в основное меню", "gotomain")

	ReplyKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{newIssueButton, getStatusIssueButton, backToMain})

	us.SendMessageWithButtons("Выберите необходимую функцию", chatID, ReplyKeyboardMarkup)

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)

}
