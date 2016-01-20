// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package devicefarm

import (
	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/client"
	"bluet-deps/aws-sdk-go/aws/client/metadata"
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/private/protocol/jsonrpc"
	"bluet-deps/aws-sdk-go/private/signer/v4"
)

// AWS Device Farm is a service that enables mobile app developers to test Android,
// iOS, and Fire OS apps on physical phones, tablets, and other devices in the
// cloud.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type DeviceFarm struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "devicefarm"

// New creates a new instance of the DeviceFarm client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a DeviceFarm client from just a session.
//     svc := devicefarm.New(mySession)
//
//     // Create a DeviceFarm client with additional configuration
//     svc := devicefarm.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DeviceFarm {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *DeviceFarm {
	svc := &DeviceFarm{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-06-23",
				JSONVersion:   "1.1",
				TargetPrefix:  "DeviceFarm_20150623",
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

// newRequest creates a new request for a DeviceFarm operation and runs any
// custom request initialization.
func (c *DeviceFarm) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
