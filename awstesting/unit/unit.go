// Package unit performs initialization and validation for unit tests
package unit

import (
	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/credentials"
	"bluet-deps/aws-sdk-go/aws/session"
)

// Session is a shared session for unit tests to use.
var Session = session.New(aws.NewConfig().
	WithCredentials(credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")).
	WithRegion("mock-region"))
