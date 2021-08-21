package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/gofor-little/xerror"
)

// Send builds an SMS message and sends it.
func Send(ctx context.Context, data Data) (string, error) {
	if SNSClient == nil {
		return "", xerror.New("SNSClient is nil")
	}

	input := &sns.PublishInput{
		Message: aws.String(data.Message),
		MessageAttributes: map[string]types.MessageAttributeValue{
			"AWS.SNS.SMS.SenderID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(data.SenderID),
			},
		},
		PhoneNumber: aws.String(data.PhoneNumber),
	}

	output, err := SNSClient.Publish(ctx, input)
	if err != nil {
		return "", xerror.Wrap("failed to publish SMS message", err)
	}

	return *output.MessageId, nil
}
