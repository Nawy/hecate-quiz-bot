package main

type User struct {
	ID int
	Login string
	Name string
	Status string

	SelectedGameId int

	CurrentGameId int
	CurrentQuestionId int
	CurrentHintAttempt int
	CurrentAttempt int
	CurrentPoints int
}
