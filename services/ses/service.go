package ses

import (
	"context"
	"github.com/go-playground/validator"

	awsses "github.com/flip-id/aws-sdk-go/aws/ses"
	"github.com/flip-id/aws-sdk-go/client"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

type ServiceInterface interface {
	SendEmail(ctx context.Context, request RequestSendEmail) (string, error)
	SendRawEmail(ctx context.Context, request RequestSendRawEmail) (string, error)
}

type Service struct {
	SESService awsses.SESServiceInterface
	validate   *validator.Validate
}

func New(ctx context.Context, region client.Region, options ...func(*ses.Options)) (ServiceInterface, error) {

	clientConfig, err := client.New(ctx, config.WithRegion(string(region)))
	if err != nil {
		return nil, err
	}

	sesClient := ses.NewFromConfig(clientConfig.Config, options...)
	return &Service{SESService: awsses.NewSES(sesClient), validate: validator.New()}, nil
}
