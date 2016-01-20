//Package simpledb provides gucumber integration tests support.
package simpledb

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/simpledb"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@simpledb", func() {
		World["client"] = simpledb.New(smoke.Session)
	})
}
