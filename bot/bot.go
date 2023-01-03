package bot

import (
	"github/mrqwer/slangTelBot/database"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
)

func initialUpdateConf(bot tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updateConfig.Limit = 1
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}
	return updates
}

func Bot() {

	// Get BotToken
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	// end

	// Connecting and getting channel of updates
	log.Printf("Authorized on account %s", bot.Self.UserName)
	updates := initialUpdateConf(*bot)
	// end

	// traversing through channels of updates
	for update := range updates {

		if update.Message != nil {
			log.Printf("Chat ID %d", update.Message.Chat.ID)
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				switch cmdText {
				case "start":
					startCommand(bot, update.Message)
				case "menu":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "меню возможных действий")
					menuCommand(bot, &msg)
				case "stop":
					flag = 0
					stopCommand(bot, update.Message)
				}
			} else {
				if flag == 0 {
					defaultBehaviour(bot, update.Message)
				} else if flag == 1 {
					flag = 0
				} else if flag == 2 {
					//filterDoc := bson.M{"$text": bson.M{"$search": "комп"}}
					//temp, err := database.GetMongoDoc(database.Dictionary, filterDoc)
					//if err != nil {
					//	log.Fatal(err)
					//}
					//msg := tgbotapi.NewMessage(update.Message.Chat.ID, temp.Standard+": "+formatPrint(temp.Slang)+"\nЕсли хотите остановить, то введите команду /stop")

					s, err := findStandard(update.Message.Text)
					if err != nil {
						log.Printf("Cannot run find algorithm\n%v", err)
					}

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, s)
					if _, err := bot.Send(msg); err != nil {
						log.Fatal(err)
					}
				}
			}

		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.AnswerCallbackQuery(callback); err != nil {
				log.Printf("Problem when getting answer of callbackquery \n%v", err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			switch update.CallbackQuery.Data {
			case "dict":
				msg.Text = d
				flag = 1
			case "slang":
				msg.Text = "Пожалуйста, напишите сленг\n" +
					"Я попробую найти"
				flag = 2
			case "data":
				msg.Text = "Краткая академическая информация про it сленги...\nСписок статей для детального ознакомления"
				msg.ReplyMarkup = optionKeyboard1
			case "count":
				countFilter := bson.M{}
				countFilterData := database.CountCollection(database.Dictionary, countFilter)
				msg.Text = "На данный момент у нас " + strconv.Itoa(countFilterData) + " записей"
				flag = 0
			}
			if _, err := bot.Send(msg); err != nil {
				log.Printf("This is the error from callbacksend part \n%v", err)
			}
		}
	}
}
