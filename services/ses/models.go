package ses

import "io"

type TypeEmail string

const (
	CHARSET = "UTF-8"

	HTMLTypeEmail  = TypeEmail("html")
	TEXTTypeEmail  = TypeEmail("text")
	MaxMessageSize = 10485760
)

type RequestSendEmail struct {
	To         []string `validate:"min=1,dive,email"`
	Cc         []string `validate:"dive,email"`
	Bcc        []string `validate:"dive,email"`
	From       string   `validate:"required"`
	FromName   string
	Subject    string `validate:"required"`
	Body       string `validate:"required"`
	Type       TypeEmail
	ReturnPath string
}

// RequestSendRawEmail is a request to send raw email.
type RequestSendRawEmail struct {
	RequestSendEmail
	AttachmentPaths   []string
	AttachmentReaders []AttachmentReader
}

// AttachmentReader is an attachment with the io.Reader and name of file.
type AttachmentReader struct {
	Name   string
	Reader io.Reader
}

type ServiceOption struct {
	Region      string
	ServiceCode string
}
