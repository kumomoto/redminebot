package mongodb

import (
	"log"

	"github.com/kumomoto/redminebot/app/conf"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Сonnection DatabaseConnection

type User struct {
	Chat_ID         int64  //Chat ID number
	Phone_Number    string //Phone number users
	Mail            string
	StatusOfChannel bool
	Step            string //Place of user
	Issues          map[string]bool
	Author          *Author
}

type DatabaseConnection struct {
	Session *mgo.Session  //Server connetcion
	DB      *mgo.Database //Connect to DB
}

type Author struct {
	ID   int64
	Name string
}

// Init DB Connection
func (connection *DatabaseConnection) Init() {
	session, err := mgo.Dial(conf.MONGODB_CONNECTION_URL) //Init server connection
	if err != nil {
		log.Fatal(err) //Print err end stop program
	}
	connection.Session = session
	db := session.DB(conf.MONGODB_DATABASE_NAME) //DB connection
	connection.DB = db
}

// Find User
func (connection *DatabaseConnection) FindUser(chat_id int64) (bool, error) {
	//Find Users
	collection := connection.DB.C(conf.MONGODB_COLLECTION_USERS)      //Get Named List
	count, err := collection.Find(bson.M{"chat_id": chat_id}).Count() //Find User
	if err != nil || count == 0 {
		return false, err
	} else {
		return true, err
	}
}

// Get user
func (connection *DatabaseConnection) GetUser(chat_id int64) (User, error) {
	//Get Users
	var result User
	find, err := connection.FindUser(chat_id) //Find chat_id of users
	if err != nil {
		return result, err
	}
	if find {
		collection := connection.DB.C(conf.MONGODB_COLLECTION_USERS) // If find is true - get collection
		err = collection.Find(bson.M{"chat_id": chat_id}).One(&result)
		return result, err
	} else {
		return result, mgo.ErrNotFound
	}
}

// Create new User
func (connection *DatabaseConnection) CreateUser(user User) error {
	collection := connection.DB.C(conf.MONGODB_COLLECTION_USERS)
	err := collection.Insert(user)
	return err
}

// Change User Phone Number
func (connection *DatabaseConnection) UpdateUser(user User, chatID int64) error {
	collection := connection.DB.C(conf.MONGODB_COLLECTION_USERS)
	err := collection.Update(bson.M{"chat_id": chatID}, &user)
	return err
}

func (connection *DatabaseConnection) DeleteIssue(chat_id int64, iss string) {
	oldusr, err := connection.GetUser(chat_id)
	if err != nil {
		log.Panic(err)
	}
	oldusr.Issues[iss] = false
	err = connection.UpdateUser(oldusr, chat_id)
	if err != nil {
		log.Panic(err)
	}
}

func (connection *DatabaseConnection) GetIssues(chat_id int64) []string {
	var arr []string
	usr, err := Сonnection.GetUser(chat_id)
	if err != nil {
		log.Panic(err)
	}
	for str, v := range usr.Issues {
		if v == true {
			arr = append(arr, str)
		}
	}
	return arr
}
