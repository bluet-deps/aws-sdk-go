// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package elasticsearchservice provides a client for Amazon Elasticsearch Service.
package elasticsearchservice

import (
	"time"

	"bluet-deps/aws-sdk-go/aws/awsutil"
	"bluet-deps/aws-sdk-go/aws/request"
)

const opAddTags = "AddTags"

// AddTagsRequest generates a request for the AddTags operation.
func (c *ElasticsearchService) AddTagsRequest(input *AddTagsInput) (req *request.Request, output *AddTagsOutput) {
	op := &request.Operation{
		Name:       opAddTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/tags",
	}

	if input == nil {
		input = &AddTagsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &AddTagsOutput{}
	req.Data = output
	return
}

// Attaches tags to an existing Elasticsearch domain. Tags are a set of case-sensitive
// key value pairs. An Elasticsearch domain may have up to 10 tags. See  Tagging
// Amazon Elasticsearch Service Domains for more information. (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-awsresorcetagging"
// target="_blank)
func (c *ElasticsearchService) AddTags(input *AddTagsInput) (*AddTagsOutput, error) {
	req, out := c.AddTagsRequest(input)
	err := req.Send()
	return out, err
}

const opCreateElasticsearchDomain = "CreateElasticsearchDomain"

// CreateElasticsearchDomainRequest generates a request for the CreateElasticsearchDomain operation.
func (c *ElasticsearchService) CreateElasticsearchDomainRequest(input *CreateElasticsearchDomainInput) (req *request.Request, output *CreateElasticsearchDomainOutput) {
	op := &request.Operation{
		Name:       opCreateElasticsearchDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/domain",
	}

	if input == nil {
		input = &CreateElasticsearchDomainInput{}
	}

	req = c.newRequest(op, input, output)
	output = &CreateElasticsearchDomainOutput{}
	req.Data = output
	return
}

// Creates a new Elasticsearch domain. For more information, see Creating Elasticsearch
// Domains (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains"
// target="_blank) in the Amazon Elasticsearch Service Developer Guide.
func (c *ElasticsearchService) CreateElasticsearchDomain(input *CreateElasticsearchDomainInput) (*CreateElasticsearchDomainOutput, error) {
	req, out := c.CreateElasticsearchDomainRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteElasticsearchDomain = "DeleteElasticsearchDomain"

// DeleteElasticsearchDomainRequest generates a request for the DeleteElasticsearchDomain operation.
func (c *ElasticsearchService) DeleteElasticsearchDomainRequest(input *DeleteElasticsearchDomainInput) (req *request.Request, output *DeleteElasticsearchDomainOutput) {
	op := &request.Operation{
		Name:       opDeleteElasticsearchDomain,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}",
	}

	if input == nil {
		input = &DeleteElasticsearchDomainInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DeleteElasticsearchDomainOutput{}
	req.Data = output
	return
}

// Permanently deletes the specified Elasticsearch domain and all of its data.
// Once a domain is deleted, it cannot be recovered.
func (c *ElasticsearchService) DeleteElasticsearchDomain(input *DeleteElasticsearchDomainInput) (*DeleteElasticsearchDomainOutput, error) {
	req, out := c.DeleteElasticsearchDomainRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeElasticsearchDomain = "DescribeElasticsearchDomain"

// DescribeElasticsearchDomainRequest generates a request for the DescribeElasticsearchDomain operation.
func (c *ElasticsearchService) DescribeElasticsearchDomainRequest(input *DescribeElasticsearchDomainInput) (req *request.Request, output *DescribeElasticsearchDomainOutput) {
	op := &request.Operation{
		Name:       opDescribeElasticsearchDomain,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}",
	}

	if input == nil {
		input = &DescribeElasticsearchDomainInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeElasticsearchDomainOutput{}
	req.Data = output
	return
}

// Returns domain configuration information about the specified Elasticsearch
// domain, including the domain ID, domain endpoint, and domain ARN.
func (c *ElasticsearchService) DescribeElasticsearchDomain(input *DescribeElasticsearchDomainInput) (*DescribeElasticsearchDomainOutput, error) {
	req, out := c.DescribeElasticsearchDomainRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeElasticsearchDomainConfig = "DescribeElasticsearchDomainConfig"

// DescribeElasticsearchDomainConfigRequest generates a request for the DescribeElasticsearchDomainConfig operation.
func (c *ElasticsearchService) DescribeElasticsearchDomainConfigRequest(input *DescribeElasticsearchDomainConfigInput) (req *request.Request, output *DescribeElasticsearchDomainConfigOutput) {
	op := &request.Operation{
		Name:       opDescribeElasticsearchDomainConfig,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}/config",
	}

	if input == nil {
		input = &DescribeElasticsearchDomainConfigInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeElasticsearchDomainConfigOutput{}
	req.Data = output
	return
}

// Provides cluster configuration information about the specified Elasticsearch
// domain, such as the state, creation date, update version, and update date
// for cluster options.
func (c *ElasticsearchService) DescribeElasticsearchDomainConfig(input *DescribeElasticsearchDomainConfigInput) (*DescribeElasticsearchDomainConfigOutput, error) {
	req, out := c.DescribeElasticsearchDomainConfigRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeElasticsearchDomains = "DescribeElasticsearchDomains"

// DescribeElasticsearchDomainsRequest generates a request for the DescribeElasticsearchDomains operation.
func (c *ElasticsearchService) DescribeElasticsearchDomainsRequest(input *DescribeElasticsearchDomainsInput) (req *request.Request, output *DescribeElasticsearchDomainsOutput) {
	op := &request.Operation{
		Name:       opDescribeElasticsearchDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/domain-info",
	}

	if input == nil {
		input = &DescribeElasticsearchDomainsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &DescribeElasticsearchDomainsOutput{}
	req.Data = output
	return
}

// Returns domain configuration information about the specified Elasticsearch
// domains, including the domain ID, domain endpoint, and domain ARN.
func (c *ElasticsearchService) DescribeElasticsearchDomains(input *DescribeElasticsearchDomainsInput) (*DescribeElasticsearchDomainsOutput, error) {
	req, out := c.DescribeElasticsearchDomainsRequest(input)
	err := req.Send()
	return out, err
}

const opListDomainNames = "ListDomainNames"

// ListDomainNamesRequest generates a request for the ListDomainNames operation.
func (c *ElasticsearchService) ListDomainNamesRequest(input *ListDomainNamesInput) (req *request.Request, output *ListDomainNamesOutput) {
	op := &request.Operation{
		Name:       opListDomainNames,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/domain",
	}

	if input == nil {
		input = &ListDomainNamesInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListDomainNamesOutput{}
	req.Data = output
	return
}

// Returns the name of all Elasticsearch domains owned by the current user's
// account.
func (c *ElasticsearchService) ListDomainNames(input *ListDomainNamesInput) (*ListDomainNamesOutput, error) {
	req, out := c.ListDomainNamesRequest(input)
	err := req.Send()
	return out, err
}

const opListTags = "ListTags"

// ListTagsRequest generates a request for the ListTags operation.
func (c *ElasticsearchService) ListTagsRequest(input *ListTagsInput) (req *request.Request, output *ListTagsOutput) {
	op := &request.Operation{
		Name:       opListTags,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-01-01/tags/",
	}

	if input == nil {
		input = &ListTagsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &ListTagsOutput{}
	req.Data = output
	return
}

// Returns all tags for the given Elasticsearch domain.
func (c *ElasticsearchService) ListTags(input *ListTagsInput) (*ListTagsOutput, error) {
	req, out := c.ListTagsRequest(input)
	err := req.Send()
	return out, err
}

const opRemoveTags = "RemoveTags"

// RemoveTagsRequest generates a request for the RemoveTags operation.
func (c *ElasticsearchService) RemoveTagsRequest(input *RemoveTagsInput) (req *request.Request, output *RemoveTagsOutput) {
	op := &request.Operation{
		Name:       opRemoveTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/tags-removal",
	}

	if input == nil {
		input = &RemoveTagsInput{}
	}

	req = c.newRequest(op, input, output)
	output = &RemoveTagsOutput{}
	req.Data = output
	return
}

// Removes the specified set of tags from the specified Elasticsearch domain.
func (c *ElasticsearchService) RemoveTags(input *RemoveTagsInput) (*RemoveTagsOutput, error) {
	req, out := c.RemoveTagsRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateElasticsearchDomainConfig = "UpdateElasticsearchDomainConfig"

// UpdateElasticsearchDomainConfigRequest generates a request for the UpdateElasticsearchDomainConfig operation.
func (c *ElasticsearchService) UpdateElasticsearchDomainConfigRequest(input *UpdateElasticsearchDomainConfigInput) (req *request.Request, output *UpdateElasticsearchDomainConfigOutput) {
	op := &request.Operation{
		Name:       opUpdateElasticsearchDomainConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/domain/{DomainName}/config",
	}

	if input == nil {
		input = &UpdateElasticsearchDomainConfigInput{}
	}

	req = c.newRequest(op, input, output)
	output = &UpdateElasticsearchDomainConfigOutput{}
	req.Data = output
	return
}

// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
func (c *ElasticsearchService) UpdateElasticsearchDomainConfig(input *UpdateElasticsearchDomainConfigInput) (*UpdateElasticsearchDomainConfigOutput, error) {
	req, out := c.UpdateElasticsearchDomainConfigRequest(input)
	err := req.Send()
	return out, err
}

// The configured access rules for the domain's document and search endpoints,
// and the current status of those rules.
type AccessPoliciesStatus struct {
	_ struct{} `type:"structure"`

	// The access policy configured for the Elasticsearch domain. Access policies
	// may be resource-based, IP-based, or IAM-based. See  Configuring Access Policies
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-access-policies"
	// target="_blank)for more information.
	Options *string `type:"string" required:"true"`

	// The status of the access policy for the Elasticsearch domain. See OptionStatus
	// for the status information that's included.
	Status *OptionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s AccessPoliciesStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessPoliciesStatus) GoString() string {
	return s.String()
}

// Container for the parameters to the AddTags operation. Specify the tags that
// you want to attach to the Elasticsearch domain.
type AddTagsInput struct {
	_ struct{} `type:"structure"`

	// Specify the ARN for which you want to add the tags.
	ARN *string `type:"string" required:"true"`

	// List of Tag that need to be added for the Elasticsearch domain.
	TagList []*Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsInput) GoString() string {
	return s.String()
}

type AddTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsOutput) GoString() string {
	return s.String()
}

// Status of the advanced options for the specified Elasticsearch domain. Currently,
// the following advanced options are available:
//
//  Option to allow references to indices in an HTTP request body. Must be
// false when configuring access to individual sub-resources. By default, the
// value is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options"
// target="_blank) for more information. Option to specify the percentage of
// heap space that is allocated to field data. By default, this setting is unbounded.
//  For more information, see Configuring Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options).
type AdvancedOptionsStatus struct {
	_ struct{} `type:"structure"`

	// Specifies the status of advanced options for the specified Elasticsearch
	// domain.
	Options map[string]*string `type:"map" required:"true"`

	// Specifies the status of OptionStatus for advanced options for the specified
	// Elasticsearch domain.
	Status *OptionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s AdvancedOptionsStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AdvancedOptionsStatus) GoString() string {
	return s.String()
}

type CreateElasticsearchDomainInput struct {
	_ struct{} `type:"structure"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `type:"string"`

	// Option to allow references to indices in an HTTP request body. Must be false
	// when configuring access to individual sub-resources. By default, the value
	// is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options"
	// target="_blank) for more information.
	AdvancedOptions map[string]*string `type:"map"`

	// The name of the Elasticsearch domain that you are creating. Domain names
	// are unique across the domains owned by an account within an AWS region. Domain
	// names must start with a letter or number and can contain the following characters:
	// a-z (lowercase), 0-9, and - (hyphen).
	DomainName *string `min:"3" type:"string" required:"true"`

	// Options to enable, disable and specify the type and size of EBS storage volumes.
	EBSOptions *EBSOptions `type:"structure"`

	// Configuration options for an Elasticsearch domain. Specifies the instance
	// type and number of instances in the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `type:"structure"`

	// Option to set time, in UTC format, of the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *SnapshotOptions `type:"structure"`
}

// String returns the string representation
func (s CreateElasticsearchDomainInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateElasticsearchDomainInput) GoString() string {
	return s.String()
}

// The result of a CreateElasticsearchDomain operation. Contains the status
// of the newly created Elasticsearch domain.
type CreateElasticsearchDomainOutput struct {
	_ struct{} `type:"structure"`

	// The status of the newly created Elasticsearch domain.
	DomainStatus *ElasticsearchDomainStatus `type:"structure"`
}

// String returns the string representation
func (s CreateElasticsearchDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateElasticsearchDomainOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the DeleteElasticsearchDomain operation.
// Specifies the name of the Elasticsearch domain that you want to delete.
type DeleteElasticsearchDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the Elasticsearch domain that you want to permanently delete.
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteElasticsearchDomainInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteElasticsearchDomainInput) GoString() string {
	return s.String()
}

// The result of a DeleteElasticsearchDomain request. Contains the status of
// the pending deletion, or no status if the domain and all of its resources
// have been deleted.
type DeleteElasticsearchDomainOutput struct {
	_ struct{} `type:"structure"`

	// The status of the Elasticsearch domain being deleted.
	DomainStatus *ElasticsearchDomainStatus `type:"structure"`
}

// String returns the string representation
func (s DeleteElasticsearchDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteElasticsearchDomainOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the DescribeElasticsearchDomainConfig operation.
// Specifies the domain name for which you want configuration information.
type DescribeElasticsearchDomainConfigInput struct {
	_ struct{} `type:"structure"`

	// The Elasticsearch domain that you want to get information about.
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainConfigInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainConfigInput) GoString() string {
	return s.String()
}

// The result of a DescribeElasticsearchDomainConfig request. Contains the configuration
// information of the requested domain.
type DescribeElasticsearchDomainConfigOutput struct {
	_ struct{} `type:"structure"`

	// The configuration information of the domain requested in the DescribeElasticsearchDomainConfig
	// request.
	DomainConfig *ElasticsearchDomainConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainConfigOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the DescribeElasticsearchDomain operation.
type DescribeElasticsearchDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the Elasticsearch domain for which you want information.
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainInput) GoString() string {
	return s.String()
}

// The result of a DescribeElasticsearchDomain request. Contains the status
// of the domain specified in the request.
type DescribeElasticsearchDomainOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the Elasticsearch domain.
	DomainStatus *ElasticsearchDomainStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the DescribeElasticsearchDomains operation.
// By default, the API returns the status of all Elasticsearch domains.
type DescribeElasticsearchDomainsInput struct {
	_ struct{} `type:"structure"`

	// The Elasticsearch domains for which you want information.
	DomainNames []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainsInput) GoString() string {
	return s.String()
}

// The result of a DescribeElasticsearchDomains request. Contains the status
// of the specified domains or all domains owned by the account.
type DescribeElasticsearchDomainsOutput struct {
	_ struct{} `type:"structure"`

	// The status of the domains requested in the DescribeElasticsearchDomains request.
	DomainStatusList []*ElasticsearchDomainStatus `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeElasticsearchDomainsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeElasticsearchDomainsOutput) GoString() string {
	return s.String()
}

type DomainInfo struct {
	_ struct{} `type:"structure"`

	// Specifies the DomainName.
	DomainName *string `min:"3" type:"string"`
}

// String returns the string representation
func (s DomainInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainInfo) GoString() string {
	return s.String()
}

// Options to enable, disable, and specify the properties of EBS storage volumes.
// For more information, see  Configuring EBS-based Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs"
// target="_blank).
type EBSOptions struct {
	_ struct{} `type:"structure"`

	// Specifies whether EBS-based storage is enabled.
	EBSEnabled *bool `type:"boolean"`

	// Specifies the IOPD for a Provisioned IOPS EBS volume (SSD).
	Iops *int64 `type:"integer"`

	// Integer to specify the size of an EBS volume.
	VolumeSize *int64 `type:"integer"`

	// Specifies the volume type for EBS-based storage.
	VolumeType *string `type:"string" enum:"VolumeType"`
}

// String returns the string representation
func (s EBSOptions) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EBSOptions) GoString() string {
	return s.String()
}

// Status of the EBS options for the specified Elasticsearch domain.
type EBSOptionsStatus struct {
	_ struct{} `type:"structure"`

	// Specifies the EBS options for the specified Elasticsearch domain.
	Options *EBSOptions `type:"structure" required:"true"`

	// Specifies the status of the EBS options for the specified Elasticsearch domain.
	Status *OptionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s EBSOptionsStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EBSOptionsStatus) GoString() string {
	return s.String()
}

// Specifies the configuration for the domain cluster, such as the type and
// number of instances.
type ElasticsearchClusterConfig struct {
	_ struct{} `type:"structure"`

	// Total number of dedicated master nodes, active and on standby, for the cluster.
	DedicatedMasterCount *int64 `type:"integer"`

	// A boolean value to indicate whether a dedicated master node is enabled. See
	// About Dedicated Master Nodes (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-dedicatedmasternodes"
	// target="_blank) for more information.
	DedicatedMasterEnabled *bool `type:"boolean"`

	// The instance type for a dedicated master node.
	DedicatedMasterType *string `type:"string" enum:"ESPartitionInstanceType"`

	// The number of instances in the specified domain cluster.
	InstanceCount *int64 `type:"integer"`

	// The instance type for an Elasticsearch cluster.
	InstanceType *string `type:"string" enum:"ESPartitionInstanceType"`

	// A boolean value to indicate whether zone awareness is enabled. See About
	// Zone Awareness (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-zoneawareness"
	// target="_blank) for more information.
	ZoneAwarenessEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s ElasticsearchClusterConfig) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ElasticsearchClusterConfig) GoString() string {
	return s.String()
}

// Specifies the configuration status for the specified Elasticsearch domain.
type ElasticsearchClusterConfigStatus struct {
	_ struct{} `type:"structure"`

	// Specifies the cluster configuration for the specified Elasticsearch domain.
	Options *ElasticsearchClusterConfig `type:"structure" required:"true"`

	// Specifies the status of the configuration for the specified Elasticsearch
	// domain.
	Status *OptionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s ElasticsearchClusterConfigStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ElasticsearchClusterConfigStatus) GoString() string {
	return s.String()
}

// The configuration of an Elasticsearch domain.
type ElasticsearchDomainConfig struct {
	_ struct{} `type:"structure"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *AccessPoliciesStatus `type:"structure"`

	// Specifies the AdvancedOptions for the domain. See Configuring Advanced Options
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options"
	// target="_blank) for more information.
	AdvancedOptions *AdvancedOptionsStatus `type:"structure"`

	// Specifies the EBSOptions for the Elasticsearch domain.
	EBSOptions *EBSOptionsStatus `type:"structure"`

	// Specifies the ElasticsearchClusterConfig for the Elasticsearch domain.
	ElasticsearchClusterConfig *ElasticsearchClusterConfigStatus `type:"structure"`

	// Specifies the SnapshotOptions for the Elasticsearch domain.
	SnapshotOptions *SnapshotOptionsStatus `type:"structure"`
}

// String returns the string representation
func (s ElasticsearchDomainConfig) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ElasticsearchDomainConfig) GoString() string {
	return s.String()
}

// The current status of an Elasticsearch domain.
type ElasticsearchDomainStatus struct {
	_ struct{} `type:"structure"`

	// The Amazon resource name (ARN) of an Elasticsearch domain. See Identifiers
	// for IAM Entities (http://docs.aws.amazon.com/IAM/latest/UserGuide/index.html?Using_Identifiers.html"
	// target="_blank) in Using AWS Identity and Access Management for more information.
	ARN *string `type:"string" required:"true"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `type:"string"`

	// Specifies the status of the AdvancedOptions
	AdvancedOptions map[string]*string `type:"map"`

	// The domain creation status. True if the creation of an Elasticsearch domain
	// is complete. False if domain creation is still in progress.
	Created *bool `type:"boolean"`

	// The domain deletion status. True if a delete request has been received for
	// the domain but resource cleanup is still in progress. False if the domain
	// has not been deleted. Once domain deletion is complete, the status of the
	// domain is no longer returned.
	Deleted *bool `type:"boolean"`

	// The unique identifier for the specified Elasticsearch domain.
	DomainId *string `min:"1" type:"string" required:"true"`

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter
	// or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	DomainName *string `min:"3" type:"string" required:"true"`

	// The EBSOptions for the specified domain. See Configuring EBS-based Storage
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs"
	// target="_blank) for more information.
	EBSOptions *EBSOptions `type:"structure"`

	// The type and number of instances in the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `type:"structure" required:"true"`

	// The Elasticsearch domain endpoint that you use to submit index and search
	// requests.
	Endpoint *string `type:"string"`

	// The status of the Elasticsearch domain configuration. True if Amazon Elasticsearch
	// Service is processing configuration changes. False if the configuration is
	// active.
	Processing *bool `type:"boolean"`

	// Specifies the status of the SnapshotOptions
	SnapshotOptions *SnapshotOptions `type:"structure"`
}

// String returns the string representation
func (s ElasticsearchDomainStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ElasticsearchDomainStatus) GoString() string {
	return s.String()
}

type ListDomainNamesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListDomainNamesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainNamesInput) GoString() string {
	return s.String()
}

// The result of a ListDomainNames operation. Contains the names of all Elasticsearch
// domains owned by this account.
type ListDomainNamesOutput struct {
	_ struct{} `type:"structure"`

	// List of Elasticsearch domain names.
	DomainNames []*DomainInfo `type:"list"`
}

// String returns the string representation
func (s ListDomainNamesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainNamesOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the ListTags operation. Specify the ARN for
// the Elasticsearch domain to which the tags are attached that you want to
// view are attached.
type ListTagsInput struct {
	_ struct{} `type:"structure"`

	// Specify the ARN for the Elasticsearch domain to which the tags are attached
	// that you want to view.
	ARN *string `location:"querystring" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsInput) GoString() string {
	return s.String()
}

// The result of a ListTags operation. Contains tags for all requested Elasticsearch
// domains.
type ListTagsOutput struct {
	_ struct{} `type:"structure"`

	// List of Tag for the requested Elasticsearch domain.
	TagList []*Tag `type:"list"`
}

// String returns the string representation
func (s ListTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsOutput) GoString() string {
	return s.String()
}

// Provides the current status of the entity.
type OptionStatus struct {
	_ struct{} `type:"structure"`

	// Timestamp which tells the creation date for the entity.
	CreationDate *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// Indicates whether the Elasticsearch domain is being deleted.
	PendingDeletion *bool `type:"boolean"`

	// Provides the OptionState for the Elasticsearch domain.
	State *string `type:"string" required:"true" enum:"OptionState"`

	// Timestamp which tells the last updated time for the entity.
	UpdateDate *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// Specifies the latest version for the entity.
	UpdateVersion *int64 `type:"integer"`
}

// String returns the string representation
func (s OptionStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s OptionStatus) GoString() string {
	return s.String()
}

// Container for the parameters to the RemoveTags operation. Specify the ARN
// for the Elasticsearch domain from which you want to remove the specified
// TagKey.
type RemoveTagsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the ARN for the Elasticsearch domain from which you want to delete
	// the specified tags.
	ARN *string `type:"string" required:"true"`

	// Specifies the TagKey list which you want to remove from the Elasticsearch
	// domain.
	TagKeys []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsInput) GoString() string {
	return s.String()
}

type RemoveTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsOutput) GoString() string {
	return s.String()
}

// Specifies the time, in UTC format, when the service takes a daily automated
// snapshot of the specified Elasticsearch domain. Default value is 0 hours.
type SnapshotOptions struct {
	_ struct{} `type:"structure"`

	// Specifies the time, in UTC format, when the service takes a daily automated
	// snapshot of the specified Elasticsearch domain. Default value is 0 hours.
	AutomatedSnapshotStartHour *int64 `type:"integer"`
}

// String returns the string representation
func (s SnapshotOptions) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SnapshotOptions) GoString() string {
	return s.String()
}

// Status of a daily automated snapshot.
type SnapshotOptionsStatus struct {
	_ struct{} `type:"structure"`

	// Specifies the daily snapshot options specified for the Elasticsearch domain.
	Options *SnapshotOptions `type:"structure" required:"true"`

	// Specifies the status of a daily automated snapshot.
	Status *OptionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s SnapshotOptionsStatus) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SnapshotOptionsStatus) GoString() string {
	return s.String()
}

// Specifies a key value pair for a resource tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// Specifies the TagKey, the name of the tag. Tag keys must be unique for the
	// Elasticsearch domain to which they are attached.
	Key *string `min:"1" type:"string" required:"true"`

	// Specifies the TagValue, the value assigned to the corresponding tag key.
	// Tag values can be null and do not have to be unique in a tag set. For example,
	// you can have a key value pair in a tag set of project : Trinity and cost-center
	// : Trinity
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Tag) GoString() string {
	return s.String()
}

// Container for the parameters to the UpdateElasticsearchDomain operation.
// Specifies the type and number of instances in the domain cluster.
type UpdateElasticsearchDomainConfigInput struct {
	_ struct{} `type:"structure"`

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `type:"string"`

	// Modifies the advanced option to allow references to indices in an HTTP request
	// body. Must be false when configuring access to individual sub-resources.
	// By default, the value is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options"
	// target="_blank) for more information.
	AdvancedOptions map[string]*string `type:"map"`

	// The name of the Elasticsearch domain that you are updating.
	DomainName *string `location:"uri" locationName:"DomainName" min:"3" type:"string" required:"true"`

	// Specify the type and size of the EBS volume that you want to use.
	EBSOptions *EBSOptions `type:"structure"`

	// The type and number of instances to instantiate for the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `type:"structure"`

	// Option to set the time, in UTC format, for the daily automated snapshot.
	// Default value is 0 hours.
	SnapshotOptions *SnapshotOptions `type:"structure"`
}

// String returns the string representation
func (s UpdateElasticsearchDomainConfigInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateElasticsearchDomainConfigInput) GoString() string {
	return s.String()
}

// The result of an UpdateElasticsearchDomain request. Contains the status of
// the Elasticsearch domain being updated.
type UpdateElasticsearchDomainConfigOutput struct {
	_ struct{} `type:"structure"`

	// The status of the updated Elasticsearch domain.
	DomainConfig *ElasticsearchDomainConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateElasticsearchDomainConfigOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateElasticsearchDomainConfigOutput) GoString() string {
	return s.String()
}

const (
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeM3MediumElasticsearch = "m3.medium.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeM3LargeElasticsearch = "m3.large.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeM3XlargeElasticsearch = "m3.xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeM32xlargeElasticsearch = "m3.2xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeT2MicroElasticsearch = "t2.micro.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeT2SmallElasticsearch = "t2.small.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeT2MediumElasticsearch = "t2.medium.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeR3LargeElasticsearch = "r3.large.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeR3XlargeElasticsearch = "r3.xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeR32xlargeElasticsearch = "r3.2xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeR34xlargeElasticsearch = "r3.4xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeR38xlargeElasticsearch = "r3.8xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeI2XlargeElasticsearch = "i2.xlarge.elasticsearch"
	// @enum ESPartitionInstanceType
	ESPartitionInstanceTypeI22xlargeElasticsearch = "i2.2xlarge.elasticsearch"
)

// The state of a requested change. One of the following:
//
//  Processing: The request change is still in-process. Active: The request
// change is processed and deployed to the Elasticsearch domain.
const (
	// @enum OptionState
	OptionStateRequiresIndexDocuments = "RequiresIndexDocuments"
	// @enum OptionState
	OptionStateProcessing = "Processing"
	// @enum OptionState
	OptionStateActive = "Active"
)

// The type of EBS volume, standard, gp2, or io1. See Configuring EBS-based
// Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs"
// target="_blank)for more information.
const (
	// @enum VolumeType
	VolumeTypeStandard = "standard"
	// @enum VolumeType
	VolumeTypeGp2 = "gp2"
	// @enum VolumeType
	VolumeTypeIo1 = "io1"
)
