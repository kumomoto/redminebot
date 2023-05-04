package uploadFiles

import (
	"fmt"
	"os"
)

func (f *FileFlow) createFileObject(userPath string, fileName string) *os.File {

	nameOfFiles = append(nameOfFiles, fileName)

	photoPath := fmt.Sprintf("%s/%s", userPath, fileName)

	fileDeskriptor, err := os.Create(photoPath)

	if err != nil {
		fmt.Print("File is not create!")
	}

	return fileDeskriptor

}
