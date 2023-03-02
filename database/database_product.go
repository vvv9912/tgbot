package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type T_productDatabase struct {
	Article     int    `json:"article"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string
	Price       float64 `json:"price"`
	Category    string  `json:"category"` //имя таблицы!
	Id_tg       int     `json:"id_tg"`    //id кто добавляет!
	Instock     int     //bool?
}

type T_SettingDbProduct struct {
	database   *sql.DB
	Table_name string
}

func NewBD() dbaser {
	return &T_SettingDbProduct{}
}

type dbaser interface {
	Open() (T_SettingDbProduct, error)
}

/*
type dbaseuserer interface {
	//T_SettingDbProduct
	T_SettingDbProduct
	All_table() ([]string, error)
	Check_table(check_Table_name string) error
	Close()
	Add(db T_productDatabase) error
	Add_any(name string, a interface{}) error
	Edit(id int, change_dirr string, a interface{}) error
	Edit_idtg(id int, change_dirr string, a interface{}) error
	Read(id_bd int, db *T_productDatabase) (int, error)
	ReadAll_idtg(in_id_tg int, db *T_productDatabase) (int, error)
	ReadAll() ([]T_productDatabase, error)
	Delete(Article int)
	Delete_idtg(id_tg int)
}*/

func Create(Table_name string) error {
	// driverName := set.driverName
	// dataSourceName := set.dataSourceName
	driverName := "sqlite3"
	dataSourceName := "testdatabase/test_product1.db"
	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	var query string
	var variable string
	//id INTEGER PRIMARY KEY,
	if Table_name == "temp" {
		variable = "(id_tg INTEGER, Article INTEGER, name text, description text, photo string, price FLOAT, category text)"
	} else {
		variable = "(Article INTEGER, name text, description text, photo string, price FLOAT)"
	}
	query = "CREATE TABLE IF NOT EXISTS " + "`" + Table_name + "`" + " " + variable + ""

	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	defer database.Close()
	return nil
}
func (set *T_SettingDbProduct) Open() (T_SettingDbProduct, error) {
	// driverName := set.driverName
	// dataSourceName := set.dataSourceName
	driverName := "sqlite3"
	dataSourceName := "testdatabase/test_product1.db"
	database, err := sql.Open(driverName, dataSourceName)
	set.database = database
	return *set, err

}
func (set T_SettingDbProduct) All_table() ([]string, error) {

	var name_table string
	var table []string

	r, err := set.database.Query("SELECT name FROM sqlite_master WHERE type='table'")
	for r.Next() {
		r.Scan(&name_table)
		if name_table != "temp" {
			table = append(table, name_table)
		}
	}
	_ = err
	//to do pars error
	return table, nil

}
func (set T_SettingDbProduct) Check_table(check_Table_name string) error {

	//query := "show tables like " + "\"" + set.Table_name + "\"" //в sqlite3 так низя

	query := "SELECT * FROM `" + check_Table_name + "`"
	res, err := set.database.Query(query) //Подгтовленный запрос.
	if err != nil {
		return err
	}
	defer res.Close()
	return err

}
func (set T_SettingDbProduct) Close() {
	defer set.database.Close()
}

func (set T_SettingDbProduct) Add(db T_productDatabase) error {
	query := "INSERT INTO " + "`" + set.Table_name + "`" + " (Article,name, description, photo, price) VALUES (?,?,?,?,?)"
	statement, err := set.database.Prepare(query) //Подгтовленный запрос.
	if err != nil {
		fmt.Printf(err.Error())
	}

	statement.Exec(db.Article, db.Name, db.Description, db.Photo, db.Price) //to do
	defer statement.Close()
	return err

}
func (set T_SettingDbProduct) Addany(name string, a interface{}) error {
	query := "INSERT INTO `" + set.Table_name + "` (" + name + ") VALUES (?)"
	statement, err := set.database.Prepare(query) //Подгтовленный запрос.
	//var json_key []byte

	//json_key, _ = json.Marshal(db.name)

	statement.Exec(a) //to do
	defer statement.Close()
	return err

}

// delete corzine
// example change_dir = name
func (set T_SettingDbProduct) Edit(id int, change_dirr string, a interface{}) error {
	statement, err := set.database.Prepare("update " + set.Table_name + " set " + change_dirr + "=? where Article=?") //Подгтовленный запрос.
	statement.Exec(a, id)
	defer statement.Close()
	return err
}
func (set T_SettingDbProduct) Edit_idtg(id int, change_dirr string, a interface{}) error {
	statement, err := set.database.Prepare("update " + set.Table_name + " set " + change_dirr + "=? where id_tg=?") //Подгтовленный запрос.

	aa, err := statement.Exec(a, id)
	_ = aa

	defer statement.Close()
	return err
}

// id in bd
// id_bd - its Articleber in bd
func (set T_SettingDbProduct) Read(id_bd int, db *T_productDatabase) (int, error) {

	row, err := set.database.Query("SELECT  * FROM `" + set.Table_name + "`")
	//Прочитать с id user
	//var id int
	id_cout_err := 0
	var Article int
	var name string
	var description string
	var photo string
	var prices float64
	//var instock int
	var category string
	var id_tg int

	defer row.Close()
	for row.Next() {
		//defer database.Close()
		if set.Table_name == "temp" {
			row.Scan(&id_tg, &Article, &name, &description, &photo, &prices, &category)
		} else {
			row.Scan(&Article, &name, &description, &photo, &prices)
		}

		if id_bd == Article {

			db.Article = Article
			db.Description = description
			//db.Instock = instock
			db.Name = name
			db.Photo = photo
			db.Price = prices
			if set.Table_name == "temp" {
				db.Category = category
				db.Id_tg = id_tg
			}
			//id_cout = db.Article
			return db.Article, err
		}

	}
	var errNOfoundusr = errors.New("user not founded")

	return id_cout_err, errNOfoundusr
}

// считаем все из бд по айдишнику из тг
func (set T_SettingDbProduct) ReadAll_idtg(in_id_tg int, db *T_productDatabase) (int, error) {
	Table_name := "temp"
	row, err := set.database.Query("SELECT * FROM `" + Table_name + "`")
	//Прочитать с id user // id_tg, Article, name, description,photo,price,category
	//var id int
	id_cout_err := 0
	var Article int
	var name string
	var description string
	var photo string
	var prices float64
	//var instock int
	//var category string

	var id_tg int
	var aa string

	defer row.Close()
	for row.Next() {
		//defer database.Close()

		row.Scan(&id_tg, &Article, &name, &description, &photo, &prices, &aa)

		if in_id_tg == id_tg {

			db.Article = Article
			db.Description = description
			//db.Instock = instock
			db.Name = name
			db.Photo = photo
			db.Price = prices
			db.Category = aa
			db.Id_tg = id_tg
			return db.Article, err
		}

	}
	var errNOfoundusr = errors.New("user not founded")
	//errors.New()
	//err = 1
	return id_cout_err, errNOfoundusr
}

//

func (set T_SettingDbProduct) ReadAll() ([]T_productDatabase, error) {
	// rows_name, err := set.database.Query("SELECT *  FROM " + Table_name + "")
	// s, _ := rows_name.Columns() //SELECT *  FROM "
	// _ = s
	// rows_name.Next()
	// rows_name.NextResultSet()
	rows, err := set.database.Query("SELECT count(*)  FROM `" + set.Table_name + "`")
	var count int
	//Счетчик кол-ва строк
	for rows.Next() {
		rows.Scan(&count)
	}
	rows.Close()

	row, err := set.database.Query("SELECT  * FROM `" + set.Table_name + "`")
	var Article int
	var name string
	var description string
	var photo string
	var prices float64
	var instock int

	ddb := make([]T_productDatabase, count)

	defer row.Close()
	i := 0
	for row.Next() {
		//defer database.Close()
		if set.Table_name == "temp" {
			row.Scan(&ddb[i].Id_tg, &Article, &name, &description, &photo, &prices, &ddb[i].Category)
		} else {
			row.Scan(&Article, &name, &description, &photo, &prices)
		}

		ddb[i].Article = Article
		ddb[i].Description = description
		ddb[i].Instock = instock
		ddb[i].Name = name
		ddb[i].Photo = photo
		ddb[i].Price = prices
		i++

	}
	return ddb, err
}
func (set T_SettingDbProduct) Delete(Article int) {
	// удаляем строку с id=1
	/*
		result, err := database.Exec("delete from user where id = $1", 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected()) // количество удаленных строк*/

	statement, err := set.database.Prepare("delete from `" + set.Table_name + "` where Article=?") //Подгтовленный запрос.
	defer statement.Close()
	if err != nil {
		panic(err)
	}
	statement.Exec(Article)
	//fmt.Println("delete in table %s , id := %s", Table_name, id) // количество удаленных строк
}
func (set T_SettingDbProduct) Delete_idtg(id_tg int) {

	statement, err := set.database.Prepare("delete from `" + set.Table_name + "` where id_tg=?") //Подгтовленный запрос.
	defer statement.Close()
	if err != nil {
		panic(err)
	}
	statement.Exec(id_tg)
	//fmt.Println("delete in table %s , id := %s", Table_name, id) // количество удаленных строк
}
