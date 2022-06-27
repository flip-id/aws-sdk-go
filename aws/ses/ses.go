package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type SESServiceInterface interface {
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
}

type sesService struct {
	client *ses.Client
}

func NewSES(client *ses.Client) SESServiceInterface {
	return &sesService{client: client}
}

func (s *sesService) SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	return s.client.SendEmail(ctx, params, optFns...)
}
