package main

import (
	api "github.com/kumomoto/redminebot/app/bot"
	"github.com/kumomoto/redminebot/app/mongodb"
	"github.com/kumomoto/redminebot/app/redmine"
	"github.com/kumomoto/redminebot/app/worker"
)

func main() {
	mongodb.Сonnection.Init()
	redmine.Init()

	telegramBot := &api.Bot{
		Worker: worker.NewWorker(&mongodb.Сonnection),
	}

	telegramBot.Init()
	telegramBot.Start()
}
