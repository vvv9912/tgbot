package sdek

import (
	"log"
	"time"
)

const (
	UrlAuth          = "https://api.edu.cdek.ru/v2/oauth/token?parameters"
	UrlCalc          = "https://api.edu.cdek.ru/v2/calculator/tarifflist"
	UrlOffice        = "https://api.edu.cdek.ru/v2/deliverypoints"
	UrlCalcCodeTarif = "https://api.edu.cdek.ru/v2/calculator/tariff"
)

type ServiceAcessConf interface {
	SetAuth() ServiceSdek
}

type ServiceSdek interface {
	ServiceAcessConf

	SdekCalcAllTariff(msg SdekAllTariffMsg) (SdekAllTariffAnaswerMsg, error)
	SdekCalcCodeTariff(msg SdekCodeTariffMsg) (SdekCodeTariff, error)
	PostOffice(msg SdekOffice, coordUser CoordUser) ([]SdekAnswOffice, error)
}

// constructor
func NewAuth() ServiceSdek {
	return &SdekAcess{}
}
func Test() {
	MytestSdekCalcAllTariff()
	// MytestSdekCalcCodeTariff()
	// MytestSearchAllOffice()
}
func MytestSdekCalcAllTariff() {
	var msg SdekAllTariffMsg
	msg.Type = 1
	msg.Date = time.Now().Format("2006-01-02T15:04:05-0700")
	msg.Currency = 1
	msg.From_location.Code = 270
	msg.Lang = "rus"
	msg.To_location.Code = 44
	//msg.To_location.City = "Москва"
	//msg.From_location.City = "Москва"
	msg.Packages = make([]SdekInfPacakege, 1)
	msg.Packages[0].Height = 10
	msg.Packages[0].Length = 10
	msg.Packages[0].Weight = 4000
	msg.Packages[0].Width = 10

	client := NewAuth().SetAuth()
	ans, err := client.SdekCalcAllTariff(msg)
	log.Println(ans)
	log.Println(err)
}
func MytestSdekCalcCodeTariff() {
	var msg SdekCodeTariffMsg
	msg.Type = 1
	msg.Currency = 1
	msg.Tariff_code = 136
	msg.From_location.City = "Москва"
	msg.From_location.Address = "Россия, Москва, Москва, ул. Шаболовка, 38"

	msg.From_location.Code = 44
	//
	msg.To_location.City = "Москва"
	msg.To_location.Address = "Россия, Москва, Москва, ш. Измайловское, 29"
	msg.To_location.Code = 44

	msg.To_location.City = "Железнодорожный"
	msg.To_location.Address = "Россия, Московская область, Железнодорожный, ул. Речная, 16"
	msg.To_location.Code = 391

	msg.Packages = make([]SdekInfPacakege, 1)
	msg.Packages[0].Height = 10
	msg.Packages[0].Length = 10
	msg.Packages[0].Weight = 4000
	msg.Packages[0].Width = 10
	client := NewAuth().SetAuth()

	ans, err := client.SdekCalcCodeTariff(msg)
	log.Println(ans)
	log.Println(err)
}
func MytestSearchAllOffice() {
	var msg SdekOffice
	msg.Country_code = "643"
	msg.Type = "ALL"
	client := NewAuth().SetAuth()

	var coordUser CoordUser
	coordUser.Long = 37.617
	coordUser.Latit = 55.755
	coordUser.Dcoord = 1500
	off, err := client.PostOffice(msg, coordUser)
	_ = off
	_ = err
}
