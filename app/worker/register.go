package worker

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kumomoto/redminebot/app/ldap"
	"github.com/kumomoto/redminebot/app/mongodb"
	"github.com/kumomoto/redminebot/app/redmine"
)

func (w *Worker) tryRegister(BotAPI *tgbotapi.BotAPI, UpdatesChannel chan tgbotapi.Update, chatID int64) error {

	w.sendRequestNumberButton(BotAPI, chatID)

	w.users.CreateUser(mongodb.User{
		Chat_ID:         chatID,
		Phone_Number:    "",
		StatusOfChannel: true,
		Step:            "reg",
		Issues:          nil,
		Author:          &mongodb.Author{},
	})

	time.Sleep(10)

	number := w.whaitContact(UpdatesChannel, BotAPI)

	fmt.Println(number)

	mongoUser, err := w.users.GetUser(chatID)

	if err != nil {
		fmt.Println(err)
	}

	ad := ldap.NewLdapConnection(number)

	isAuth := ad.GetAuthResult()

	mail := ad.GetMail()

	rmID, name := redmine.GetRMUserID(mail)

	if isAuth {
		mongoUser.Author.ID = int64(rmID)
		mongoUser.Author.Name = name
		mongoUser.Phone_Number = mongoUser.Phone_Number + number
		mongoUser.Step = "main"
		mongoUser.StatusOfChannel = false
		mongoUser.Mail = mail
		w.users.UpdateUser(mongoUser, chatID)
		return nil
	} else {
		return fmt.Errorf("Error create user!")
	}
}

func (w *Worker) sendRequestNumberButton(BotAPI *tgbotapi.BotAPI, chatID int64) {

	button := tgbotapi.NewKeyboardButtonContact("Предоставить номер")

	markup := tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{button})

	mess := tgbotapi.NewMessage(chatID, "Просим предоставить номер телефона для регистрации в данном боте")

	markup.OneTimeKeyboard = true

	mess.ReplyMarkup = markup

	BotAPI.Send(mess)

}

func (w *Worker) whaitContact(Updates chan tgbotapi.Update, api *tgbotapi.BotAPI) string {
	var number string
	time.Sleep(100)
	for updates := range Updates {
		if updates.Message != nil {
			number = updates.Message.Contact.PhoneNumber
			break
		} else {
			w.sendMessage(api, updates.Message.Chat.ID, "Для предоставления номера просим нажать на кнопку ниже")
		}
	}
	return number
}

func (w *Worker) sendMessage(api *tgbotapi.BotAPI, chatID int64, mess string) {
	messageConf := tgbotapi.NewMessage(chatID, mess)
	api.Send(messageConf)
}
