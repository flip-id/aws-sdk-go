package ses

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/go-playground/validator"

	awsSes "github.com/flip-id/aws-sdk-go/aws/ses"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
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
		arguments RequestSendEmail
		mockings  mocking
		wantError bool
	}{
		{
			name: "success send email with text type",
			arguments: RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "This is Body of Email",
				Type:    TEXTTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &sesv2.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    1,
				},
			},
			wantError: false,
		},
		{
			name: "success send email with html type",
			arguments: RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &sesv2.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    1,
				},
			},
			wantError: false,
		},
		{
			name: "error send email because invalid payload",
			arguments: RequestSendEmail{
				To:      []string{},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &sesv2.SendEmailOutput{MessageId: &messageID},
					"error":    nil,
					"times":    0,
				},
			},
			wantError: true,
		},
		{
			name: "error send email because error when send email with ses service",
			arguments: RequestSendEmail{
				To:      []string{"your_email@gmail.com"},
				From:    "your_email@gmail.com",
				Subject: "Testing Email",
				Body:    "<h1>Hello world</h1>",
				Type:    HTMLTypeEmail,
			},
			mockings: mocking{
				sendEmail: map[string]interface{}{
					"response": &sesv2.SendEmailOutput{MessageId: &messageID},
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
			mockingSESService.EXPECT().SendEmailV2(gomock.Any(), gomock.Any()).Return(
				tc.mockings.sendEmail["response"], tc.mockings.sendEmail["error"]).Times(tc.mockings.sendEmail["times"].(int))

			sesService := Service{
				SESService: mockingSESService,
				validate:   validator.New(),
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

func TestService_SendRawEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		SESService awsSes.SESServiceInterface
		validate   *validator.Validate
	}
	type args struct {
		ctx     context.Context
		request RequestSendRawEmail
	}
	tests := []struct {
		name    string
		fields  func() fields
		args    func() args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "no destination",
			fields: func() fields {
				return fields{
					SESService: awsSes.NewMockSESServiceInterface(ctrl),
					validate:   validator.New(),
				}
			},
			args: func() args {
				return args{
					ctx:     context.TODO(),
					request: RequestSendRawEmail{},
				}
			},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name: "error sending text email",
			fields: func() fields {
				mockSESService := awsSes.NewMockSESServiceInterface(ctrl)
				mockSESService.EXPECT().SendEmailV2(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("error sending text email"))
				return fields{
					SESService: mockSESService,
					validate:   validator.New(),
				}
			},
			args: func() args {
				return args{
					ctx: context.TODO(),
					request: RequestSendRawEmail{
						RequestSendEmail: RequestSendEmail{
							To:      []string{"hello@test.id"},
							From:    "no-reply@test.id",
							Subject: "Test",
							Body:    "Test",
							Type:    TEXTTypeEmail,
						},
					},
				}
			},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name: "error sending html email",
			fields: func() fields {
				mockSESService := awsSes.NewMockSESServiceInterface(ctrl)
				mockSESService.EXPECT().SendEmailV2(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("error sending html email"))
				return fields{
					SESService: mockSESService,
					validate:   validator.New(),
				}
			},
			args: func() args {
				return args{
					ctx: context.TODO(),
					request: RequestSendRawEmail{
						RequestSendEmail: RequestSendEmail{
							To:      []string{"hello@test.id"},
							From:    "no-reply@test.id",
							Subject: "Test",
							Body:    "Test",
							Type:    HTMLTypeEmail,
						},
					},
				}
			},
			want:    "",
			wantErr: assert.Error,
		},
		{
			name: "success sending text email with attachments",
			fields: func() fields {
				mockSESService := awsSes.NewMockSESServiceInterface(ctrl)
				mockSESService.EXPECT().SendEmailV2(gomock.Any(), gomock.Any()).
					Return(&sesv2.SendEmailOutput{
						MessageId: aws.String("5123"),
					}, nil)
				return fields{
					SESService: mockSESService,
					validate:   validator.New(),
				}
			},
			args: func() args {
				filePathTest := "./attachment.txt"
				if _, err := os.Stat(filePathTest); errors.Is(err, os.ErrNotExist) {
					filePathTest = "./services/ses/attachment.txt"
					_, err := os.Stat(filePathTest)
					if err != nil {
						t.Errorf("Error checking file attachment.txt for testing: %v", err)
					}
				}
				buff := strings.NewReader(`hello world!`)
				return args{
					ctx: context.TODO(),
					request: RequestSendRawEmail{
						RequestSendEmail: RequestSendEmail{
							To:      []string{"hello@test.id"},
							Cc:      []string{"hello-cc@test.id"},
							Bcc:     []string{"hello-bcc@test.id"},
							From:    "no-reply@test.id",
							Subject: "Test",
							Body:    "Test",
							Type:    TEXTTypeEmail,
						},
						AttachmentPaths: []string{filePathTest},
						AttachmentReaders: []AttachmentReader{
							{
								Name:   "attachment-world.txt",
								Reader: buff,
							},
						},
					},
				}
			},
			want:    "5123",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.fields()
			s := &Service{
				SESService: fields.SESService,
				validate:   fields.validate,
			}
			args := tt.args()
			got, err := s.SendRawEmail(args.ctx, args.request)
			if !tt.wantErr(t, err, fmt.Sprintf("SendRawEmail(%v, %v)", args.ctx, args.request)) {
				return
			}
			assert.Equalf(t, tt.want, got, "SendRawEmail(%v, %v)", args.ctx, args.request)
		})
	}
}
