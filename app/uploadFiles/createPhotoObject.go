package uploadFiles

import (
	"fmt"
	"os"
)

func (f *FileFlow) createPhotoObject(userPath string, fileName string) *os.File {

	nameOfFile := fmt.Sprintf("%s.jpg", fileName)

	nameOfFiles = append(nameOfFiles, nameOfFile)

	photoPath := fmt.Sprintf("%s/%s", userPath, nameOfFile)

	fileDeskriptor, err := os.Create(photoPath)

	if err != nil {
		fmt.Print("File is not create!")
	}

	return fileDeskriptor

}
