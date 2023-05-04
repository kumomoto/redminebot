package Flow

import (
	"fmt"

	"github.com/kumomoto/redminebot/app/mongodb"
	redmineAPI "github.com/kumomoto/redminebot/app/redmine"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (us *FlowUserStruct) FlowProcessCreateIssue(chatID int64, BotAPI *tgbotapi.BotAPI, UpdatesChannel chan tgbotapi.Update) {

	//id, name := us.getRedmineUser(chatID)

	//task := us.getNewIssStruct(id, name)
	API.Send(us.faqAboutNewIss(chatID))

	us.SendMessage("Напишите пожалуйста название вашего обращения", chatID)
	us.SetStepUser(chatID, "waitSubject")
	us.Task.Subject = us.waitSubjectMessage(UpdatesChannel)

	if us.Cancel == true {
		us.SetStepUser(chatID, "newiss")
		us.SendMessage("Отмена создания обращения", chatID)
		return
	}

	us.SendMessage("В следующем обращении просим вас подробно описать возникшую проблему", chatID)
	us.SetStepUser(chatID, "waitDesc")
	us.Task.Description += us.waitDeskMessage(UpdatesChannel)

	if us.Cancel == true {
		us.SetStepUser(chatID, "newiss")
		us.SendMessage("Отмена создания обращения", chatID)
		return
	}

	us.SendMessageWithButtons("Теперь у вас есть возможность загрузить файлы, которые могут помочь в решение проблемы, или же вы можете дополнить описание отправив тестовое сообщение.\n\n Прикрепить файлы можно нажав на скрепку и выбрав необходимые файлы. \n\nПросим вас при загрузке изображений не сжимать их, или загружать в формате jpg! \n\nЕсли вы закончили создавать обращение, или хотите сразу отправить его без изменений и дополнений нажмите кнопку 'Отправить'", chatID, us.getSendButton())
	us.SetStepUser(chatID, "waitAttach")
	us.waitResponceWithFiles(UpdatesChannel)

	if us.Cancel == true {
		us.SetStepUser(chatID, "newiss")
		us.SendMessage("Отмена создания обращения", chatID)
		return
	}

	us.SetChannelClosedState(chatID)
	us.SetStepUser(chatID, "newiss")
	us.setAuthors(chatID)

	if len(us.Task.Attachments) != 0 && us.Task.Description != "" {
		us.ErrOfNewIss, us.IdNewIss = redmineAPI.AddNewTaskWithFiles(us.Task, chatID)

		if us.ErrOfNewIss != nil {
			messageOboutCreateIssue := fmt.Sprint("Ошибка создания обращения с файлами под номером - ", us.IdNewIss, "\n\n", "Ошибка - ", us.ErrOfNewIss)
			us.SendMessage(messageOboutCreateIssue, chatID)
		} else {
			messageOboutCreateIssue := fmt.Sprint("Обращение успешно создано! Номер обращения - ", us.IdNewIss)
			us.SendMessage(messageOboutCreateIssue, chatID)
		}

	} else if us.Task.Description != "" && len(us.Task.Attachments) == 0 {
		us.ErrOfNewIss, us.IdNewIss = redmineAPI.AddNewTaskWithFiles(us.Task, chatID)

		if us.ErrOfNewIss != nil {
			messageOboutCreateIssue := fmt.Sprint("Ошибка создания обращения без файлов под номером - ", us.IdNewIss, "\n\n", "Ошибка - ", us.ErrOfNewIss)
			us.SendMessage(messageOboutCreateIssue, chatID)
		} else {
			messageOboutCreateIssue := fmt.Sprint("Обращение успешно создано! Номер обращения - ", us.IdNewIss)
			us.SendMessage(messageOboutCreateIssue, chatID)
		}
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Print("Error, recover from processCreateIss")
		}
	}()

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)
}

func (us *FlowUserStruct) SetStepUser(chatID int64, st string) {
	user, _ := mongodb.Сonnection.GetUser(chatID)
	user.Step = st
	mongodb.Сonnection.UpdateUser(user, chatID)
}

func (us *FlowUserStruct) SetChannelClosedState(chatID int64) {
	user, _ := mongodb.Сonnection.GetUser(chatID)
	user.StatusOfChannel = false
	mongodb.Сonnection.UpdateUser(user, chatID)
}

func (us *FlowUserStruct) getSendButton() tgbotapi.InlineKeyboardMarkup {
	sendbutton := tgbotapi.NewInlineKeyboardButtonData("Отправить", "send")
	editButton := tgbotapi.NewInlineKeyboardButtonData("Редактировать обращение", "edit")
	return tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{sendbutton, editButton})
}

func (us *FlowUserStruct) setAuthors(chatID int64) {
	user, err := us.Object.GetUser(chatID)

	if err != nil {
		fmt.Errorf("Cant get user error!")
	}
	defer func() {
		if recoverValue := recover(); recoverValue != nil {
			fmt.Print("recover from getRedmineUser")
		}
	}()

	us.Task.Author.ID = user.Author.ID
	us.Task.Author.Name = user.Author.Name
}
