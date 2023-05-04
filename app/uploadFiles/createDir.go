package uploadFiles

import (
	"fmt"
	"os"
	"strconv"
)

func (f *FileFlow) createDirOfUser() string {

	userPath := fmt.Sprintf("/RedmineBot_data/Data/%s", strconv.Itoa(int(ChatID)))

	err := os.MkdirAll(userPath, 0755)

	if err != nil {
		fmt.Println("Error! Cant create dir of User")
	}

	DirOfUser = userPath

	return userPath

}
