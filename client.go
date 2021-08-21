package sms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var (
	// SNSClient is used to interact with AWS SNS.
	SNSClient *sns.Client
)

// Initialize will initialize the sms package. Both the profile
// and region parameters are optional if authentication can be achieved
// via another method. For example, environment variables or IAM roles.
func Initialize(ctx context.Context, profile string, region string) error {
	var cfg aws.Config
	var err error

	if profile != "" && region != "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile), config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx)
	}
	if err != nil {
		return fmt.Errorf("failed to load default config: %w", err)
	}

	SNSClient = sns.NewFromConfig(cfg)

	return nil
}
