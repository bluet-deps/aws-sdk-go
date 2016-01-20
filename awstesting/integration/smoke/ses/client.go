//Package ses provides gucumber integration tests support.
package ses

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/ses"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@ses", func() {
		World["client"] = ses.New(smoke.Session)
	})
}
