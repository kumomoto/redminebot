package uploadFiles

import (
	"fmt"
	"io"
	"net/http"
)

func (f *FileFlow) downloadDocOnServer(photoUri string, idOfFile string) error {

	pathOfUser := f.createDirOfUser()

	FileDescriptor := f.createFileObject(pathOfUser, Documents.FileName)

	fileBody, err := http.Get(photoUri)

	if err != nil {
		fmt.Print("Cant download file! Photo path is not valid")
	}

	defer fileBody.Body.Close()

	_, erroroOfCopy := io.Copy(FileDescriptor, fileBody.Body)

	defer FileDescriptor.Close()

	return erroroOfCopy
}
