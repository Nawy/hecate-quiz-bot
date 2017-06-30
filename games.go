package main

import (
	"io/ioutil"
	"encoding/json"
)

type GamesJSON []struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	HintAttempts int `json:"hint_attempts"`
	Questions []struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Text string `json:"text"`
		Image string `json:"image"`
		Audio string `json:"audio"`
		Answers []string `json:"answers"`
		Hint string `json:"hint"`
		Attempts int `json:"attempts"`
		Points int `json:"points"`
	}
}

var GAMES GamesJSON = GamesJSON{}

func LoadGames() {
	gamesFile, err := ioutil.ReadFile(conf.Bot.Resources.Games)
	if err != nil {
		panic("Not found games.json by path " + conf.Bot.Resources.Messages)
	}

	err = json.Unmarshal(gamesFile, &GAMES)
	if err != nil {
		panic("Cannot read games.json by path " + conf.Bot.Resources.Messages)
	}
}