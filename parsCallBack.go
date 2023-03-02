package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	database2 "sample-app/tg2/database"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func parscallback(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	word := strings.Split(update.CallbackQuery.Data, "\n") //проверка есть ли слово
	//todo обработка ошибок
	if word[0] == "/ucatalog" {
		parsAnswerUcatalog(update, bot)
		return
	}
	//сюда бобавить переход на командыыыыыыыы (но тк присылается ответ, тут есть прикол что айди смотрится в другом месте)
	if word[0] == "/addCorzine" {
		parsReplyCommandAddCorzine(update, bot)
		return
	}
	if word[0] == "/moredetailed" {
		moredetailed(update, bot)
		return
	}

	if word[0] == "/deleteshoppcart" {
		udeleteshoppcart(update, bot)
		return
	}
	if word[0] == "/shoppcartttttttttttt" {

		idd := update.CallbackQuery.Message.MessageID //update.CallbackQuery.From.ID

		var numericKeyboardInline22 = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Очистить корзину", "/deleteshoppcart"),
			),
		)
		ms := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, idd, numericKeyboardInline22)
		bot.Send(ms)
		return
	}
	if word[0] == "/deleteshopposition" {
		deleteshopposition(update, bot)
		return
	}
	if word[0] == "/placeanorder" {
		placeanorder(update, bot)
		return
	}
	if word[0] == "/takepvz" {
		//msg type  ("/deleteshopposition\narticle:%d\nmsg:%d", article, msgid)
		//Надо написать обрабротчик команды
		//id обновление сообщг
		//idd := update.CallbackQuery.Message.MessageID //update.CallbackQuery.From.ID

		takepvz(update, bot)
		return
	}
	if word[0] == "/addcard_article" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_ARTICLE)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Введите артикул")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Введите артикул"))
		bot.Send(msg)
		return
	}

	if word[0] == "/addcard_name" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_NAME)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Введите название")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Введите название"))
		bot.Send(msg)
		return
	}
	if word[0] == "/addcard_description" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()

		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_DESCRIPTION)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Введите описание")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Введите описание"))
		bot.Send(msg)
		return
	}
	if word[0] == "/addcard_price" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()

		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_PRICE)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Введите цену")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Введите цену"))
		bot.Send(msg)
		return

	}
	if word[0] == "/addcard_photo" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_PHOTO)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Добавьте фото")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Добавьте фото"))
		bot.Send(msg)
		return
	}
	if word[0] == "/addcard_category" {
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_ADDCARD_CATEGORY)
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Введите или выберите категорию")

		//Сообщение с категориями + добавить ф-цию где будут кнопки выбираться если не нажимались
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Введите или выберите категорию"))
		bot.Send(msg)
		//новое сообщение с выбором категории
		return
	}
	if word[0] == "/addcard_showDB" {
		//добавление в базу данныххх
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_NOTHING)
		}
		//to do ДОБАВЛЕНИЕ В БД
		dbUser, err := database2.NewBD().Open()
		_ = err
		var dbproduct_temp database2.T_productDatabase
		//записываем во временную БД
		_, _ = dbUser.ReadAll_idtg(id_user, &dbproduct_temp)
		//проверка есть ли такая категория
		text := fmt.Sprintf("Артикул товара:%d\nКатегория:%s\nНазвание:%s\nОписание:%s\nЦена:%f", dbproduct_temp.Article, dbproduct_temp.Category, dbproduct_temp.Name, dbproduct_temp.Description, dbproduct_temp.Price)
		var photoFileBytes tgbotapi.FileBytes

		if dbproduct_temp.Photo != "" {
			photoBytes, err := ioutil.ReadFile(dbproduct_temp.Photo)
			_ = err
			photoFileBytes = tgbotapi.FileBytes{
				Name:  "",
				Bytes: photoBytes,
			}
		}
		var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Добавить в БД", "/addcard_addDB"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Изменить категорию", "/addcard_category"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Изменить название", "/addcard_name"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Изменить описание", "/addcard_description"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Изменить цену", "/addcard_price"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Изменить фото", "/addcard_photo"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
			),
		)

		msg1 := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "Предварительный вид: ")
		bot.Send(msg1)
		if dbproduct_temp.Photo != "" {
			msg2 := tgbotapi.NewPhotoUpload(int64(update.CallbackQuery.From.ID), photoFileBytes)
			msg2.Caption = text
			bot.Send(msg2)
		} else {
			msg2 := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), text)
			bot.Send(msg2)
		}
		msg3 := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "Выберите действие:")
		msg3.ReplyMarkup = numericKeyboardInline
		bot.Send(msg3)
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Выберите действие:")
		bot.AnswerCallbackQuery(answercall)
		//msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Выберите действие:"))
		//bot.Send(msg)
		return
	}
	if word[0] == "/addcard_addDB" {
		//добавление в базу данныххх
		id_user := update.CallbackQuery.From.ID
		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_NOTHING)
		}
		//to do ДОБАВЛЕНИЕ В БД
		dbUser, err := database2.NewBD().Open()
		var dbproduct_temp database2.T_productDatabase
		//записываем во временную БД
		_, _ = dbUser.ReadAll_idtg(id_user, &dbproduct_temp)
		//проверка есть ли такая категория

		err = dbUser.Check_table(dbproduct_temp.Category)
		if err != nil {
			//to DO если ошибка другая
			database2.Create(dbproduct_temp.Category) //todo
			//createp()
		}
		//
		dbUser.Table_name = dbproduct_temp.Category

		dbUser.Add(dbproduct_temp)
		dbUser.Table_name = "temp"
		dbUser.Delete_idtg(id_user)
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("В БД ДОБАВЛЕНО")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("В БД ДОБАВЛЕНО"))
		bot.Send(msg)
		return
	}
	if word[0] == "/addcard_cancel" {
		//добавление в базу данныххх
		id_user := update.CallbackQuery.From.ID

		database, err := database2.NewBDUser().Open()
		if err != nil {
			//todo
		}
		database.Table_name = "user"
		status, _, err := database.Bd_checkUser(id_user)
		if err != nil {
			//todo
		}
		if status == Admin {

			_ = database.EditTGid(id_user, "state", E_STATE_NOTHING)
		}
		//to do ДОБАВЛЕНИЕ В БД
		dbUser, err := database2.NewBD().Open()
		dbUser.Table_name = "temp"

		db, err := dbUser.ReadAll()
		if db != nil {
			for i := range db {
				if db[i].Id_tg == id_user {
					dbUser.Delete_idtg(id_user)
				}
			}
		}
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Вышли из режима добавления")
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("Вышли из режима добавления"))
		bot.Send(msg)
		return
	}
}
func parsAnswerUcatalog(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //todo
	tg_id := update.CallbackQuery.From.ID
	dbproduct, err := database2.NewBD().Open()

	if err != nil {
		//
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//
	}

	if len(table) == 0 {
		msg := tgbotapi.NewMessage(int64(tg_id), "Каталог пуст. Нет таблиц.")
		bot.Send(msg)
		return
	}

	var bd database2.T_productDatabase
	var catalog []string
	catalog = append(catalog, "category:")
	//receive := update.CallbackQuery.Data,
	var new string
	word, key := search_word(update.CallbackQuery.Data, catalog)
	for i, k := range key {
		switch word[k] {
		case "category:":
			new = card_getStringJson(key, i, k, update.CallbackQuery.Data, word[k])
			json.Unmarshal([]byte(new), &bd)
		}
	}
	dbproduct.Table_name = bd.Category
	db_send, err := dbproduct.ReadAll()
	//проверка струкутры типа db_send == (T_PRoduct..{})

	if len(db_send) == 0 {
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Каталог: %s пуст", bd.Category)
		bot.AnswerCallbackQuery(answercall)
		msg := tgbotapi.NewMessage(int64(tg_id), answercall.Text)
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("Каталог:%s", bd.Category))
	bot.Send(msg)
	for i := range db_send {
		//text := fmt.Sprintf("Артикул: %d\nНазвание: %s\n%s\nЦена: %0.2fрублей\n", db_send[i].Article, db_send[i].Name, db_send[i].Description, db_send[i].Price)
		text := fmt.Sprintf("Артикул: %d\nНазвание: %s\nПодходит для: \nЦена: %0.2fрублей\n", db_send[i].Article, db_send[i].Name, db_send[i].Price)
		podrobnee := fmt.Sprintf("/moredetailed\narticle:%d", db_send[i].Article)
		if db_send[i].Photo != "" {
			photoBytes, err := ioutil.ReadFile(db_send[i].Photo)
			_ = err
			photoFileBytes := tgbotapi.FileBytes{
				Name:  "",
				Bytes: photoBytes,
			}

			msg := tgbotapi.NewPhotoUpload(update.CallbackQuery.Message.Chat.ID, photoFileBytes)
			sss := fmt.Sprintf("/addCorzine\narticle:%d\ncategory:%s", db_send[i].Article, bd.Category) //todo
			msg.Caption = text

			var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Подробнее", podrobnee),
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
				),
			)
			/*//example how do button
			var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonURL("Добавить в корзину", "/addCorzine\n"+string(db[i].Num)),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("В наличии:"+string(db[i].Instock), ""),
				),
			)*/

			msg.ReplyMarkup = numericKeyboardInline
			bot.Send(msg)
		} else { //если нет фото

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
			sss := fmt.Sprintf("/addCorzine\narticle:%d\ncategory:%s", db_send[i].Article, bd.Category) //todo
			var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Подробнее", podrobnee),
					tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
				),
			)
			msg.ReplyMarkup = numericKeyboardInline
			bot.Send(msg)
		}
	}
	var answercall tgbotapi.CallbackConfig
	answercall.CallbackQueryID = update.CallbackQuery.ID
	answercall.Text = fmt.Sprintf("Каталог:'%s'", bd.Category)
	bot.AnswerCallbackQuery(answercall)
	//update.Message.Text = "больше сюда не тыкай"
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//bot.Send(msg)
}
