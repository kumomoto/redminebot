package uploadFiles

import (
	"fmt"

	"github.com/kumomoto/redminebot/app/redmine"

	redmineAPI "github.com/nixys/nxs-go-redmine/v4"
)

func (f *FileFlow) PushPhotoToRedmine() redmineAPI.AttachmentUploadObject {

	var AttachmentObject redmineAPI.AttachmentUploadObject
	var err error

	f.downloadPhotosFromUserToServer()

	for _, file := range nameOfFiles {

		fullPath := fmt.Sprintf("%s/%s", DirOfUser, file)

		AttachmentObject, _, err = redmine.DownloadFilesForIssue(fullPath)

		if err != nil {
			fmt.Print("Error of upload files. Error - ", err)
		}
	}

	return AttachmentObject
}
