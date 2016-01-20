//Package cloudwatch provides gucumber integration tests support.
package cloudwatch

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/cloudwatch"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudwatch", func() {
		World["client"] = cloudwatch.New(smoke.Session)
	})
}
