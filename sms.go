package sms

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gofor-little/xerror"
)

// Send builds a SMS message and sends it.
func Send(ctx context.Context, data Data) (string, error) {
	if SNSClient == nil {
		return "", xerror.Newf("SNSClient is nil")
	}

	input := &sns.PublishInput{
		Message: aws.String(data.Message),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.SenderID),
			},
		},
		PhoneNumber: aws.String(data.PhoneNumber),
	}

	if err := input.Validate(); err != nil {
		return "", xerror.New("failed to validate sns.PublishInput", err)
	}

	output, err := SNSClient.PublishWithContext(ctx, input)
	if err != nil {
		return "", xerror.New("failed to publish SMS message", err)
	}

	return *output.MessageId, nil
}
