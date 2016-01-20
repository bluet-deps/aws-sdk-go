//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/dynamodbstreams"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodbstreams", func() {
		World["client"] = dynamodbstreams.New(smoke.Session)
	})
}
