//Package kms provides gucumber integration tests support.
package kms

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/kms"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@kms", func() {
		World["client"] = kms.New(smoke.Session)
	})
}
