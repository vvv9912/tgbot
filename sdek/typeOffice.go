package sdek

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// Структура офисов сдэк
type SdekOffice struct {
	Country_code string `json:"country_code"`
	Type         string `json:"code"`
}

type SdekAnswOffice struct {
	Code       string             `json:"code"`
	Locat      SdekLocationOffice `json:"location"`
	Is_handout bool               `json:"is_handout"`
	Range      float64
}
type SdekLocationOffice struct {
	City         string  `json:"city"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Address_full string  `json:"address_full"`
	City_code    int     `json:"city_code"`
}
type CoordUser struct {
	Long   float64
	Latit  float64
	Dcoord float64
}

// Поиск ближайшего офиса //seacrhOffice
func (ath *SdekAcess) PostOffice(msg SdekOffice, coordUser CoordUser) ([]SdekAnswOffice, error) {
	Url := UrlOffice

	Long := coordUser.Long
	Latit := coordUser.Latit
	Dcoord := coordUser.Dcoord //m

	jsonresp, err := json.Marshal(msg)

	if err != nil {
		log.Fatalln(err)
	}

	bearer := "Bearer " + ath.Access_token
	r, err := http.NewRequest(http.MethodGet, Url, bytes.NewReader(jsonresp))
	r.Header.Set("Authorization", bearer)
	r.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var ans []SdekAnswOffice
	//var Msc []AnswOffice
	//json.NewDecoder(resp.Body).Decode(&ans)
	var office []SdekAnswOffice
	txt := string([]byte(body))
	err = json.Unmarshal([]byte(body), &ans)
	//log.Println(ans)

	for i := range ans {
		//coordUser	if ans[i].Is_handout {
		r, ok := searhCoord(Long, Latit, ans[i].Locat.Longitude, ans[i].Locat.Latitude, Dcoord)
		if ok {
			ans[i].Range = r
			office = append(office, ans[i])
		}
		//	}
	}
	// m := make(map[string][]string)
	// for i := range ans {
	// 	m[ans[i].Locat.City] = append(m[ans[i].Locat.City], ans[i].Locat.Address_full)

	// }
	// a := m["Бутово"]
	// a1 := m["Москва"]
	// // a2 := m["Коммунарка"]
	// _ = a
	// _ = a1
	// _ = a2
	_ = txt
	_ = body
	return office, err
}

// func SearhCoord(longist, latiist, long2, lati2, dcoord float64) bool {
// 	dphiist := math.Abs((latiist - lati2) * 60 * 1852) //latitude широта
// 	dlymbdaist := math.Abs((longist - long2) * 60 * 1852 * math.Cos(latiist*math.Pi/180))
// 	if dphiist <= dcoord {
// 		if dlymbdaist <= dcoord {
// 			if math.Sqrt(dphiist*dphiist+dlymbdaist*dlymbdaist) < float64(dcoord) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

func searhCoord(longist, latiist, long2, lati2, dcoord float64) (float64, bool) {
	dphiist := math.Abs((latiist - lati2) * 60 * 1852) //latitude широта
	dlymbdaist := math.Abs((longist - long2) * 60 * 1852 * math.Cos(latiist*math.Pi/180))
	if dphiist <= dcoord {
		if dlymbdaist <= dcoord {
			r := math.Sqrt(dphiist*dphiist + dlymbdaist*dlymbdaist)
			if r < float64(dcoord) {
				return r, true
			}
		}
	}
	return 0, false
}
