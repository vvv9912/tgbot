package sdek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// расчет по коду тарифа
type SdekCodeTariffMsg struct {
	Type        int    `json:"type"`
	Date        string `json:"date"`
	Currency    int    `json:"currency"`
	Tariff_code int    `json:"tariff_code"`

	From_location sdekLocation      `json:"from_location"`
	To_location   sdekLocation      `json:"to_location"`
	Services      SdekService       `json:"operation"`
	Packages      []SdekInfPacakege `json:"packages"`
}

type SdekCodeTariffAnaswerMsg struct {
	Tariff_codes []SdekCodeTariff `json:"tariff_codes"`
	Errors       []SdekInfError   `json:"error"`
}

type SdekCodeTariff struct {
	Tariff_code        int     `json:"tariff_code"`
	Tariff_name        string  `json:"tariff_name"`
	Tariff_description string  `json:"tariff_description"`
	Delivery_mode      int     `json:"delivery_mode"`
	Delivery_sum       float64 `json:"delivery_sum"`
	Period_min         int     `json:"period_min"`
	Period_max         int     `json:"period_max"`
	Calendar_min       int     `json:"calendar_min"`
	Calendar_max       int     `json:"calendar_max"`
	//add
	Total_sum   float64     `json:"total_sum"`
	Weight_calc int         `json:"weight_calc"`
	Services    SdekService `json:"services"`
}

// по широтам долготам определили ближайший адресс сдэк. Взяли оттуда город, код, адресса.
// Далее рассчитываем сумму
func (ath *SdekAcess) SdekCalcCodeTariff(msg SdekCodeTariffMsg) (SdekCodeTariff, error) {
	//To_location.Code = 391 сумма зависит от кода
	jsonresp, err := json.Marshal(msg)
	if err != nil {
		log.Fatalln(err)
		return SdekCodeTariff{}, err
	}

	bearer := "Bearer " + ath.Access_token
	r, err := http.NewRequest(http.MethodPost, UrlCalcCodeTarif, bytes.NewReader(jsonresp))
	r.Header.Set("Authorization", bearer)
	r.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var ans SdekCodeTariff
	txt := string([]byte(body))
	err = json.Unmarshal([]byte(body), &ans)
	_ = txt
	_ = body

	return ans, err
}
