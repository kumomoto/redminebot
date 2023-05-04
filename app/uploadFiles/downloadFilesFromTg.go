package uploadFiles

import (
	"fmt"

	"github.com/kumomoto/redminebot/app/conf"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (f *FileFlow) downloadDocsFromUserToServer() {

	filesConf := tgbotapi.FileConfig{}

	filesConf.FileID = Documents.FileID

	Doc, err := API.GetFile(filesConf)

	if err != nil {
		fmt.Print("Error! Cant get file from telegram")
	}

	photoLink := Doc.Link(conf.TELEGRAM_BOT_API_KEY)

	idOfPhoto := Doc.FileID

	errorOfDownload := f.downloadDocOnServer(photoLink, idOfPhoto)

	if errorOfDownload != nil {
		fmt.Print("Error of download")
	}
}
