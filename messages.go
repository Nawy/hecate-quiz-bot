package main

import (
	"io/ioutil"
	"encoding/json"
)

type MessagesJSON map[string] struct {
	Hello string `json:"hello"`
	Wrong string `json:"wrong"`
	Success string `json:"success"`
	Bye string `json:"bye"`
}

var MESSAGES MessagesJSON = MessagesJSON{}

func LoadMessages() {
	messageFile, err := ioutil.ReadFile(conf.Bot.Resources.Messages)
	if err != nil {
		panic("Not found message.json by path " + conf.Bot.Resources.Messages)
	}

	err = json.Unmarshal(messageFile, &MESSAGES)
	if err != nil {
		panic("Cannot read message.json by path " + conf.Bot.Resources.Messages)
	}
}