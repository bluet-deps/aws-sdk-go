//Package machinelearning provides gucumber integration tests support.
package machinelearning

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/machinelearning"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@machinelearning", func() {
		World["client"] = machinelearning.New(smoke.Session)
	})
}
