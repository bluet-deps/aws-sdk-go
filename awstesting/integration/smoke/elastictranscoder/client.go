//Package elastictranscoder provides gucumber integration tests support.
package elastictranscoder

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/elastictranscoder"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@elastictranscoder", func() {
		World["client"] = elastictranscoder.New(smoke.Session)
	})
}
