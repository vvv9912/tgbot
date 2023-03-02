package sdek

//расчет по доступным тарифам
import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// https://api-docs.cdek.ru/63345519.html

type SdekAllTariffMsg struct {
	Type          int               `json:"type"`
	Date          string            `json:"date"`
	Currency      int               `json:"currency"`
	Lang          string            `json:"lang"`
	From_location sdekLocation      `json:"from_location"`
	To_location   sdekLocation      `json:"to_location"`
	Services      SdekService       `json:"operation"`
	Packages      []SdekInfPacakege `json:"packages"`
}

// private
type sdekLocation struct {
	Code         int    `json:"code"`
	Postal_code  string `json:"postal_code"`
	Country_code string `json:"country_code"`
	City         string `json:"city"` //= city_code
	Address      string `json:"address"`
	//	City_code    int    `json:"city_code"`
}

//	type AnaswerMsg struct {
//		Delivery_sum float64  `json:"operation"`
//		Period_min   int      `json:"operation"`
//		Period_max   int      `json:"operation"`
//		Weight_calc  int      `json:"operation"`
//		Calendar_min int      `json:"operation"`
//		Calendar_max int      `json:"operation"`
//		Services     Service  `json:"operation"`
//		Total_sum    float64  `json:"operation"`
//		Currency     string   `json:"operation"`
//		Errors       InfError `json:"operation"`
//	}
type SdekAllTariffAnaswerMsg struct {
	Tariff_codes []sdekAllTariff `json:"tariff_codes"`
	Errors       []SdekInfError  `json:"error"`
}

// private
type sdekAllTariff struct {
	Tariff_code        int     `json:"tariff_code"`
	Tariff_name        string  `json:"tariff_name"`
	Tariff_description string  `json:"tariff_description"`
	Delivery_mode      int     `json:"delivery_mode"`
	Delivery_sum       float64 `json:"delivery_sum"`
	Period_min         int     `json:"period_min"`
	Period_max         int     `json:"period_max"`
	Calendar_min       int     `json:"calendar_min"`
	Calendar_max       int     `json:"calendar_max"`
}

// По доступным тарифам калькулятор
// func (msg SdekDeliveryMsg) PostInfo() {

func (ath *SdekAcess) SdekCalcAllTariff(msg SdekAllTariffMsg) (SdekAllTariffAnaswerMsg, error) {

	jsonresp, err := json.Marshal(msg)

	if err != nil {
		log.Fatalln(err)
	}

	bearer := "Bearer " + ath.Access_token
	_ = bearer
	//метод1
	// req, err := http.NewRequest("GET", UrlCalc, bytes.NewReader(jsonresp))
	// req.Header.Set("Authorization", bearer)
	// req.Header.Add("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println("Error on response.\n[ERROR] -", err)
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("Error while reading the response bytes:", err)
	// }
	// log.Println(string([]byte(body)))
	//метод2
	r, err := http.NewRequest(http.MethodPost, UrlCalc, bytes.NewReader(jsonresp))
	r.Header.Set("Authorization", bearer)
	r.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ans SdekAllTariffAnaswerMsg
	//_ = jsonrespp
	//json.NewDecoder(resp.Body).Decode(&ans)

	txt := string([]byte(body))
	err = json.Unmarshal([]byte(body), &ans)
	_ = txt
	_ = body

	return ans, err

}
