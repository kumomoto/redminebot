package Flow

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kumomoto/redminebot/app/uploadFiles"
	redmine "github.com/nixys/nxs-go-redmine/v4"
)

func (us *FlowUserStruct) pushDocuments(chatID int64, dock *tgbotapi.Document, AttachmentObjects []redmine.AttachmentUploadObject) []redmine.AttachmentUploadObject {
	push := uploadFiles.NewDocsFlow(API, dock, chatID)
	AttacheDocs := push.PushDocsToRedmine()
	AttachmentObjects = append(AttachmentObjects, AttacheDocs)
	return AttachmentObjects
}

func (us *FlowUserStruct) pushPhotos(chatID int64, Photo *[]tgbotapi.PhotoSize, AttachmentObjects []redmine.AttachmentUploadObject) []redmine.AttachmentUploadObject {
	push := uploadFiles.NewFileFlow(API, Photo, chatID)
	attachments := push.PushPhotoToRedmine()
	AttachmentObjects = append(AttachmentObjects, attachments)
	return AttachmentObjects
}
