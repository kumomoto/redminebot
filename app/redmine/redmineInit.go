package redmine

import (
	"fmt"
	"strings"

	"github.com/kumomoto/redminebot/app/conf"

	redmine "github.com/nixys/nxs-go-redmine/v4"
)

var redmineContext redmine.Context
var issue redmine.IssueCreateObject

type Task struct {
	ProjectID   int64
	StatusID    int64
	Subject     string
	Description string
	PriorityID  int64
	Author      *Author
	Attachments []redmine.AttachmentUploadObject
}

type Author struct {
	ID   int64
	Name string
}

func Init() {
	redmineContext.SetEndpoint(conf.REDMINE_ENDPOINT)
	redmineContext.SetAPIKey(conf.REDMINE_APIKEY)
}

func DownloadFilesForIssue(path string) (redmine.AttachmentUploadObject, int, error) {
	return redmineContext.AttachmentUpload(path)
}

func GetInfoTask(chatID int64, iss int64) redmine.IssueObject {
	res, _, err := redmineContext.IssueSingleGet(int(iss), redmine.IssueSingleGetRequest{
		Includes: []string{"relations"},
	})
	if err != nil {
		fmt.Print(err)
	}

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Print("Error, recovered from redmineInit")
		}
	}()
	return res
}

func issueCheckStatus(chatID int64, idIss int64) bool {

	var isClosed bool

	stat := GetInfoTask(chatID, int64(idIss)).Status.Name
	if strings.Compare(stat, "Закрыт") == 0 {
		isClosed = true
	}
	return isClosed
}
