// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package autoscaling

import (
	"github.com/bluet-deps/aws-sdk-go/aws"
	"github.com/bluet-deps/aws-sdk-go/aws/client"
	"github.com/bluet-deps/aws-sdk-go/aws/client/metadata"
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/private/protocol/query"
	"github.com/bluet-deps/aws-sdk-go/private/signer/v4"
)

// Auto Scaling is designed to automatically launch or terminate EC2 instances
// based on user-defined policies, schedules, and health checks. Use this service
// in conjunction with the Amazon CloudWatch and Elastic Load Balancing services.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type AutoScaling struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "autoscaling"

// New creates a new instance of the AutoScaling client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a AutoScaling client from just a session.
//     svc := autoscaling.New(mySession)
//
//     // Create a AutoScaling client with additional configuration
//     svc := autoscaling.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *AutoScaling {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *AutoScaling {
	svc := &AutoScaling{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2011-01-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(query.Build)
	svc.Handlers.Unmarshal.PushBack(query.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a AutoScaling operation and runs any
// custom request initialization.
func (c *AutoScaling) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
