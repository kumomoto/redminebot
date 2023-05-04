package Flow

import (
	"fmt"
	"strconv"

	redmineAPI "github.com/kumomoto/redminebot/app/redmine"

	redmine "github.com/nixys/nxs-go-redmine/v4"
)

func (us *FlowUserStruct) getSelectedTask(chatID int64, button string) redmine.IssueObject {

	redmineAPI.DeleteClosedIssuesFromCollection(chatID)

	iss, err := strconv.ParseInt(button, 10, 64)

	if err != nil {
		fmt.Print(err)
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("Error, recover from getTaskFromRM")
		}

	}()

	return redmineAPI.GetInfoTask(chatID, iss)
}
