package Flow

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (us *FlowUserStruct) waitSubjectMessage(Updates chan tgbotapi.Update) string {
	var message string
	for update := range Updates {
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "cancel" {
				us.Cancel = true
				return ""
			}
		} else if update.Message != nil {
			if update.Message.Document != nil || update.Message.Photo != nil {
				us.SendMessage("Внимание! Просим вас предоставить только название вашего обращения!", update.Message.Chat.ID)
				continue
			} else if update.Message.Text != "" && update.Message.Document == nil && update.Message.Photo == nil {
				message = update.Message.Text
				break
			}
		}
	}
	return message
}

func (us *FlowUserStruct) waitDeskMessage(Updates chan tgbotapi.Update) string {
	var message string
	for update := range Updates {
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "cancel" {
				us.Cancel = true
				return ""
			}
		} else if update.Message != nil {
			if update.Message.Document != nil || update.Message.Photo != nil {
				us.SendMessage("Внимание! Просим вас предоставить только описание вашего обращения!", update.Message.Chat.ID)
				continue
			} else if update.Message.Text != "" && update.Message.Document == nil && update.Message.Photo == nil {
				message = update.Message.Text
				break
			}
		}
	}
	return message
}

func (us *FlowUserStruct) waitResponceCallback(Updates chan tgbotapi.Update) string {
	var data string
	for update := range Updates {
		if update.CallbackQuery.Data != "" {
			data = update.CallbackQuery.Data
			break
		}
	}
	return data
}

func (us *FlowUserStruct) waitResponceWithFiles(Updates chan tgbotapi.Update) {

	for updates := range Updates {

		if updates.CallbackQuery != nil {
			if updates.CallbackQuery.Data == "edit" {
				us.SetStepUser(updates.CallbackQuery.Message.Chat.ID, "edit")
				us.editNewIss(updates.CallbackQuery.Message.Chat.ID, Updates)
				return
			} else if updates.CallbackQuery.Data == "send" {
				return
			} else if updates.CallbackQuery.Data == "cancel" {
				us.Cancel = true
				return
			}
		} else if updates.Message != nil {
			if updates.Message.Text != "" {
				us.Task.Description = fmt.Sprintf("%s %s \n \n", us.Task.Description, updates.Message.Text)
			} else if updates.Message.Caption != "" && updates.Message.Document == nil && updates.Message.Photo == nil {
				us.Task.Description = fmt.Sprintf("%s %s \n", us.Task.Description, updates.Message.Caption)
			} else if updates.Message.Photo != nil && updates.Message.Caption == "" {
				us.Task.Attachments = us.pushPhotos(int64(updates.Message.Chat.ID), updates.Message.Photo, us.Task.Attachments)
			} else if updates.Message.Document != nil && updates.Message.Caption == "" {
				us.Task.Attachments = us.pushDocuments(int64(updates.Message.Chat.ID), updates.Message.Document, us.Task.Attachments)
			} else if updates.Message.Document != nil && updates.Message.Caption != "" {
				us.Task.Attachments = us.pushDocuments(int64(updates.Message.Chat.ID), updates.Message.Document, us.Task.Attachments)
				us.Task.Description = fmt.Sprintf("%s %s \n", us.Task.Description, updates.Message.Caption)
			} else if updates.Message.Photo != nil && updates.Message.Caption != "" {
				us.Task.Attachments = us.pushPhotos(int64(updates.Message.Chat.ID), updates.Message.Photo, us.Task.Attachments)
				us.Task.Description = fmt.Sprintf("%s %s \n", us.Task.Description, updates.Message.Caption)
			}
		}
	}
}
