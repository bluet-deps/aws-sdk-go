//Package rds provides gucumber integration tests support.
package rds

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/rds"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@rds", func() {
		World["client"] = rds.New(smoke.Session)
	})
}
