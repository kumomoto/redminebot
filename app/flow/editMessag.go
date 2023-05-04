package Flow

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	redmineAPI "github.com/kumomoto/redminebot/app/redmine"
)

func (us *FlowUserStruct) editNewIss(chatID int64, UpdatesChannel chan tgbotapi.Update) {

	us.Task.Attachments = nil

	currentIss := fmt.Sprintf("Название обращения:\n\n %s\n\n Описание обращения:\n\n %s", us.Task.Subject, us.Task.Description)

	us.SendMessage(currentIss, chatID)

	us.SendMessage("Введите пожалуйста новое название обращения", chatID)
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

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Print("Error, recover from processCreateIss")
		}
	}()

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)
}
