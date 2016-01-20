//Package cloudfront provides gucumber integration tests support.
package cloudfront

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/cloudfront"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudfront", func() {
		World["client"] = cloudfront.New(smoke.Session)
	})
}
