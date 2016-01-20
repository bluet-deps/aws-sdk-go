// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package codedeploy

import (
	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/client"
	"bluet-deps/aws-sdk-go/aws/client/metadata"
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/private/protocol/jsonrpc"
	"bluet-deps/aws-sdk-go/private/signer/v4"
)

// Overview This is the AWS CodeDeploy API Reference. This guide provides descriptions
// of the AWS CodeDeploy APIs. For additional information, see the AWS CodeDeploy
// User Guide (http://docs.aws.amazon.com/codedeploy/latest/userguide).
//
// Using the APIs You can use the AWS CodeDeploy APIs to work with the following
// items:
//
//   Applications are unique identifiers that AWS CodeDeploy uses to ensure
// that the correct combinations of revisions, deployment configurations, and
// deployment groups are being referenced during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
// applications.
//
//   Deployment configurations are sets of deployment rules and deployment
// success and failure conditions that AWS CodeDeploy uses during deployments.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, and list deployment
// configurations.
//
//   Deployment groups are groups of instances to which application revisions
// can be deployed.
//
// You can use the AWS CodeDeploy APIs to create, delete, get, list, and update
// deployment groups.
//
//   Instances represent Amazon EC2 instances to which application revisions
// are deployed. Instances are identified by their Amazon EC2 tags or Auto Scaling
// group names. Instances belong to deployment groups.
//
// You can use the AWS CodeDeploy APIs to get and list instances.
//
//   Deployments represent the process of deploying revisions to instances.
//
// You can use the AWS CodeDeploy APIs to create, get, list, and stop deployments.
//
//   Application revisions are archive files that are stored in Amazon S3 buckets
// or GitHub repositories. These revisions contain source content (such as source
// code, web pages, executable files, any deployment scripts, and similar) along
// with an Application Specification file (AppSpec file). (The AppSpec file
// is unique to AWS CodeDeploy; it defines a series of deployment actions that
// you want AWS CodeDeploy to execute.) An application revision is uniquely
// identified by its Amazon S3 object key and its ETag, version, or both (for
// application revisions that are stored in Amazon S3 buckets) or by its repository
// name and commit ID (for applications revisions that are stored in GitHub
// repositories). Application revisions are deployed through deployment groups.
//
// You can use the AWS CodeDeploy APIs to get, list, and register application
// revisions.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CodeDeploy struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "codedeploy"

// New creates a new instance of the CodeDeploy client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CodeDeploy client from just a session.
//     svc := codedeploy.New(mySession)
//
//     // Create a CodeDeploy client with additional configuration
//     svc := codedeploy.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CodeDeploy {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CodeDeploy {
	svc := &CodeDeploy{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-10-06",
				JSONVersion:   "1.1",
				TargetPrefix:  "CodeDeploy_20141006",
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

// newRequest creates a new request for a CodeDeploy operation and runs any
// custom request initialization.
func (c *CodeDeploy) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
