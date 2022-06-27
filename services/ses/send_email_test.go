package ses_test

import (
	"context"
	"errors"
	awsSes "flip/aws-sdk-go/aws/ses"
	servicesSes "flip/aws-sdk-go/services/ses"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {

	ctrl := gomock.NewController(t)

	messageID := uuid.New().String()

	type mocking struct {
		sendEmail map[string]interface{}
	}

	testCases := []struct {
		name      string
		arguments servicesSes.RequestSendEmail
		mockings  mocking
		wantError bool
	}{
		{
			name: "success send email with text type",
			arguments: servicesSes.RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "This is Body of Email",
				Type:    servicesSes.TEXTTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &ses.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    1,
				},
			},
			wantError: false,
		},
		{
			name: "success send email with html type",
			arguments: servicesSes.RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    servicesSes.HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &ses.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    1,
				},
			},
			wantError: false,
		},
		{
			name: "error send email because invalid payload",
			arguments: servicesSes.RequestSendEmail{
				To:      []string{},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    servicesSes.HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &ses.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    0,
				},
			},
			wantError: true,
		},
		{
			name: "error send email because error when send email with ses service",
			arguments: servicesSes.RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    servicesSes.HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &ses.SendEmailOutput{MessageId: &messageID},
					"error":    errors.New("something wrong"),
					"times":    1,
				},
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			mockingSESService := awsSes.NewMockSESServiceInterface(ctrl)
			mockingSESService.EXPECT().SendEmail(gomock.Any(), gomock.Any()).Return(
				tc.mockings.sendEmail["response"], tc.mockings.sendEmail["error"]).Times(tc.mockings.sendEmail["times"].(int))

			sesService := servicesSes.Service{
				SESService: mockingSESService,
			}

			result, err := sesService.SendEmail(context.TODO(), tc.arguments)

			if tc.wantError {
				assert.Equal(t, "", result)
				assert.Error(t, err)
			} else {
				assert.Equal(t, messageID, result)
				assert.NoError(t, err)
			}
		})
	}

}
