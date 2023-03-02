package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func parsTextCommand(update *tgbotapi.Update, status int, bot *tgbotapi.BotAPI) (bool, string) {
	switch status {
	case User:
		msg_status, msg := textUser(update, bot)
		return msg_status, msg
	case Admin:
		msg_status, msg := textAdmin(update, bot)
		return msg_status, msg
	}

	return true, "ERR STATUS 2"
}

// state_addcard - 100-199

func textUser(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {
	switch update.Message.Text {
	case "Очистить корзину":
		udeleteshoppcart(update, bot)
		return false, ""
	case "Корзина":
		ushoppcart(update, bot) //for user
		return false, ""
	case "Каталог":
		ucatalog(update, bot)
		return false, ""
	case "Мои заказы":
		umyorders(update, bot)
		return false, ""
	case "FAQ":
		ufaq(update, bot) //for user
		return false, ""
	case "HELP":
		uhelp(update, bot) //for user
		return false, ""
	default:
		msg := "Такой команды нет"
		return true, msg
	}
}
func textAdmin(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {
	//Добавить автоматически команнды юзера,чтоб не дублировать

	switch update.Message.Text {
	case "Настройки админа":
		asetting(update, bot)
		return false, ""
	case "Очистить корзину":
		udeleteshoppcart(update, bot)
		return false, ""
	case "Корзина":
		ushoppcart(update, bot) //for user
		return false, ""
	case "Каталог":
		ucatalog(update, bot)
		return false, ""
	case "Мои заказы":
		umyorders(update, bot)
		return false, ""
	case "FAQ":
		ufaq(update, bot) //for user
		return false, ""
	case "HELP":
		uhelp(update, bot) //for user
		return false, ""
	case "Как пользоваться?":
		ahowuse(update, bot)
		//uhelp(update, bot) //for user
		return false, ""
	case "Info:/addcardv1":
		//uhelp(update, bot) //for user
		return false, ""

	default:

		msg := "Такой команды нет"
		return true, msg
	}
}

// var numericKeyboard = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Как пользоваться?"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Info:/addcardv1"),
// 		tgbotapi.NewKeyboardButton("Info:/addcard"),
// 		tgbotapi.NewKeyboardButton("Info:/deletecard"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Info:/doingadmin"),
// 		tgbotapi.NewKeyboardButton("Info:/deleteadmin"),
// 		//tgbotapi.NewKeyboardButton("6"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("/button"),
// 		//tgbotapi.NewKeyboardButton("Мои заказы"),
// 		//tgbotapi.NewKeyboardButton("Корзина"),
// 	),
// )
