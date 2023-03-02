package mydadata

import (
	"fmt"

	"gopkg.in/webdeskltd/dadata.v2"
)

func Dadata(adr, apikey, secretkey string) []string {

	daData := dadata.NewDaData(apikey, secretkey)

	addresses, err := daData.SuggestAddresses(dadata.SuggestRequestParams{Query: adr, Count: 10})
	if nil != err {
		fmt.Println(err)
	}
	if len(addresses) != 0 {
		var aa []string

		for _, address := range addresses {
			aa = append(aa, address.Value)
		}
		return aa
	} else {
		return nil
	}
}
