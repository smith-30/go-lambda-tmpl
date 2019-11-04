package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/smith-30/go-lambda-tmpl/service"
)

type EventSlice struct {
	Events      []Event `json:"events"`
	Destination string  `json:"source"`
}

type Event struct {
	Type       string     `json:"type"`
	ReplyToken string     `json:"replyToken"`
	Message    MessageObj `json:"message"`
	Source     Src        `json:"source"`
}

type Src struct {
	Type    string `json:"type"`
	UserID  string `json:"userId"`
	GroupID string `json:"groupId"`
	RoomID  string `json:"roomId"`
}

type MessageObj struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) error {
	fmt.Printf("%v\n", req.Body)

	e := &EventSlice{}
	json.Unmarshal([]byte(req.Body), e)
	fmt.Printf("%v\n", e)
	at, err := service.GetAccessToken()
	if err != nil {
		return err
	}
	bot, err := linebot.New(service.GetChannelSecret(), at.AccessToken)
	if err != nil {
		return err
	}

	if _, err := bot.ReplyMessage(e.Events[0].ReplyToken, linebot.NewTextMessage("hi")).Do(); err != nil {
		return err
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
