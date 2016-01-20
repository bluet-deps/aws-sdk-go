//Package directconnect provides gucumber integration tests support.
package directconnect

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/directconnect"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@directconnect", func() {
		World["client"] = directconnect.New(smoke.Session)
	})
}
