//Package ses provides gucumber integration tests support.
package ses

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/ses"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ses", func() {
		World["client"] = ses.New(smoke.Session)
	})
}
