// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package emr_test

import (
	"bytes"
	"fmt"
	"time"

	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/session"
	"bluet-deps/aws-sdk-go/service/emr"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleEMR_AddInstanceGroups() {
	svc := emr.New(session.New())

	params := &emr.AddInstanceGroupsInput{
		InstanceGroups: []*emr.InstanceGroupConfig{ // Required
			{ // Required
				InstanceCount: aws.Int64(1),                   // Required
				InstanceRole:  aws.String("InstanceRoleType"), // Required
				InstanceType:  aws.String("InstanceType"),     // Required
				BidPrice:      aws.String("XmlStringMaxLen256"),
				Configurations: []*emr.Configuration{
					{ // Required
						Classification: aws.String("String"),
						Configurations: []*emr.Configuration{
						// Recursive values...
						},
						Properties: map[string]*string{
							"Key": aws.String("String"), // Required
							// More values...
						},
					},
					// More values...
				},
				Market: aws.String("MarketType"),
				Name:   aws.String("XmlStringMaxLen256"),
			},
			// More values...
		},
		JobFlowId: aws.String("XmlStringMaxLen256"), // Required
	}
	resp, err := svc.AddInstanceGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_AddJobFlowSteps() {
	svc := emr.New(session.New())

	params := &emr.AddJobFlowStepsInput{
		JobFlowId: aws.String("XmlStringMaxLen256"), // Required
		Steps: []*emr.StepConfig{ // Required
			{ // Required
				HadoopJarStep: &emr.HadoopJarStepConfig{ // Required
					Jar: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
					MainClass: aws.String("XmlString"),
					Properties: []*emr.KeyValue{
						{ // Required
							Key:   aws.String("XmlString"),
							Value: aws.String("XmlString"),
						},
						// More values...
					},
				},
				Name:            aws.String("XmlStringMaxLen256"), // Required
				ActionOnFailure: aws.String("ActionOnFailure"),
			},
			// More values...
		},
	}
	resp, err := svc.AddJobFlowSteps(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_AddTags() {
	svc := emr.New(session.New())

	params := &emr.AddTagsInput{
		ResourceId: aws.String("ResourceId"), // Required
		Tags: []*emr.Tag{ // Required
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
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

func ExampleEMR_DescribeCluster() {
	svc := emr.New(session.New())

	params := &emr.DescribeClusterInput{
		ClusterId: aws.String("ClusterId"), // Required
	}
	resp, err := svc.DescribeCluster(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_DescribeJobFlows() {
	svc := emr.New(session.New())

	params := &emr.DescribeJobFlowsInput{
		CreatedAfter:  aws.Time(time.Now()),
		CreatedBefore: aws.Time(time.Now()),
		JobFlowIds: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		JobFlowStates: []*string{
			aws.String("JobFlowExecutionState"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeJobFlows(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_DescribeStep() {
	svc := emr.New(session.New())

	params := &emr.DescribeStepInput{
		ClusterId: aws.String("ClusterId"), // Required
		StepId:    aws.String("StepId"),    // Required
	}
	resp, err := svc.DescribeStep(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ListBootstrapActions() {
	svc := emr.New(session.New())

	params := &emr.ListBootstrapActionsInput{
		ClusterId: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
	}
	resp, err := svc.ListBootstrapActions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ListClusters() {
	svc := emr.New(session.New())

	params := &emr.ListClustersInput{
		ClusterStates: []*string{
			aws.String("ClusterState"), // Required
			// More values...
		},
		CreatedAfter:  aws.Time(time.Now()),
		CreatedBefore: aws.Time(time.Now()),
		Marker:        aws.String("Marker"),
	}
	resp, err := svc.ListClusters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ListInstanceGroups() {
	svc := emr.New(session.New())

	params := &emr.ListInstanceGroupsInput{
		ClusterId: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
	}
	resp, err := svc.ListInstanceGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ListInstances() {
	svc := emr.New(session.New())

	params := &emr.ListInstancesInput{
		ClusterId:       aws.String("ClusterId"), // Required
		InstanceGroupId: aws.String("InstanceGroupId"),
		InstanceGroupTypes: []*string{
			aws.String("InstanceGroupType"), // Required
			// More values...
		},
		Marker: aws.String("Marker"),
	}
	resp, err := svc.ListInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ListSteps() {
	svc := emr.New(session.New())

	params := &emr.ListStepsInput{
		ClusterId: aws.String("ClusterId"), // Required
		Marker:    aws.String("Marker"),
		StepIds: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		StepStates: []*string{
			aws.String("StepState"), // Required
			// More values...
		},
	}
	resp, err := svc.ListSteps(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_ModifyInstanceGroups() {
	svc := emr.New(session.New())

	params := &emr.ModifyInstanceGroupsInput{
		InstanceGroups: []*emr.InstanceGroupModifyConfig{
			{ // Required
				InstanceGroupId: aws.String("XmlStringMaxLen256"), // Required
				EC2InstanceIdsToTerminate: []*string{
					aws.String("InstanceId"), // Required
					// More values...
				},
				InstanceCount: aws.Int64(1),
			},
			// More values...
		},
	}
	resp, err := svc.ModifyInstanceGroups(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_RemoveTags() {
	svc := emr.New(session.New())

	params := &emr.RemoveTagsInput{
		ResourceId: aws.String("ResourceId"), // Required
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

func ExampleEMR_RunJobFlow() {
	svc := emr.New(session.New())

	params := &emr.RunJobFlowInput{
		Instances: &emr.JobFlowInstancesConfig{ // Required
			AdditionalMasterSecurityGroups: []*string{
				aws.String("XmlStringMaxLen256"), // Required
				// More values...
			},
			AdditionalSlaveSecurityGroups: []*string{
				aws.String("XmlStringMaxLen256"), // Required
				// More values...
			},
			Ec2KeyName:                    aws.String("XmlStringMaxLen256"),
			Ec2SubnetId:                   aws.String("XmlStringMaxLen256"),
			EmrManagedMasterSecurityGroup: aws.String("XmlStringMaxLen256"),
			EmrManagedSlaveSecurityGroup:  aws.String("XmlStringMaxLen256"),
			HadoopVersion:                 aws.String("XmlStringMaxLen256"),
			InstanceCount:                 aws.Int64(1),
			InstanceGroups: []*emr.InstanceGroupConfig{
				{ // Required
					InstanceCount: aws.Int64(1),                   // Required
					InstanceRole:  aws.String("InstanceRoleType"), // Required
					InstanceType:  aws.String("InstanceType"),     // Required
					BidPrice:      aws.String("XmlStringMaxLen256"),
					Configurations: []*emr.Configuration{
						{ // Required
							Classification: aws.String("String"),
							Configurations: []*emr.Configuration{
							// Recursive values...
							},
							Properties: map[string]*string{
								"Key": aws.String("String"), // Required
								// More values...
							},
						},
						// More values...
					},
					Market: aws.String("MarketType"),
					Name:   aws.String("XmlStringMaxLen256"),
				},
				// More values...
			},
			KeepJobFlowAliveWhenNoSteps: aws.Bool(true),
			MasterInstanceType:          aws.String("InstanceType"),
			Placement: &emr.PlacementType{
				AvailabilityZone: aws.String("XmlString"), // Required
			},
			ServiceAccessSecurityGroup: aws.String("XmlStringMaxLen256"),
			SlaveInstanceType:          aws.String("InstanceType"),
			TerminationProtected:       aws.Bool(true),
		},
		Name:           aws.String("XmlStringMaxLen256"), // Required
		AdditionalInfo: aws.String("XmlString"),
		AmiVersion:     aws.String("XmlStringMaxLen256"),
		Applications: []*emr.Application{
			{ // Required
				AdditionalInfo: map[string]*string{
					"Key": aws.String("String"), // Required
					// More values...
				},
				Args: []*string{
					aws.String("String"), // Required
					// More values...
				},
				Name:    aws.String("String"),
				Version: aws.String("String"),
			},
			// More values...
		},
		BootstrapActions: []*emr.BootstrapActionConfig{
			{ // Required
				Name: aws.String("XmlStringMaxLen256"), // Required
				ScriptBootstrapAction: &emr.ScriptBootstrapActionConfig{ // Required
					Path: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		Configurations: []*emr.Configuration{
			{ // Required
				Classification: aws.String("String"),
				Configurations: []*emr.Configuration{
				// Recursive values...
				},
				Properties: map[string]*string{
					"Key": aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},
		JobFlowRole: aws.String("XmlString"),
		LogUri:      aws.String("XmlString"),
		NewSupportedProducts: []*emr.SupportedProductConfig{
			{ // Required
				Args: []*string{
					aws.String("XmlString"), // Required
					// More values...
				},
				Name: aws.String("XmlStringMaxLen256"),
			},
			// More values...
		},
		ReleaseLabel: aws.String("XmlStringMaxLen256"),
		ServiceRole:  aws.String("XmlString"),
		Steps: []*emr.StepConfig{
			{ // Required
				HadoopJarStep: &emr.HadoopJarStepConfig{ // Required
					Jar: aws.String("XmlString"), // Required
					Args: []*string{
						aws.String("XmlString"), // Required
						// More values...
					},
					MainClass: aws.String("XmlString"),
					Properties: []*emr.KeyValue{
						{ // Required
							Key:   aws.String("XmlString"),
							Value: aws.String("XmlString"),
						},
						// More values...
					},
				},
				Name:            aws.String("XmlStringMaxLen256"), // Required
				ActionOnFailure: aws.String("ActionOnFailure"),
			},
			// More values...
		},
		SupportedProducts: []*string{
			aws.String("XmlStringMaxLen256"), // Required
			// More values...
		},
		Tags: []*emr.Tag{
			{ // Required
				Key:   aws.String("String"),
				Value: aws.String("String"),
			},
			// More values...
		},
		VisibleToAllUsers: aws.Bool(true),
	}
	resp, err := svc.RunJobFlow(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_SetTerminationProtection() {
	svc := emr.New(session.New())

	params := &emr.SetTerminationProtectionInput{
		JobFlowIds: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
		TerminationProtected: aws.Bool(true), // Required
	}
	resp, err := svc.SetTerminationProtection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_SetVisibleToAllUsers() {
	svc := emr.New(session.New())

	params := &emr.SetVisibleToAllUsersInput{
		JobFlowIds: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
		VisibleToAllUsers: aws.Bool(true), // Required
	}
	resp, err := svc.SetVisibleToAllUsers(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleEMR_TerminateJobFlows() {
	svc := emr.New(session.New())

	params := &emr.TerminateJobFlowsInput{
		JobFlowIds: []*string{ // Required
			aws.String("XmlString"), // Required
			// More values...
		},
	}
	resp, err := svc.TerminateJobFlows(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
