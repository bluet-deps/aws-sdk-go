//Package ssm provides gucumber integration tests support.
package ssm

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/ssm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ssm", func() {
		World["client"] = ssm.New(smoke.Session)
	})
}
