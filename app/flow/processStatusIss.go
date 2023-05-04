package Flow

import (
	"github.com/kumomoto/redminebot/app/mongodb"
	redmineAPI "github.com/kumomoto/redminebot/app/redmine"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (us *FlowUserStruct) FlowProcessStatusIssue(BotAPI *tgbotapi.BotAPI, chatID int64, UpdatesChannel chan tgbotapi.Update) bool {

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)

	markup, ok := us.getReplMarkupStatusIssue(chatID)

	if !ok {
		us.SetStepUser(chatID, "getstatus")
		us.SetChannelClosedState(chatID)
		return ok
	}

	us.SetStepUser(chatID, "search")

	us.SendMessageWithButtons("Пожалуйста, выберите задачу, по которой хотите получить текущий статус", chatID, markup)

	data := us.waitResponceCallback(UpdatesChannel)

	issObj := us.getSelectedTask(chatID, data)

	us.sendInfoAboutTask(issObj, chatID)

	us.SetChannelClosedState(chatID)

	us.SetStepUser(chatID, "getstatus")

	return ok
}

func (us *FlowUserStruct) checkAmountIssues(chatID int64, arrInlineKeyboard []tgbotapi.InlineKeyboardButton) bool {

	if len(arrInlineKeyboard) == 0 {
		us.SendMessage("У вас пока что нет отслеживаемых задач", chatID)
		return false
	}

	return true
}

func (us *FlowUserStruct) getReplMarkupStatusIssue(chatID int64) (tgbotapi.InlineKeyboardMarkup, bool) {

	var keyboard []tgbotapi.InlineKeyboardButton

	mongArrIssues := mongodb.Сonnection.GetIssues(chatID)

	for _, str := range mongArrIssues {
		task := us.getSelectedTask(chatID, str)
		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData(task.Subject, str))
	}

	if !us.checkAmountIssues(chatID, keyboard) {
		return tgbotapi.NewInlineKeyboardMarkup(keyboard), false
	}

	tgbotapi.NewInlineKeyboardMarkup(keyboard)

	return tgbotapi.NewInlineKeyboardMarkup(keyboard), true
}
