package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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
