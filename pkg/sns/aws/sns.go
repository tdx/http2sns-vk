package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"github.com/tdx/http2sns-vk/pkg/config"
)

//
type Publisher struct {
	sns *sns.Client
}

//
func NewPublisher(cfg *config.Config) (*Publisher, error) {

	var (
		err    error
		awsCfg aws.Config

		snsApiEndpoint = cfg.SnsApiEndpoint
	)

	if snsApiEndpoint == "" {
		// default SNS endpoint
		awsCfg, err = awsConfig.LoadDefaultConfig(context.Background(),
			awsConfig.WithRegion(cfg.SnsRegion),
		)
	} else {
		snsResolver := aws.EndpointResolverWithOptionsFunc(
			func(service, reg string, opts ...interface{}) (aws.Endpoint, error) {
				if service == sns.ServiceID && reg == cfg.SnsRegion {
					return aws.Endpoint{
						PartitionID:   "aws",
						URL:           snsApiEndpoint,
						SigningRegion: reg,
					}, nil
				}
				return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
			})

		awsCfg, err = awsConfig.LoadDefaultConfig(
			context.Background(),
			awsConfig.WithRegion(cfg.SnsRegion),
			awsConfig.WithEndpointResolverWithOptions(snsResolver),
		)
	}

	if err != nil {
		return nil, fmt.Errorf("aws:NewPublisher() failed: %w", err)
	}

	p := &Publisher{
		sns: sns.NewFromConfig(awsCfg),
	}

	return p, nil
}

//
func (p *Publisher) Publish(topic, message string) error {

	params := &sns.PublishInput{
		TopicArn: aws.String(topic),
		Message:  aws.String(message),
	}

	_, err := p.sns.Publish(context.Background(), params)

	return err
}
