//Package redshift provides gucumber integration tests support.
package redshift

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/redshift"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@redshift", func() {
		World["client"] = redshift.New(smoke.Session)
	})
}
