package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/go-playground/validator"
)

func (s *Service) SendEmail(ctx context.Context, request RequestSendEmail) (string, error) {

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return "", err
	}

	bodyEmail := &types.Body{}
	if request.Type == HTMLTypeEmail {
		bodyEmail.Html = &types.Content{
			Charset: aws.String(CHARSET),
			Data:    aws.String(request.Body),
		}
	} else if request.Type == TEXTTypeEmail {
		bodyEmail.Text = &types.Content{
			Charset: aws.String(CHARSET),
			Data:    aws.String(request.Body),
		}
	}

	response, err := s.SESService.SendEmail(ctx, &ses.SendEmailInput{
		Source: aws.String(request.From),
		Destination: &types.Destination{
			ToAddresses:  request.To,
			CcAddresses:  request.Cc,
			BccAddresses: request.Bcc,
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String(CHARSET),
				Data:    aws.String(request.Subject),
			},
			Body: bodyEmail,
		},
	})
	if err != nil {
		return "", err
	}

	return *response.MessageId, nil
}
