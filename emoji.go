package main

import "strings"

type EmojiType struct {
	Name string
	Code string
}

var EMOJIS []EmojiType = make([]EmojiType, 10)

func InitEmojis() {
	EMOJIS[0] = EmojiType{":wink:", "\xF0\x9F\x98\x89"}
	EMOJIS[1] = EmojiType{":devil:", "\xF0\x9F\x98\x88"}
	EMOJIS[2] = EmojiType{":msquare:", "\xE2\x97\xBC"}
	EMOJIS[3] = EmojiType{":uppoint:", "\xE2\x98\x9D"}
	EMOJIS[4] = EmojiType{":party:", "\xF0\x9F\x8E\x89"}
	EMOJIS[5] = EmojiType{":oksign:", "\xF0\x9F\x91\x8C"}
	EMOJIS[6] = EmojiType{":pensivef:", "\xF0\x9F\x98\x94"}
	EMOJIS[7] = EmojiType{":smile:", "\xF0\x9F\x98\x8A"}
	EMOJIS[8] = EmojiType{":question:", "\xE2\x9D\x93"}
	EMOJIS[9] = EmojiType{":smirkingf:", "\xF0\x9F\x98\x8F"}
}

func EmojiReplace(value string) string {
	result := value
	for _, emoji := range EMOJIS {
		result = strings.Replace(result, emoji.Name, emoji.Code, -1)
	}
	return result
}
