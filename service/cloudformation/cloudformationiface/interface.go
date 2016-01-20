// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudformationiface provides an interface for the AWS CloudFormation.
package cloudformationiface

import (
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/service/cloudformation"
)

// CloudFormationAPI is the interface type for cloudformation.CloudFormation.
type CloudFormationAPI interface {
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) (*request.Request, *cloudformation.CancelUpdateStackOutput)

	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)

	CreateStackRequest(*cloudformation.CreateStackInput) (*request.Request, *cloudformation.CreateStackOutput)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)

	DeleteStackRequest(*cloudformation.DeleteStackInput) (*request.Request, *cloudformation.DeleteStackOutput)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)

	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) (*request.Request, *cloudformation.DescribeAccountLimitsOutput)

	DescribeAccountLimits(*cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)

	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) (*request.Request, *cloudformation.DescribeStackEventsOutput)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error

	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) (*request.Request, *cloudformation.DescribeStackResourceOutput)

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)

	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) (*request.Request, *cloudformation.DescribeStackResourcesOutput)

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)

	DescribeStacksRequest(*cloudformation.DescribeStacksInput) (*request.Request, *cloudformation.DescribeStacksOutput)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error

	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) (*request.Request, *cloudformation.EstimateTemplateCostOutput)

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)

	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) (*request.Request, *cloudformation.GetStackPolicyOutput)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)

	GetTemplateRequest(*cloudformation.GetTemplateInput) (*request.Request, *cloudformation.GetTemplateOutput)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)

	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) (*request.Request, *cloudformation.GetTemplateSummaryOutput)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)

	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) (*request.Request, *cloudformation.ListStackResourcesOutput)

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error

	ListStacksRequest(*cloudformation.ListStacksInput) (*request.Request, *cloudformation.ListStacksOutput)

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error

	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) (*request.Request, *cloudformation.SetStackPolicyOutput)

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)

	SignalResourceRequest(*cloudformation.SignalResourceInput) (*request.Request, *cloudformation.SignalResourceOutput)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)

	UpdateStackRequest(*cloudformation.UpdateStackInput) (*request.Request, *cloudformation.UpdateStackOutput)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)

	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) (*request.Request, *cloudformation.ValidateTemplateOutput)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
}

var _ CloudFormationAPI = (*cloudformation.CloudFormation)(nil)
