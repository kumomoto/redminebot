package Flow

import (
	"fmt"
	"strings"

	redmine "github.com/nixys/nxs-go-redmine/v4"
)

func (us *FlowUserStruct) sendInfoAboutTask(object redmine.IssueObject, chatID int64) {

	if strings.Compare(object.Status.Name, "Новый") == 0 {

		infoOboutTask := fmt.Sprintln("Задача была создана ⏰: ", object.StartDate, "\n", "Задача назначена на ⛑: ", object.AssignedTo.Name, "\n", "Название задачи 👉: ", object.Subject, "\n", "Статус задачи 🔧:", object.Status.Name, "\n", "📌 Ссылка на задачу: ", fmt.Sprintf("http://domain/issues/%d", object.ID))

		us.SendMessage(infoOboutTask, chatID)

	} else {

		infoOboutTask := fmt.Sprintln("Задача была создана ⏰: ", object.StartDate, "\n", "Задача назначена на ⛑: ", object.AssignedTo.Name, "\n", "Название задачи 👉: ", object.Subject, "\n", "Статус задачи 🔧:", object.Status.Name, "\n", "📌 Ссылка на задачу: ", fmt.Sprintf("http://domain/issues/%d", object.ID))

		us.SendMessage(infoOboutTask, chatID)
	}

}
