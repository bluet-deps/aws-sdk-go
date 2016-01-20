// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package efsiface provides an interface for the Amazon Elastic File System.
package efsiface

import (
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/service/efs"
)

// EFSAPI is the interface type for efs.EFS.
type EFSAPI interface {
	CreateFileSystemRequest(*efs.CreateFileSystemInput) (*request.Request, *efs.FileSystemDescription)

	CreateFileSystem(*efs.CreateFileSystemInput) (*efs.FileSystemDescription, error)

	CreateMountTargetRequest(*efs.CreateMountTargetInput) (*request.Request, *efs.MountTargetDescription)

	CreateMountTarget(*efs.CreateMountTargetInput) (*efs.MountTargetDescription, error)

	CreateTagsRequest(*efs.CreateTagsInput) (*request.Request, *efs.CreateTagsOutput)

	CreateTags(*efs.CreateTagsInput) (*efs.CreateTagsOutput, error)

	DeleteFileSystemRequest(*efs.DeleteFileSystemInput) (*request.Request, *efs.DeleteFileSystemOutput)

	DeleteFileSystem(*efs.DeleteFileSystemInput) (*efs.DeleteFileSystemOutput, error)

	DeleteMountTargetRequest(*efs.DeleteMountTargetInput) (*request.Request, *efs.DeleteMountTargetOutput)

	DeleteMountTarget(*efs.DeleteMountTargetInput) (*efs.DeleteMountTargetOutput, error)

	DeleteTagsRequest(*efs.DeleteTagsInput) (*request.Request, *efs.DeleteTagsOutput)

	DeleteTags(*efs.DeleteTagsInput) (*efs.DeleteTagsOutput, error)

	DescribeFileSystemsRequest(*efs.DescribeFileSystemsInput) (*request.Request, *efs.DescribeFileSystemsOutput)

	DescribeFileSystems(*efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error)

	DescribeMountTargetSecurityGroupsRequest(*efs.DescribeMountTargetSecurityGroupsInput) (*request.Request, *efs.DescribeMountTargetSecurityGroupsOutput)

	DescribeMountTargetSecurityGroups(*efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error)

	DescribeMountTargetsRequest(*efs.DescribeMountTargetsInput) (*request.Request, *efs.DescribeMountTargetsOutput)

	DescribeMountTargets(*efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error)

	DescribeTagsRequest(*efs.DescribeTagsInput) (*request.Request, *efs.DescribeTagsOutput)

	DescribeTags(*efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error)

	ModifyMountTargetSecurityGroupsRequest(*efs.ModifyMountTargetSecurityGroupsInput) (*request.Request, *efs.ModifyMountTargetSecurityGroupsOutput)

	ModifyMountTargetSecurityGroups(*efs.ModifyMountTargetSecurityGroupsInput) (*efs.ModifyMountTargetSecurityGroupsOutput, error)
}

var _ EFSAPI = (*efs.EFS)(nil)
