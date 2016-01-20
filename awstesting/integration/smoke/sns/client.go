//Package sns provides gucumber integration tests support.
package sns

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/sns"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@sns", func() {
		World["client"] = sns.New(smoke.Session)
	})
}
