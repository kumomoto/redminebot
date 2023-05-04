<span dir="">#</span>RedmineBot

Conf - Has a conf file with global variables:

Telegram Bot Settings:

- TELEGRAM_BOT_API_KEY: API key for connect to bot
- TLEGRAM_BOT_UPDATE_OFFSET: time of long polls
- TELEGRAM_BOT_UPDATE_TIMEOUT: timeout of new updates

Monbodb settings:

- MONGODB_CONNECTION_URL: End Point to connect DB
- MONGODB_DATABASE_NAME: Name of DBs
- MONGODB_COLLECTION_USERS: Name collection of users

Redmine settings:

- REDMINE_ENDPOINT: End Point for connect
- REDMINE_APIKEY: API key of user which has credentials to create new issue
- REDMINE_PROJECTID: ID of project IT outsorsing

---

Files:

/src/api:

- TGBotApi.go: 
- -- Bot struct: Contains API, connect parameters and link on Worker of user 
- -- Init(): Initial connect to bot and set parameters of updates, Api and timeouts 
- -- Start(): Start process user messages

/src/bot: First version of TGBotApi.go /src/flow:

- Flow.go: 
- -- FlowUserStruct struct: User Object from Mongo 
- -- NewFlow(): Open new flow for User 
- -- getNewIssStruct(): Link on new redmine Issue object 
- -- SetApiAndUpdates(): Installing Api and Updates for Flow Variables 
- -- FlowProcessCreateIssue(): Process request "Create new Iss" 
- -- getSelectedTask(): Return object of selected task 
- -- sendInfoAboutTask(): Sending a message to the user with data that contains the selected problem 
- -- sendMessageForGetTaskInfo(): Sending message with markup of number issues for get selected issues 
- -- getReplMarkupStatusIssue(): Get buttons with Numbers of issues 
- -- checkAmountIssues(): Checks the number of user problems created 
- -- setStatus(): Set a custom spot in the bot 
- -- SendMessage(): Sending message for user 
- -- SendMessageWithButtons(): Sending message with buttons 
- -- CreateButtonsStatIssMenu(): Create buttons if user in "Status Issue" menu 
- -- CreateButtonsNewIssMenu(): Create buttons if user in "New Issue" menu 
- -- CreateButtonsMainMenu(): Create buttons if user in "Main" menu 
- -- WaitResponce(): Waiting update from user

/src/mongo:

- mongodb.go: 
- -- User sruct: Contain mongo user data 
- -- DatabaseConnection struct: DB connection data 
- -- Init(): Init connection 
- -- Find(): Find user in DB 
- -- GetUser(): Get collection of user 
- -- CreateUser(): Create new user in DB 
- -- UpdateUser(): Update user collection 
- -- DeleteIssue(): Delete issue monitoring 
- -- GetIssues(): Get all issues of user

/src/redmine:

- redmine.go: 
- -- Task struct: Info obout issue 
- -- Init(): Init new connection 
- -- AddNewTask(): Adding new iss to redmine 
- -- GetInfoTask(): Get selected issue object 
- -- DeleteClosedIssuesFromCollection(): Кemoval of closed issues from the collection 
- -- issueCheckStatus(): Check status of issue

/src/worker:

- worker.go 
- -- Worker struct: Contain data of user 
- -- NewWorker(): Create new worker for Process Message of user 
- -- ProcessMessage(): Process new message of user 
- -- userExists(): Check user exist 
- -- addNewUserToMongo(): Add new collection of new user 
- -- createRegisterButton(): Create button for get number of user 
- -- sendRegisterMessage(): Sent register message with buttons

Bot work scheme:

```
                 main (start bot)
                         |
                         |
                         |
                         |
                         v
                 Create new Worker
                         |
                         |
                         |
                         |
                         v
                 Process Message <--------------------------<|
                         |                                   |
                         |                                   |
                         |                                   |
                         |                                   |
                         v                                   |
         Exist <--- If user is ---> Not exist                |
           |                            |                    |
           |                            |                    |
           |                            |                    |
           |                            |                    |
           v                            v                    |
           |                         Register                |
           |                            |                    |
           v                            v                    |
           ------------------------------                    |
                         |                                   |
                         |                                   |
                         |                                   |
                         |                                   |
                         v                                   |
                  New Flow of User                           |
       _________________________________________             |
      |                   |                     |            |
      |                   |                     |            |
      |                   |                     |            |
      |                   |                     |            |
      v                   v                     v            |
Create new issue  See status of issue     Back to main       |
|                                                     |      |
Another New Iss? See Status Another Iss? Back to main?       |
                           |                                 |
                           |                                 |
                           |                                 |
                           v                                 |
                           IF                                |
                           |                                 |
                           |                                 |
                           |                                 |
                       v--------v                            |
Wait new message <----NO         YES ------------------------^
```
