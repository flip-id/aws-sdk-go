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
	UserAgent  string
	SESService awsses.SESServiceInterface
	validate   *validator.Validate
}

func New(ctx context.Context, serviceOption *ServiceOption, options ...func(*ses.Options)) (ServiceInterface, error) {
	var (
		region    string
		userAgent string
	)

	if serviceOption.Region != "" {
		region = serviceOption.Region
	}
	if serviceOption.ServiceCode != "" {
		userAgent = serviceOption.ServiceCode
	}

	clientConfig, err := client.New(ctx, config.WithRegion(string(region)), config.WithHTTPClient(serviceOption.Client))
	if err != nil {
		return nil, err
	}

	sesClient := ses.NewFromConfig(clientConfig.Config, options...)
	return &Service{
		UserAgent:  userAgent,
		SESService: awsses.NewSES(sesClient),
		validate:   validator.New(),
	}, nil
}
