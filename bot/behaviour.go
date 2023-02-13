package bot

import (
	"fmt"
	"github/mrqwer/slangTelBot/database"
	"io/ioutil"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
)

func startCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "")
	msg.Text = "Меня зовут SlangITBot\n" +
		"Я занимаюсь переводом сленговых\n" +
		"IT терминов в формальную речь\n" +
		"Выбери один из вариантов ⬇⬇⬇"
	menuCommand(bot, &msg)
}

func stopCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Пожалуйста, выберите один из нижеперечисленных\nвариантов для дальнейшего взаимодействия")
	//msg.Text = "Пожалуйста, введите команду /start, чтобы увидеть меню возможных действий"
	menuCommand(bot, &msg)
}

func menuCommand(bot *tgbotapi.BotAPI, MsgConfig *tgbotapi.MessageConfig) {
	MsgConfig.ReplyMarkup = optionKeyboard

	if _, err := bot.Send(MsgConfig); err != nil {
		log.Fatal(err)
	}
}

func defaultBehaviour(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Пожалуйста, введи команду /menu и выберите один из вариантов")
	if _, err := bot.Send(msg); err != nil {
		log.Fatal(err)
	}
}

func sendDocs(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//file, err := os.Open("article/a1.pdf")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//defer file.Close()

	//fileData, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//var pdfFile multipart.File = ioutil.NopCloser(bytes.NewBuffer(fileData))

	//pdfFile := tgbotapi.FileBytes{Name: "a1.pdf", Bytes: fileData}
	//_, err = bot.Send(tgbotapi.NewDocumentUpload(update.CallbackQuery.Message.Chat.ID, pdfFile))
	//if err != nil {
	//	panic(err)
	//}
	//media := tgbotapi.NewMediaGroup(update.CallbackQuery.Message.Chat.ID, pdfFile)

	for i := 1; i <= 5; i++ {

		docBytes, err := ioutil.ReadFile(fmt.Sprintf("articles/a%v.pdf", i))
		if err != nil {
			log.Fatal(err)
		}
		docFileBytes := tgbotapi.FileBytes{Name: fmt.Sprintf("Статья%v", i), Bytes: docBytes}
		_, err = bot.Send(tgbotapi.NewDocumentUpload(update.CallbackQuery.Message.Chat.ID, docFileBytes))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func countMessage(inputMessage *tgbotapi.MessageConfig) {

	countFilter := bson.M{}
	countFilterData := database.CountCollection(database.Dictionary, countFilter)
	inputMessage.Text = "*На данный момент у нас " + strconv.Itoa(countFilterData) + " записи*"
}
