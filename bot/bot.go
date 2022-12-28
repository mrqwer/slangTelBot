package bot

import (
	"fmt"
	"github/mrqwer/slangTelBot/database"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var numericKeyboard = tgbotapi.NewReplyKeyboard(
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("1"),
//		tgbotapi.NewKeyboardButton("2"),
//		tgbotapi.NewKeyboardButton("3"),
//	),
//	tgbotapi.NewKeyboardButtonRow(
//		tgbotapi.NewKeyboardButton("4"),
//		tgbotapi.NewKeyboardButton("5"),
//		tgbotapi.NewKeyboardButton("6"),
//	),
//)

var (
	dictionary = map[string][]string{
		"компьютер": []string{"компухтер", "комп", "копи", "железо", "машина"},
		"линукс":    []string{"лин", "пингвин", "гнулинукс"},
	}

	d = func(map[string][]string) string {
		s := ""
		for k, v := range dictionary {
			s += string(k) + ": "
			for i := range v {
				if i == len(v)-1 {
					s += v[i]
				} else {
					s += v[i] + ","
				}
			}
			s += "\n"
		}
		return s
	}(dictionary)
)

var optionKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Словарь 10 самых популярных формальных слов c сленгами", "dict"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Cленг -> Формальная речь", "slang"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Об источниках и по дополнительной информацией", "data"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сколько записей сейчас в словаре", "count"),
	),
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🖊 Сленг -> Формальная речь"),
		tgbotapi.NewKeyboardButton("🖍 Формальная речь -> Сленг"),
	),
)

var optionKeyboard1 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("1 статья", "https://astanait-my.sharepoint.com/personal/211533_astanait_edu_kz/_layouts/15/onedrive.aspx?login_hint=211533%40astanait%2Eedu%2Ekz&id=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9%2Fslang5%2Epdf&parent=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("2 статья", "https://astanait-my.sharepoint.com/:w:/r/personal/211533_astanait_edu_kz/_layouts/15/Doc.aspx?sourcedoc=%7B6CA57C31-021C-4CE7-9C48-A983BDF1A52B%7D&file=slang2.doc&action=default&mobileredirect=true"),
	),

	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("3 статья", "https://astanait-my.sharepoint.com/personal/211533_astanait_edu_kz/_layouts/15/onedrive.aspx?login_hint=211533%40astanait%2Eedu%2Ekz&id=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9%2Fslang5%2Epdf&parent=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("4 статья", "https://astanait-my.sharepoint.com/personal/211533_astanait_edu_kz/_layouts/15/onedrive.aspx?login_hint=211533%40astanait%2Eedu%2Ekz&id=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9%2Fslang4%2Epdf&parent=%2Fpersonal%2F211533%5Fastanait%5Fedu%5Fkz%2FDocuments%2F%D1%81%D0%BF%D0%B8%D1%81%D0%BE%D0%BA%20%D1%81%D1%82%D0%B0%D1%82%D0%B5%D0%B9"),
	),
)

var (
	flag int
)

func formatPrint(arr []string) string {
	s := ""
	for _, v := range arr {
		s += (v + " ")
	}
	return s
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

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updateConfig.Limit = 1
	updates := bot.GetUpdatesChan(updateConfig)
	// end

	// traversing through channels of updates
	for update := range updates {

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы ввели непонятное сообщение")
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				switch cmdText {
				case "start":
					msg.Text = "Меня зовут SlangITBot\n" +
						"Я занимаюсь переводом сленговых\n" +
						"IT терминов в формальную речь\n" +
						"Выбери один из вариантов ⬇⬇⬇"
					msg.ReplyMarkup = optionKeyboard
					if _, err := bot.Send(msg); err != nil {
						log.Fatal(err)
					}
				case "menu":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msg.ReplyMarkup = mainMenu

					if _, err := bot.Send(msg); err != nil {
						log.Fatal(err)
					}
				case "test":
					findOptions := options.Find()
					temp, err := database.GetMongoDocs(database.Dictionary, bson.M{}, findOptions)
					//temp, err := database.GetMongoDoc(database.Dictionary, bson.M{"name": "компьютер"})
					if err != nil {
						log.Fatal(err)
					}
					str := ""
					for _, v := range *temp {
						str += v.Standard + ":\t" + formatPrint(v.Slang) + "\n"
					}
					msg.Text = fmt.Sprintf("%v \n", str)
					if _, err := bot.Send(msg); err != nil {
						log.Fatal(err)
					}
				case "stop":
					flag = 0
				}
			} else {
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				if flag == 0 {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, введи команду /start и выберите один из вариантов")
				} else if flag == 1 {
					flag = 0
				} else if flag == 2 {
					filterDoc := bson.M{"$text": bson.M{"$search": "комп"}}
					temp, err := database.GetMongoDoc(database.Dictionary, filterDoc)
					if err != nil {
						log.Fatal(err)
					}
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, temp.Standard+": "+formatPrint(temp.Slang)+"\nЕсли хотите остановить, то введите команду /stop")
					// logic is here
					// if /stop then flag = 0
				}
				//switch update.Message.Text {
				//case "start":
				//	msg.ReplyMarkup = optionKeyboard
				//	msg.Text = "Меня зовут SlangITBot\n" +
				//		"Я занимаюсь переводом сленговых\n" +
				//		"IT терминов в формальную речь\n" +
				//		"Выбери один из вариантов ⬇⬇⬇"
				//default:
				//	if update.Message.Text == mainMenu.Keyboard[0][1].Text {
				//		msg.Text = "что еще хотите быстро сьебались отсюда"
				//	} else if update.Message.Text == mainMenu.Keyboard[0][0].Text {
				//		msg.Text = "ok"
				//	}
				//}
				if _, err := bot.Send(msg); err != nil {
					log.Fatal(err)
				}
			}

		} else if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			if update.CallbackQuery.Data == "dict" {
				msg.Text = d
				flag = 1
			} else if update.CallbackQuery.Data == "slang" {
				msg.Text = "Пожалуйста, напишите сленг\n" +
					"Я попробую найти"
				flag = 2
			} else if update.CallbackQuery.Data == "data" {
				msg.Text = "Краткая академическая информация про it сленги...\nСписок статей для детального ознакомления"
				msg.ReplyMarkup = optionKeyboard1
			} else if update.CallbackQuery.Data == "count" {
				countFilter := bson.M{}
				countFilterData := database.CountCollection(database.Dictionary, countFilter)
				msg.Text = "На данный момент у нас " + strconv.Itoa(countFilterData) + " записей"
				flag = 0
			}
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}
