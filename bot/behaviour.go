package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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
