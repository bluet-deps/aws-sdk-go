//Package dynamodb provides gucumber integration tests support.
package dynamodb

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/dynamodb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@dynamodb", func() {
		World["client"] = dynamodb.New(smoke.Session)
	})
}
