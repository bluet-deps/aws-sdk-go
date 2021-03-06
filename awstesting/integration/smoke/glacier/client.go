//Package glacier provides gucumber integration tests support.
package glacier

import (
	"github.com/bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"github.com/bluet-deps/aws-sdk-go/service/glacier"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@glacier", func() {
		World["client"] = glacier.New(smoke.Session)
	})
}
