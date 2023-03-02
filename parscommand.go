package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	database2 "sample-app/tg2/database"
	"sample-app/tg2/sdek"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// ///////
func parsCommand(update *tgbotapi.Update, status int, bot *tgbotapi.BotAPI) (bool, string) {
	switch status {
	case User:
		msg_status, msg := commandUser(update, bot)
		return msg_status, msg
	case Admin:
		msg_status, msg := commandAdmin(update, bot)
		return msg_status, msg
	}
	return true, "ERR STATUS 1, пишите админу"
}
func commandUser(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {
	command := update.Message.Command()
	switch command {
	case "button":
		getState(update)
		buttonUser(update, bot)

		return false, ""
	case "start":
		//Кнопки юзера
		return false, ""
	case "deleteshoppcart":
		udeleteshoppcart(update, bot)
		return false, ""
	default:
		msg := "Такой команды нет"
		return true, msg
		//msg.Text = "I don't know that command"
	}

}

func commandAdmin(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {
	command := update.Message.Command()
	switch command {
	case "button":
		getState(update)
		buttonAdmin(update, bot)
		return false, ""
	case "start":
		//Кнопки админа
		return false, ""
	case "doingadmin":
		doingadmin(update, bot) //todo
		return false, ""
	case "deleteadmin":
		deleteadmin(update, bot) //todo
		return false, ""
	case "addcardv1":
		addcardv1(update, bot)
		//надо узнать фото или нет, если фото смотреть в капшионе
		return false, ""
	case "addcard":
		addcard(update, bot) //todo
		//надо узнать фото или нет, если фото смотреть в капшионе
		return false, ""
	case "deletecard":
		deletecard(update, bot)
		return false, ""
		// case "login":
	// 	var i = 1
	// 	_ = i
	case "addcardphoto":
		//
		return false, ""
	case "status":
		//
		return false, ""
	case "deleteshoppcart":
		udeleteshoppcart(update, bot)
		return false, ""
		////////////////////////////
	// case "addcard_article":
	// 	Add_card.Article = update.Message.M
	// 	return false, ""
	default:
		msg := "Такой команды нет"
		return true, msg
	}
}

/*
	func buttonLocation(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
		var numericKeyboard = tgbotapi.NewKeyboardButtonLocation("Геолокация")
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Геолокация отправлена")
		msg.ReplyMarkup = numericKeyboard
		bot.Send(msg)

}
*/

func getState(update *tgbotapi.Update) {
	dbUser, err := database2.NewBDUser().Open()
	tg_id := update.Message.From.ID //id tg user
	dbUser.Table_name = "user"
	dbUser.EditTGid(tg_id, "state", E_STATE_NOTHING) //todo
	_ = err                                          // todo
}

func buttonUser(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог"),
			tgbotapi.NewKeyboardButton("Мои заказы"),
			tgbotapi.NewKeyboardButton("Корзина"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("FAQ"),
			tgbotapi.NewKeyboardButton("HELP"),
			//tgbotapi.NewKeyboardButton("6"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню добавлено")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)

}
func buttonAdmin(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Каталог"),
			tgbotapi.NewKeyboardButton("Мои заказы"),
			tgbotapi.NewKeyboardButton("Корзина"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("FAQ"),
			tgbotapi.NewKeyboardButton("HELP"),
			//tgbotapi.NewKeyboardButton("6"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Настройки админа"),
			//tgbotapi.NewKeyboardButton("Мои заказы"),
			//tgbotapi.NewKeyboardButton("Корзина"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)

}

func asetting(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Как пользоваться?"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Info:/addcardv1"),
			tgbotapi.NewKeyboardButton("Info:/addcard"),
			tgbotapi.NewKeyboardButton("Info:/deletecard"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Info:/doingadmin"),
			tgbotapi.NewKeyboardButton("Info:/deleteadmin"),
			//tgbotapi.NewKeyboardButton("6"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/button"),
			//tgbotapi.NewKeyboardButton("Мои заказы"),
			//tgbotapi.NewKeyboardButton("Корзина"),
		),
	)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню админа")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}
func doingadmin(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	//тоже наверное делать через кэллбэк лучше
	//update.Message.Text
	// /doingadmin
	// id_tg:
	var slovo []string
	slovo = append(slovo, "id_tg:")
	var new string
	var db database2.T_Database
	word, key := search_word(update.Message.Text, slovo)
	for i, k := range key {
		switch word[k] {
		case "id_tg:":
			new = card_getIntJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		}

	}
	//получили айдишник
	// Добавить условие, если пользователь есть в БД to
	database, _ := database2.NewBDUser().Open()
	database.Table_name = "user"
	id, err := database.Read(db.Id_user, &db)
	_ = err //to do
	db.Status = Admin

	database.Edit(id, "status", db.Status)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Администратор %d добавлен", db.Id_user))
	bot.Send(msg)
	msg = tgbotapi.NewMessage(int64(db.Id_user), "Поздравляем, теперь вы являетсь модератором! \nЗайдите в меню и выберите комманду:\n/button")
	bot.Send(msg)
}
func deleteadmin(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var slovo []string
	slovo = append(slovo, "id_tg:")
	var new string
	var db database2.T_Database
	word, key := search_word(update.Message.Text, slovo)
	for i, k := range key {
		switch word[k] {
		case "id_tg:":
			new = card_getIntJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		}

	}
	//получили айдишник

	database, _ := database2.NewBDUser().Open()
	database.Table_name = "user"
	id, err := database.Read(db.Id_user, &db)
	_ = err //to do
	db.Status = User
	database.Edit(id, "status", db.Status)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Администратор %d удален", db.Id_user))
	bot.Send(msg)
	msg = tgbotapi.NewMessage(int64(db.Id_user), "Поздравляем, теперь вы не модератор!")
	bot.Send(msg)
}
func addcard(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить артикул", "/addcard_article"),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Добавить название", "/addcard_name"),
		// ),

		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Добавить описание", "/addcard_description"),
		// ),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Добавить цену", "/addcard_price"),
		// ),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Добавить фото", "/addcard_photo"),
		// ),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Добавить в БД", "/addcard_addDB"),
		// ),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выйти из режима добавления", "/addcard_cancel"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добавить карточку товара: ")
	msg.ReplyMarkup = numericKeyboardInline
	aa, err := bot.Send(msg)

	dbaseproduct, err := database2.NewBD().Open()

	//table_name := "temp" //записываем во временную БД
	dbaseproduct.Table_name = "temp"
	//toDO to do //Если вызываешь второй раз то обновляется
	dbaseproduct.Addany("id_tg", update.Message.From.ID)
	_ = aa
	_ = err
	_ = numericKeyboardInline
}
func addcardv1(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	var db database2.T_productDatabase

	if update.Message.Photo != nil {
		a := (update.Message.Photo)
		b := len(*a)
		c := (*a)[b-1]
		var file tgbotapi.FileConfig
		file.FileID = c.FileID
		outfile, _ := bot.GetFile(file)

		filename, err := savefile(outfile)
		db.Photo = filename
		if err != nil {
			////todo
		}
	}

	var new string
	article := "article:"
	description := "description:"
	price := "price:"
	name := "name:"
	category := "category:"
	var slova []string
	slova = append(slova, price)
	slova = append(slova, article)
	slova = append(slova, name)
	slova = append(slova, description)
	slova = append(slova, category)
	word, key := search_word(update.Message.Text, slova)
	for i, k := range key {
		switch word[k] {
		case article:
			new = card_getIntJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		case price:
			new = card_getIntJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		case description:
			new = card_getStringJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		case name:
			new = card_getStringJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		case category:
			new = card_getStringJson(key, i, k, update.Message.Text, word[k])
			json.Unmarshal([]byte(new), &db)
		}
	}
	fmt.Println(db)
	dbaseproduct, err := database2.NewBD().Open()

	if err != nil {
		// to do
	}
	err = dbaseproduct.Check_table(db.Category)
	if err != nil {
		//createp(db.Category)
		database2.Create(db.Category) //todo
		err = nil
	}
	dbaseproduct.Table_name = db.Category
	dbaseproduct.Add(db)
	_ = err //todo
}
func deletecard(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //to do category
	var db database2.T_productDatabase

	var new string
	word := strings.Split(update.Message.Text, "\n")
	for i := 1; i < len(word); i++ {
		w := strings.Split(word[i], ":")
		switch w[0] {
		//todo new algoritgm
		case "article":
			new = "{" + string(0x22) + string(w[0]) + string(0x22) + ":" + string(w[1]) + "}" //только когда число
			json.Unmarshal([]byte(new), &db)
		}

	}
	// Было так start
	// var dbaseproduct T_SettingDbProduct
	// err := dbaseproduct.openp()
	// dbaseproduct.table_name = "product"
	// dbaseproduct.deletep(db.Article)
	// //проверка нет айдишника
	// _ = err //todo
	//end
	chat_id := update.Message.From.ID
	dbproduct, err := database2.NewBD().Open()

	if err != nil {
		//
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//
	}

	if len(table) == 0 {
		msg := tgbotapi.NewMessage(int64(chat_id), "Каталог пуст. Нет таблиц.")
		bot.Send(msg)
		return
	}
	for i := range table {
		dbproduct.Table_name = table[i]
		dbproduct.Delete(db.Article)
	}
}

// asetting

func udeleteshoppcart(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	dbUser, err := database2.NewBDUser().Open()
	tg_id := update.CallbackQuery.From.ID
	var dbcorzineold database2.T_Database
	//Передалть на изменение сразу по ID
	dbUser.Table_name = "user"
	nid, _ := dbUser.Read(tg_id, &dbcorzineold)
	dbUser.Edit(nid, "Corz", "")
	_ = err //todo
	//answer

	//answer
	//Change message
	//
	//Основное сообщ с описанием в deleteshoppcart.Msg

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Корзина пуста")
	bot.Send(msg)
	//Ответ toDo
	var answercall tgbotapi.CallbackConfig
	answercall.CallbackQueryID = update.CallbackQuery.ID
	answercall.Text = fmt.Sprintf("Корзина удалена")
	bot.AnswerCallbackQuery(answercall)
}

type T_Deleteshoppcarter struct {
	Article int `json:"article"`
	Msg     int `json:"msg"`
}

func deleteshopposition(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	//msg type  ("/deleteshopposition\narticle:%d\nmsg:%d", article, msgid)
	//parsmsg
	msg := update.CallbackQuery.Data
	//
	dbUser, err := database2.NewBDUser().Open()
	tg_id := update.CallbackQuery.From.ID //id tg user

	var dbcorzineold database2.T_Database
	//Передалть на изменение сразу по ID
	dbUser.Table_name = "user"
	nid, err := dbUser.Read(tg_id, &dbcorzineold)
	//
	if len(dbcorzineold.Corzine.Chtokypil) == 0 {
		var answercall tgbotapi.CallbackConfig
		answercall.CallbackQueryID = update.CallbackQuery.ID
		answercall.Text = fmt.Sprintf("Корзина пуста")
		bot.AnswerCallbackQuery(answercall)
		return
	}
	var deleteshoppcart T_Deleteshoppcarter
	var new string
	article := "article:"
	msgidchange := "msg:"
	var slova []string
	slova = append(slova, article)
	slova = append(slova, msgidchange)
	word, key := search_word(msg, slova)
	for i, k := range key {
		switch word[k] {
		case article:
			new = card_getIntJson(key, i, k, msg, word[k])
			json.Unmarshal([]byte(new), &deleteshoppcart)
		case msgidchange:
			new = card_getIntJson(key, i, k, msg, word[k])
			json.Unmarshal([]byte(new), &deleteshoppcart)
		}
	}

	//

	//удаляем из артикля из смещаем и удаляем эту же позицию в что купил
	//dbcorzineold.Corzine.Chtokypil
	articleincorzin := dbcorzineold.Corzine.Chtokypil
	iddelete, err1 := find_massiv(articleincorzin, deleteshoppcart.Article)
	if err1 == 1 {
		//
	}
	dbcorzineold.Corzine.Numof[iddelete-1] -= 1
	if dbcorzineold.Corzine.Numof[iddelete-1] == 0 {
		dbcorzineold.Corzine.Chtokypil = delete_massiv(dbcorzineold.Corzine.Chtokypil, iddelete)
		dbcorzineold.Corzine.Numof = delete_massiv(dbcorzineold.Corzine.Numof, iddelete)
		if len(dbcorzineold.Corzine.Chtokypil) == 0 {
			jsonkeyadd, _ := json.Marshal(dbcorzineold.Corzine)
			dbUser.Edit(nid, "Corz", jsonkeyadd)
			ms1 := tgbotapi.NewEditMessageText(int64(tg_id), deleteshoppcart.Msg, "Корзина пуста")
			bot.Send(ms1)
			//update.CallbackQuery.Message.MessageID
			//Второе редачить
			ms2 := tgbotapi.NewDeleteMessage(int64(tg_id), update.CallbackQuery.Message.MessageID)

			bot.Send(ms2)
			//answer сверху
			var answercall tgbotapi.CallbackConfig
			answercall.CallbackQueryID = update.CallbackQuery.ID
			answercall.Text = fmt.Sprintf("Корзина изменена")
			bot.AnswerCallbackQuery(answercall)
			_ = err
		}
	}
	jsonkeyadd, _ := json.Marshal(dbcorzineold.Corzine)
	dbUser.Edit(nid, "Corz", jsonkeyadd)

	inlinekb, outmsg, err2 := answer(tg_id, deleteshoppcart.Msg)
	if err2 {
		//todo
	}
	//answer
	//Change message
	//
	//Основное сообщ с описанием в deleteshoppcart.Msg

	ms1 := tgbotapi.NewEditMessageText(int64(tg_id), deleteshoppcart.Msg, outmsg)
	ms1.ReplyMarkup = &inlinekb[0]
	bot.Send(ms1)
	//update.CallbackQuery.Message.MessageID
	//Второе редачить
	ms2 := tgbotapi.NewEditMessageText(int64(tg_id), update.CallbackQuery.Message.MessageID, "Удалить из корзины позиции:")
	ms2.ReplyMarkup = &inlinekb[1]
	bot.Send(ms2)
	//answer сверху
	var answercall tgbotapi.CallbackConfig
	answercall.CallbackQueryID = update.CallbackQuery.ID
	answercall.Text = fmt.Sprintf("Корзина изменена")
	bot.AnswerCallbackQuery(answercall)
	_ = err

}

func answer(tgid int, msgidold int) ([2]tgbotapi.InlineKeyboardMarkup, string, bool) {
	//chat_id := update.Message.From.ID
	var rnumKeyInline [2]tgbotapi.InlineKeyboardMarkup
	var baza database2.T_Database
	database_corzz, err := database2.NewBDUser().Open() //todo
	database_corzz.Table_name = "user"
	database_corzz.Read(tgid, &baza)
	var outmsg string
	if err != nil {

	}
	if len(baza.Corzine.Chtokypil) == 0 {
		return rnumKeyInline, "Корзина пуста", true
	}

	dbproduct, err := database2.NewBD().Open()

	if err != nil {
		//todo
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//
	}

	if len(table) == 0 {
		return rnumKeyInline, "Каталог пуст. Нет таблиц.", true
	}
	var ischet = 0
	outmsg = ""
	var add_corzine_name []string
	corzinenameToAtricle := make(map[string]int)
	for i := range table {
		dbproduct.Table_name = table[i]
		dbarticle, err := dbproduct.ReadAll()
		_ = err //to do
		for k := range dbarticle {
			for iuser := range baza.Corzine.Chtokypil {
				if dbarticle[k].Article == baza.Corzine.Chtokypil[iuser] {
					ischet++
					outmsg += fmt.Sprintf("%d товар.\nАртикул товара:%d\nНазвание:%s\n Количество:%d\n", ischet, dbarticle[k].Article, dbarticle[k].Name, baza.Corzine.Numof[iuser])
					add_corzine_name = append(add_corzine_name, fmt.Sprintf("%s", dbarticle[k].Name))
					corzinenameToAtricle[dbarticle[k].Name] = dbarticle[k].Article
				}
			}
		}
	}
	if ischet == 0 {
		return rnumKeyInline, "Корзина пуста. (Корзина не сходится с артикулем из бд)", true
	}
	size_corz := int(len(add_corzine_name))
	var numKeyInline tgbotapi.InlineKeyboardMarkup
	//очистка всей корзины
	numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 1) //size_corz+2
	for i := range numKeyInline.InlineKeyboard {
		var data string
		switch i {
		case 0:
			numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1) //В поле создаем еще поле
			numKeyInline.InlineKeyboard[i][0].Text = "Очистить корзину"
			data = "/deleteshoppcart" //надо передавать команду +id что удалить(?) //todo
			numKeyInline.InlineKeyboard[i][0].CallbackData = &data
		case 1:
			numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1) //В поле создаем еще поле
			numKeyInline.InlineKeyboard[i][0].Text = "Оформить заказ"
			data = "/placeanorder" //надо передавать команду +id что удалить(?) //todo
			numKeyInline.InlineKeyboard[i][0].CallbackData = &data
		}
	}
	//очистка определенных позиций в корзине
	var numKeyInline2 tgbotapi.InlineKeyboardMarkup
	numKeyInline2.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, size_corz)

	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, outmsg)
	//msg.ReplyMarkup = numKeyInline

	//msgsend, _ := bot.Send(msg)
	//msgid := msgsend.MessageID
	for i := range numKeyInline2.InlineKeyboard {
		var data string
		numKeyInline2.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
		numKeyInline2.InlineKeyboard[i][0].Text = add_corzine_name[i] //
		article, ok := corzinenameToAtricle[add_corzine_name[i]]
		if ok {
			data = fmt.Sprintf("/deleteshopposition\narticle:%d\nmsg:%d", article, msgidold)
			numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
		} else {
			// data = "/shoppcartttttttttttt \n id:1" //надо передавать команду +id что удалить(?) //todo
			data = ""
			numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
		}
	}
	//msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, "Удалить из корзины позиции:")
	//msg2.ReplyMarkup = numKeyInline2
	//bot.Send(msg2)

	rnumKeyInline[0] = numKeyInline
	rnumKeyInline[1] = numKeyInline2
	_ = err
	return rnumKeyInline, outmsg, false
	//обращаться к БД каталога

}

func takepvz(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	tg_id := update.CallbackQuery.From.ID //id tg user
	//update.CallbackQuery.Data
	//inlinemsg := fmt.Sprintf("/takepvz\nnumber_pvz:%d", i)
	var new string
	number_pvz := "number_pvz:"
	var pvzRead database2.T_UserPvz
	var slova []string
	slova = append(slova, number_pvz)
	word, key := search_word(update.CallbackQuery.Data, slova)
	for i, k := range key {
		switch word[k] {
		case number_pvz:
			new = card_getIntJson(key, i, k, update.CallbackQuery.Data, word[k])
			json.Unmarshal([]byte(new), &pvzRead)
		}
	}
	fmt.Println(pvzRead)
	var userpvz database2.T_Database
	database, err := database2.NewBDUser().Open()
	if err != nil {
		return //todo
	}
	database.Table_name = "user"
	_, err = database.Read(tg_id, &userpvz)
	if err != nil {
		return //todo
	}
	if len(userpvz.UserPvz) == 0 {
		return //err
	}
	for i := range userpvz.UserPvz {
		if userpvz.UserPvz[i].Number_pvz == pvzRead.Number_pvz {
			//обнулить бд (либо добавить итоговую)
			//отправляем сколько будет стоить
			//txt := fmt.Sprintln(userpvz.UserPvz[i])
			tarif := calcTarif(userpvz.UserPvz[i].Sdek_pvz)
			txt := fmt.Sprintf("Доставка до ПВЗ	🚜🏫🚂⏲: %s\n", userpvz.UserPvz[i].Sdek_pvz.Locat.Address_full)
			txt += fmt.Sprintf("🙏🙏🙏🙏🙏🙏🙏🙏🙏\n")
			txt += fmt.Sprintf("Код тарифа: %d\n", 136)
			txt += fmt.Sprintf("Стоимость доставки💩: %0.2f\n", tarif.Delivery_sum)
			txt += fmt.Sprintf("Минимальное время доставки (в календарных днях)⏰: %d\n", tarif.Calendar_min)
			txt += fmt.Sprintf("Максимальное время доставки (в календарных днях)😰: %d\n", tarif.Calendar_max)
			txt += fmt.Sprintf("Расчетный вес (в граммах)💀: %d\n", tarif.Weight_calc)
			txt += fmt.Sprintf("Стоимость доставки с учетом дополнительных услуг 🚁✈🚀: %0.2f\n", tarif.Total_sum)

			msg := tgbotapi.NewMessage(int64(tg_id), txt)
			bot.Send(msg)
			var answercall tgbotapi.CallbackConfig
			answercall.CallbackQueryID = update.CallbackQuery.ID
			answercall.Text = fmt.Sprintf("Расчет представлен")
			bot.AnswerCallbackQuery(answercall)
			return
		}
	}
}
func calcTarif(to sdek.SdekAnswOffice) sdek.SdekCodeTariff {
	var msg sdek.SdekCodeTariffMsg
	msg.Type = 1
	msg.Currency = 1
	msg.Tariff_code = 136
	msg.From_location.City = "Москва"
	msg.From_location.Address = "Россия, Москва, Бутово, б-р Адмирала Ушакова, 18Б"

	msg.From_location.Code = 467
	//
	msg.To_location.City = to.Locat.City
	msg.To_location.Address = to.Locat.Address_full

	msg.To_location.Code = to.Locat.City_code

	msg.Packages = make([]sdek.SdekInfPacakege, 1)
	msg.Packages[0].Height = 10
	msg.Packages[0].Length = 10
	msg.Packages[0].Weight = 500
	msg.Packages[0].Width = 10
	client := sdek.NewAuth().SetAuth()
	ans, err := client.SdekCalcCodeTariff(msg)
	if err != nil {
		//todo
	}
	return ans
}
func placeanorder(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	dbUser, err := database2.NewBDUser().Open()
	tg_id := update.CallbackQuery.From.ID //id tg user

	var dbcorzineold database2.T_Database
	//Передалть на изменение сразу по ID
	dbUser.Table_name = "user"
	nid, err := dbUser.Read(tg_id, &dbcorzineold)          //Корзина считали
	dbUser.Edit(nid, "state", E_STATE_GETUSER_GEOLOCATION) //todo
	//"Два варианта работы.\n1. Прислать геопозицию(нажмите на кнопку моя геопозиция) или примерной точки откуда хотите забрать, бот выберет ближайшие пвз.\n2.Пришлите ближай адрес доставки или адрес сдэк и мы подберем пвз"
	txtt := "Два варианта работы.\n1. Прислать геопозицию(нажмите на кнопку моя геопозиция) или примерной точки откуда хотите забрать(скрепочка->геопозиция), бот выберет ближайшие пвз.\n2.Пришлите ближай адрес доставки или адрес сдэк и мы подберем пвз. \nАдрес можно вбить таким способом \"@toDo58292_bot [адрес]\"  \n\nДля выхода команда /button"
	msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), txtt)

	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation("Моя геопозиция"),
		),
	)
	//var numericKeyboard = tgbotapi.NewKeyboardButtonLocation("Геолокация")
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
	//todo добавляем статус геолокациии в бд
	a := 5455
	_ = a
	_ = nid
	_ = err

	//Выберите способ доставки
	////тут где то рассчитыавется цена доставки
	//Введите
}
func ushoppcart(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //корзина
	//id_user := int(update.Message.Chat.ID) //получаем айдишки его

	// var baza T_Database
	// database_corzz, err := open("sqlite3", "testdatabase/data_test1.db") //todo
	// tg_id := update.Message.From.ID
	// read(database_corzz, tg_id, &baza, "user")
	// var outmsg = ""
	// var msg tgbotapi.MessageConfig
	// var msg2 tgbotapi.MessageConfig
	// if baza.Corzine.Chtokypil != nil {
	// 	//_ = id
	// 	/*if id == 0 {
	// 		corz := create_massiv(1)
	// 		corz[0] = 0
	// 		baza.Corzine.Chtokypil = corz

	// 	}*/
	// 	//len(baza.Corzine.Chtokypil)
	// 	var dbaseproduct T_SettingDbProduct
	// 	err := dbaseproduct.openp()
	// 	_ = err //todo
	// 	dbaseproduct.table_name = "product"

	// 	var dbproduct T_productDatabase
	// 	var ischet = 0
	// 	//var add_corzine []int
	// 	var add_corzine_name []string
	// 	for i := range baza.Corzine.Chtokypil {
	// 		ischet++
	// 		fmt.Println(i)
	// 		dbaseproduct.readp(baza.Corzine.Chtokypil[i], &dbproduct)
	// 		outmsg += fmt.Sprintf("%d товар.\nid товара:%d\nНазвание:%s\nОписание:%s\n Количество:%d\n", ischet, dbproduct.Article, dbproduct.Name, dbproduct.Description, baza.Corzine.Numof[i])
	// 		add_corzine_name = append(add_corzine_name, fmt.Sprintf("%s", dbproduct.Name))

	// 		//формируем сообщение и отправляем
	// 		// первое сообщ ваша корзина:
	// 	}
	// 	var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("Очистить корзину", "/deleteshoppcart"),
	// 		),
	// 	)
	// 	var numericKeyboardInline2 = tgbotapi.NewInlineKeyboardMarkup(
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("Очистить корзину", "/deleteshoppcart"),
	// 		),
	// 		tgbotapi.NewInlineKeyboardRow(
	// 			tgbotapi.NewInlineKeyboardButtonData("sss", "ss"),
	// 			tgbotapi.NewInlineKeyboardButtonData("ff", "d"),
	// 		),
	// 	)
	// 	size_corz := int(len(add_corzine_name))
	// 	var numKeyInline tgbotapi.InlineKeyboardMarkup
	// 	//очистка всей корзины
	// 	numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 1) //size_corz+2
	// 	for i := range numKeyInline.InlineKeyboard {
	// 		var data string
	// 		switch i {
	// 		case 0:
	// 			numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1) //В поле создаем еще поле
	// 			numKeyInline.InlineKeyboard[i][0].Text = "Очистить корзину"
	// 			data = "/deleteshoppcart" //надо передавать команду +id что удалить(?) //todo
	// 			numKeyInline.InlineKeyboard[i][0].CallbackData = &data

	// 		}
	// 	}
	// 	//очистка определенных позиций в корзине
	// 	var numKeyInline2 tgbotapi.InlineKeyboardMarkup
	// 	numKeyInline2.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, size_corz)
	// 	for i := range numKeyInline2.InlineKeyboard {
	// 		var data string
	// 		numKeyInline2.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
	// 		numKeyInline2.InlineKeyboard[i][0].Text = add_corzine_name[i] //
	// 		data = "/shoppcartttttttttttt"                                //надо передавать команду +id что удалить(?) //todo
	// 		numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
	// 	}

	//Попытка 2 Работает start
	/*var numKeyInline tgbotapi.InlineKeyboardMarkup
	numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 2)
	numKeyInline.InlineKeyboard[0] = make([]tgbotapi.InlineKeyboardButton, 1)
	numKeyInline.InlineKeyboard[1] = make([]tgbotapi.InlineKeyboardButton, 2)
	//[0][0] - 0 строка, 0 столбец
	//[0][1] - 0 строка 1 столбец
	numKeyInline.InlineKeyboard[0][0].Text = "aaa"
	aaaaa := "aaaaaaaaaaa"
	numKeyInline.InlineKeyboard[0][0].CallbackData = &aaaaa
	*/
	//var key tgbotapi.InlineKeyboardButton
	//s := make(, size_corz)

	////sizenumericKeyboardInline := len(numericKeyboardInline)
	//sizenumericKeyboardInline2 := len(numericKeyboardInline2)
	// _ = sizenumericKeyboardInline
	// _ = sizenumericKeyboardInline2
	//end
	// 	_ = numericKeyboardInline2 //todo
	// 	_ = numericKeyboardInline  //todo
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, outmsg)
	// 	msg.ReplyMarkup = numKeyInline
	// 	msg2 = tgbotapi.NewMessage(update.Message.Chat.ID, "Удалить из корзины позиции:")
	// 	msg2.ReplyMarkup = numKeyInline2
	// 	//msg.ReplyMarkup = numKeyInline

	// } else {
	// 	outmsg = "Корзина пуста"
	// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, outmsg)
	// }

	// bot.Send(msg)
	// bot.Send(msg2)
	// //обращаться к БД каталога
	// _ = err

	tgid := update.Message.From.ID
	var baza database2.T_Database
	database_corzz, err := database2.NewBDUser().Open() //todo
	database_corzz.Table_name = "user"
	_, _ = database_corzz.Read(tgid, &baza)
	//if len(baza.Corzine.Chtokypil) == 0
	var outmsg string
	if err != nil {
		//
	}
	if len(baza.Corzine.Chtokypil) == 0 {
		msg := tgbotapi.NewMessage(int64(tgid), "Корзина пуста")
		bot.Send(msg)
		return
	}

	dbproduct, err := database2.NewBD().Open()
	if err != nil {
		//
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//
	}

	if len(table) == 0 {
		msg := tgbotapi.NewMessage(int64(tgid), "Каталог пуст. Нет таблиц.")
		bot.Send(msg)
		return
	}
	var ischet = 0
	outmsg = ""
	var add_corzine_name []string
	corzinenameToAtricle := make(map[string]int)
	//Сделать связанный список и переделать //todo
	for i := range table {
		dbproduct.Table_name = table[i]
		dbarticle, err := dbproduct.ReadAll()
		_ = err //to do
		for k := range dbarticle {
			for iuser := range baza.Corzine.Chtokypil {
				if dbarticle[k].Article == baza.Corzine.Chtokypil[iuser] {
					ischet++
					//outmsg += fmt.Sprintf("%d товар.\nАртикул товара:%d\nНазвание:%s\nОписание:%s\n Количество:%d\n", ischet, dbarticle[k].Article, dbarticle[k].Name, dbarticle[k].Description, baza.Corzine.Numof[iuser])
					outmsg += fmt.Sprintf("%d товар.\nАртикул товара:%d\nНазвание:%s\n Количество:%d\n", ischet, dbarticle[k].Article, dbarticle[k].Name, baza.Corzine.Numof[iuser])
					add_corzine_name = append(add_corzine_name, fmt.Sprintf("%s", dbarticle[k].Name))
					corzinenameToAtricle[dbarticle[k].Name] = dbarticle[k].Article
				}
			}
		}
	}
	if ischet == 0 {
		msg := tgbotapi.NewMessage(int64(tgid), "Корзина пуста. (Корзина не сходится с артикулем из бд)")
		bot.Send(msg)
		return
	}
	size_corz := int(len(add_corzine_name))
	var numKeyInline tgbotapi.InlineKeyboardMarkup
	//очистка всей корзины
	numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, 2) //size_corz+2
	for i := range numKeyInline.InlineKeyboard {
		var data string
		switch i {
		case 0:
			numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1) //В поле создаем еще поле
			numKeyInline.InlineKeyboard[i][0].Text = "Очистить корзину"
			data = "/deleteshoppcart" //надо передавать команду +id что удалить(?) //todo
			numKeyInline.InlineKeyboard[i][0].CallbackData = &data
		case 1:
			numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1) //В поле создаем еще поле
			numKeyInline.InlineKeyboard[i][0].Text = "Оформить заказ"
			data = "/placeanorder" //надо передавать команду +id что удалить(?) //todo
			numKeyInline.InlineKeyboard[i][0].CallbackData = &data
		}
	}
	//очистка определенных позиций в корзине
	var numKeyInline2 tgbotapi.InlineKeyboardMarkup
	numKeyInline2.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, size_corz)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, outmsg)
	msg.ReplyMarkup = numKeyInline

	msgsend, _ := bot.Send(msg)
	msgid := msgsend.MessageID
	for i := range numKeyInline2.InlineKeyboard {
		var data string
		numKeyInline2.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
		numKeyInline2.InlineKeyboard[i][0].Text = add_corzine_name[i] //
		article, ok := corzinenameToAtricle[add_corzine_name[i]]
		if ok {
			data = fmt.Sprintf("/deleteshopposition\narticle:%d\nmsg:%d", article, msgid)
			numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
		} else {
			// data = "/shoppcartttttttttttt \n id:1" //надо передавать команду +id что удалить(?) //todo
			data = ""
			numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
		}
	}
	msg2 := tgbotapi.NewMessage(update.Message.Chat.ID, "Удалить из корзины позиции:")
	msg2.ReplyMarkup = numKeyInline2
	bot.Send(msg2)
	//обращаться к БД каталога
	_ = err

}
func ucatalog(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //каталог
	// var numKeyInline2 tgbotapi.InlineKeyboardMarkup
	// numKeyInline2.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, size_corz)
	// for i := range numKeyInline2.InlineKeyboard {
	// 	var data string
	// 	numKeyInline2.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
	// 	numKeyInline2.InlineKeyboard[i][0].Text = add_corzine_name[i] //
	// 	data = "/shoppcartttttttttttt"                                //надо передавать команду +id что удалить(?) //todo
	// 	numKeyInline2.InlineKeyboard[i][0].CallbackData = &data
	// }
	dbproduct, err := database2.NewBD().Open()

	if err != nil {
		//
	}

	table, err := dbproduct.All_table()
	if err != nil {
		//
	}

	if len(table) == 0 {
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Каталог пуст. Нет таблиц.")
		bot.Send(msg)
		return
	}
	var numKeyInline tgbotapi.InlineKeyboardMarkup
	numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, len(table))
	for i := range numKeyInline.InlineKeyboard {
		var data string
		numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
		numKeyInline.InlineKeyboard[i][0].Text = table[i]      //
		data = fmt.Sprintf("/ucatalog\ncategory:%s", table[i]) //надо передавать команду +id что удалить(?) //todo
		numKeyInline.InlineKeyboard[i][0].CallbackData = &data
	}
	msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Выберите каталог")
	msg.ReplyMarkup = numKeyInline
	bot.Send(msg)
	// var dbaseproduct T_SettingDbProduct
	// err := dbaseproduct.openp()
	// _ = err //todo

	// dbaseproduct.table_name = "product"
	// db, _ := dbaseproduct.readAllp()
	// a := len(db)
	// for i := range db {
	// 	text := fmt.Sprintf("Артикул: %d\nНазвание: %s\n%s\nЦена: %0.2fрублей\nВ наличии: %d", db[i].Article, db[i].Name, db[i].Description, db[i].Price, db[i].Instock)
	// 	update.Message.Text = text
	// 	if db[i].Photo != "" {
	// 		photoBytes, err := ioutil.ReadFile(db[i].Photo)
	// 		_ = err
	// 		photoFileBytes := tgbotapi.FileBytes{
	// 			Name:  "",
	// 			Bytes: photoBytes,
	// 		}
	// 		msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, photoFileBytes)
	// 		sss := fmt.Sprintf("/addCorzine\nid:%d", db[i].Article) //todo
	// 		msg.Caption = text
	// 		var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	// 			tgbotapi.NewInlineKeyboardRow(
	// 				tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
	// 			),
	// 		)
	// 		/*//example how do button
	// 		var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	// 			tgbotapi.NewInlineKeyboardRow(
	// 				tgbotapi.NewInlineKeyboardButtonURL("Добавить в корзину", "/addCorzine\n"+string(db[i].Num)),
	// 			),
	// 			tgbotapi.NewInlineKeyboardRow(
	// 				tgbotapi.NewInlineKeyboardButtonData("В наличии:"+string(db[i].Instock), ""),
	// 			),
	// 		)*/

	// 		msg.ReplyMarkup = numericKeyboardInline
	// 		bot.Send(msg)

	// 	} else { //если нет фото
	// 		update.Message.Text = text
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		sss := fmt.Sprintf("/addCorzine\nid:%d", db[i].Article) //todo

	// 		var numericKeyboardInline = tgbotapi.NewInlineKeyboardMarkup(
	// 			tgbotapi.NewInlineKeyboardRow(
	// 				tgbotapi.NewInlineKeyboardButtonData("Добавить в корзину", sss),
	// 			),
	// 		)
	// 		msg.ReplyMarkup = numericKeyboardInline
	// 		bot.Send(msg)
	// 	}
	// 	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 	//bot.Send(msg)
	// 	//формируем сообщение и отправляем
	// 	// первое сообщ ваша корзина:
	// }
	// _ = a  //todo
	// _ = db //todo
	// //обращение к БД каталога
}
func umyorders(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	id_user := int(update.Message.Chat.ID) //получаем айдишки его
	var baza database2.T_Database

	database, err := database2.NewBDUser().Open()
	database.Table_name = "user"
	id, err := database.Read(id_user, &baza)
	txt := fmt.Sprintln(baza.UserPvz)
	_ = id
	msg := tgbotapi.NewMessage(int64(id_user), txt)
	bot.Send(msg)
	// if id == 0 {
	// 	// 	corz := create_massiv(1)
	// 	// 	corz[0] = 0
	// 	// 	baza.Uorders.Uorders = corz
	// }
	// for i := range baza.Uorders.Uorders {
	// 	fmt.Println(i)
	// 	//формируем сообщ аналогично ushoppcart
	// }
	_ = err //todo
}
func ufaq(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //todo
	update.Message.Text = "больше сюда не тыкай"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	bot.Send(msg)
}
func uhelp(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //todo
	update.Message.Text = "И сюда тоже"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	bot.Send(msg)
}

func ahowuse(update *tgbotapi.Update, bot *tgbotapi.BotAPI) { //todo

	//"testdatabase/photo/ac6bd93f-1386-4028-9cf8-298c72084d46.jpg"
	nameFile := "testdatabase/photo/tutorial/Как пользоваться ботом.pdf"
	FileBytes, err := ioutil.ReadFile(nameFile)
	_ = err
	DocFileBytes := tgbotapi.FileBytes{
		Name:  "Как пользоваться ботом.pdf",
		Bytes: FileBytes,
	}
	msg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, DocFileBytes)
	msg.Caption = "Туториал"
	bot.Send(msg)

}

func sendsdeklocation(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {
	//Задаем нашу начальную точку и точку Юзера
	//К точке юзера находим ближайшие пвз
	var coordUser sdek.CoordUser
	//		msgLocation := &update.Message.Location
	//Настройки пвз
	var pvz sdek.SdekOffice
	pvz.Country_code = "643"
	pvz.Type = "ALL"
	//
	//Коорд Юзера
	coordUser.Latit = update.Message.Location.Latitude
	coordUser.Long = update.Message.Location.Longitude
	coordUser.Dcoord = 5000 //5km

	//Получаем пвз
	client := sdek.NewAuth().SetAuth()
	allpvz, err := client.PostOffice(pvz, coordUser)

	//var usrpvz []T_UserPvz

	//Формируем сообщение для выбора пвз
	if err != nil {
		return false, "" //todo
	}
	//Добавляем как клаву (кнопк)
	if len(allpvz) == 0 {
		return false, "" //todo
	}
	usrpvz := make([]database2.T_UserPvz, len(allpvz))
	//var numKeyInline tgbotapi.InlineKeyboardMarkup
	//numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, len(allpvz))
	//это все в обертку
	msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Выберите ПВЗ")
	bot.Send(msg)
	for i := range allpvz {
		usrpvz[i].Sdek_pvz = allpvz[i]
		usrpvz[i].Number_pvz = i
		//Сохранить в кэш по айдишникам
		var userlocation = tgbotapi.NewLocation(int64(update.Message.From.ID), allpvz[i].Locat.Latitude, allpvz[i].Locat.Longitude)
		bot.Send(userlocation)
		txt := fmt.Sprintf("Адрес: %s\nРасстоние от вашей геоточки: %0.2f км", allpvz[i].Locat.Address_full, allpvz[i].Range/1000)
		//txt := fmt.Sprint("Адрес:", allpvz[i].Locat.Address_full, "\nРасстояние от вас: %0.2f км", allpvz[i].Range/1000)
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), txt)
		//idmsg := fmt.Sprintf("/takepvz\ncode:%s\ncitycode:%d\nadress:%s", allpvz[i].Code, allpvz[i].Locat.City_code, allpvz[i].Locat.Address_full)
		//Добавить в бд пвз!
		inlinemsg := fmt.Sprintf("/takepvz\nnumber_pvz:%d", i)
		var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				//tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
				tgbotapi.NewInlineKeyboardButtonData("Выбрать", inlinemsg),
			),
		)
		msg.ReplyMarkup = numericKeyboard
		bot.Send(msg)
	}

	// for i := range numKeyInline.InlineKeyboard {
	// 	var data string
	// 	numKeyInline.InlineKeyboard[i] = make([]tgbotapi.InlineKeyboardButton, 1)
	// 	txt := fmt.Sprint("Адрес:", allpvz[i].Locat.Address_full, "\nРасстояние от вас:", allpvz[i].Range/1000, " км")
	// 	numKeyInline.InlineKeyboard[i][0].Text = txt //
	// 	data = fmt.Sprintf("/alalal:")               //надо передавать команду +id что удалить(?) //todo
	// 	numKeyInline.InlineKeyboard[i][0].CallbackData = &data
	// }
	jsonk, _ := json.Marshal(usrpvz)
	fmt.Printf("%s", jsonk)
	database, err := database2.NewBDUser().Open()
	database.Table_name = "user"
	_ = database.EditTGid(update.Message.From.ID, "pvz", jsonk)
	//
	//возвращаем 0 состояние
	// dbUser, err := open("sqlite3", "testdatabase/data_test1.db")
	// id_tg := update.Message.From.ID
	// editTGid(dbUser, id_tg, "state", "user", E_STATE_NOTHING) //todo

	return false, "" //todo
}

func sendyalocation(update *tgbotapi.Update, bot *tgbotapi.BotAPI) (bool, string) {

	//
	adr := update.Message.Text

	coordUser, err := YaGeocoder(adr)
	if err != nil {
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Адрес не корректный")
		bot.Send(msg)
		return false, "" //todo
	}
	//Задаем нашу начальную точку и точку Юзера
	//К точке юзера находим ближайшие пвз

	//Настройки пвз
	var pvz sdek.SdekOffice
	pvz.Country_code = "643"
	pvz.Type = "ALL"
	//
	//Коорд Юзера, радиус
	coordUser.Dcoord = 5000 //5km

	//Получаем пвз
	client := sdek.NewAuth().SetAuth()
	allpvz, err := client.PostOffice(pvz, *coordUser)

	//Формируем сообщение для выбора пвз
	if err != nil {
		return false, "" //todo
	}
	//Добавляем как клаву (кнопк)
	if len(allpvz) == 0 {
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Адрес не корректный,либо нет офисов сдэк, адрес по мнению яндекса:")
		bot.Send(msg)
		var userlocation = tgbotapi.NewLocation(int64(update.Message.From.ID), coordUser.Latit, coordUser.Long)
		bot.Send(userlocation)
		return false, "" //todo
	}
	usrpvz := make([]database2.T_UserPvz, len(allpvz))

	msgUserAdr := tgbotapi.NewMessage(int64(update.Message.From.ID), "Ваш адрес:")
	bot.Send(msgUserAdr)
	var userlocation = tgbotapi.NewLocation(int64(update.Message.From.ID), coordUser.Latit, coordUser.Long)
	bot.Send(userlocation)
	//var numKeyInline tgbotapi.InlineKeyboardMarkup
	//numKeyInline.InlineKeyboard = make([][]tgbotapi.InlineKeyboardButton, len(allpvz))
	msg := tgbotapi.NewMessage(int64(update.Message.From.ID), "Выберите ПВЗ")
	bot.Send(msg)
	for i := range allpvz {
		usrpvz[i].Sdek_pvz = allpvz[i]
		usrpvz[i].Number_pvz = i
		var userlocation = tgbotapi.NewLocation(int64(update.Message.From.ID), allpvz[i].Locat.Latitude, allpvz[i].Locat.Longitude)
		bot.Send(userlocation)
		txt := fmt.Sprintf("Адрес: %s\nРасстоние от вашей геоточки: %0.2f км", allpvz[i].Locat.Address_full, allpvz[i].Range/1000)
		msg := tgbotapi.NewMessage(int64(update.Message.From.ID), txt)
		inlinemsg := fmt.Sprintf("/takepvz\nnumber_pvz:%d", i)
		var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(

			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Выбрать", inlinemsg),
			),
		)

		msg.ReplyMarkup = numericKeyboard
		bot.Send(msg)
	}

	jsonk, _ := json.Marshal(usrpvz)
	fmt.Printf("%s", jsonk)
	database, err := database2.NewBDUser().Open()
	database.Table_name = "user"
	_ = database.EditTGid(update.Message.From.ID, "pvz", jsonk)
	_ = err
	//
	//возвращаем 0 состояние
	// dbUser, err := open("sqlite3", "testdatabase/data_test1.db")
	// id_tg := update.Message.From.ID
	// editTGid(dbUser, id_tg, "state", "user", E_STATE_NOTHING) //todo
	return false, "" //todo
}
