package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Region based on this thread
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
type Region string

const (
	APSoutheast1 = Region("ap-southeast-1")
	APSoutheast2 = Region("ap-southeast-2")
	APSoutheast3 = Region("ap-southeast-3")
)

type Config struct {
	Config aws.Config
}

func New(ctx context.Context, configs ...func(*config.LoadOptions) error) (Config, error) {

	awsConfig, err := config.LoadDefaultConfig(ctx, configs...)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Config: awsConfig,
	}, nil
}
