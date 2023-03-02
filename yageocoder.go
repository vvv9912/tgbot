package main

import (
	"errors"
	"sample-app/tg2/sdek"
	"sample-app/tg2/yageocoder"
)

func YaGeocoder(adr string) (*sdek.CoordUser, error) {
	api := YaGeocoderapi
	geocoder := yageocoder.Geocoder(api)
	location, err := geocoder.Geocode(adr)
	_ = err // todo
	if location != nil {
		return &sdek.CoordUser{Long: location.Lng, Latit: location.Lat}, nil
	} else {
		return &sdek.CoordUser{}, errors.New("Error geocoder")
	}
}
