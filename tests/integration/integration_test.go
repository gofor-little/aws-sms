package sms_test

import (
	"strings"
	"testing"

	"github.com/gofor-little/env"
	"github.com/stretchr/testify/require"

	sms "github.com/gofor-little/aws-sms"
)

func setup(t *testing.T) []string {
	if err := env.Load("../../.env"); err != nil {
		t.Logf("failed to load .env file, ignore if running in CI/CD: %v", err)
	}

	require.NoError(t, sms.Initialize(env.Get("AWS_PROFILE", ""), env.Get("AWS_REGION", "")))

	recipients, err := env.MustGet("TEST_SMS_RECIPIENTS")
	require.NoError(t, err)

	return strings.Split(recipients, ",")
}
