package sms_test

import (
	"context"
	"fmt"
	"testing"

	mock "github.com/gofor-little/aws-sdk-mock"
	"github.com/stretchr/testify/require"

	sms "github.com/gofor-little/aws-sms"
)

func TestSend(t *testing.T) {
	sms.SNSClient = &mock.SNSClient{}

	testCases := []struct {
		name string
		data sms.Data
		want error
	}{
		{"TestSend_Unit", sms.Data{SenderID: "TestSMS", PhoneNumber: "+61412345678", Message: "Test SMS"}, nil},
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("%s_%d", tc.name, i)

		t.Run(name, func(t *testing.T) {
			_, err := sms.Send(context.Background(), tc.data)
			require.Equal(t, tc.want, err)
		})
	}
}
