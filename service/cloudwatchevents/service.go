// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchevents

import (
	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/client"
	"bluet-deps/aws-sdk-go/aws/client/metadata"
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/private/protocol/jsonrpc"
	"bluet-deps/aws-sdk-go/private/signer/v4"
)

// Amazon CloudWatch Events helps you to respond to state changes in your AWS
// resources. When your resources change state they automatically send events
// into an event stream. You can create rules that match selected events in
// the stream and route them to targets to take action. You can also use rules
// to take action on a pre-determined schedule. For example, you can configure
// rules to:
//
//  Automatically invoke an AWS Lambda function to update DNS entries when
// an event notifies you that Amazon EC2 instance enters the running state.
// Direct specific API records from CloudTrail to an Amazon Kinesis stream for
// detailed analysis of potential security or availability risks. Periodically
// invoke a built-in target to create a snapshot of an Amazon EBS volume.
// For more information about Amazon CloudWatch Events features, see the Amazon
// CloudWatch Developer Guide (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CloudWatchEvents struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "events"

// New creates a new instance of the CloudWatchEvents client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudWatchEvents client from just a session.
//     svc := cloudwatchevents.New(mySession)
//
//     // Create a CloudWatchEvents client with additional configuration
//     svc := cloudwatchevents.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CloudWatchEvents {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CloudWatchEvents {
	svc := &CloudWatchEvents{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-02-03",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSEvents",
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

// newRequest creates a new request for a CloudWatchEvents operation and runs any
// custom request initialization.
func (c *CloudWatchEvents) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
