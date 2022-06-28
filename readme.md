# AWS SDK GO

![This is an image](./assets/aws-services.webp)

## Introduction

AWS SDK GO is a wrapper library of official library [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2). With this library will help you to make it easier for you to use it!

## How to use

Before you using this library, please read the explanation "How to set the credentials" in this article Before you running this service, please set up your credentials with following this instruction [Get your AWS access keys](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys) & [Where are configuration settings stored?](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

### SES Example

```go
package main

import (
	"context"
	"github.com/flip-id/aws-sdk-go/client"
	"github.com/flip-id/aws-sdk-go/services/ses"
	"fmt"
)

func main() {

	service, err := ses.New(context.TODO(), client.APSoutheast1)
	if err != nil {
		panic(err)
	}

	messageID, err := service.SendEmail(context.TODO(), ses.RequestSendEmail{
		To:      []string{"your_email@gmail.com"},
		From:    "your_email@gmail.com",
		Subject: "Testing",
		Body:    "This is body of email",
		Type:    ses.TEXTTypeEmail,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(messageID)
}
```
