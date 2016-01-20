// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudtrailiface provides an interface for the AWS CloudTrail.
package cloudtrailiface

import (
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/service/cloudtrail"
)

// CloudTrailAPI is the interface type for cloudtrail.CloudTrail.
type CloudTrailAPI interface {
	AddTagsRequest(*cloudtrail.AddTagsInput) (*request.Request, *cloudtrail.AddTagsOutput)

	AddTags(*cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error)

	CreateTrailRequest(*cloudtrail.CreateTrailInput) (*request.Request, *cloudtrail.CreateTrailOutput)

	CreateTrail(*cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)

	DeleteTrailRequest(*cloudtrail.DeleteTrailInput) (*request.Request, *cloudtrail.DeleteTrailOutput)

	DeleteTrail(*cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)

	DescribeTrailsRequest(*cloudtrail.DescribeTrailsInput) (*request.Request, *cloudtrail.DescribeTrailsOutput)

	DescribeTrails(*cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)

	GetTrailStatusRequest(*cloudtrail.GetTrailStatusInput) (*request.Request, *cloudtrail.GetTrailStatusOutput)

	GetTrailStatus(*cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)

	ListPublicKeysRequest(*cloudtrail.ListPublicKeysInput) (*request.Request, *cloudtrail.ListPublicKeysOutput)

	ListPublicKeys(*cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error)

	ListTagsRequest(*cloudtrail.ListTagsInput) (*request.Request, *cloudtrail.ListTagsOutput)

	ListTags(*cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error)

	LookupEventsRequest(*cloudtrail.LookupEventsInput) (*request.Request, *cloudtrail.LookupEventsOutput)

	LookupEvents(*cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)

	RemoveTagsRequest(*cloudtrail.RemoveTagsInput) (*request.Request, *cloudtrail.RemoveTagsOutput)

	RemoveTags(*cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error)

	StartLoggingRequest(*cloudtrail.StartLoggingInput) (*request.Request, *cloudtrail.StartLoggingOutput)

	StartLogging(*cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)

	StopLoggingRequest(*cloudtrail.StopLoggingInput) (*request.Request, *cloudtrail.StopLoggingOutput)

	StopLogging(*cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)

	UpdateTrailRequest(*cloudtrail.UpdateTrailInput) (*request.Request, *cloudtrail.UpdateTrailOutput)

	UpdateTrail(*cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
}

var _ CloudTrailAPI = (*cloudtrail.CloudTrail)(nil)
