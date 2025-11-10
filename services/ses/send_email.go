package ses

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	sesv2types "github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/wneessen/go-mail"
)

const (
	HeaderReturnPath = "Return-Path"
)

func (s *Service) SendEmail(ctx context.Context, request RequestSendEmail) (string, error) {
	if err := s.validate.Struct(request); err != nil {
		return "", err
	}

	bodyEmail := &sesv2types.Body{}
	if request.Type == HTMLTypeEmail {
		bodyEmail.Html = &sesv2types.Content{
			Charset: aws.String(CHARSET),
			Data:    aws.String(request.Body),
		}
	} else if request.Type == TEXTTypeEmail {
		bodyEmail.Text = &sesv2types.Content{
			Charset: aws.String(CHARSET),
			Data:    aws.String(request.Body),
		}
	}

	// Use SES v2 API for better support and consistency
	response, err := s.SESService.SendEmailV2(ctx, &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(request.From),
		Destination: &sesv2types.Destination{
			ToAddresses:  request.To,
			CcAddresses:  request.Cc,
			BccAddresses: request.Bcc,
		},
		Content: &sesv2types.EmailContent{
			Simple: &sesv2types.Message{
				Subject: &sesv2types.Content{
					Charset: aws.String(CHARSET),
					Data:    aws.String(request.Subject),
				},
				Body: bodyEmail,
			},
		},
	})
	if err != nil {
		return "", err
	}

	return *response.MessageId, nil
}

func (s *Service) SendRawEmail(ctx context.Context, request RequestSendRawEmail) (messageID string, err error) {
	if err := s.validate.Struct(request); err != nil {
		return "", err
	}

	m := mail.NewMsg()
	if request.Type == HTMLTypeEmail {
		m.SetBodyString(mail.TypeTextHTML, request.Body)
	} else if request.Type == TEXTTypeEmail {
		m.SetBodyString(mail.TypeTextPlain, request.Body)
	}
	m.FromFormat(request.FromName, request.From)
	m.SetUserAgent(s.UserAgent)
	returnPath := request.From
	if request.ReturnPath != "" {
		returnPath = request.ReturnPath
	}
	m.SetHeader(HeaderReturnPath, returnPath)

	err = m.To(request.To...)
	if err != nil {
		return
	}

	err = m.Cc(request.Cc...)
	if err != nil {
		return
	}

	err = m.Bcc(request.Bcc...)
	if err != nil {
		return
	}

	for _, a := range request.AttachmentPaths {
		m.AttachFile(a)
	}

	for _, a := range request.AttachmentReaders {
		m.AttachReader(a.Name, a.Reader)
	}

	m.Subject(request.Subject)
	var buff = new(bytes.Buffer)

	// Write the email message to the buffer
	_, err = m.WriteTo(buff)
	if err != nil {
		return "", err
	}

	// Use SES v2 API for 40MB attachment support
	response, err := s.SESService.SendEmailV2(ctx, &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(request.From),
		Content: &sesv2types.EmailContent{
			Raw: &sesv2types.RawMessage{
				Data: buff.Bytes(),
			},
		},
		Destination: &sesv2types.Destination{
			ToAddresses:  request.To,
			CcAddresses:  request.Cc,
			BccAddresses: request.Bcc,
		},
	})
	if err != nil {
		return
	}

	return *response.MessageId, nil
}
