//Package kinesis provides gucumber integration tests support.
package kinesis

import (
	"bluet-deps/aws-sdk-go/awstesting/integration/smoke"
	"bluet-deps/aws-sdk-go/service/kinesis"
	. "github.com/lsegal/gucumber"
)

func init() {
	Before("@kinesis", func() {
		World["client"] = kinesis.New(smoke.Session)
	})
}
