//Package autoscaling provides gucumber integration tests support.
package autoscaling

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/autoscaling"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@autoscaling", func() {
		World["client"] = autoscaling.New(smoke.Session)
	})
}
