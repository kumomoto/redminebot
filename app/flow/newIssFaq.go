package Flow

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (us *FlowUserStruct) faqAboutNewIss(chatID int64) tgbotapi.MessageConfig {

	cancelButton := tgbotapi.NewInlineKeyboardButtonData("Отмена создания обращения", "cancel")

	markup := tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{cancelButton})

	text := fmt.Sprintf("Как создавать задачи через данного бота 🛠🛠🛠: \n \n ✉️✉️✉️ Следующим сообщением просим вас отправить только название обращения. \n \n 🧷📁📒 Далее просим описать проблему как можно подробнее и отправить его в канал. \nПосле того, как вы отправили в канал полное описание вашей проблемы вы можете прикрепить файлы, которые могут помочь решить проблему и также отправить их в канал. \n\n❗️❗️❗️ При загрузке фотографий просим вас не сжимать изображения, или же отправлять сжатые фотографии ТОЛЬКО формата JPG! Файлы же поддерживаются с любым расширением. \n\n Если вы забыли что-то написать в описании проблемы, то вы можете дописать описание задачи отправив новое текстовое сообщение в канал \n\n ✏️✏️✏️ Чтобы завершить создание обращения нажмите кнопку 'Отправить' в конце создания обращения \n\n Чтобы отменить создание нового обращения вы можете в любой момент нажать на кнопку 'Отмена создания обращения'.")

	message := tgbotapi.NewMessage(chatID, text)

	message.ReplyMarkup = markup

	return message

}
