package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"log"
)

var client *as.Client

func InitAerospike() {
	var err error
	client, err = as.NewClient(conf.Bot.Aerospike.Host, conf.Bot.Aerospike.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func GenKey(telegramName string) (*as.Key, error) {
	return as.NewKey(conf.Bot.Aerospike.Namespace, "hecate_users", telegramName)
}

func UpsertUser(user User) {
	key, err := GenKey(user.telegramName)
	if err != nil {
		log.Fatal(err)
	}

	bins := as.BinMap{
		"name":          user.name,
		"telegram_name": user.telegramName,
		"state":         user.state,
		"game_id":       user.gameState.GameId,
		"question_id":   user.gameState.QuestionId,
		"attempts":      user.gameState.Attempt,
	}

	err = client.Put(nil, key, bins)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUser(telegramName string) *User {
	key, err := GenKey(telegramName)
	if err != nil {
		log.Fatal(err)
	}

	results, err := client.Get(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	return &User{
		string(results.Bins["name"]),
		string(results.Bins["telegram_name"]),
		string(results.Bins["state"]),
		&GameState{
			int32(results.Bins["game_id"]),
			int32(results.Bins["question_id"]),
			int32(results.Bins["attempts"]),
		},
	}
}
