//Package codepipeline provides gucumber integration tests support.
package codepipeline

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/codepipeline"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@codepipeline", func() {
		World["client"] = codepipeline.New(smoke.Session)
	})
}
