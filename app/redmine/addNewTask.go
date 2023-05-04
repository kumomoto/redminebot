package redmine

import (
	"strconv"

	"github.com/kumomoto/redminebot/app/mongodb"
)

func AddNewTask(task *Task, chatID int64) (error, int64) {

	var errorOfCreateIssue error

	issue.ProjectID = int(task.ProjectID)
	issue.Description = task.Description
	issue.Subject = task.Subject
	issue.PriorityID = int(task.PriorityID)
	issue.StatusID = int(task.StatusID)

	arr, _, err := redmineContext.IssueCreate(issue)

	if err != nil {
		errorOfCreateIssue = err
	}

	getuser, er := mongodb.Сonnection.GetUser(chatID)

	if er != nil {
		errorOfCreateIssue = er
	}

	getuser.Issues[strconv.Itoa(arr.ID)] = true

	mongodb.Сonnection.UpdateUser(getuser, chatID)

	return errorOfCreateIssue, int64(arr.ID)
}
