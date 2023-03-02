package yageocoder

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Response struct {
			GeoObjectCollection struct {
				MetaDataProperty struct {
					GeocoderResponseMetaData struct {
						Request string `json:"request"`
						Found   string `json:"found"`
						Results string `json:"results"`
					} `json:"GeocoderResponseMetaData"`
				} `json:"metaDataProperty"`
				FeatureMember []*yandexFeatureMember `json:"featureMember"`
			} `json:"GeoObjectCollection"`
		} `json:"response"`
	}

	yandexFeatureMember struct {
		GeoObject struct {
			MetaDataProperty struct {
				GeocoderMetaData struct {
					Kind      string `json:"kind"`
					Text      string `json:"text"`
					Precision string `json:"precision"`
					Address   struct {
						CountryCode string `json:"country_code"`
						PostalCode  string `json:"postal_code"`
						Formatted   string `json:"formatted"`
						Components  []struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
						} `json:"Components"`
					} `json:"Address"`
				} `json:"GeocoderMetaData"`
			} `json:"metaDataProperty"`
			Description string `json:"description"`
			Name        string `json:"name"`
			BoundedBy   struct {
				Envelope struct {
					LowerCorner string `json:"lowerCorner"`
					UpperCorner string `json:"upperCorner"`
				} `json:"Envelope"`
			} `json:"boundedBy"`
			Point struct {
				Pos string `json:"pos"`
			} `json:"Point"`
		} `json:"GeoObject"`
	}
)

const (
	componentTypeHouseNumber   = "house"
	componentTypeStreetName    = "street"
	componentTypeLocality      = "locality"
	componentTypeStateDistrict = "area"
	componentTypeState         = "province"
	componentTypeCountry       = "country"
)

// Geocoder constructs Yandex geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	return geo.HTTPGeocoder{
		EndpointBuilder:       baseURL(getURL(apiKey, baseURLs...)),
		ResponseParserFactory: func() geo.ResponseParser { return &geocodeResponse{} },
	}
}

func getURL(apiKey string, baseURLs ...string) string {
	if len(baseURLs) > 0 {
		return baseURLs[0]
	}
	return fmt.Sprintf("https://geocode-maps.yandex.ru/1.x/?results=1&lang=en_US&format=json&apikey=%s&", apiKey)
}

func (b baseURL) GeocodeURL(address string) string {
	return string(b) + "geocode=" + address
}

func (b baseURL) ReverseGeocodeURL(l geo.Location) string {
	return string(b) + fmt.Sprintf("sco=latlong&geocode=%f,%f", l.Lat, l.Lng)
}

func (r *geocodeResponse) Location() (*geo.Location, error) {
	if r.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Found == "0" {
		return nil, nil
	}
	if len(r.Response.GeoObjectCollection.FeatureMember) == 0 {
		return nil, nil
	}
	featureMember := r.Response.GeoObjectCollection.FeatureMember[0]
	result := &geo.Location{}
	latLng := strings.Split(featureMember.GeoObject.Point.Pos, " ")
	if len(latLng) > 1 {
		// Yandex return geo coord in format "long lat"
		result.Lat, _ = strconv.ParseFloat(latLng[1], 64)
		result.Lng, _ = strconv.ParseFloat(latLng[0], 64)
	}

	return result, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	if r.Response.GeoObjectCollection.MetaDataProperty.GeocoderResponseMetaData.Found == "0" {
		return nil, nil
	}
	if len(r.Response.GeoObjectCollection.FeatureMember) == 0 {
		return nil, nil
	}

	return parseYandexResult(r.Response.GeoObjectCollection.FeatureMember[0]), nil
}

func parseYandexResult(r *yandexFeatureMember) *geo.Address {
	addr := &geo.Address{}
	res := r.GeoObject.MetaDataProperty.GeocoderMetaData

	for _, comp := range res.Address.Components {
		switch comp.Kind {
		case componentTypeHouseNumber:
			addr.HouseNumber = comp.Name
			continue
		case componentTypeStreetName:
			addr.Street = comp.Name
			continue
		case componentTypeLocality:
			addr.City = comp.Name
			continue
		case componentTypeStateDistrict:
			addr.StateDistrict = comp.Name
			continue
		case componentTypeState:
			addr.State = comp.Name
			continue
		case componentTypeCountry:
			addr.Country = comp.Name
			continue
		}
	}

	addr.Postcode = res.Address.PostalCode
	addr.CountryCode = res.Address.CountryCode
	addr.FormattedAddress = res.Address.Formatted

	return addr
}

// func Geocoder() {
// 	url := "https://geocode-maps.yandex.ru/1.x"
// 	api := "5b44a9c7-a890-452c-8b20-0b3432ce513d"
// 	adr := "Москва,тверская+6"
// 	format := "json"
// 	fulladr := fmt.Sprintf("%s?apikey=%s&format=%s&geocode=%s", url, api, format, adr)
// 	//fulladr = "https://geocode-maps.yandex.ru/1.x/?apikey=5b44a9c7-a890-452c-8b20-0b3432ce513d&geocode=Москва,Тверская+6"
// 	// resp, err := http.Get(fulladr)
// 	r, err := http.NewRequest(http.MethodGet, fulladr, nil)
// 	_ = err
// 	resp, err := http.DefaultClient.Do(r)
// 	_ = resp
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)

// 	txt := string([]byte(body))
// 	fmt.Println(txt)
// 	//fmt.Println(body)
// 	_ = err
// 	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
// 		fmt.Println("HTTP Status is in the 2xx range")
// 	} else {
// 		fmt.Println("Argh! Broken")
// 		err := errors.New(fmt.Sprintf(resp.Status))
// 		//return []SdekAnswOffice{}, err
// 		fmt.Println(err)
// 	}

// }
