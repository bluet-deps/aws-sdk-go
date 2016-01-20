//Package support provides gucumber integration tests support.
package support

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/support"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@support", func() {
		World["client"] = support.New(smoke.Session)
	})
}
