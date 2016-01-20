//Package efs provides gucumber integration tests support.
package efs

import (
	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/efs"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@efs", func() {
		// FIXME remove custom region
		World["client"] = efs.New(smoke.Session,
			aws.NewConfig().WithRegion("us-west-2"))
	})
}
