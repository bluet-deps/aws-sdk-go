// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudwatchevents_test

import (
	"bytes"
	"fmt"
	"time"

	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/session"
	"bluet-deps/aws-sdk-go/service/cloudwatchevents"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudWatchEvents_DeleteRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.DeleteRuleInput{
		Name: aws.String("RuleName"), // Required
	}
	resp, err := svc.DeleteRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_DescribeRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.DescribeRuleInput{
		Name: aws.String("RuleName"), // Required
	}
	resp, err := svc.DescribeRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_DisableRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.DisableRuleInput{
		Name: aws.String("RuleName"), // Required
	}
	resp, err := svc.DisableRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_EnableRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.EnableRuleInput{
		Name: aws.String("RuleName"), // Required
	}
	resp, err := svc.EnableRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_ListRuleNamesByTarget() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.ListRuleNamesByTargetInput{
		TargetArn: aws.String("TargetArn"), // Required
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListRuleNamesByTarget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_ListRules() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.ListRulesInput{
		Limit:      aws.Int64(1),
		NamePrefix: aws.String("RuleName"),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListRules(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_ListTargetsByRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.ListTargetsByRuleInput{
		Rule:      aws.String("RuleName"), // Required
		Limit:     aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListTargetsByRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_PutEvents() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.PutEventsInput{
		Entries: []*cloudwatchevents.PutEventsRequestEntry{ // Required
			{ // Required
				Detail:     aws.String("String"),
				DetailType: aws.String("String"),
				Resources: []*string{
					aws.String("EventResource"), // Required
					// More values...
				},
				Source: aws.String("String"),
				Time:   aws.Time(time.Now()),
			},
			// More values...
		},
	}
	resp, err := svc.PutEvents(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_PutRule() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.PutRuleInput{
		Name:               aws.String("RuleName"), // Required
		Description:        aws.String("RuleDescription"),
		EventPattern:       aws.String("EventPattern"),
		RoleArn:            aws.String("RoleArn"),
		ScheduleExpression: aws.String("ScheduleExpression"),
		State:              aws.String("RuleState"),
	}
	resp, err := svc.PutRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_PutTargets() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.PutTargetsInput{
		Rule: aws.String("RuleName"), // Required
		Targets: []*cloudwatchevents.Target{ // Required
			{ // Required
				Arn:       aws.String("TargetArn"), // Required
				Id:        aws.String("TargetId"),  // Required
				Input:     aws.String("TargetInput"),
				InputPath: aws.String("TargetInputPath"),
			},
			// More values...
		},
	}
	resp, err := svc.PutTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_RemoveTargets() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.RemoveTargetsInput{
		Ids: []*string{ // Required
			aws.String("TargetId"), // Required
			// More values...
		},
		Rule: aws.String("RuleName"), // Required
	}
	resp, err := svc.RemoveTargets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudWatchEvents_TestEventPattern() {
	svc := cloudwatchevents.New(session.New())

	params := &cloudwatchevents.TestEventPatternInput{
		Event:        aws.String("String"),       // Required
		EventPattern: aws.String("EventPattern"), // Required
	}
	resp, err := svc.TestEventPattern(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
