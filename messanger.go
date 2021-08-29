package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func sendTelegramMessage(chatId int, text string) (string, error) {

	var bot_token string = "insert your bot_token here"
	var telegramUrl string = "https://api.telegram.org/" + bot_token + "/sendMessage"

	response, err := http.PostForm(
		telegramUrl,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	result, resultErr := ioutil.ReadAll(response.Body)
	if resultErr != nil {
		return "", resultErr
	}

	return string(result), nil
}
