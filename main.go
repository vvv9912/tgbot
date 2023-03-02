package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sample-app/tg2/database"
	"sample-app/tg2/mydadata"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	Nouser int = -1
	Admin  int = 1
	User   int = 2
)
const (
	E_STATE_NOTHING             int = 0
	E_STATE_ADDCARD_ARTICLE         = 100
	E_STATE_ADDCARD_NAME            = 101
	E_STATE_ADDCARD_DESCRIPTION     = 102
	E_STATE_ADDCARD_PRICE           = 103
	E_STATE_ADDCARD_PHOTO           = 104
	E_STATE_ADDCARD_CATEGORY        = 105
	E_STATE_ADDCARD_SHOW            = 106
	E_STATE_ADMIN_MAX               = 499
)
const (
	E_STATE_GETUSER_GEOLOCATION int = 500
	E_STATE_GETUSER_DOINGPAY    int = 501
)

//var adminap int

func main() {
	//yageocoder.Geocoder()
	//Добавляю себя админом в бд
	adminap := 143616120
	err := database.CreateDBUser("user")
	if err != nil {
		fmt.Println("create:", err)
	}
	database.Create("temp")
	dbUser, _ := database.NewBDUser().Open()
	//_ = dbUser.EditTGid(adminap, "state", E_STATE_NOTHING) //выставляю состояние //todo
	//добавляю себя админом begin
	var a database.T_Database
	dbUser.Table_name = "user"
	id, err := dbUser.Read(adminap, &a)
	a.Status = Admin
	dbUser.Edit(id, "status", a.Status)
	//end

	bot, err := tgbotapi.NewBotAPI(get_token())
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//todo переход на хуки
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1 //60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		//каждый запрос в свою рутину
		//ограничение поставить 1 запрос 1 рутина и ждем освобождение юзера

		go func(update tgbotapi.Update) {
			if update.Message == nil { // ignore any non-Message Updates
				if update.CallbackQuery != nil {
					parscallback(&update, bot)
					//todo обработка ошибок
					return
				} else if update.InlineQuery != nil {
					//todo обработка ошибок
					textFrom := update.InlineQuery.Query
					adr := mydadata.Dadata(textFrom, Dadataapikey, Dadatasecretkey)
					if adr == nil {
						return //todo
					}
					var inlineconf tgbotapi.InlineConfig
					var articles []interface{}
					inlineconf.CacheTime = 0
					inlineconf.InlineQueryID = update.InlineQuery.ID
					inlineconf.IsPersonal = true //?
					for i := range adr {
						msg := tgbotapi.NewInlineQueryResultArticleMarkdown(fmt.Sprintf("%d", i), adr[i], adr[i])
						msg.Description = adr[i]
						articles = append(articles, msg)

					}

					inlineconf.Results = append(inlineconf.Results, articles...)

					//res, err := bot.AnswerInlineQuery(inlineconf)
					//_ = res //todo
					//_ = err //todo
					bot.AnswerInlineQuery(inlineconf)

					return

				} else {
					return
				}

			}
			status, state, err := checkUser(&update, bot) //вся логика
			_ = err
			msg_status, msg_err := checkCommand(&update, status, state, bot)
			if msg_status {
				//формировать сообщение вывода
				update.Message.Text = msg_err
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				//msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			}
			return
		}(update)
	}
}

var errNOfoundusr = errors.New("uesr not founded")

// /////////////////////////////////////////////////////////////
func parsReplyCommandAddCorzine(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var dbcorzinenew database.T_productDatabase
	var new string
	var slova []string

	slova = append(slova, "article:")
	slova = append(slova, "category:")

	word, key := search_word(update.CallbackQuery.Data, slova)
	for i, k := range key {
		switch word[k] {
		case "article:":
			new = card_getIntJson(key, i, k, update.CallbackQuery.Data, word[k])
			json.Unmarshal([]byte(new), &dbcorzinenew)
		case "category:":
			new = card_getStringJson(key, i, k, update.CallbackQuery.Data, word[k])
			json.Unmarshal([]byte(new), &dbcorzinenew)
		}
	}
	var dbcorzineold database.T_Database
	dbUser, err := database.NewBDUser().Open()
	dbUser.Table_name = "user"
	nid, _ := dbUser.Read(update.CallbackQuery.From.ID, &dbcorzineold)

	if dbcorzineold.Corzine.Chtokypil != nil { //Проверка есть ли в корзине что то //todo это так не работает возможно
		var check_tovar bool
		check_tovar = false
		for i := range dbcorzineold.Corzine.Chtokypil { //проверка повторяется ли товар
			if dbcorzineold.Corzine.Chtokypil[i] == dbcorzinenew.Article {
				dbcorzineold.Corzine.Numof[i] += 1
				check_tovar = true
			}
		}
		if !check_tovar {
			dbcorzineold.Corzine.Chtokypil = add_massiv(dbcorzineold.Corzine.Chtokypil, dbcorzinenew.Article)
			dbcorzineold.Corzine.Numof = add_massiv(dbcorzineold.Corzine.Numof, 1)

		}

	} else { //не повторяется
		dbcorzineold.Corzine.Chtokypil = create_massiv(1)
		dbcorzineold.Corzine.Chtokypil[0] = dbcorzinenew.Article
		dbcorzineold.Corzine.Numof = create_massiv(1)
		dbcorzineold.Corzine.Numof[0] = 1
	}

	jsonkeyadd, _ := json.Marshal(dbcorzineold.Corzine)
	dbUser.Table_name = "user"
	dbUser.Edit(nid, "Corz", jsonkeyadd)

	var product_add database.T_productDatabase
	//var dbaseproduct database.T_SettingDbProduct
	dbaseproduct, err := database.NewBD().Open()
	dbaseproduct.Table_name = dbcorzinenew.Category
	dbaseproduct.Read(dbcorzinenew.Article, &product_add) //err
	var answercall tgbotapi.CallbackConfig
	answercall.CallbackQueryID = update.CallbackQuery.ID
	answercall.Text = fmt.Sprintf("Добавлено '%s'", product_add.Name)
	bot.AnswerCallbackQuery(answercall)
	_ = err
}
func moredetailed(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	tg_id := update.CallbackQuery.From.ID
	var dbcorzinenew database.T_productDatabase
	var new string
	var slova []string

	slova = append(slova, "article:")
	word, key := search_word(update.CallbackQuery.Data, slova)
	for i, k := range key {
		switch word[k] {
		case "article:":
			new = card_getIntJson(key, i, k, update.CallbackQuery.Data, word[k])
			json.Unmarshal([]byte(new), &dbcorzinenew)
		}
	}
	//Считываем все таблицы
	//var  T_SettingDbProduct
	dbproduct, err := database.NewBD().Open()

	if err != nil {
		//todo
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//todo
	}

	if len(table) == 0 {
		msg := tgbotapi.NewMessage(int64(tg_id), "Каталог пуст. Нет таблиц.")
		bot.Send(msg)
		return
	}
	//dbcorzinenew.Category
	for i := range table {
		dbproduct.Table_name = table[i]
		dbarticle, err := dbproduct.ReadAll()
		_ = err //to do
		for k := range dbarticle {
			if dbarticle[k].Article == dbcorzinenew.Article {
				if dbarticle[k].Photo != "" {
					//outmsg += fmt.Sprintf("%d товар.\nАртикул товара:%d\nНазвание:%s\nОписание:%s\n Количество:%d\n", ischet, dbarticle[k].Article, dbarticle[k].Name, dbarticle[k].Description, baza.Corzine.Numof[iuser])
					text := fmt.Sprintf("Артикул: %d\nНазвание: %s\n%s\nЦена: %0.2fрублей\n", dbarticle[k].Article, dbarticle[k].Name, dbarticle[k].Description, dbarticle[k].Price)
					ms1 := tgbotapi.NewEditMessageCaption(int64(tg_id), update.CallbackQuery.Message.MessageID, text)
					//	ms1 := tgbotapi.NewEditMessageText(int64(tg_id), update.CallbackQuery.Message.MessageID, text)
					sss := fmt.Sprintf("/addCorzine\narticle:%d\ncategory:%s", dbarticle[i].Article, table[i]) //todo
					var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
						),
					)
					ms1.ReplyMarkup = &numericKeyboardInline
					bot.Send(ms1)
				} else {
					text := fmt.Sprintf("Артикул: %d\nНазвание: %s\n%s\nЦена: %0.2fрублей\n", dbarticle[k].Article, dbarticle[k].Name, dbarticle[k].Description, dbarticle[k].Price)
					ms1 := tgbotapi.NewEditMessageText(int64(tg_id), update.CallbackQuery.Message.MessageID, text)
					//	ms1 := tgbotapi.NewEditMessageText(int64(tg_id), update.CallbackQuery.Message.MessageID, text)
					sss := fmt.Sprintf("/addCorzine\narticle:%d\ncategory:%s", dbarticle[i].Article, table[i]) //todo
					var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
						tgbotapi.NewInlineKeyboardRow(
							tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
						),
					)
					ms1.ReplyMarkup = &numericKeyboardInline
					bot.Send(ms1)
				}
			}

		}
	}
	var answercall tgbotapi.CallbackConfig
	answercall.CallbackQueryID = update.CallbackQuery.ID
	answercall.Text = fmt.Sprintf("Выполнено")
	bot.AnswerCallbackQuery(answercall)
}

// /////////////////////////////////////////////////////////////////////////
func checkUser(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (int, int, error) {
	//	var baza T_Database
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "") //update.Message.Text
	data_base, err := database.NewBDUser().Open()

	if err != nil {
		return 0, 0, err
	}

	data_base.Table_name = "user"
	status, state, err := data_base.Bd_checkUser(update.Message.From.ID)
	_ = state
	a := errNOfoundusr
	_ = a
	if err != nil {
		if err.Error() == errNOfoundusr.Error() {
			err = nil
		} else {
			return 0, 0, err
		}
	}

	switch status {
	case Nouser:
		var baza database.T_Database
		baza.Id_user = update.Message.From.ID
		baza.Status = User
		baza.State = E_STATE_NOTHING
		data_base.Adduser(baza)
		msg.Text = "Хеллоу нью юзер, ай эм бот, бай vladixxxa"
		bot.Send(msg)
		return User, 0, err
	case Admin:
		return Admin, state, err
	case User:
		return User, state, err
	}

	return 0, 0, err
}

func checkCommand(update *tgbotapi.Update, status int, state int, bot *tgbotapi.BotAPI) (bool, string) {
	//ПРОВЕРКА НА СОСТОЯНИЕ
	if status == Admin {
		if state < E_STATE_ADMIN_MAX && state > 0 { //todo
			msg_status, msg := aparsState(state, update, bot)
			return msg_status, msg
		}
	}
	if state == E_STATE_GETUSER_GEOLOCATION {
		if update.Message.Location != nil {
			sendsdeklocation(update, bot) //todo get err
			return false, ""
		} else {
			Iscommand := update.Message.IsCommand()
			if !Iscommand {
				sendyalocation(update, bot)
				return false, ""
			}
		}

	}
	//подделываем месседж для упрощения (при добавлении в бд через команды) //todo мб убрать
	if update.Message.Photo != nil {
		//подделываем для упрощения
		update.Message.Text = update.Message.Caption
		word := strings.Split(update.Message.Text, "\n")
		lencomandd := len(word[0])
		a := make([]tgbotapi.MessageEntity, 1)
		a[0].Type = "bot_command"
		a[0].Length = lencomandd
		update.Message.Entities = &a

	}

	Iscommand := update.Message.IsCommand()
	if Iscommand {
		msg_status, msg := parsCommand(update, status, bot)
		return msg_status, msg
	} else {
		msg_status, msg := parsTextCommand(update, status, bot)
		return msg_status, msg
	}

}
