//Package iam provides gucumber integration tests support.
package iam

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/iam"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@iam", func() {
		World["client"] = iam.New(smoke.Session)
	})
}
