// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatcheventsiface provides an interface for the Amazon CloudWatch Events.
package cloudwatcheventsiface

import (
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/service/cloudwatchevents"
)

// CloudWatchEventsAPI is the interface type for cloudwatchevents.CloudWatchEvents.
type CloudWatchEventsAPI interface {
	DeleteRuleRequest(*cloudwatchevents.DeleteRuleInput) (*request.Request, *cloudwatchevents.DeleteRuleOutput)

	DeleteRule(*cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error)

	DescribeRuleRequest(*cloudwatchevents.DescribeRuleInput) (*request.Request, *cloudwatchevents.DescribeRuleOutput)

	DescribeRule(*cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error)

	DisableRuleRequest(*cloudwatchevents.DisableRuleInput) (*request.Request, *cloudwatchevents.DisableRuleOutput)

	DisableRule(*cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error)

	EnableRuleRequest(*cloudwatchevents.EnableRuleInput) (*request.Request, *cloudwatchevents.EnableRuleOutput)

	EnableRule(*cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error)

	ListRuleNamesByTargetRequest(*cloudwatchevents.ListRuleNamesByTargetInput) (*request.Request, *cloudwatchevents.ListRuleNamesByTargetOutput)

	ListRuleNamesByTarget(*cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error)

	ListRulesRequest(*cloudwatchevents.ListRulesInput) (*request.Request, *cloudwatchevents.ListRulesOutput)

	ListRules(*cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error)

	ListTargetsByRuleRequest(*cloudwatchevents.ListTargetsByRuleInput) (*request.Request, *cloudwatchevents.ListTargetsByRuleOutput)

	ListTargetsByRule(*cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error)

	PutEventsRequest(*cloudwatchevents.PutEventsInput) (*request.Request, *cloudwatchevents.PutEventsOutput)

	PutEvents(*cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error)

	PutRuleRequest(*cloudwatchevents.PutRuleInput) (*request.Request, *cloudwatchevents.PutRuleOutput)

	PutRule(*cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error)

	PutTargetsRequest(*cloudwatchevents.PutTargetsInput) (*request.Request, *cloudwatchevents.PutTargetsOutput)

	PutTargets(*cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error)

	RemoveTargetsRequest(*cloudwatchevents.RemoveTargetsInput) (*request.Request, *cloudwatchevents.RemoveTargetsOutput)

	RemoveTargets(*cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error)

	TestEventPatternRequest(*cloudwatchevents.TestEventPatternInput) (*request.Request, *cloudwatchevents.TestEventPatternOutput)

	TestEventPattern(*cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error)
}

var _ CloudWatchEventsAPI = (*cloudwatchevents.CloudWatchEvents)(nil)
