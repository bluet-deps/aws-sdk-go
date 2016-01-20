//Package elasticbeanstalk provides gucumber integration tests support.
package elasticbeanstalk

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/elasticbeanstalk"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elasticbeanstalk", func() {
		World["client"] = elasticbeanstalk.New(smoke.Session)
	})
}
