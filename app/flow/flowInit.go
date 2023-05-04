package Flow

import (
	"github.com/kumomoto/redminebot/app/conf"
	"github.com/kumomoto/redminebot/app/mongodb"
	redmineAPI "github.com/kumomoto/redminebot/app/redmine"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var API *tgbotapi.BotAPI

type FlowUserStruct struct {
	Object      *mongodb.DatabaseConnection
	IdNewIss    int64
	ErrOfNewIss error
	Cancel      bool
	Task        *redmineAPI.Task
}

func NewFlow(user *mongodb.DatabaseConnection) *FlowUserStruct {

	// TODO: Сзодить в storage и посмотреть, нет ли активного процесса, если есть то какой

	return &FlowUserStruct{
		Object:      user,
		IdNewIss:    0,
		ErrOfNewIss: nil,
		Cancel:      false,
		Task: &redmineAPI.Task{
			ProjectID:   conf.REDMINE_PROJECTID,
			PriorityID:  4,
			StatusID:    1,
			Subject:     "",
			Description: "",
			Attachments: nil,
			Author:      &redmineAPI.Author{ID: 0, Name: ""},
		},
	}
}

func (us *FlowUserStruct) SetApiAndUpdates(botAPI *tgbotapi.BotAPI) {
	API = botAPI
}

func (us *FlowUserStruct) Close() {
	us = nil
}
