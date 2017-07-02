package main

import (
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type GameJSON struct {
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
	} `json:"questions"`
}

type GamesJSON []GameJSON

var GAMES GamesJSON = GamesJSON{}
var GAMES_MAP map[string] int = make(map[string] int)

func LoadGames() {
	gamesFile, err := ioutil.ReadFile(conf.Bot.Resources.Games)
	if err != nil {
		panic("Not found games.json by path " + conf.Bot.Resources.Messages)
	}

	err = json.Unmarshal(gamesFile, &GAMES)
	if err != nil {
		panic("Cannot read games.json by path " + conf.Bot.Resources.Messages)
	}

	for i, game := range GAMES {
		gameIdString := strconv.Itoa(game.Id)
		GAMES_MAP[gameIdString] = i
	}
}

func GetGame(gameId string) *GameJSON {
	return &GAMES[GAMES_MAP[gameId]]
}