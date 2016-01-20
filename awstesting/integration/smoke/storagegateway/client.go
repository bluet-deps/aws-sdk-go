//Package storagegateway provides gucumber integration tests support.
package storagegateway

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/storagegateway"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@storagegateway", func() {
		World["client"] = storagegateway.New(smoke.Session)
	})
}
