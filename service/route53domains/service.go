// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53domains

import (
	"github.com/bluet-deps/aws-sdk-go/aws"
	"github.com/bluet-deps/aws-sdk-go/aws/client"
	"github.com/bluet-deps/aws-sdk-go/aws/client/metadata"
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/bluet-deps/aws-sdk-go/private/signer/v4"
)

// Route53Domains is a client for Amazon Route 53 Domains.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type Route53Domains struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "route53domains"

// New creates a new instance of the Route53Domains client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Route53Domains client from just a session.
//     svc := route53domains.New(mySession)
//
//     // Create a Route53Domains client with additional configuration
//     svc := route53domains.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Route53Domains {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *Route53Domains {
	svc := &Route53Domains{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-05-15",
				JSONVersion:   "1.1",
				TargetPrefix:  "Route53Domains_v20140515",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(jsonrpc.Build)
	svc.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Route53Domains operation and runs any
// custom request initialization.
func (c *Route53Domains) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
