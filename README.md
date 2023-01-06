## A package for sending SMS message via AWS SNS

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/aws-sms?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/aws-sms)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/aws-sms/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/gofor-little/aws-sms/ci.yaml?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/aws-sms)](https://goreportcard.com/report/github.com/gofor-little/aws-sms)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/aws-sms)](https://pkg.go.dev/github.com/gofor-little/aws-sms)

### Introduction
* Send emails via AWS SNS

### Example
```go
package main

import (
	"context"

	sms "github.com/gofor-little/aws-sms"
)

func main() {
	// Initialize the SMS package.
	if err := sms.Initialize("AWS_PROFILE", "AWS_REGION"); err != nil {
		panic(err)
	}

	// Build the SMS data.
	data := sms.Data{
		SenderID: "ExampleApp",
		PhoneNumber: "+61412345678",
		Message: "Example Message",
	}

	// Send the SMS message.
	if _, err := sms.Send(context.Background(), data); err != nil {
		panic(err)
	}
}
```

### Testing
Ensure the following environment variables are set, usually with a .env file.
* ```AWS_PROFILE``` (an AWS CLI profile name)
* ```AWS_REGION``` (a valid AWS region)
* ```TEST_SMS_RECIPIENTS``` (a comma separated list of valid phone numbers including country code)
