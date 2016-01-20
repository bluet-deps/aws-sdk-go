//Package sts provides gucumber integration tests support.
package sts

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/sts"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@sts", func() {
		World["client"] = sts.New(smoke.Session)
	})
}
