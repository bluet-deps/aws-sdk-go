//Package cloudhsm provides gucumber integration tests support.
package cloudhsm

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/cloudhsm"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cloudhsm", func() {
		World["client"] = cloudhsm.New(smoke.Session)
	})
}
