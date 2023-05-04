package uploadFiles

import (
	"fmt"

	"github.com/kumomoto/redminebot/app/conf"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (f *FileFlow) downloadPhotosFromUserToServer() {

	filesConf := tgbotapi.FileConfig{}

	filesConf.FileID = f.getPhotoPhileID()

	Photo, err := API.GetFile(filesConf)

	if err != nil {
		fmt.Print("Error! Cant get file from telegram")
	}

	photoLink := Photo.Link(conf.TELEGRAM_BOT_API_KEY)

	idOfPhoto := Photo.FileID

	errorOfDownload := f.downloadPhotoOnServer(photoLink, idOfPhoto)

	if errorOfDownload != nil {
		fmt.Print("Error of download")
	}
}

func (f *FileFlow) getPhotoPhileID() string {

	var photo tgbotapi.PhotoSize

	var size int = 0

	for _, object := range *PhotosFromUser {

		if object.FileSize > size {
			photo = object
			size = object.FileSize
		}
	}

	return photo.FileID
}
