package sqs

import "github.com/bluet-deps/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
