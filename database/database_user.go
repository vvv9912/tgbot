package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"sample-app/tg2/sdek"
)

type T_CORZINE struct {
	Id        string `json:"id"`
	Chtokypil []int  `json:"chtokypil"`
	Numof     []int  `json:"numof"` //ctokypil[i] = numof[i] //article
}
type T_uorders struct {
	Uorders []int
}
type T_UserPvz struct {
	Sdek_pvz   sdek.SdekAnswOffice `json:"pvz"`
	Number_pvz int                 `json:"number_pvz"` //задаю сам для callback
}
type T_Database struct {
	//	t_file  T_file
	//id_base int
	Id_user int `json:"id_tg"`
	Status  int //admin or user //-1 дефолт
	State   int //состояние (что делает) //-1 ничего
	// login    string
	// password string
	// toDO     string
	Corzine T_CORZINE
	Uorders T_uorders
	UserPvz []T_UserPvz
}

type T_file struct {
	driverName     string
	dataSourceName string
}

type T_SettingDbUser struct {
	database   *sql.DB
	Table_name string
}

func NewBDUser() dbOpenUser {
	return &T_SettingDbUser{}
}

type dbOpenUser interface {
	Open() (T_SettingDbUser, error)
}

/*
type dbaseproducter interface {
	T_SettingDbUser
	//Create(driverName, dataSourceName string, table_name string)
	Open() (T_SettingDbUser, error)
	Close()
	// Функции
	Adduser(db T_Database) error
	Edit(id int, change_dirr string, a interface{}) error
	EditTGid(id_tg int, change_dirr string, a interface{}) error
	Read(id_user_tg int, db *T_Database) (int, error)
	Bd_checkUser(id_user_tg int) (int, int, error)
	Delete(id int) error

	//delete_all()

}*/

func CreateDBUser(table_name string) error {
	driverName := "sqlite3"
	//log:botuser
	//pass:botuser
	//connStr := "user=botuser password=botuser dbname=data_test1 sslmode=disable"
	//driverName = "postgres"
	dataSourceName := "testdatabase/data_test1.db"
	database, err := sql.Open(driverName, dataSourceName)

	//database, err := sql.Open(driverName, connStr)
	if err != nil {
		return err
	}
	var query string
	variable := "(id INTEGER PRIMARY KEY, id_user INTEGER,status INTEGER, state INTEGER, Corz BLOB,pvz BLOB, Orderr BLOB)"
	query = "CREATE TABLE IF NOT EXISTS " + table_name + " " + variable + ""
	//query = strconv.Itoa(db.id_user)

	statement, err := database.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	defer database.Close()
	return nil
}
func (set *T_SettingDbUser) Open() (T_SettingDbUser, error) {
	driverName := "sqlite3"
	dataSourceName := "testdatabase/data_test1.db"
	dataabase, err := sql.Open(driverName, dataSourceName)
	set.database = dataabase
	return *set, err
}
func (set *T_SettingDbUser) Close() {
	defer set.database.Close()
}
func (set *T_SettingDbUser) Adduser(db T_Database) error {
	statement, err := set.database.Prepare("INSERT INTO " + set.Table_name + " (id_user, status,state, Corz,pvz,Orderr) VALUES (?,?,?,?,?,?)") //Подгтовленный запрос.
	statement.Exec(db.Id_user, db.Status, db.State, -1, -1, -1)
	defer statement.Close()
	return err

}

// редактирование по айди из бд
func (set *T_SettingDbUser) Edit(id int, change_dirr string, a interface{}) error {
	statement, err := set.database.Prepare("update " + set.Table_name + " set " + change_dirr + "=? where id=?") //Подгтовленный запрос.
	statement.Exec(a, id)
	defer statement.Close()
	return err
}

// редактирование по ТГ айди
func (set *T_SettingDbUser) EditTGid(id_tg int, change_dirr string, a interface{}) error {
	statement, err := set.database.Prepare("update " + set.Table_name + " set " + change_dirr + "=? where id_user=?") //Подгтовленный запрос.
	statement.Exec(a, id_tg)
	defer statement.Close()
	return err
}

// variable := "(id INTEGER PRIMARY KEY, id_user INTEGER,status INTEGER, state INTEGER, pvz BLOB,Corz BLOB, Orderr BLOB)"
func (set *T_SettingDbUser) Read(id_user_tg int, db *T_Database) (int, error) {

	row, err := set.database.Query("SELECT id,id_user, status,state, Corz,pvz,Orderr FROM " + set.Table_name + "")
	//Прочитать с id user
	var id int
	id_cout := 0
	var id_user int
	//var login string
	var bufPvz []byte
	var bufCorz []byte
	var bufOrderr []byte
	var status int
	var state int
	defer row.Close()
	for row.Next() {
		//defer database.Close()
		row.Scan(&id, &id_user, &status, &state, &bufCorz, &bufPvz, &bufOrderr)
		if id_user == id_user_tg {
			_ = json.Unmarshal(bufCorz, &db.Corzine)   //todo
			_ = json.Unmarshal(bufOrderr, &db.Uorders) //todo
			_ = json.Unmarshal(bufPvz, &db.UserPvz)    //todo

			//row.Close()
			//ПРИРАВНИВАНЕИ СТРУКТУРЕ
			db.Id_user = id_user
			db.Status = status
			db.State = state
			id_cout = id //айди в таблице
			return id_cout, err
		}

	}
	var errNOfoundusr = errors.New("uesr not founded")

	return id_cout, errNOfoundusr
}

// func (set *T_SettingDbUser) Bd_checkUser(id_user_tg int, table_name string) (int, int, error) {
func (set *T_SettingDbUser) Bd_checkUser(id_user_tg int) (int, int, error) {
	row, err := set.database.Query("SELECT id_user, status,state FROM " + set.Table_name + "")
	if err != nil {
		return -1, -1, err
	}

	var id_user int
	var status int
	status = -1
	var state int
	state = -1
	defer row.Close()
	for row.Next() {
		row.Scan(&id_user, &status, &state)
		if id_user == id_user_tg {
			return status, state, err
		}

	}

	errNOfoundusr := errors.New("uesr not founded")

	return -1, -1, errNOfoundusr
}
func (set *T_SettingDbUser) Delete(id int) error {
	// удаляем строку с id=1
	/*
		result, err := database.Exec("delete from user where id = $1", 1)
		if err != nil {
			panic(err)
		}
		fmt.Println(result.RowsAffected()) // количество удаленных строк*/

	statement, err := set.database.Prepare("delete from " + set.Table_name + " where id=?") //Подгтовленный запрос.
	defer statement.Close()
	if err != nil {
		panic(err)
	}
	statement.Exec(id)
	return err
}

// _auth_user
// _auth_pass //https://github.com/mattn/go-sqlite3#user-authentication
