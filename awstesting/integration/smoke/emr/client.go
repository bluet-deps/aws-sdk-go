//Package emr provides gucumber integration tests support.
package emr

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/emr"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@emr", func() {
		World["client"] = emr.New(smoke.Session)
	})
}
