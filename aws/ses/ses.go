package ses

//go:generate mockgen -source=ses.go -destination=ses_mock.go -package=ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type SESServiceInterface interface {
	SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error)
	SendRawEmail(ctx context.Context, params *ses.SendRawEmailInput, optFns ...func(*ses.Options)) (*ses.SendRawEmailOutput, error)
	SendEmailV2(ctx context.Context, params *sesv2.SendEmailInput, optFns ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error)
}

type sesService struct {
	client   *ses.Client
	clientV2 *sesv2.Client
}

func NewSES(client *ses.Client) SESServiceInterface {
	return &sesService{client: client}
}

func NewSESV2(clientV2 *sesv2.Client) SESServiceInterface {
	return &sesService{clientV2: clientV2}
}

func NewSESWithBoth(client *ses.Client, clientV2 *sesv2.Client) SESServiceInterface {
	return &sesService{client: client, clientV2: clientV2}
}

func (s *sesService) SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	return s.client.SendEmail(ctx, params, optFns...)
}

// SendRawEmail sends email using raw input data to SES v1 (deprecated, use SendEmailV2 for 40MB support).
func (s *sesService) SendRawEmail(ctx context.Context, params *ses.SendRawEmailInput, optFns ...func(*ses.Options)) (*ses.SendRawEmailOutput, error) {
	return s.client.SendRawEmail(ctx, params, optFns...)
}

// SendEmailV2 sends email using SES v2 API with support for 40MB attachments.
func (s *sesService) SendEmailV2(ctx context.Context, params *sesv2.SendEmailInput, optFns ...func(*sesv2.Options)) (*sesv2.SendEmailOutput, error) {
	return s.clientV2.SendEmail(ctx, params, optFns...)
}
