//Package cloudwatchlogs provides gucumber integration tests support.
package cloudwatchlogs

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/cloudwatchlogs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudwatchlogs", func() {
		World["client"] = cloudwatchlogs.New(smoke.Session)
	})
}
