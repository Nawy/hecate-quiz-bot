package main

type GameState struct {
	GameId int32
	QuestionId int32
	Attempt int32
}

type User struct {
	name string
	telegramName string
	status UserStatus
	gameState *GameState
}
