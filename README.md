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

/app/bot:

bot.go:
-- Bot struct: Contains API, connect parameters and link on Worker of user
-- Init(): Initial connect to bot and set parameters of updates, Api and timeouts
channelsDistrib.go
-- addChannelsToMap(): add new channels to map for monitoring (open or close)
-- checkOpenChannels(): check open channels else close them
-- getStatusChannels(): get status if channel from user mongo collection
-- distributionCallbackBetweenUsersChannels(): distribution of User Callbacks(update of channel when user click on the InlineButton)
-- distributionMessagesBetweenUsersChannels(): distribution of User Messages(update of channel when user click on 'Предоставить номер')
processUpdate.go
-- Start(): start of handle updates of channel
-- processCallback(): handle input Callback of user. Send an update to a specific channel from the channel map. "waitSubject", "waitDesc" , etc - step where needet data from user  
-- processMessage(): handle input Message of user. Send an update to a specific channel from the channel map. "reg", "waitAttach", etc - step where needet data from user
-- getStatusUser(): get current step of user


/app/flow:

buttonsMenu.go:
-- CreateButtonsMainMenu(): Send mess with buttons for main menu
-- CreateButtonsStatIssMenu(): Send mess with buttons for menu after getting status of iss
-- CreateButtonsNewIssMenu(): Send mess with buttons for menu after create new iss

editMessag.go:
-- editNewIss(): rewrite created iss

flowinit.go:
-- NewFlow(): create new flow for user request
-- SetApiAndUpdates(): set API for communication with user
-- Close(): close active flow

getTaskFromRM.go:
-- getSelectedTask(): get specific iss object form RM

MessagesWithButtons.go:
-- SendMessage(): send specific string as message to user
-- SendMessageWithButtons(): send cpecific message with specific buttons to user

messagesWithIssInfo.go:
-- sendInfoAboutTask(): send info about specific iss form RM

newIssFaq.go:
-- faqAboutNewIss(): send FAQ "How to create new iss"

processCreateIssue.go:
-- FlowProcessCreateIssue(): create new iss in RM. 
Steps:
1. Title
2. Description of new iss
3. Upload Files

-- SetStepUser(): set cpecific step of user in mongo collection
-- SetChannelclosedState(): set status "close" in mongo collection
-- getSendButton(): create marcup with buttons for send of reedit iss
-- setAuthors(): set author RM ID in created iss

processStatusIss.go:
-- FlowProcessStatusIssue(): get info about specific iss from RM
Steps:
1. Get all iss of current user
2. Whait data from user(the number of necessary iss)
3. Send info in chat with user

-- checkAmountIssues(): if amount iss = zero, then send mess obout zero iss
-- getReplMarkupStatusIssue(): get markup with iss as buttons

pushFiles.go:
-- pushDocuments(): push docs to RM server
-- pushPhotos(): push photos to RM server

whaitResponce.go:
-- whaitSubjectMessage(): wait data from user on "waitSubject" step
-- waitDeskMessage(): wait message with desk created iss
-- waitResponceCallback(): wait Callback update
-- waitResponceWithFiles(): wait message with files

/app/ldap
ldap.go:
-- NewLdapConnection(): create new conn with DC
-- userSearchRequest(): create new search request for get user data
-- GetAuthResult(): return auth result
-- GetMail(): return mail of user

/app/mongo:

mongodb.go:
-- User sruct: Contain mongo user data
-- DatabaseConnection struct: DB connection data
-- Init(): Init connection
-- FindUser(): Find user in DB
-- GetUser(): Get collection of user
-- CreateUser(): Create new user in DB
-- UpdateUser(): Update user collection
-- DeleteIssue(): Delete issue monitoring
-- GetIssues(): Get all issues of user

/app/redmine:
addNewTask.go:
-- AddNewTask(): create new iss in RM

deleteIssues.go:
-- DeleteClosedIssuesFromCollection(): check closed iss and set 'false' in mongo collection

newTaskWithAttachments.go
-- AddNewTaskWithFiles(): create iss with attachments from user

redmineInit.go:
-- Init(): new RM connection
-- DownloadfilesForIssue(): download files for created new iss
-- GetInfoTask(): return Iss object
-- issueCheckStatus(): return false if iss is close

users.go:
-- GetRMUserID(): return ID and name of user with specific mail address

/app/uploadFiles:
createDir.go:
-- createDirOfUser(): create dir for store files from user

createFileObject.go:
-- createFileObject(): save file in file system of host

createPhotoObject.go:
-- createPhotoObject(): save photo in file system of host

downloadFilesFromTg.go:
-- dowloadDocsFromUserToServer(): download file from TG URL on 555 host

downloadPhotosFromTg.go:
-- downloadPhotosFromUserToServer(): download photo from TG URL on 555 host
-- getPhotoFileID(): get ID photo from TG

flowFileInit.go:
-- NewFileFlow(): create flow for process photos
-- NewDocsFlow(): create flow for process docs
-- setDocsSizeAndApi(): set doc data in flow

pushFiles.go:
-- PushDocsToRedmine(): return attachments with files for create iss

pushPhoto.go:
-- PushPhotoToRedmine(): return attachments with photos for create iss

uploadDocsToRm.go:
-- downloadDocOnServer(): download files on RM server

uploadPhotoOnRm.go:
-- downloadPhotosOnServer(): download photos on RM server

/app/worker:

processCallback.go:
-- ProcessCallback(): handle request for create iss or get iss info

processStart.go:
-- ProcessStartMessage(): hande start message

register.go:
-- tryRegister(): create new user in mongo if ad auth ok
-- whaitContact(): wait data with telephone number form user
-- sendMessage(): send specific message to user

worker.go:
-- NewWorker(): create new worker for user
-- userExist(): check user in mongo collections

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
                         v
                     AD Auth ok? NO >---------------------- Deny
                         | OK
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
