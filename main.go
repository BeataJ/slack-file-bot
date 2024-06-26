package main

import (
	"os"

	"github.com/slack-go/slack"
)

func main() {

	tokenbot := os.Getenv("BOT_TOKEN")
	channelid := os.Getenv("CHANNEL_TOKEN")
	os.Setenv("SLACK_BOT_TOKEN", tokenbot)
	os.Setenv("CHANNEL_ID", channelid)

	channelArr := []string{os.Getenv("CHANNEL_ID")}

	api := slack.New(os.Getenv("SLACK_APP_TOKEN"))
}
