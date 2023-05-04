package redmine

import (
	"log"
	"strconv"

	"github.com/kumomoto/redminebot/app/mongodb"
)

func DeleteClosedIssuesFromCollection(chatID int64) {

	usr, err := mongodb.Сonnection.GetUser(chatID)

	if err != nil {
		log.Panic(err)
	}

	arr := mongodb.Сonnection.GetIssues(chatID)

	for _, iss := range arr {

		intiss, err := strconv.Atoi(iss)

		if err != nil {
			log.Panic(err)
		}

		if issueCheckStatus(chatID, int64(intiss)) {
			usr.Issues[iss] = false
			mongodb.Сonnection.UpdateUser(usr, chatID)
		}
	}
}
