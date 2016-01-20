// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package supportiface provides an interface for the AWS Support.
package supportiface

import (
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/service/support"
)

// SupportAPI is the interface type for support.Support.
type SupportAPI interface {
	AddAttachmentsToSetRequest(*support.AddAttachmentsToSetInput) (*request.Request, *support.AddAttachmentsToSetOutput)

	AddAttachmentsToSet(*support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error)

	AddCommunicationToCaseRequest(*support.AddCommunicationToCaseInput) (*request.Request, *support.AddCommunicationToCaseOutput)

	AddCommunicationToCase(*support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error)

	CreateCaseRequest(*support.CreateCaseInput) (*request.Request, *support.CreateCaseOutput)

	CreateCase(*support.CreateCaseInput) (*support.CreateCaseOutput, error)

	DescribeAttachmentRequest(*support.DescribeAttachmentInput) (*request.Request, *support.DescribeAttachmentOutput)

	DescribeAttachment(*support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error)

	DescribeCasesRequest(*support.DescribeCasesInput) (*request.Request, *support.DescribeCasesOutput)

	DescribeCases(*support.DescribeCasesInput) (*support.DescribeCasesOutput, error)

	DescribeCasesPages(*support.DescribeCasesInput, func(*support.DescribeCasesOutput, bool) bool) error

	DescribeCommunicationsRequest(*support.DescribeCommunicationsInput) (*request.Request, *support.DescribeCommunicationsOutput)

	DescribeCommunications(*support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error)

	DescribeCommunicationsPages(*support.DescribeCommunicationsInput, func(*support.DescribeCommunicationsOutput, bool) bool) error

	DescribeServicesRequest(*support.DescribeServicesInput) (*request.Request, *support.DescribeServicesOutput)

	DescribeServices(*support.DescribeServicesInput) (*support.DescribeServicesOutput, error)

	DescribeSeverityLevelsRequest(*support.DescribeSeverityLevelsInput) (*request.Request, *support.DescribeSeverityLevelsOutput)

	DescribeSeverityLevels(*support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error)

	DescribeTrustedAdvisorCheckRefreshStatusesRequest(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput)

	DescribeTrustedAdvisorCheckRefreshStatuses(*support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error)

	DescribeTrustedAdvisorCheckResultRequest(*support.DescribeTrustedAdvisorCheckResultInput) (*request.Request, *support.DescribeTrustedAdvisorCheckResultOutput)

	DescribeTrustedAdvisorCheckResult(*support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error)

	DescribeTrustedAdvisorCheckSummariesRequest(*support.DescribeTrustedAdvisorCheckSummariesInput) (*request.Request, *support.DescribeTrustedAdvisorCheckSummariesOutput)

	DescribeTrustedAdvisorCheckSummaries(*support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error)

	DescribeTrustedAdvisorChecksRequest(*support.DescribeTrustedAdvisorChecksInput) (*request.Request, *support.DescribeTrustedAdvisorChecksOutput)

	DescribeTrustedAdvisorChecks(*support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error)

	RefreshTrustedAdvisorCheckRequest(*support.RefreshTrustedAdvisorCheckInput) (*request.Request, *support.RefreshTrustedAdvisorCheckOutput)

	RefreshTrustedAdvisorCheck(*support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error)

	ResolveCaseRequest(*support.ResolveCaseInput) (*request.Request, *support.ResolveCaseOutput)

	ResolveCase(*support.ResolveCaseInput) (*support.ResolveCaseOutput, error)
}

var _ SupportAPI = (*support.Support)(nil)
