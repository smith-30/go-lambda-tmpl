package service

import "os"

var (
	channelSecret string
	channelID     string
)

func init() {
	channelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	channelID = os.Getenv("LINE_CHANNEL_ID")
}

func GetChannelSecret() string {
	return channelSecret
}

func GetChannelID() string {
	return channelID
}
