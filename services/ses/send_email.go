package ses

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/wneessen/go-mail"
)

func (s *Service) SendEmail(ctx context.Context, request RequestSendEmail) (string, error) {
	if err := s.validate.Struct(request); err != nil {
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

func (s *Service) SendRawEmail(ctx context.Context, request RequestSendRawEmail) (string, error) {
	if err := s.validate.Struct(request); err != nil {
		return "", err
	}

	m := mail.NewMsg()
	if request.Type == HTMLTypeEmail {
		m.SetBodyString(mail.TypeTextHTML, request.Body)
	} else if request.Type == TEXTTypeEmail {
		m.SetBodyString(mail.TypeTextPlain, request.Body)
	}

	err := m.To(request.To...)
	if err != nil {
		return "", err
	}

	err = m.Cc(request.Cc...)
	if err != nil {
		return "", err
	}

	err = m.Bcc(request.Bcc...)
	if err != nil {
		return "", err
	}

	for _, a := range request.AttachmentPaths {
		m.AttachFile(a)
	}

	for _, a := range request.AttachmentReaders {
		m.AttachReader(a.Name, a.Reader)
	}

	m.Subject(request.Subject)
	var buff = new(bytes.Buffer)
	_, err = m.WriteTo(buff)
	if err != nil {
		return "", err
	}

	response, err := s.SESService.SendRawEmail(ctx, &ses.SendRawEmailInput{
		RawMessage: &types.RawMessage{
			Data: buff.Bytes(),
		},
		Destinations: append(
			request.To,
			append(request.Cc, request.Bcc...)...,
		),
		Source: aws.String(request.From),
	})
	if err != nil {
		return "", err
	}

	return *response.MessageId, nil
}
