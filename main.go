package main

import (
	"context"
	"fmt"

	"github.com/flip-id/aws-sdk-go/client"
	"github.com/flip-id/aws-sdk-go/services/ses"
)

func main() {

	service, err := ses.New(context.TODO(), client.APSoutheast1)
	if err != nil {
		panic(err)
	}

	messageID, err := service.SendEmail(context.TODO(), ses.RequestSendEmail{
		To:      []string{"muhammadrivaldy16@gmail.com"},
		From:    "muhammadrivaldy16@gmail.com",
		Subject: "Testing",
		Body:    "This is body of email",
		Type:    ses.TEXTTypeEmail,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(messageID)
}
