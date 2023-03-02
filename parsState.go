package main

import (
	"sample-app/tg2/database"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func aparsState(state int, update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {

	//Добавить автоматически команнды юзера,чтоб не дублировать
	switch state {
	case E_STATE_ADDCARD_ARTICLE:
		//asetting(update, bot)
		state_addcard_article(update, bot)
		return false, ""
	case E_STATE_ADDCARD_CATEGORY: //будет выдавать ответ с категориями которые щас есть, либо добавьте новую
		state_addcard_category(update, bot)
		//ufaq(update, bot) //for user
		return false, ""
	case E_STATE_ADDCARD_NAME:
		//udeleteshoppcart(update, bot)
		state_addcard_name(update, bot)
		return false, ""
	case E_STATE_ADDCARD_DESCRIPTION:
		//ushoppcart(update, bot) //for user
		state_addcard_description(update, bot)
		return false, ""
	case E_STATE_ADDCARD_PRICE:
		//ucatalog(update, bot)
		state_addcard_price(update, bot)
		return false, ""
	case E_STATE_ADDCARD_PHOTO:
		//umyorders(update)
		state_addcard_photo(update, bot)
		return false, ""

	default:
		msg := "Такой команды нет"
		return true, msg
	}
}
func state_addcard_article(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	arcticle, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		//Если ошибка попробовать //todo
	}
	var db database.T_productDatabase
	db.Article = arcticle
	db.Id_tg = update.Message.From.ID
	dbUser, err := database.NewBD().Open()
	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	//toDO to do //Если вызываешь второй раз то обновляется

	dbUser.Edit_idtg(db.Id_tg, "Article", db.Article)

	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить категорию", "/addcard_category"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Артикул добавлен")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
}
func state_addcard_category(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	id_tg := update.Message.From.ID
	dbUser, err := database.NewBD().Open()
	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	var dbread database.T_productDatabase
	_, err = dbUser.ReadAll_idtg(id_tg, &dbread)
	if err != nil {
		//
	}

	dbread.Category = update.Message.Text //ne nado
	dbUser.Edit_idtg(id_tg, "category", dbread.Category)
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить название", "/addcard_name"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Категория добавлена")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
	//return nil

}

func state_addcard_price(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	id_tg := update.Message.From.ID
	dbUser, err := database.NewBD().Open()

	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	var dbread database.T_productDatabase
	_, err = dbUser.ReadAll_idtg(id_tg, &dbread)
	if err != nil {
		//
	}

	price, _ := strconv.ParseFloat(update.Message.Text, 64)
	//article := dbread[len(dbread)-1].Article
	dbread.Price = price //ne nado
	dbUser.Edit_idtg(id_tg, "price", price)
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить фото", "/addcard_photo"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Цена добавлена")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
	//return nil

}

func state_addcard_description(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	id_tg := update.Message.From.ID
	dbUser, err := database.NewBD().Open()
	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	var dbread database.T_productDatabase
	_, err = dbUser.ReadAll_idtg(id_tg, &dbread)
	if err != nil {
		//
	}

	dbread.Description = update.Message.Text
	dbUser.Edit_idtg(id_tg, "description", dbread.Description)
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить цену", "/addcard_price"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Описание добавлено")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
	//return nil

}
func state_addcard_name(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	id_tg := update.Message.From.ID
	dbUser, err := database.NewBD().Open()
	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	var dbread database.T_productDatabase
	_, err = dbUser.ReadAll_idtg(id_tg, &dbread)
	if err != nil {
		//
	}

	//article := dbread[len(dbread)-1].Article
	dbread.Name = update.Message.Text
	dbUser.Edit_idtg(id_tg, "name", dbread.Name)
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить описание", "/addcard_description"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Название добавлено")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
	//return nil

}
func state_addcard_photo(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {

	//будем присылать кнопку для добавления следующих позиций с нужным id
	id_tg := update.Message.From.ID
	if update.Message.Photo == nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фото нет (можно добавить без него), выберите действие:")
		var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить фото", "/addcard_photo"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Предварительный просмотр", "/addcard_showDB"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
			),
		)
		msg.ReplyMarkup = numericKeyboardInline
		bot.Send(msg)
		return
	}
	dbUser, err := database.NewBD().Open()
	//table_name := "temp" //записываем во временную БД
	dbUser.Table_name = "temp"
	var dbread database.T_productDatabase
	_, err = dbUser.ReadAll_idtg(id_tg, &dbread)
	//_ = dbread
	if err != nil {
		//
	}

	a := (update.Message.Photo)
	b := len(*a)
	c := (*a)[b-1]
	var file tgbotapi.FileConfig
	file.FileID = c.FileID
	outfile, _ := bot.GetFile(file)
	filename, err := savefile(outfile)
	dbUser.Edit_idtg(id_tg, "photo", filename)
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Предварительный просмотр", "/addcard_showDB"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фото добавлено")
	msg.ReplyMarkup = numericKeyboardInline
	bot.Send(msg)
	//return nil

}

// var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Добавить категорию", "/addcard_category"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Добавить название", "/addcard_name"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Добавить описание", "/addcard_description"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Добавить цену", "/addcard_price"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Добавить фото", "/addcard_photo"),
// 	),
// )

//
////////////////
