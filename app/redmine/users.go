package redmine

import (
	"fmt"
	"strings"

	redmine "github.com/nixys/nxs-go-redmine/v4"
)

func GetRMUserID(mail string) (int, string) {

	var rmID int = 0
	var name string
	for {
		users, _, err := redmineContext.UserAllGet(redmine.UserAllGetRequest{
			redmine.UserGetRequestFilters{
				Status: 1,
				Name:   mail,
			},
		})

		if err != nil {
			fmt.Println(err)
		}

		for _, user := range users.Users {
			if strings.Compare(user.Mail, mail) == 0 {
				rmID = user.ID
				name = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
			}
		}

		if rmID != 0 {
			break
		}

		defer func() {
			if panicValue := recover(); panicValue != nil {
				fmt.Printf("Get user err - ", err)
			}
		}()

	}

	return rmID, name
}
