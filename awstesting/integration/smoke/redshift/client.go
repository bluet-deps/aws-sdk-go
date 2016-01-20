//Package redshift provides gucumber integration tests support.
package redshift

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/redshift"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@redshift", func() {
		World["client"] = redshift.New(smoke.Session)
	})
}
