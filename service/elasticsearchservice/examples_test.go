// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package elasticsearchservice_test

import (
	"bytes"
	"fmt"
	"time"

	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/session"
	"bluet-deps/aws-sdk-go/service/elasticsearchservice"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleElasticsearchService_AddTags() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.AddTagsInput{
		ARN: aws.String("ARN"), // Required
		TagList: []*elasticsearchservice.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.AddTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_CreateElasticsearchDomain() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.CreateElasticsearchDomainInput{
		DomainName:     aws.String("DomainName"), // Required
		AccessPolicies: aws.String("PolicyDocument"),
		AdvancedOptions: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		EBSOptions: &elasticsearchservice.EBSOptions{
			EBSEnabled: aws.Bool(true),
			Iops:       aws.Int64(1),
			VolumeSize: aws.Int64(1),
			VolumeType: aws.String("VolumeType"),
		},
		ElasticsearchClusterConfig: &elasticsearchservice.ElasticsearchClusterConfig{
			DedicatedMasterCount:   aws.Int64(1),
			DedicatedMasterEnabled: aws.Bool(true),
			DedicatedMasterType:    aws.String("ESPartitionInstanceType"),
			InstanceCount:          aws.Int64(1),
			InstanceType:           aws.String("ESPartitionInstanceType"),
			ZoneAwarenessEnabled:   aws.Bool(true),
		},
		SnapshotOptions: &elasticsearchservice.SnapshotOptions{
			AutomatedSnapshotStartHour: aws.Int64(1),
		},
	}
	resp, err := svc.CreateElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DeleteElasticsearchDomain() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.DeleteElasticsearchDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeleteElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomain() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.DescribeElasticsearchDomainInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeElasticsearchDomain(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomainConfig() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.DescribeElasticsearchDomainConfigInput{
		DomainName: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeElasticsearchDomainConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_DescribeElasticsearchDomains() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.DescribeElasticsearchDomainsInput{
		DomainNames: []*string{ // Required
			aws.String("DomainName"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeElasticsearchDomains(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListDomainNames() {
	svc := elasticsearchservice.New(session.New())

	var params *elasticsearchservice.ListDomainNamesInput
	resp, err := svc.ListDomainNames(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_ListTags() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.ListTagsInput{
		ARN: aws.String("ARN"), // Required
	}
	resp, err := svc.ListTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_RemoveTags() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.RemoveTagsInput{
		ARN: aws.String("ARN"), // Required
		TagKeys: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.RemoveTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleElasticsearchService_UpdateElasticsearchDomainConfig() {
	svc := elasticsearchservice.New(session.New())

	params := &elasticsearchservice.UpdateElasticsearchDomainConfigInput{
		DomainName:     aws.String("DomainName"), // Required
		AccessPolicies: aws.String("PolicyDocument"),
		AdvancedOptions: map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		EBSOptions: &elasticsearchservice.EBSOptions{
			EBSEnabled: aws.Bool(true),
			Iops:       aws.Int64(1),
			VolumeSize: aws.Int64(1),
			VolumeType: aws.String("VolumeType"),
		},
		ElasticsearchClusterConfig: &elasticsearchservice.ElasticsearchClusterConfig{
			DedicatedMasterCount:   aws.Int64(1),
			DedicatedMasterEnabled: aws.Bool(true),
			DedicatedMasterType:    aws.String("ESPartitionInstanceType"),
			InstanceCount:          aws.Int64(1),
			InstanceType:           aws.String("ESPartitionInstanceType"),
			ZoneAwarenessEnabled:   aws.Bool(true),
		},
		SnapshotOptions: &elasticsearchservice.SnapshotOptions{
			AutomatedSnapshotStartHour: aws.Int64(1),
		},
	}
	resp, err := svc.UpdateElasticsearchDomainConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
