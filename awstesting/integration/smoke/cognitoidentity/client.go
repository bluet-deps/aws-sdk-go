//Package cognitoidentity provides gucumber integration tests support.
package cognitoidentity

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/cognitoidentity"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@cognitoidentity", func() {
		World["client"] = cognitoidentity.New(smoke.Session)
	})
}
