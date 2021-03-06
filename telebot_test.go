package gotelebot

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	bot := InitTeleBot(token)
	me, err := bot.GetMe()
	if err != nil {
		assert.Fail("Bot getMe error")
	}
	assert.NotEmpty(me.Id)
}

func TestSendMessage(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	testMsg := "Test Msg"
	msg, err := bot.SendMessage(chatid, testMsg, nil)
	if err != nil {
		assert.Fail("Bot send message error")
	}
	assert.EqualValues(testMsg, msg.Text)
}

func SendMsgAndGetId() int {
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	testMsg := "Test Msg"
	msg, _ := bot.SendMessage(chatid, testMsg, nil)
	return int(msg.Message_Id)
}

func TestSendMessageWithOpt(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	testMsg := "Test Msg"
	dis := true
	rid := SendMsgAndGetId()
	opt := &SendMessageOptional{DisableWebPagePreview: &dis, ReplyToMessageId: &rid}
	msg, err := bot.SendMessage(chatid, testMsg, opt)
	if err != nil {
		assert.Fail("Bot send message error")
	}
	assert.EqualValues(testMsg, msg.Text)
}

func TestForwardMessage(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	msg, err := bot.ForwardMessage(chatid, chatid, SendMsgAndGetId())
	if err != nil {
		assert.Fail("Bot forwardMessage error")
	}
	assert.NotEmpty(msg.ForwardFrom)

}

func TestSendPhoto(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/go.png"
	msg, err := bot.SendPhoto(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendPhoto error")
	}
	assert.NotEmpty(msg.Photo)
}

func TestSendPhotoWithOpt(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/go.png"
	rsi := SendMsgAndGetId()
	opt := &SendPhotoOptional{ReplyToMessageId: &rsi}
	msg, err := bot.SendPhoto(chatid, filePath, opt)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendPhoto error")
	}
	assert.NotEmpty(msg.Photo)
}

func TestSendAudio(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/record.mp3"
	msg, err := bot.SendAudio(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendAudio error")
	}
	assert.NotEmpty(msg)
}

func TestSendAudioWithOpt(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	performer := "tele"
	title := "gram"
	rsi := SendMsgAndGetId()
	opt := &SendAudioOptional{Performer: &performer, Title: &title, ReplyToMessageId: &rsi}
	filePath := "./test_data/record.mp3"
	msg, err := bot.SendAudio(chatid, filePath, opt)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendAudio error")
	}
	assert.Equal(msg.Audio.Title, title)
}

func TestSendDocument(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/go.png"
	msg, err := bot.SendDocument(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendDocument error")
	}
	assert.NotEmpty(msg.Document.FileId)
}

func TestSendSticker(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/go.webp"
	msg, err := bot.SendSticker(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendStick error")
	}
	assert.NotEmpty(msg.Sticker.FileId)
}

func TestSendVideo(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/test_video.mp4"
	msg, err := bot.SendVideo(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendVideo error")
	}
	assert.NotEmpty(msg.Video.FileId)
}

func TestSendVoice(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	filePath := "./test_data/record.ogg"
	msg, err := bot.SendVoice(chatid, filePath, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot sendVOice error")
	}
	assert.NotEmpty(msg.Voice.FileId)
}

func TestSendLocation(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	lat := 26.3875591
	lon := -161.2901042
	msg, err := bot.SendLocation(chatid, lat, lon, nil)
	if err != nil {
		assert.Fail("Bot send Location error")
	}
	assert.NotEmpty(msg.Location.Latitude)
}

func TestSendChatAction(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	msg, err := bot.SendChatAction(chatid, "typing")
	if err != nil {
		assert.Fail("Bot send ChatAction error")
	}
	fmt.Println(msg)
}

func TestGetUserProfilePhotos(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	photos, err := bot.GetUserProfilePhotos(chatid, nil)
	if err != nil {
		fmt.Println(err)
		assert.Fail("Bot get User Profile phtots fail.")
	}
	assert.True(len(*photos.Photos) > 0)
	assert.True(photos.TotalCount > 0)
}

func TestChatType(t *testing.T) {
	assert := assert.New(t)
	token := os.Getenv("TOKEN")
	chatid, _ := strconv.Atoi(os.Getenv("CHAT"))
	bot := InitTeleBot(token)
	msg, err := bot.SendMessage(chatid, "Test", nil)
	if err != nil {
		assert.Fail("Bot send message error")
	}
	assert.EqualValues(int(msg.Chat.Id), chatid)
	assert.EqualValues(msg.Chat.Type, "private")
}
