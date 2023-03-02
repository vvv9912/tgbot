package sdek

import "time"

// type Money struct {
// 	Value float64
// 	//Vat_sum  float64
// 	//Vat_rate float64
// }

// // threshold - дополнительный сбор за доставку в зависимости от стоимости товара
// type Threshold struct {
// 	Threshold int
// 	Sum       float64
// 	//Vat_sum   float64
// 	//Vat_rate  int
// }

// // contact - данные контрагента (отправителя, получателя)
// type Contact struct {
// 	//Company string
// 	Name   string
// 	Email  string
// 	Phones Phone
// 	//
// }

// // phone - номер телефона (мобильный/городской)
// type Phone struct {
// 	Number string
// }

// // данные истинного продавца
// type Seller struct {
// 	Name           string
// 	Inn            string
// 	Phone          string
// 	Ownership_form int
// 	Adress         string
// }

// // location - адрес местоположения контрагента (отправителя, получателя), включая геолокационные данные
// type SdekLocation struct {
// 	Address string
// }

// service - данные о дополнительных услугах
type SdekService struct {
	Code int     `json:"code"`
	Sum  float64 `json:"sum"`
	//	Parameter string
}

// package - информация о местах заказа
type SdekInfPacakege struct {
	//Number  string
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width  int `json:"width"`
	Height int `json:"height"`
	//Comment string
	//Items   Item
}

// item - информация о товарах места заказа (только для заказа типа "интернет-магазин")
// type SdekItem struct {
// 	Name     string
// 	Ware_key string
// 	Payment  Money
// 	Cost     float32
// 	Weight   int
// 	//Weight_gross int
// 	Amount int
// 	// Name_i18n    string
// 	// Brand        string
// 	// Country_code string
// 	// Material     string
// 	// Wifi_gsm     bool
// 	// Url          string
// }

type SdekInfError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SdekStatus struct {
	Code      []string
	Name      string
	Date_time time.Time //?dateTime  дата и время в формате ISO 8601: YYYY-MM-DDThh:mm:ss±hhmm.
	//	Reason_code string
	City string
}
