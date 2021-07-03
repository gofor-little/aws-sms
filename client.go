package sms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/gofor-little/xerror"
)

var (
	// SNSClient is used to interact with AWS SNS.
	SNSClient snsiface.SNSAPI
)

// Initialize will initialize the sms package. Both the profile
// and region parameters are optional if authentication can be achieved
// via another method. For example, environment variables or IAM roles.
func Initialize(profile string, region string) error {
	var sess *session.Session
	var err error

	if profile != "" && region != "" {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String(region),
				CredentialsChainVerboseErrors: aws.Bool(true),
			},
			Profile: profile,
		})
	} else {
		sess, err = session.NewSession()
	}
	if err != nil {
		return xerror.Wrap("failed to create session.Session", err)
	}

	SNSClient = sns.New(sess)

	return nil
}
