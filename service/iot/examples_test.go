// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package iot_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/bluet-deps/aws-sdk-go/aws"
	"github.com/bluet-deps/aws-sdk-go/aws/session"
	"github.com/bluet-deps/aws-sdk-go/service/iot"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleIoT_AcceptCertificateTransfer() {
	svc := iot.New(session.New())

	params := &iot.AcceptCertificateTransferInput{
		CertificateId: aws.String("CertificateId"), // Required
		SetAsActive:   aws.Bool(true),
	}
	resp, err := svc.AcceptCertificateTransfer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_AttachPrincipalPolicy() {
	svc := iot.New(session.New())

	params := &iot.AttachPrincipalPolicyInput{
		PolicyName: aws.String("PolicyName"), // Required
		Principal:  aws.String("Principal"),  // Required
	}
	resp, err := svc.AttachPrincipalPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_AttachThingPrincipal() {
	svc := iot.New(session.New())

	params := &iot.AttachThingPrincipalInput{
		Principal: aws.String("Principal"), // Required
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.AttachThingPrincipal(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CancelCertificateTransfer() {
	svc := iot.New(session.New())

	params := &iot.CancelCertificateTransferInput{
		CertificateId: aws.String("CertificateId"), // Required
	}
	resp, err := svc.CancelCertificateTransfer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreateCertificateFromCsr() {
	svc := iot.New(session.New())

	params := &iot.CreateCertificateFromCsrInput{
		CertificateSigningRequest: aws.String("CertificateSigningRequest"), // Required
		SetAsActive:               aws.Bool(true),
	}
	resp, err := svc.CreateCertificateFromCsr(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreateKeysAndCertificate() {
	svc := iot.New(session.New())

	params := &iot.CreateKeysAndCertificateInput{
		SetAsActive: aws.Bool(true),
	}
	resp, err := svc.CreateKeysAndCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreatePolicy() {
	svc := iot.New(session.New())

	params := &iot.CreatePolicyInput{
		PolicyDocument: aws.String("PolicyDocument"), // Required
		PolicyName:     aws.String("PolicyName"),     // Required
	}
	resp, err := svc.CreatePolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreatePolicyVersion() {
	svc := iot.New(session.New())

	params := &iot.CreatePolicyVersionInput{
		PolicyDocument: aws.String("PolicyDocument"), // Required
		PolicyName:     aws.String("PolicyName"),     // Required
		SetAsDefault:   aws.Bool(true),
	}
	resp, err := svc.CreatePolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreateThing() {
	svc := iot.New(session.New())

	params := &iot.CreateThingInput{
		ThingName: aws.String("ThingName"), // Required
		AttributePayload: &iot.AttributePayload{
			Attributes: map[string]*string{
				"Key": aws.String("AttributeValue"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.CreateThing(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_CreateTopicRule() {
	svc := iot.New(session.New())

	params := &iot.CreateTopicRuleInput{
		RuleName: aws.String("RuleName"), // Required
		TopicRulePayload: &iot.TopicRulePayload{ // Required
			Actions: []*iot.Action{ // Required
				{ // Required
					DynamoDB: &iot.DynamoDBAction{
						HashKeyField:  aws.String("HashKeyField"),  // Required
						HashKeyValue:  aws.String("HashKeyValue"),  // Required
						RangeKeyField: aws.String("RangeKeyField"), // Required
						RangeKeyValue: aws.String("RangeKeyValue"), // Required
						RoleArn:       aws.String("AwsArn"),        // Required
						TableName:     aws.String("TableName"),     // Required
						PayloadField:  aws.String("PayloadField"),
					},
					Firehose: &iot.FirehoseAction{
						DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
						RoleArn:            aws.String("AwsArn"),             // Required
					},
					Kinesis: &iot.KinesisAction{
						RoleArn:      aws.String("AwsArn"),     // Required
						StreamName:   aws.String("StreamName"), // Required
						PartitionKey: aws.String("PartitionKey"),
					},
					Lambda: &iot.LambdaAction{
						FunctionArn: aws.String("FunctionArn"), // Required
					},
					Republish: &iot.RepublishAction{
						RoleArn: aws.String("AwsArn"),       // Required
						Topic:   aws.String("TopicPattern"), // Required
					},
					S3: &iot.S3Action{
						BucketName: aws.String("BucketName"), // Required
						Key:        aws.String("Key"),        // Required
						RoleArn:    aws.String("AwsArn"),     // Required
					},
					Sns: &iot.SnsAction{
						RoleArn:   aws.String("AwsArn"), // Required
						TargetArn: aws.String("AwsArn"), // Required
					},
					Sqs: &iot.SqsAction{
						QueueUrl:  aws.String("QueueUrl"), // Required
						RoleArn:   aws.String("AwsArn"),   // Required
						UseBase64: aws.Bool(true),
					},
				},
				// More values...
			},
			Sql:          aws.String("SQL"), // Required
			Description:  aws.String("Description"),
			RuleDisabled: aws.Bool(true),
		},
	}
	resp, err := svc.CreateTopicRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DeleteCertificate() {
	svc := iot.New(session.New())

	params := &iot.DeleteCertificateInput{
		CertificateId: aws.String("CertificateId"), // Required
	}
	resp, err := svc.DeleteCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DeletePolicy() {
	svc := iot.New(session.New())

	params := &iot.DeletePolicyInput{
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.DeletePolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DeletePolicyVersion() {
	svc := iot.New(session.New())

	params := &iot.DeletePolicyVersionInput{
		PolicyName:      aws.String("PolicyName"),      // Required
		PolicyVersionId: aws.String("PolicyVersionId"), // Required
	}
	resp, err := svc.DeletePolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DeleteThing() {
	svc := iot.New(session.New())

	params := &iot.DeleteThingInput{
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.DeleteThing(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DeleteTopicRule() {
	svc := iot.New(session.New())

	params := &iot.DeleteTopicRuleInput{
		RuleName: aws.String("RuleName"), // Required
	}
	resp, err := svc.DeleteTopicRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DescribeCertificate() {
	svc := iot.New(session.New())

	params := &iot.DescribeCertificateInput{
		CertificateId: aws.String("CertificateId"), // Required
	}
	resp, err := svc.DescribeCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DescribeEndpoint() {
	svc := iot.New(session.New())

	var params *iot.DescribeEndpointInput
	resp, err := svc.DescribeEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DescribeThing() {
	svc := iot.New(session.New())

	params := &iot.DescribeThingInput{
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.DescribeThing(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DetachPrincipalPolicy() {
	svc := iot.New(session.New())

	params := &iot.DetachPrincipalPolicyInput{
		PolicyName: aws.String("PolicyName"), // Required
		Principal:  aws.String("Principal"),  // Required
	}
	resp, err := svc.DetachPrincipalPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_DetachThingPrincipal() {
	svc := iot.New(session.New())

	params := &iot.DetachThingPrincipalInput{
		Principal: aws.String("Principal"), // Required
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.DetachThingPrincipal(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_GetLoggingOptions() {
	svc := iot.New(session.New())

	var params *iot.GetLoggingOptionsInput
	resp, err := svc.GetLoggingOptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_GetPolicy() {
	svc := iot.New(session.New())

	params := &iot.GetPolicyInput{
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.GetPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_GetPolicyVersion() {
	svc := iot.New(session.New())

	params := &iot.GetPolicyVersionInput{
		PolicyName:      aws.String("PolicyName"),      // Required
		PolicyVersionId: aws.String("PolicyVersionId"), // Required
	}
	resp, err := svc.GetPolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_GetTopicRule() {
	svc := iot.New(session.New())

	params := &iot.GetTopicRuleInput{
		RuleName: aws.String("RuleName"), // Required
	}
	resp, err := svc.GetTopicRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListCertificates() {
	svc := iot.New(session.New())

	params := &iot.ListCertificatesInput{
		AscendingOrder: aws.Bool(true),
		Marker:         aws.String("Marker"),
		PageSize:       aws.Int64(1),
	}
	resp, err := svc.ListCertificates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListPolicies() {
	svc := iot.New(session.New())

	params := &iot.ListPoliciesInput{
		AscendingOrder: aws.Bool(true),
		Marker:         aws.String("Marker"),
		PageSize:       aws.Int64(1),
	}
	resp, err := svc.ListPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListPolicyVersions() {
	svc := iot.New(session.New())

	params := &iot.ListPolicyVersionsInput{
		PolicyName: aws.String("PolicyName"), // Required
	}
	resp, err := svc.ListPolicyVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListPrincipalPolicies() {
	svc := iot.New(session.New())

	params := &iot.ListPrincipalPoliciesInput{
		Principal:      aws.String("Principal"), // Required
		AscendingOrder: aws.Bool(true),
		Marker:         aws.String("Marker"),
		PageSize:       aws.Int64(1),
	}
	resp, err := svc.ListPrincipalPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListPrincipalThings() {
	svc := iot.New(session.New())

	params := &iot.ListPrincipalThingsInput{
		Principal:  aws.String("Principal"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("NextToken"),
	}
	resp, err := svc.ListPrincipalThings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListThingPrincipals() {
	svc := iot.New(session.New())

	params := &iot.ListThingPrincipalsInput{
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.ListThingPrincipals(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListThings() {
	svc := iot.New(session.New())

	params := &iot.ListThingsInput{
		AttributeName:  aws.String("AttributeName"),
		AttributeValue: aws.String("AttributeValue"),
		MaxResults:     aws.Int64(1),
		NextToken:      aws.String("NextToken"),
	}
	resp, err := svc.ListThings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ListTopicRules() {
	svc := iot.New(session.New())

	params := &iot.ListTopicRulesInput{
		MaxResults:   aws.Int64(1),
		NextToken:    aws.String("NextToken"),
		RuleDisabled: aws.Bool(true),
		Topic:        aws.String("Topic"),
	}
	resp, err := svc.ListTopicRules(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_RejectCertificateTransfer() {
	svc := iot.New(session.New())

	params := &iot.RejectCertificateTransferInput{
		CertificateId: aws.String("CertificateId"), // Required
	}
	resp, err := svc.RejectCertificateTransfer(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_ReplaceTopicRule() {
	svc := iot.New(session.New())

	params := &iot.ReplaceTopicRuleInput{
		RuleName: aws.String("RuleName"), // Required
		TopicRulePayload: &iot.TopicRulePayload{ // Required
			Actions: []*iot.Action{ // Required
				{ // Required
					DynamoDB: &iot.DynamoDBAction{
						HashKeyField:  aws.String("HashKeyField"),  // Required
						HashKeyValue:  aws.String("HashKeyValue"),  // Required
						RangeKeyField: aws.String("RangeKeyField"), // Required
						RangeKeyValue: aws.String("RangeKeyValue"), // Required
						RoleArn:       aws.String("AwsArn"),        // Required
						TableName:     aws.String("TableName"),     // Required
						PayloadField:  aws.String("PayloadField"),
					},
					Firehose: &iot.FirehoseAction{
						DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
						RoleArn:            aws.String("AwsArn"),             // Required
					},
					Kinesis: &iot.KinesisAction{
						RoleArn:      aws.String("AwsArn"),     // Required
						StreamName:   aws.String("StreamName"), // Required
						PartitionKey: aws.String("PartitionKey"),
					},
					Lambda: &iot.LambdaAction{
						FunctionArn: aws.String("FunctionArn"), // Required
					},
					Republish: &iot.RepublishAction{
						RoleArn: aws.String("AwsArn"),       // Required
						Topic:   aws.String("TopicPattern"), // Required
					},
					S3: &iot.S3Action{
						BucketName: aws.String("BucketName"), // Required
						Key:        aws.String("Key"),        // Required
						RoleArn:    aws.String("AwsArn"),     // Required
					},
					Sns: &iot.SnsAction{
						RoleArn:   aws.String("AwsArn"), // Required
						TargetArn: aws.String("AwsArn"), // Required
					},
					Sqs: &iot.SqsAction{
						QueueUrl:  aws.String("QueueUrl"), // Required
						RoleArn:   aws.String("AwsArn"),   // Required
						UseBase64: aws.Bool(true),
					},
				},
				// More values...
			},
			Sql:          aws.String("SQL"), // Required
			Description:  aws.String("Description"),
			RuleDisabled: aws.Bool(true),
		},
	}
	resp, err := svc.ReplaceTopicRule(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_SetDefaultPolicyVersion() {
	svc := iot.New(session.New())

	params := &iot.SetDefaultPolicyVersionInput{
		PolicyName:      aws.String("PolicyName"),      // Required
		PolicyVersionId: aws.String("PolicyVersionId"), // Required
	}
	resp, err := svc.SetDefaultPolicyVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_SetLoggingOptions() {
	svc := iot.New(session.New())

	params := &iot.SetLoggingOptionsInput{
		LoggingOptionsPayload: &iot.LoggingOptionsPayload{
			RoleArn:  aws.String("AwsArn"), // Required
			LogLevel: aws.String("LogLevel"),
		},
	}
	resp, err := svc.SetLoggingOptions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_TransferCertificate() {
	svc := iot.New(session.New())

	params := &iot.TransferCertificateInput{
		CertificateId:    aws.String("CertificateId"), // Required
		TargetAwsAccount: aws.String("AwsAccountId"),  // Required
	}
	resp, err := svc.TransferCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_UpdateCertificate() {
	svc := iot.New(session.New())

	params := &iot.UpdateCertificateInput{
		CertificateId: aws.String("CertificateId"),     // Required
		NewStatus:     aws.String("CertificateStatus"), // Required
	}
	resp, err := svc.UpdateCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleIoT_UpdateThing() {
	svc := iot.New(session.New())

	params := &iot.UpdateThingInput{
		AttributePayload: &iot.AttributePayload{ // Required
			Attributes: map[string]*string{
				"Key": aws.String("AttributeValue"), // Required
				// More values...
			},
		},
		ThingName: aws.String("ThingName"), // Required
	}
	resp, err := svc.UpdateThing(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
