package worker

import (
	"github.com/kumomoto/redminebot/app/mongodb"
)

type Worker struct {
	users *mongodb.DatabaseConnection
}

func NewWorker(usersRepository *mongodb.DatabaseConnection) *Worker {
	return &Worker{
		users: usersRepository,
	}
}

func (w *Worker) userExists(chat_id int64) (bool, error) {
	return w.users.FindUser(chat_id)
}
