package RandomGoodHanime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Takes images from akaneko-api.herokuapp.com
func RetrieveHentai(parameter string) string {
	api := "https://akaneko-api.herokuapp.com/api/"

	response, err := http.Get(api + parameter)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data HentaiAPIs
	json.Unmarshal(responseData, &data)

	return data.Url
}
