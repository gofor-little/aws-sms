package sms_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	sms "github.com/gofor-little/aws-sms"
)

func TestSend(t *testing.T) {
	recipients := setup(t)

	var testCases []struct {
		name string
		data sms.Data
		want error
	}

	for _, r := range recipients {
		testCases = append(testCases, struct {
			name string
			data sms.Data
			want error
		}{"TestSend_Integration", sms.Data{SenderID: "TestSMS", PhoneNumber: r, Message: "Test SMS"}, nil})
	}

	for i, tc := range testCases {
		name := fmt.Sprintf("%s_%d", tc.name, i)

		t.Run(name, func(t *testing.T) {
			_, err := sms.Send(context.Background(), tc.data)
			if err != nil {
				fmt.Println(err)
			}
			require.Equal(t, tc.want, err, err.Error())
		})
	}
}
