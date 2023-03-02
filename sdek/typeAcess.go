package sdek

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// для авторизации
type SdekAccount struct {
	Grant_type string `json:"grant_type"`
	Login      string `json:"client_id"`
	Password   string `json:"client_secret"`
}
type SdekAcess struct {
	//setauth      *SdekAccount
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   string `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

func (ans *SdekAcess) SetAuth() ServiceSdek {
	//Метод 1

	// data := url.Values{}

	// msg := GetLP()
	// data.Set("grant_type", msg.Grant_type)
	// data.Set("client_id", msg.Login)
	// data.Set("client_secret", msg.Password)
	// client := &http.Client{}
	// r, _ := http.NewRequest(http.MethodPost, UrlAuth, strings.NewReader(data.Encode())) // URL-encoded payload
	// r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// resp, _ := client.Do(r)
	// //var ans Acess

	// //_ = jsonrespp
	// json.NewDecoder(resp.Body).Decode(&ans)
	// fmt.Println(resp.Status)
	//Метод 2
	data := url.Values{}
	msg := getLP()

	data.Set("grant_type", msg.Grant_type)
	data.Set("client_id", msg.Login)
	data.Set("client_secret", msg.Password)
	r, err := http.Post(UrlAuth, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalln(err)
	}
	//var ans Acess
	json.NewDecoder(r.Body).Decode(&ans) //todo
	fmt.Println(r.Status)
	return ans
}
func getLP() (a SdekAccount) {
	//todo переделать хранение апи, вынести в конструктор
	Gr := "client_credentials"
	ALogin := "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI" //тестовый логин
	Securepassword := "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	return SdekAccount{Grant_type: Gr, Login: ALogin, Password: Securepassword}
}
