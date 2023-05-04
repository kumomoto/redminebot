package uploadFiles

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type FileFlow struct{}

var PhotosFromUser *[]tgbotapi.PhotoSize
var Documents *tgbotapi.Document
var API *tgbotapi.BotAPI
var ChatID int64
var DirOfUser string
var nameOfFiles []string

func NewFileFlow(api *tgbotapi.BotAPI, size *[]tgbotapi.PhotoSize, chatid int64) *FileFlow {
	setPhotoSizeAndApi(api, size, chatid)
	return &FileFlow{}
}

func NewDocsFlow(api *tgbotapi.BotAPI, dock *tgbotapi.Document, chatid int64) *FileFlow {
	setDocsSizeAndApi(api, dock, chatid)
	return &FileFlow{}
}

func setPhotoSizeAndApi(api *tgbotapi.BotAPI, size *[]tgbotapi.PhotoSize, idOfChat int64) {
	API = api
	PhotosFromUser = size
	ChatID = idOfChat
}

func setDocsSizeAndApi(api *tgbotapi.BotAPI, doc *tgbotapi.Document, idOfChat int64) {
	API = api
	Documents = doc
	ChatID = idOfChat
}
