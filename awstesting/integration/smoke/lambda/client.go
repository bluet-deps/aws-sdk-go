//Package lambda provides gucumber integration tests support.
package lambda

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/lambda"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@lambda", func() {
		World["client"] = lambda.New(smoke.Session)
	})
}
