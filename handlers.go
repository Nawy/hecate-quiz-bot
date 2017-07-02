package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"io/ioutil"
	"strconv"
	"strings"
)

// UserStatus for right cmdHandlers
type UserStatus string

const (
	RENAME = "rename"
	QUESTION = "question"
	CHOOSE_GAME = "choose_game"
	START_GAME = "start_game"
	IDLE = "idle"
)

const (
	AGREE_BUTTON = "1"
	DISAGREE_BUTTON = "2"
)

var cmdHandlers map[string]func(*tgbotapi.Update, *User, int64) *tgbotapi.MessageConfig = make(map[string]func(*tgbotapi.Update, *User, int64) *tgbotapi.MessageConfig)
var helpMsg string

func InitHandle() {
	cmdHandlers["callme"] = cmdCallme
	cmdHandlers["rename"] = cmdRename
	cmdHandlers["games"] = cmdGames
	cmdHandlers["continue"] = cmdContinueGame
	cmdHandlers["help"] = cmdHelp

	data, err := ioutil.ReadFile(conf.Bot.Resources.Help)
	if err != nil {
		log.Fatal(err)
	}
	helpMsg = string(data)
}

func Handle(update *tgbotapi.Update) *tgbotapi.MessageConfig {

	logMessage(update)
	userID, chatID := getUpdateParam(update)

	user := handleUser(update, userID)

	if update.Message != nil {
		handler := cmdHandlers[update.Message.Command()]
		if handler != nil {
			return handler(update, user, chatID)
		}
	}

	switch user.Status {
	case IDLE : return sendMessage(update, MESSAGES["idle_msg"], chatID)
	case RENAME : return cmdRenameProcess(update, user, chatID)
	case CHOOSE_GAME : return cmdGamesProcess(update, user, chatID)
	case START_GAME : return cmdStartGame(update, user, chatID)
	case QUESTION : return cmdQuestion(update, user, chatID)
	default : return sendMessage(update, MESSAGES["wrong_request"], chatID)
	}
}

func handleUser(update *tgbotapi.Update, userID int) *User {
	user := GetUser(userID)
	if user != nil {
		return user
	}

	user = &User{
		update.Message.From.ID,
		update.Message.From.UserName,
		update.Message.From.FirstName,
		IDLE,
		-1,
		-1,
		-1,
		-1,
		1,
		0,
	}

	InsertUser(user)
	return user
}

func cmdRename(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	user.Status = RENAME
	UpdateUser(user)

	msgProcessed := strings.Replace(MESSAGES["rename_hello"], "$name$", user.Name, 1)
	msg := tgbotapi.NewMessage(chatID, msgProcessed)
	return &msg
}

func cmdRenameProcess(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	if !nameRegexp.MatchString(update.Message.Text) {
		return sendMessage(update, MESSAGES["invalid_name_chars"], chatID)
	}
	if len(update.Message.Text) > 0 {
		user.Name = update.Message.Text
		user.Status = IDLE
		UpdateUser(user)
	}
	msgProcessed := strings.Replace(MESSAGES["rename_finish"], "$name$", user.Name, 1)
	return sendMessage(update, msgProcessed, chatID)
}

func cmdCallme(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	msgProcessed := strings.Replace(MESSAGES["callme_msg"], "$name$", user.Name, 1)
	return sendMessage(update, msgProcessed, chatID)
}

func cmdHelp(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
	return &msg
}

func cmdGames(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {

	var buttons [][]tgbotapi.InlineKeyboardButton = make([][]tgbotapi.InlineKeyboardButton, len(GAMES))
	for i, game := range GAMES {
		row := tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(EmojiReplace(game.Name), strconv.Itoa(game.Id)),
		)
		buttons[i] = row
	}

	user.Status = CHOOSE_GAME
	UpdateUser(user)

	var gamesButtons = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return sendMessageWithMarkup(update, MESSAGES["game_list_topic"], chatID, gamesButtons)
}

func cmdGamesProcess(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	if update.CallbackQuery == nil {
		return sendMessage(update, MESSAGES["choose_game_wrong"], chatID)
	}

	log.Println("CALLBACK ID: ", update.CallbackQuery.Data)
	choosenGame := GetGame(update.CallbackQuery.Data)

	user.Status = START_GAME
	user.SelectedGameId = choosenGame.Id
	UpdateUser(user)

	resultString := MESSAGES["game_hello"]
	resultString = strings.Replace(resultString, "$game_name$", choosenGame.Name, -1)
	resultString = strings.Replace(resultString, "$game_description$", choosenGame.Description, -1)
	resultString = strings.Replace(resultString, "$attempts$", strconv.Itoa(choosenGame.HintAttempts), -1)

	var yesNoButtons = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(MESSAGES["agree_button"], AGREE_BUTTON),
			tgbotapi.NewInlineKeyboardButtonData(MESSAGES["disagree_button"], DISAGREE_BUTTON),
		),
	)

	return sendMessageWithMarkup(update, resultString, chatID, yesNoButtons)
}

func cmdStartGame(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	if update.CallbackQuery == nil {
		return sendMessage(update, MESSAGES["game_wrong_reply"], chatID)
	}

	userDecision := update.CallbackQuery.Data

	if userDecision == AGREE_BUTTON {

		game := GetGame(strconv.Itoa(user.CurrentGameId))
		user.Status = QUESTION
		user.SelectedGameId = 0

		user.CurrentHintAttempt = game.HintAttempts
		user.CurrentAttempt = 0
		user.CurrentPoints = 0
		user.CurrentQuestionId = 0
		return cmdAskQuestion(update, user, MESSAGES["game_start"], chatID, game)
	} else if userDecision == DISAGREE_BUTTON {
		user.Status = IDLE
		user.SelectedGameId = 0
		UpdateUser(user)
		return sendMessage(update, MESSAGES["game_start_disagree"], chatID)
	}

	return sendMessage(update, MESSAGES["game_wrong_reply"], chatID)
}

func cmdContinueGame(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	if user.CurrentGameId == -1 {
		return sendMessage(update, MESSAGES["wrong_continue"], chatID)
	}

	game := GetGame(strconv.Itoa(user.CurrentGameId))

	result := MESSAGES["normal_continue"]
	result = strings.Replace(result, "$name$", user.Name, -1)
	return cmdAskQuestion(update, user, result, chatID, game)
}

func cmdQuestion(update *tgbotapi.Update, user *User, chatID int64) *tgbotapi.MessageConfig {
	game := GetGame(strconv.Itoa(user.CurrentGameId))
	currentQuestion := game.Questions[user.CurrentQuestionId]
	userAnswer := strings.ToLower(update.Message.Text)

	if isRightAnswer(userAnswer, currentQuestion.Answers) {
		responseText := MESSAGES["right_answer"]
		user.CurrentQuestionId = user.CurrentQuestionId + 1
		user.CurrentPoints = user.CurrentPoints + 1
		user.CurrentAttempt = 0
		return cmdAskQuestion(update, user, responseText, chatID, game)
	} else {
		user.CurrentAttempt = user.CurrentAttempt + 1
		UpdateUser(user)
		if user.CurrentAttempt == currentQuestion.Attempts {
			user.CurrentQuestionId = user.CurrentQuestionId + 1
			user.CurrentAttempt = user.CurrentAttempt + 1
			responseText := MESSAGES["wrong_answer"]
			return cmdAskQuestion(update, user, responseText, chatID, game)
		}
		resultString := MESSAGES["wrong_answer_attempt"]
		resultString = strings.Replace(resultString, "$attempt$", strconv.Itoa(currentQuestion.Attempts - user.CurrentAttempt), -1)
		return sendMessage(update, resultString, chatID)
	}
}

func cmdAskQuestion(update *tgbotapi.Update, user *User, text string, chatID int64, game *GameJSON) *tgbotapi.MessageConfig {

	// check game is finished?
	if user.CurrentQuestionId >= len(game.Questions) {
		return finishGame(update, user, text, game, chatID)
	}

	question := game.Questions[user.CurrentQuestionId]

	user.Status = QUESTION
	UpdateUser(user)

	result := text + MESSAGES["game_question"]
	result = strings.Replace(result, "$name$", question.Name, -1)
	result = strings.Replace(result, "$text$", question.Text, -1)
	return sendMessage(update, result, chatID)
}

func finishGame(update *tgbotapi.Update, user *User, text string, game *GameJSON, chatID int64) *tgbotapi.MessageConfig {
	result := text + MESSAGES["game_is_finished"]
	result = strings.Replace(result, "$points$", strconv.Itoa(user.CurrentPoints), -1)
	result = strings.Replace(result, "$length$", strconv.Itoa(len(game.Questions)), -1)

	user.Status = IDLE
	user.CurrentGameId = -1
	user.CurrentQuestionId = -1
	user.CurrentPoints = 0
	user.CurrentHintAttempt = 0
	UpdateUser(user)

	return sendMessage(update, result, chatID)
}

