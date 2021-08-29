package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Data struct {
		Base     string `json:"base"`
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"data"`
}

func main() {
	for {
		getBtcPrice()
		time.Sleep(time.Second * 60)
	}
}

func getBtcPrice() {

	var tooLowThreeshold float64 = 30000

	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=EUR")

	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		output, _ := ioutil.ReadAll(response.Body)

		var response Response
		err = json.Unmarshal(output, &response)
		if err != nil {
			fmt.Println(PrettyPrint(response))
		} else {
			amount, err := strconv.ParseFloat(response.Data.Amount, 64)
			if err != nil {
				fmt.Println(err)
			}

			if amount >= tooLowThreeshold {
				sendTelegramMessage(1877098271, string(response.Data.Amount))
			}
		}
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
