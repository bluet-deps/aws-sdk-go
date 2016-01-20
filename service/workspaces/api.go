// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package workspaces provides a client for Amazon WorkSpaces.
package workspaces

import (
	"github.com/bluet-deps/aws-sdk-go/aws/awsutil"
	"github.com/bluet-deps/aws-sdk-go/aws/request"
)

const opCreateWorkspaces = "CreateWorkspaces"

// CreateWorkspacesRequest generates a request for the CreateWorkspaces operation.
func (c *WorkSpaces) CreateWorkspacesRequest(input *CreateWorkspacesInput) (req *request.Request, output *CreateWorkspacesOutput) {
	op := &request.Operation{
		Name:       opCreateWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWorkspacesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateWorkspacesOutput{}
	req.Data = output
	return
}

// Creates one or more WorkSpaces.
//
//  This operation is asynchronous and returns before the WorkSpaces are created.
func (c *WorkSpaces) CreateWorkspaces(input *CreateWorkspacesInput) (*CreateWorkspacesOutput, error) {
	req, out := c.CreateWorkspacesRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeWorkspaceBundles = "DescribeWorkspaceBundles"

// DescribeWorkspaceBundlesRequest generates a request for the DescribeWorkspaceBundles operation.
func (c *WorkSpaces) DescribeWorkspaceBundlesRequest(input *DescribeWorkspaceBundlesInput) (req *request.Request, output *DescribeWorkspaceBundlesOutput) {
	op := &request.Operation{
		Name:       opDescribeWorkspaceBundles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeWorkspaceBundlesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeWorkspaceBundlesOutput{}
	req.Data = output
	return
}

// Obtains information about the WorkSpace bundles that are available to your
// account in the specified region.
//
// You can filter the results with either the BundleIds parameter, or the Owner
// parameter, but not both.
//
// This operation supports pagination with the use of the NextToken request
// and response parameters. If more results are available, the NextToken response
// member contains a token that you pass in the next call to this operation
// to retrieve the next set of items.
func (c *WorkSpaces) DescribeWorkspaceBundles(input *DescribeWorkspaceBundlesInput) (*DescribeWorkspaceBundlesOutput, error) {
	req, out := c.DescribeWorkspaceBundlesRequest(input)
	err := req.Send()
	return out, err
}

func (c *WorkSpaces) DescribeWorkspaceBundlesPages(input *DescribeWorkspaceBundlesInput, fn func(p *DescribeWorkspaceBundlesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.DescribeWorkspaceBundlesRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*DescribeWorkspaceBundlesOutput), lastPage)
	})
}

const opDescribeWorkspaceDirectories = "DescribeWorkspaceDirectories"

// DescribeWorkspaceDirectoriesRequest generates a request for the DescribeWorkspaceDirectories operation.
func (c *WorkSpaces) DescribeWorkspaceDirectoriesRequest(input *DescribeWorkspaceDirectoriesInput) (req *request.Request, output *DescribeWorkspaceDirectoriesOutput) {
	op := &request.Operation{
		Name:       opDescribeWorkspaceDirectories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeWorkspaceDirectoriesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeWorkspaceDirectoriesOutput{}
	req.Data = output
	return
}

// Retrieves information about the AWS Directory Service directories in the
// region that are registered with Amazon WorkSpaces and are available to your
// account.
//
// This operation supports pagination with the use of the NextToken request
// and response parameters. If more results are available, the NextToken response
// member contains a token that you pass in the next call to this operation
// to retrieve the next set of items.
func (c *WorkSpaces) DescribeWorkspaceDirectories(input *DescribeWorkspaceDirectoriesInput) (*DescribeWorkspaceDirectoriesOutput, error) {
	req, out := c.DescribeWorkspaceDirectoriesRequest(input)
	err := req.Send()
	return out, err
}

func (c *WorkSpaces) DescribeWorkspaceDirectoriesPages(input *DescribeWorkspaceDirectoriesInput, fn func(p *DescribeWorkspaceDirectoriesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.DescribeWorkspaceDirectoriesRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*DescribeWorkspaceDirectoriesOutput), lastPage)
	})
}

const opDescribeWorkspaces = "DescribeWorkspaces"

// DescribeWorkspacesRequest generates a request for the DescribeWorkspaces operation.
func (c *WorkSpaces) DescribeWorkspacesRequest(input *DescribeWorkspacesInput) (req *request.Request, output *DescribeWorkspacesOutput) {
	op := &request.Operation{
		Name:       opDescribeWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &request.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeWorkspacesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeWorkspacesOutput{}
	req.Data = output
	return
}

// Obtains information about the specified WorkSpaces.
//
// Only one of the filter parameters, such as BundleId, DirectoryId, or WorkspaceIds,
// can be specified at a time.
//
// This operation supports pagination with the use of the NextToken request
// and response parameters. If more results are available, the NextToken response
// member contains a token that you pass in the next call to this operation
// to retrieve the next set of items.
func (c *WorkSpaces) DescribeWorkspaces(input *DescribeWorkspacesInput) (*DescribeWorkspacesOutput, error) {
	req, out := c.DescribeWorkspacesRequest(input)
	err := req.Send()
	return out, err
}

func (c *WorkSpaces) DescribeWorkspacesPages(input *DescribeWorkspacesInput, fn func(p *DescribeWorkspacesOutput, lastPage bool) (shouldContinue bool)) error {
	page, _ := c.DescribeWorkspacesRequest(input)
	page.Handlers.Build.PushBack(request.MakeAddToUserAgentFreeFormHandler("Paginator"))
	return page.EachPage(func(p interface{}, lastPage bool) bool {
		return fn(p.(*DescribeWorkspacesOutput), lastPage)
	})
}

const opRebootWorkspaces = "RebootWorkspaces"

// RebootWorkspacesRequest generates a request for the RebootWorkspaces operation.
func (c *WorkSpaces) RebootWorkspacesRequest(input *RebootWorkspacesInput) (req *request.Request, output *RebootWorkspacesOutput) {
	op := &request.Operation{
		Name:       opRebootWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootWorkspacesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RebootWorkspacesOutput{}
	req.Data = output
	return
}

// Reboots the specified WorkSpaces.
//
// To be able to reboot a WorkSpace, the WorkSpace must have a State of AVAILABLE,
// IMPAIRED, or INOPERABLE.
//
//  This operation is asynchronous and will return before the WorkSpaces have
// rebooted.
func (c *WorkSpaces) RebootWorkspaces(input *RebootWorkspacesInput) (*RebootWorkspacesOutput, error) {
	req, out := c.RebootWorkspacesRequest(input)
	err := req.Send()
	return out, err
}

const opRebuildWorkspaces = "RebuildWorkspaces"

// RebuildWorkspacesRequest generates a request for the RebuildWorkspaces operation.
func (c *WorkSpaces) RebuildWorkspacesRequest(input *RebuildWorkspacesInput) (req *request.Request, output *RebuildWorkspacesOutput) {
	op := &request.Operation{
		Name:       opRebuildWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebuildWorkspacesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RebuildWorkspacesOutput{}
	req.Data = output
	return
}

// Rebuilds the specified WorkSpaces.
//
// Rebuilding a WorkSpace is a potentially destructive action that can result
// in the loss of data. Rebuilding a WorkSpace causes the following to occur:
//
//  The system is restored to the image of the bundle that the WorkSpace is
// created from. Any applications that have been installed, or system settings
// that have been made since the WorkSpace was created will be lost. The data
// drive (D drive) is re-created from the last automatic snapshot taken of the
// data drive. The current contents of the data drive are overwritten. Automatic
// snapshots of the data drive are taken every 12 hours, so the snapshot can
// be as much as 12 hours old.  To be able to rebuild a WorkSpace, the WorkSpace
// must have a State of AVAILABLE or ERROR.
//
//  This operation is asynchronous and will return before the WorkSpaces have
// been completely rebuilt.
func (c *WorkSpaces) RebuildWorkspaces(input *RebuildWorkspacesInput) (*RebuildWorkspacesOutput, error) {
	req, out := c.RebuildWorkspacesRequest(input)
	err := req.Send()
	return out, err
}

const opTerminateWorkspaces = "TerminateWorkspaces"

// TerminateWorkspacesRequest generates a request for the TerminateWorkspaces operation.
func (c *WorkSpaces) TerminateWorkspacesRequest(input *TerminateWorkspacesInput) (req *request.Request, output *TerminateWorkspacesOutput) {
	op := &request.Operation{
		Name:       opTerminateWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateWorkspacesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &TerminateWorkspacesOutput{}
	req.Data = output
	return
}

// Terminates the specified WorkSpaces.
//
// Terminating a WorkSpace is a permanent action and cannot be undone. The
// user's data is not maintained and will be destroyed. If you need to archive
// any user data, contact Amazon Web Services before terminating the WorkSpace.
//
// You can terminate a WorkSpace that is in any state except SUSPENDED.
//
//  This operation is asynchronous and will return before the WorkSpaces have
// been completely terminated.
func (c *WorkSpaces) TerminateWorkspaces(input *TerminateWorkspacesInput) (*TerminateWorkspacesOutput, error) {
	req, out := c.TerminateWorkspacesRequest(input)
	err := req.Send()
	return out, err
}

// Contains information about the compute type of a WorkSpace bundle.
type ComputeType struct {
	_ struct{} `type:"structure"`

	// The name of the compute type for the bundle.
	Name *string `type:"string" enum:"Compute"`
}

// String returns the string representation
func (s ComputeType) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ComputeType) GoString() string {
	return s.String()
}

// Contains the inputs for the CreateWorkspaces operation.
type CreateWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// An array of structures that specify the WorkSpaces to create.
	Workspaces []*WorkspaceRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateWorkspacesInput) GoString() string {
	return s.String()
}

// Contains the result of the CreateWorkspaces operation.
type CreateWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that represent the WorkSpaces that could not be created.
	FailedRequests []*FailedCreateWorkspaceRequest `type:"list"`

	// An array of structures that represent the WorkSpaces that were created.
	//
	// Because this operation is asynchronous, the identifier in WorkspaceId is
	// not immediately available. If you immediately call DescribeWorkspaces with
	// this identifier, no information will be returned.
	PendingRequests []*Workspace `type:"list"`
}

// String returns the string representation
func (s CreateWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateWorkspacesOutput) GoString() string {
	return s.String()
}

// Contains default WorkSpace creation information.
type DefaultWorkspaceCreationProperties struct {
	_ struct{} `type:"structure"`

	// The identifier of any custom security groups that are applied to the WorkSpaces
	// when they are created.
	CustomSecurityGroupId *string `type:"string"`

	// The organizational unit (OU) in the directory that the WorkSpace machine
	// accounts are placed in.
	DefaultOu *string `type:"string"`

	// A public IP address will be attached to all WorkSpaces that are created or
	// rebuilt.
	EnableInternetAccess *bool `type:"boolean"`

	// Specifies if the directory is enabled for Amazon WorkDocs.
	EnableWorkDocs *bool `type:"boolean"`

	// The WorkSpace user is an administrator on the WorkSpace.
	UserEnabledAsLocalAdministrator *bool `type:"boolean"`
}

// String returns the string representation
func (s DefaultWorkspaceCreationProperties) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DefaultWorkspaceCreationProperties) GoString() string {
	return s.String()
}

// Contains the inputs for the DescribeWorkspaceBundles operation.
type DescribeWorkspaceBundlesInput struct {
	_ struct{} `type:"structure"`

	// An array of strings that contains the identifiers of the bundles to retrieve.
	// This parameter cannot be combined with any other filter parameter.
	BundleIds []*string `min:"1" type:"list"`

	// The NextToken value from a previous call to this operation. Pass null if
	// this is the first call.
	NextToken *string `min:"1" type:"string"`

	// The owner of the bundles to retrieve. This parameter cannot be combined with
	// any other filter parameter.
	//
	// This contains one of the following values:
	//
	//  null - Retrieves the bundles that belong to the account making the call.
	//  AMAZON - Retrieves the bundles that are provided by AWS.
	Owner *string `type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceBundlesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspaceBundlesInput) GoString() string {
	return s.String()
}

// Contains the results of the DescribeWorkspaceBundles operation.
type DescribeWorkspaceBundlesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that contain information about the bundles.
	Bundles []*WorkspaceBundle `type:"list"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to this operation to retrieve the next set
	// of items. This token is valid for one day and must be used within that timeframe.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceBundlesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspaceBundlesOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DescribeWorkspaceDirectories operation.
type DescribeWorkspaceDirectoriesInput struct {
	_ struct{} `type:"structure"`

	// An array of strings that contains the directory identifiers to retrieve information
	// for. If this member is null, all directories are retrieved.
	DirectoryIds []*string `min:"1" type:"list"`

	// The NextToken value from a previous call to this operation. Pass null if
	// this is the first call.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceDirectoriesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspaceDirectoriesInput) GoString() string {
	return s.String()
}

// Contains the results of the DescribeWorkspaceDirectories operation.
type DescribeWorkspaceDirectoriesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that contain information about the directories.
	Directories []*WorkspaceDirectory `type:"list"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to this operation to retrieve the next set
	// of items. This token is valid for one day and must be used within that timeframe.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeWorkspaceDirectoriesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspaceDirectoriesOutput) GoString() string {
	return s.String()
}

// Contains the inputs for the DescribeWorkspaces operation.
type DescribeWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of a bundle to obtain the WorkSpaces for. All WorkSpaces that
	// are created from this bundle will be retrieved. This parameter cannot be
	// combined with any other filter parameter.
	BundleId *string `type:"string"`

	// Specifies the directory identifier to which to limit the WorkSpaces. Optionally,
	// you can specify a specific directory user with the UserName parameter. This
	// parameter cannot be combined with any other filter parameter.
	DirectoryId *string `type:"string"`

	// The maximum number of items to return.
	Limit *int64 `min:"1" type:"integer"`

	// The NextToken value from a previous call to this operation. Pass null if
	// this is the first call.
	NextToken *string `min:"1" type:"string"`

	// Used with the DirectoryId parameter to specify the directory user for which
	// to obtain the WorkSpace.
	UserName *string `min:"1" type:"string"`

	// An array of strings that contain the identifiers of the WorkSpaces for which
	// to retrieve information. This parameter cannot be combined with any other
	// filter parameter.
	//
	// Because the CreateWorkspaces operation is asynchronous, the identifier returned
	// by CreateWorkspaces is not immediately available. If you immediately call
	// DescribeWorkspaces with this identifier, no information will be returned.
	WorkspaceIds []*string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspacesInput) GoString() string {
	return s.String()
}

// Contains the results for the DescribeWorkspaces operation.
type DescribeWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to this operation to retrieve the next set
	// of items. This token is valid for one day and must be used within that timeframe.
	NextToken *string `min:"1" type:"string"`

	// An array of structures that contain the information about the WorkSpaces.
	//
	// Because the CreateWorkspaces operation is asynchronous, some of this information
	// may be incomplete for a newly-created WorkSpace.
	Workspaces []*Workspace `type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeWorkspacesOutput) GoString() string {
	return s.String()
}

// Contains information about a WorkSpace that could not be created.
type FailedCreateWorkspaceRequest struct {
	_ struct{} `type:"structure"`

	// The error code.
	ErrorCode *string `type:"string"`

	// The textual error message.
	ErrorMessage *string `type:"string"`

	// A WorkspaceRequest object that contains the information about the WorkSpace
	// that could not be created.
	WorkspaceRequest *WorkspaceRequest `type:"structure"`
}

// String returns the string representation
func (s FailedCreateWorkspaceRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FailedCreateWorkspaceRequest) GoString() string {
	return s.String()
}

// Contains information about a WorkSpace that could not be rebooted (RebootWorkspaces),
// rebuilt (RebuildWorkspaces), or terminated (TerminateWorkspaces).
type FailedWorkspaceChangeRequest struct {
	_ struct{} `type:"structure"`

	// The error code.
	ErrorCode *string `type:"string"`

	// The textual error message.
	ErrorMessage *string `type:"string"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s FailedWorkspaceChangeRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s FailedWorkspaceChangeRequest) GoString() string {
	return s.String()
}

// Contains information used with the RebootWorkspaces operation to reboot a
// WorkSpace.
type RebootRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace to reboot.
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebootRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootRequest) GoString() string {
	return s.String()
}

// Contains the inputs for the RebootWorkspaces operation.
type RebootWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// An array of structures that specify the WorkSpaces to reboot.
	RebootWorkspaceRequests []*RebootRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RebootWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootWorkspacesInput) GoString() string {
	return s.String()
}

// Contains the results of the RebootWorkspaces operation.
type RebootWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that represent any WorkSpaces that could not be rebooted.
	FailedRequests []*FailedWorkspaceChangeRequest `type:"list"`
}

// String returns the string representation
func (s RebootWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebootWorkspacesOutput) GoString() string {
	return s.String()
}

// Contains information used with the RebuildWorkspaces operation to rebuild
// a WorkSpace.
type RebuildRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace to rebuild.
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RebuildRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebuildRequest) GoString() string {
	return s.String()
}

// Contains the inputs for the RebuildWorkspaces operation.
type RebuildWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// An array of structures that specify the WorkSpaces to rebuild.
	RebuildWorkspaceRequests []*RebuildRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s RebuildWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebuildWorkspacesInput) GoString() string {
	return s.String()
}

// Contains the results of the RebuildWorkspaces operation.
type RebuildWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that represent any WorkSpaces that could not be rebuilt.
	FailedRequests []*FailedWorkspaceChangeRequest `type:"list"`
}

// String returns the string representation
func (s RebuildWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RebuildWorkspacesOutput) GoString() string {
	return s.String()
}

// Contains information used with the TerminateWorkspaces operation to terminate
// a WorkSpace.
type TerminateRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the WorkSpace to terminate.
	WorkspaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TerminateRequest) GoString() string {
	return s.String()
}

// Contains the inputs for the TerminateWorkspaces operation.
type TerminateWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// An array of structures that specify the WorkSpaces to terminate.
	TerminateWorkspaceRequests []*TerminateRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TerminateWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TerminateWorkspacesInput) GoString() string {
	return s.String()
}

// Contains the results of the TerminateWorkspaces operation.
type TerminateWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures that represent any WorkSpaces that could not be terminated.
	FailedRequests []*FailedWorkspaceChangeRequest `type:"list"`
}

// String returns the string representation
func (s TerminateWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TerminateWorkspacesOutput) GoString() string {
	return s.String()
}

// Contains information about the user storage for a WorkSpace bundle.
type UserStorage struct {
	_ struct{} `type:"structure"`

	// The amount of user storage for the bundle.
	Capacity *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UserStorage) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UserStorage) GoString() string {
	return s.String()
}

// Contains information about a WorkSpace.
type Workspace struct {
	_ struct{} `type:"structure"`

	// The identifier of the bundle that the WorkSpace was created from.
	BundleId *string `type:"string"`

	// The name of the WorkSpace as seen by the operating system.
	ComputerName *string `type:"string"`

	// The identifier of the AWS Directory Service directory that the WorkSpace
	// belongs to.
	DirectoryId *string `type:"string"`

	// If the WorkSpace could not be created, this contains the error code.
	ErrorCode *string `type:"string"`

	// If the WorkSpace could not be created, this contains a textual error message
	// that describes the failure.
	ErrorMessage *string `type:"string"`

	// The IP address of the WorkSpace.
	IpAddress *string `type:"string"`

	// Specifies whether the data stored on the root volume, or C: drive, is encrypted.
	RootVolumeEncryptionEnabled *bool `type:"boolean"`

	// The operational state of the WorkSpace.
	State *string `type:"string" enum:"WorkspaceState"`

	// The identifier of the subnet that the WorkSpace is in.
	SubnetId *string `type:"string"`

	// The user that the WorkSpace is assigned to.
	UserName *string `min:"1" type:"string"`

	// Specifies whether the data stored on the user volume, or D: drive, is encrypted.
	UserVolumeEncryptionEnabled *bool `type:"boolean"`

	// The KMS key used to encrypt data stored on your WorkSpace.
	VolumeEncryptionKey *string `type:"string"`

	// The identifier of the WorkSpace.
	WorkspaceId *string `type:"string"`
}

// String returns the string representation
func (s Workspace) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Workspace) GoString() string {
	return s.String()
}

// Contains information about a WorkSpace bundle.
type WorkspaceBundle struct {
	_ struct{} `type:"structure"`

	// The bundle identifier.
	BundleId *string `type:"string"`

	// A ComputeType object that specifies the compute type for the bundle.
	ComputeType *ComputeType `type:"structure"`

	// The bundle description.
	Description *string `type:"string"`

	// The name of the bundle.
	Name *string `min:"1" type:"string"`

	// The owner of the bundle. This contains the owner's account identifier, or
	// AMAZON if the bundle is provided by AWS.
	Owner *string `type:"string"`

	// A UserStorage object that specifies the amount of user storage that the bundle
	// contains.
	UserStorage *UserStorage `type:"structure"`
}

// String returns the string representation
func (s WorkspaceBundle) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s WorkspaceBundle) GoString() string {
	return s.String()
}

// Contains information about an AWS Directory Service directory for use with
// Amazon WorkSpaces.
type WorkspaceDirectory struct {
	_ struct{} `type:"structure"`

	// The directory alias.
	Alias *string `type:"string"`

	// The user name for the service account.
	CustomerUserName *string `min:"1" type:"string"`

	// The directory identifier.
	DirectoryId *string `type:"string"`

	// The name of the directory.
	DirectoryName *string `type:"string"`

	// The directory type.
	DirectoryType *string `type:"string" enum:"WorkspaceDirectoryType"`

	// An array of strings that contains the IP addresses of the DNS servers for
	// the directory.
	DnsIpAddresses []*string `type:"list"`

	// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces
	// to make calls to other services, such as Amazon EC2, on your behalf.
	IamRoleId *string `type:"string"`

	// The registration code for the directory. This is the code that users enter
	// in their Amazon WorkSpaces client application to connect to the directory.
	RegistrationCode *string `min:"1" type:"string"`

	// The state of the directory's registration with Amazon WorkSpaces
	State *string `type:"string" enum:"WorkspaceDirectoryState"`

	// An array of strings that contains the identifiers of the subnets used with
	// the directory.
	SubnetIds []*string `type:"list"`

	// A structure that specifies the default creation properties for all WorkSpaces
	// in the directory.
	WorkspaceCreationProperties *DefaultWorkspaceCreationProperties `type:"structure"`

	// The identifier of the security group that is assigned to new WorkSpaces.
	WorkspaceSecurityGroupId *string `type:"string"`
}

// String returns the string representation
func (s WorkspaceDirectory) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s WorkspaceDirectory) GoString() string {
	return s.String()
}

// Contains information about a WorkSpace creation request.
type WorkspaceRequest struct {
	_ struct{} `type:"structure"`

	// The identifier of the bundle to create the WorkSpace from. You can use the
	// DescribeWorkspaceBundles operation to obtain a list of the bundles that are
	// available.
	BundleId *string `type:"string" required:"true"`

	// The identifier of the AWS Directory Service directory to create the WorkSpace
	// in. You can use the DescribeWorkspaceDirectories operation to obtain a list
	// of the directories that are available.
	DirectoryId *string `type:"string" required:"true"`

	// Specifies whether the data stored on the root volume, or C: drive, is encrypted.
	RootVolumeEncryptionEnabled *bool `type:"boolean"`

	// The username that the WorkSpace is assigned to. This username must exist
	// in the AWS Directory Service directory specified by the DirectoryId member.
	UserName *string `min:"1" type:"string" required:"true"`

	// Specifies whether the data stored on the user volume, or D: drive, is encrypted.
	UserVolumeEncryptionEnabled *bool `type:"boolean"`

	// The KMS key used to encrypt data stored on your WorkSpace.
	VolumeEncryptionKey *string `type:"string"`
}

// String returns the string representation
func (s WorkspaceRequest) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s WorkspaceRequest) GoString() string {
	return s.String()
}

const (
	// @enum Compute
	ComputeValue = "VALUE"
	// @enum Compute
	ComputeStandard = "STANDARD"
	// @enum Compute
	ComputePerformance = "PERFORMANCE"
)

const (
	// @enum WorkspaceDirectoryState
	WorkspaceDirectoryStateRegistering = "REGISTERING"
	// @enum WorkspaceDirectoryState
	WorkspaceDirectoryStateRegistered = "REGISTERED"
	// @enum WorkspaceDirectoryState
	WorkspaceDirectoryStateDeregistering = "DEREGISTERING"
	// @enum WorkspaceDirectoryState
	WorkspaceDirectoryStateDeregistered = "DEREGISTERED"
	// @enum WorkspaceDirectoryState
	WorkspaceDirectoryStateError = "ERROR"
)

const (
	// @enum WorkspaceDirectoryType
	WorkspaceDirectoryTypeSimpleAd = "SIMPLE_AD"
	// @enum WorkspaceDirectoryType
	WorkspaceDirectoryTypeAdConnector = "AD_CONNECTOR"
)

const (
	// @enum WorkspaceState
	WorkspaceStatePending = "PENDING"
	// @enum WorkspaceState
	WorkspaceStateAvailable = "AVAILABLE"
	// @enum WorkspaceState
	WorkspaceStateImpaired = "IMPAIRED"
	// @enum WorkspaceState
	WorkspaceStateUnhealthy = "UNHEALTHY"
	// @enum WorkspaceState
	WorkspaceStateRebooting = "REBOOTING"
	// @enum WorkspaceState
	WorkspaceStateRebuilding = "REBUILDING"
	// @enum WorkspaceState
	WorkspaceStateTerminating = "TERMINATING"
	// @enum WorkspaceState
	WorkspaceStateTerminated = "TERMINATED"
	// @enum WorkspaceState
	WorkspaceStateSuspended = "SUSPENDED"
	// @enum WorkspaceState
	WorkspaceStateError = "ERROR"
)
