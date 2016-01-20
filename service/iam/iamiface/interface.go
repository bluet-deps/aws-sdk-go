// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package iamiface provides an interface for the AWS Identity and Access Management.
package iamiface

import (
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/service/iam"
)

// IAMAPI is the interface type for iam.IAM.
type IAMAPI interface {
	AddClientIDToOpenIDConnectProviderRequest(*iam.AddClientIDToOpenIDConnectProviderInput) (*request.Request, *iam.AddClientIDToOpenIDConnectProviderOutput)

	AddClientIDToOpenIDConnectProvider(*iam.AddClientIDToOpenIDConnectProviderInput) (*iam.AddClientIDToOpenIDConnectProviderOutput, error)

	AddRoleToInstanceProfileRequest(*iam.AddRoleToInstanceProfileInput) (*request.Request, *iam.AddRoleToInstanceProfileOutput)

	AddRoleToInstanceProfile(*iam.AddRoleToInstanceProfileInput) (*iam.AddRoleToInstanceProfileOutput, error)

	AddUserToGroupRequest(*iam.AddUserToGroupInput) (*request.Request, *iam.AddUserToGroupOutput)

	AddUserToGroup(*iam.AddUserToGroupInput) (*iam.AddUserToGroupOutput, error)

	AttachGroupPolicyRequest(*iam.AttachGroupPolicyInput) (*request.Request, *iam.AttachGroupPolicyOutput)

	AttachGroupPolicy(*iam.AttachGroupPolicyInput) (*iam.AttachGroupPolicyOutput, error)

	AttachRolePolicyRequest(*iam.AttachRolePolicyInput) (*request.Request, *iam.AttachRolePolicyOutput)

	AttachRolePolicy(*iam.AttachRolePolicyInput) (*iam.AttachRolePolicyOutput, error)

	AttachUserPolicyRequest(*iam.AttachUserPolicyInput) (*request.Request, *iam.AttachUserPolicyOutput)

	AttachUserPolicy(*iam.AttachUserPolicyInput) (*iam.AttachUserPolicyOutput, error)

	ChangePasswordRequest(*iam.ChangePasswordInput) (*request.Request, *iam.ChangePasswordOutput)

	ChangePassword(*iam.ChangePasswordInput) (*iam.ChangePasswordOutput, error)

	CreateAccessKeyRequest(*iam.CreateAccessKeyInput) (*request.Request, *iam.CreateAccessKeyOutput)

	CreateAccessKey(*iam.CreateAccessKeyInput) (*iam.CreateAccessKeyOutput, error)

	CreateAccountAliasRequest(*iam.CreateAccountAliasInput) (*request.Request, *iam.CreateAccountAliasOutput)

	CreateAccountAlias(*iam.CreateAccountAliasInput) (*iam.CreateAccountAliasOutput, error)

	CreateGroupRequest(*iam.CreateGroupInput) (*request.Request, *iam.CreateGroupOutput)

	CreateGroup(*iam.CreateGroupInput) (*iam.CreateGroupOutput, error)

	CreateInstanceProfileRequest(*iam.CreateInstanceProfileInput) (*request.Request, *iam.CreateInstanceProfileOutput)

	CreateInstanceProfile(*iam.CreateInstanceProfileInput) (*iam.CreateInstanceProfileOutput, error)

	CreateLoginProfileRequest(*iam.CreateLoginProfileInput) (*request.Request, *iam.CreateLoginProfileOutput)

	CreateLoginProfile(*iam.CreateLoginProfileInput) (*iam.CreateLoginProfileOutput, error)

	CreateOpenIDConnectProviderRequest(*iam.CreateOpenIDConnectProviderInput) (*request.Request, *iam.CreateOpenIDConnectProviderOutput)

	CreateOpenIDConnectProvider(*iam.CreateOpenIDConnectProviderInput) (*iam.CreateOpenIDConnectProviderOutput, error)

	CreatePolicyRequest(*iam.CreatePolicyInput) (*request.Request, *iam.CreatePolicyOutput)

	CreatePolicy(*iam.CreatePolicyInput) (*iam.CreatePolicyOutput, error)

	CreatePolicyVersionRequest(*iam.CreatePolicyVersionInput) (*request.Request, *iam.CreatePolicyVersionOutput)

	CreatePolicyVersion(*iam.CreatePolicyVersionInput) (*iam.CreatePolicyVersionOutput, error)

	CreateRoleRequest(*iam.CreateRoleInput) (*request.Request, *iam.CreateRoleOutput)

	CreateRole(*iam.CreateRoleInput) (*iam.CreateRoleOutput, error)

	CreateSAMLProviderRequest(*iam.CreateSAMLProviderInput) (*request.Request, *iam.CreateSAMLProviderOutput)

	CreateSAMLProvider(*iam.CreateSAMLProviderInput) (*iam.CreateSAMLProviderOutput, error)

	CreateUserRequest(*iam.CreateUserInput) (*request.Request, *iam.CreateUserOutput)

	CreateUser(*iam.CreateUserInput) (*iam.CreateUserOutput, error)

	CreateVirtualMFADeviceRequest(*iam.CreateVirtualMFADeviceInput) (*request.Request, *iam.CreateVirtualMFADeviceOutput)

	CreateVirtualMFADevice(*iam.CreateVirtualMFADeviceInput) (*iam.CreateVirtualMFADeviceOutput, error)

	DeactivateMFADeviceRequest(*iam.DeactivateMFADeviceInput) (*request.Request, *iam.DeactivateMFADeviceOutput)

	DeactivateMFADevice(*iam.DeactivateMFADeviceInput) (*iam.DeactivateMFADeviceOutput, error)

	DeleteAccessKeyRequest(*iam.DeleteAccessKeyInput) (*request.Request, *iam.DeleteAccessKeyOutput)

	DeleteAccessKey(*iam.DeleteAccessKeyInput) (*iam.DeleteAccessKeyOutput, error)

	DeleteAccountAliasRequest(*iam.DeleteAccountAliasInput) (*request.Request, *iam.DeleteAccountAliasOutput)

	DeleteAccountAlias(*iam.DeleteAccountAliasInput) (*iam.DeleteAccountAliasOutput, error)

	DeleteAccountPasswordPolicyRequest(*iam.DeleteAccountPasswordPolicyInput) (*request.Request, *iam.DeleteAccountPasswordPolicyOutput)

	DeleteAccountPasswordPolicy(*iam.DeleteAccountPasswordPolicyInput) (*iam.DeleteAccountPasswordPolicyOutput, error)

	DeleteGroupRequest(*iam.DeleteGroupInput) (*request.Request, *iam.DeleteGroupOutput)

	DeleteGroup(*iam.DeleteGroupInput) (*iam.DeleteGroupOutput, error)

	DeleteGroupPolicyRequest(*iam.DeleteGroupPolicyInput) (*request.Request, *iam.DeleteGroupPolicyOutput)

	DeleteGroupPolicy(*iam.DeleteGroupPolicyInput) (*iam.DeleteGroupPolicyOutput, error)

	DeleteInstanceProfileRequest(*iam.DeleteInstanceProfileInput) (*request.Request, *iam.DeleteInstanceProfileOutput)

	DeleteInstanceProfile(*iam.DeleteInstanceProfileInput) (*iam.DeleteInstanceProfileOutput, error)

	DeleteLoginProfileRequest(*iam.DeleteLoginProfileInput) (*request.Request, *iam.DeleteLoginProfileOutput)

	DeleteLoginProfile(*iam.DeleteLoginProfileInput) (*iam.DeleteLoginProfileOutput, error)

	DeleteOpenIDConnectProviderRequest(*iam.DeleteOpenIDConnectProviderInput) (*request.Request, *iam.DeleteOpenIDConnectProviderOutput)

	DeleteOpenIDConnectProvider(*iam.DeleteOpenIDConnectProviderInput) (*iam.DeleteOpenIDConnectProviderOutput, error)

	DeletePolicyRequest(*iam.DeletePolicyInput) (*request.Request, *iam.DeletePolicyOutput)

	DeletePolicy(*iam.DeletePolicyInput) (*iam.DeletePolicyOutput, error)

	DeletePolicyVersionRequest(*iam.DeletePolicyVersionInput) (*request.Request, *iam.DeletePolicyVersionOutput)

	DeletePolicyVersion(*iam.DeletePolicyVersionInput) (*iam.DeletePolicyVersionOutput, error)

	DeleteRoleRequest(*iam.DeleteRoleInput) (*request.Request, *iam.DeleteRoleOutput)

	DeleteRole(*iam.DeleteRoleInput) (*iam.DeleteRoleOutput, error)

	DeleteRolePolicyRequest(*iam.DeleteRolePolicyInput) (*request.Request, *iam.DeleteRolePolicyOutput)

	DeleteRolePolicy(*iam.DeleteRolePolicyInput) (*iam.DeleteRolePolicyOutput, error)

	DeleteSAMLProviderRequest(*iam.DeleteSAMLProviderInput) (*request.Request, *iam.DeleteSAMLProviderOutput)

	DeleteSAMLProvider(*iam.DeleteSAMLProviderInput) (*iam.DeleteSAMLProviderOutput, error)

	DeleteSSHPublicKeyRequest(*iam.DeleteSSHPublicKeyInput) (*request.Request, *iam.DeleteSSHPublicKeyOutput)

	DeleteSSHPublicKey(*iam.DeleteSSHPublicKeyInput) (*iam.DeleteSSHPublicKeyOutput, error)

	DeleteServerCertificateRequest(*iam.DeleteServerCertificateInput) (*request.Request, *iam.DeleteServerCertificateOutput)

	DeleteServerCertificate(*iam.DeleteServerCertificateInput) (*iam.DeleteServerCertificateOutput, error)

	DeleteSigningCertificateRequest(*iam.DeleteSigningCertificateInput) (*request.Request, *iam.DeleteSigningCertificateOutput)

	DeleteSigningCertificate(*iam.DeleteSigningCertificateInput) (*iam.DeleteSigningCertificateOutput, error)

	DeleteUserRequest(*iam.DeleteUserInput) (*request.Request, *iam.DeleteUserOutput)

	DeleteUser(*iam.DeleteUserInput) (*iam.DeleteUserOutput, error)

	DeleteUserPolicyRequest(*iam.DeleteUserPolicyInput) (*request.Request, *iam.DeleteUserPolicyOutput)

	DeleteUserPolicy(*iam.DeleteUserPolicyInput) (*iam.DeleteUserPolicyOutput, error)

	DeleteVirtualMFADeviceRequest(*iam.DeleteVirtualMFADeviceInput) (*request.Request, *iam.DeleteVirtualMFADeviceOutput)

	DeleteVirtualMFADevice(*iam.DeleteVirtualMFADeviceInput) (*iam.DeleteVirtualMFADeviceOutput, error)

	DetachGroupPolicyRequest(*iam.DetachGroupPolicyInput) (*request.Request, *iam.DetachGroupPolicyOutput)

	DetachGroupPolicy(*iam.DetachGroupPolicyInput) (*iam.DetachGroupPolicyOutput, error)

	DetachRolePolicyRequest(*iam.DetachRolePolicyInput) (*request.Request, *iam.DetachRolePolicyOutput)

	DetachRolePolicy(*iam.DetachRolePolicyInput) (*iam.DetachRolePolicyOutput, error)

	DetachUserPolicyRequest(*iam.DetachUserPolicyInput) (*request.Request, *iam.DetachUserPolicyOutput)

	DetachUserPolicy(*iam.DetachUserPolicyInput) (*iam.DetachUserPolicyOutput, error)

	EnableMFADeviceRequest(*iam.EnableMFADeviceInput) (*request.Request, *iam.EnableMFADeviceOutput)

	EnableMFADevice(*iam.EnableMFADeviceInput) (*iam.EnableMFADeviceOutput, error)

	GenerateCredentialReportRequest(*iam.GenerateCredentialReportInput) (*request.Request, *iam.GenerateCredentialReportOutput)

	GenerateCredentialReport(*iam.GenerateCredentialReportInput) (*iam.GenerateCredentialReportOutput, error)

	GetAccessKeyLastUsedRequest(*iam.GetAccessKeyLastUsedInput) (*request.Request, *iam.GetAccessKeyLastUsedOutput)

	GetAccessKeyLastUsed(*iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error)

	GetAccountAuthorizationDetailsRequest(*iam.GetAccountAuthorizationDetailsInput) (*request.Request, *iam.GetAccountAuthorizationDetailsOutput)

	GetAccountAuthorizationDetails(*iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error)

	GetAccountAuthorizationDetailsPages(*iam.GetAccountAuthorizationDetailsInput, func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool) error

	GetAccountPasswordPolicyRequest(*iam.GetAccountPasswordPolicyInput) (*request.Request, *iam.GetAccountPasswordPolicyOutput)

	GetAccountPasswordPolicy(*iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error)

	GetAccountSummaryRequest(*iam.GetAccountSummaryInput) (*request.Request, *iam.GetAccountSummaryOutput)

	GetAccountSummary(*iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error)

	GetContextKeysForCustomPolicyRequest(*iam.GetContextKeysForCustomPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)

	GetContextKeysForCustomPolicy(*iam.GetContextKeysForCustomPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)

	GetContextKeysForPrincipalPolicyRequest(*iam.GetContextKeysForPrincipalPolicyInput) (*request.Request, *iam.GetContextKeysForPolicyResponse)

	GetContextKeysForPrincipalPolicy(*iam.GetContextKeysForPrincipalPolicyInput) (*iam.GetContextKeysForPolicyResponse, error)

	GetCredentialReportRequest(*iam.GetCredentialReportInput) (*request.Request, *iam.GetCredentialReportOutput)

	GetCredentialReport(*iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error)

	GetGroupRequest(*iam.GetGroupInput) (*request.Request, *iam.GetGroupOutput)

	GetGroup(*iam.GetGroupInput) (*iam.GetGroupOutput, error)

	GetGroupPages(*iam.GetGroupInput, func(*iam.GetGroupOutput, bool) bool) error

	GetGroupPolicyRequest(*iam.GetGroupPolicyInput) (*request.Request, *iam.GetGroupPolicyOutput)

	GetGroupPolicy(*iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error)

	GetInstanceProfileRequest(*iam.GetInstanceProfileInput) (*request.Request, *iam.GetInstanceProfileOutput)

	GetInstanceProfile(*iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error)

	GetLoginProfileRequest(*iam.GetLoginProfileInput) (*request.Request, *iam.GetLoginProfileOutput)

	GetLoginProfile(*iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error)

	GetOpenIDConnectProviderRequest(*iam.GetOpenIDConnectProviderInput) (*request.Request, *iam.GetOpenIDConnectProviderOutput)

	GetOpenIDConnectProvider(*iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error)

	GetPolicyRequest(*iam.GetPolicyInput) (*request.Request, *iam.GetPolicyOutput)

	GetPolicy(*iam.GetPolicyInput) (*iam.GetPolicyOutput, error)

	GetPolicyVersionRequest(*iam.GetPolicyVersionInput) (*request.Request, *iam.GetPolicyVersionOutput)

	GetPolicyVersion(*iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error)

	GetRoleRequest(*iam.GetRoleInput) (*request.Request, *iam.GetRoleOutput)

	GetRole(*iam.GetRoleInput) (*iam.GetRoleOutput, error)

	GetRolePolicyRequest(*iam.GetRolePolicyInput) (*request.Request, *iam.GetRolePolicyOutput)

	GetRolePolicy(*iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error)

	GetSAMLProviderRequest(*iam.GetSAMLProviderInput) (*request.Request, *iam.GetSAMLProviderOutput)

	GetSAMLProvider(*iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error)

	GetSSHPublicKeyRequest(*iam.GetSSHPublicKeyInput) (*request.Request, *iam.GetSSHPublicKeyOutput)

	GetSSHPublicKey(*iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error)

	GetServerCertificateRequest(*iam.GetServerCertificateInput) (*request.Request, *iam.GetServerCertificateOutput)

	GetServerCertificate(*iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error)

	GetUserRequest(*iam.GetUserInput) (*request.Request, *iam.GetUserOutput)

	GetUser(*iam.GetUserInput) (*iam.GetUserOutput, error)

	GetUserPolicyRequest(*iam.GetUserPolicyInput) (*request.Request, *iam.GetUserPolicyOutput)

	GetUserPolicy(*iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error)

	ListAccessKeysRequest(*iam.ListAccessKeysInput) (*request.Request, *iam.ListAccessKeysOutput)

	ListAccessKeys(*iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)

	ListAccessKeysPages(*iam.ListAccessKeysInput, func(*iam.ListAccessKeysOutput, bool) bool) error

	ListAccountAliasesRequest(*iam.ListAccountAliasesInput) (*request.Request, *iam.ListAccountAliasesOutput)

	ListAccountAliases(*iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error)

	ListAccountAliasesPages(*iam.ListAccountAliasesInput, func(*iam.ListAccountAliasesOutput, bool) bool) error

	ListAttachedGroupPoliciesRequest(*iam.ListAttachedGroupPoliciesInput) (*request.Request, *iam.ListAttachedGroupPoliciesOutput)

	ListAttachedGroupPolicies(*iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error)

	ListAttachedGroupPoliciesPages(*iam.ListAttachedGroupPoliciesInput, func(*iam.ListAttachedGroupPoliciesOutput, bool) bool) error

	ListAttachedRolePoliciesRequest(*iam.ListAttachedRolePoliciesInput) (*request.Request, *iam.ListAttachedRolePoliciesOutput)

	ListAttachedRolePolicies(*iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error)

	ListAttachedRolePoliciesPages(*iam.ListAttachedRolePoliciesInput, func(*iam.ListAttachedRolePoliciesOutput, bool) bool) error

	ListAttachedUserPoliciesRequest(*iam.ListAttachedUserPoliciesInput) (*request.Request, *iam.ListAttachedUserPoliciesOutput)

	ListAttachedUserPolicies(*iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error)

	ListAttachedUserPoliciesPages(*iam.ListAttachedUserPoliciesInput, func(*iam.ListAttachedUserPoliciesOutput, bool) bool) error

	ListEntitiesForPolicyRequest(*iam.ListEntitiesForPolicyInput) (*request.Request, *iam.ListEntitiesForPolicyOutput)

	ListEntitiesForPolicy(*iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error)

	ListEntitiesForPolicyPages(*iam.ListEntitiesForPolicyInput, func(*iam.ListEntitiesForPolicyOutput, bool) bool) error

	ListGroupPoliciesRequest(*iam.ListGroupPoliciesInput) (*request.Request, *iam.ListGroupPoliciesOutput)

	ListGroupPolicies(*iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error)

	ListGroupPoliciesPages(*iam.ListGroupPoliciesInput, func(*iam.ListGroupPoliciesOutput, bool) bool) error

	ListGroupsRequest(*iam.ListGroupsInput) (*request.Request, *iam.ListGroupsOutput)

	ListGroups(*iam.ListGroupsInput) (*iam.ListGroupsOutput, error)

	ListGroupsPages(*iam.ListGroupsInput, func(*iam.ListGroupsOutput, bool) bool) error

	ListGroupsForUserRequest(*iam.ListGroupsForUserInput) (*request.Request, *iam.ListGroupsForUserOutput)

	ListGroupsForUser(*iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error)

	ListGroupsForUserPages(*iam.ListGroupsForUserInput, func(*iam.ListGroupsForUserOutput, bool) bool) error

	ListInstanceProfilesRequest(*iam.ListInstanceProfilesInput) (*request.Request, *iam.ListInstanceProfilesOutput)

	ListInstanceProfiles(*iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error)

	ListInstanceProfilesPages(*iam.ListInstanceProfilesInput, func(*iam.ListInstanceProfilesOutput, bool) bool) error

	ListInstanceProfilesForRoleRequest(*iam.ListInstanceProfilesForRoleInput) (*request.Request, *iam.ListInstanceProfilesForRoleOutput)

	ListInstanceProfilesForRole(*iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error)

	ListInstanceProfilesForRolePages(*iam.ListInstanceProfilesForRoleInput, func(*iam.ListInstanceProfilesForRoleOutput, bool) bool) error

	ListMFADevicesRequest(*iam.ListMFADevicesInput) (*request.Request, *iam.ListMFADevicesOutput)

	ListMFADevices(*iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error)

	ListMFADevicesPages(*iam.ListMFADevicesInput, func(*iam.ListMFADevicesOutput, bool) bool) error

	ListOpenIDConnectProvidersRequest(*iam.ListOpenIDConnectProvidersInput) (*request.Request, *iam.ListOpenIDConnectProvidersOutput)

	ListOpenIDConnectProviders(*iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error)

	ListPoliciesRequest(*iam.ListPoliciesInput) (*request.Request, *iam.ListPoliciesOutput)

	ListPolicies(*iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error)

	ListPoliciesPages(*iam.ListPoliciesInput, func(*iam.ListPoliciesOutput, bool) bool) error

	ListPolicyVersionsRequest(*iam.ListPolicyVersionsInput) (*request.Request, *iam.ListPolicyVersionsOutput)

	ListPolicyVersions(*iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error)

	ListRolePoliciesRequest(*iam.ListRolePoliciesInput) (*request.Request, *iam.ListRolePoliciesOutput)

	ListRolePolicies(*iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error)

	ListRolePoliciesPages(*iam.ListRolePoliciesInput, func(*iam.ListRolePoliciesOutput, bool) bool) error

	ListRolesRequest(*iam.ListRolesInput) (*request.Request, *iam.ListRolesOutput)

	ListRoles(*iam.ListRolesInput) (*iam.ListRolesOutput, error)

	ListRolesPages(*iam.ListRolesInput, func(*iam.ListRolesOutput, bool) bool) error

	ListSAMLProvidersRequest(*iam.ListSAMLProvidersInput) (*request.Request, *iam.ListSAMLProvidersOutput)

	ListSAMLProviders(*iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error)

	ListSSHPublicKeysRequest(*iam.ListSSHPublicKeysInput) (*request.Request, *iam.ListSSHPublicKeysOutput)

	ListSSHPublicKeys(*iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error)

	ListServerCertificatesRequest(*iam.ListServerCertificatesInput) (*request.Request, *iam.ListServerCertificatesOutput)

	ListServerCertificates(*iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error)

	ListServerCertificatesPages(*iam.ListServerCertificatesInput, func(*iam.ListServerCertificatesOutput, bool) bool) error

	ListSigningCertificatesRequest(*iam.ListSigningCertificatesInput) (*request.Request, *iam.ListSigningCertificatesOutput)

	ListSigningCertificates(*iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error)

	ListSigningCertificatesPages(*iam.ListSigningCertificatesInput, func(*iam.ListSigningCertificatesOutput, bool) bool) error

	ListUserPoliciesRequest(*iam.ListUserPoliciesInput) (*request.Request, *iam.ListUserPoliciesOutput)

	ListUserPolicies(*iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error)

	ListUserPoliciesPages(*iam.ListUserPoliciesInput, func(*iam.ListUserPoliciesOutput, bool) bool) error

	ListUsersRequest(*iam.ListUsersInput) (*request.Request, *iam.ListUsersOutput)

	ListUsers(*iam.ListUsersInput) (*iam.ListUsersOutput, error)

	ListUsersPages(*iam.ListUsersInput, func(*iam.ListUsersOutput, bool) bool) error

	ListVirtualMFADevicesRequest(*iam.ListVirtualMFADevicesInput) (*request.Request, *iam.ListVirtualMFADevicesOutput)

	ListVirtualMFADevices(*iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error)

	ListVirtualMFADevicesPages(*iam.ListVirtualMFADevicesInput, func(*iam.ListVirtualMFADevicesOutput, bool) bool) error

	PutGroupPolicyRequest(*iam.PutGroupPolicyInput) (*request.Request, *iam.PutGroupPolicyOutput)

	PutGroupPolicy(*iam.PutGroupPolicyInput) (*iam.PutGroupPolicyOutput, error)

	PutRolePolicyRequest(*iam.PutRolePolicyInput) (*request.Request, *iam.PutRolePolicyOutput)

	PutRolePolicy(*iam.PutRolePolicyInput) (*iam.PutRolePolicyOutput, error)

	PutUserPolicyRequest(*iam.PutUserPolicyInput) (*request.Request, *iam.PutUserPolicyOutput)

	PutUserPolicy(*iam.PutUserPolicyInput) (*iam.PutUserPolicyOutput, error)

	RemoveClientIDFromOpenIDConnectProviderRequest(*iam.RemoveClientIDFromOpenIDConnectProviderInput) (*request.Request, *iam.RemoveClientIDFromOpenIDConnectProviderOutput)

	RemoveClientIDFromOpenIDConnectProvider(*iam.RemoveClientIDFromOpenIDConnectProviderInput) (*iam.RemoveClientIDFromOpenIDConnectProviderOutput, error)

	RemoveRoleFromInstanceProfileRequest(*iam.RemoveRoleFromInstanceProfileInput) (*request.Request, *iam.RemoveRoleFromInstanceProfileOutput)

	RemoveRoleFromInstanceProfile(*iam.RemoveRoleFromInstanceProfileInput) (*iam.RemoveRoleFromInstanceProfileOutput, error)

	RemoveUserFromGroupRequest(*iam.RemoveUserFromGroupInput) (*request.Request, *iam.RemoveUserFromGroupOutput)

	RemoveUserFromGroup(*iam.RemoveUserFromGroupInput) (*iam.RemoveUserFromGroupOutput, error)

	ResyncMFADeviceRequest(*iam.ResyncMFADeviceInput) (*request.Request, *iam.ResyncMFADeviceOutput)

	ResyncMFADevice(*iam.ResyncMFADeviceInput) (*iam.ResyncMFADeviceOutput, error)

	SetDefaultPolicyVersionRequest(*iam.SetDefaultPolicyVersionInput) (*request.Request, *iam.SetDefaultPolicyVersionOutput)

	SetDefaultPolicyVersion(*iam.SetDefaultPolicyVersionInput) (*iam.SetDefaultPolicyVersionOutput, error)

	SimulateCustomPolicyRequest(*iam.SimulateCustomPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)

	SimulateCustomPolicy(*iam.SimulateCustomPolicyInput) (*iam.SimulatePolicyResponse, error)

	SimulatePrincipalPolicyRequest(*iam.SimulatePrincipalPolicyInput) (*request.Request, *iam.SimulatePolicyResponse)

	SimulatePrincipalPolicy(*iam.SimulatePrincipalPolicyInput) (*iam.SimulatePolicyResponse, error)

	UpdateAccessKeyRequest(*iam.UpdateAccessKeyInput) (*request.Request, *iam.UpdateAccessKeyOutput)

	UpdateAccessKey(*iam.UpdateAccessKeyInput) (*iam.UpdateAccessKeyOutput, error)

	UpdateAccountPasswordPolicyRequest(*iam.UpdateAccountPasswordPolicyInput) (*request.Request, *iam.UpdateAccountPasswordPolicyOutput)

	UpdateAccountPasswordPolicy(*iam.UpdateAccountPasswordPolicyInput) (*iam.UpdateAccountPasswordPolicyOutput, error)

	UpdateAssumeRolePolicyRequest(*iam.UpdateAssumeRolePolicyInput) (*request.Request, *iam.UpdateAssumeRolePolicyOutput)

	UpdateAssumeRolePolicy(*iam.UpdateAssumeRolePolicyInput) (*iam.UpdateAssumeRolePolicyOutput, error)

	UpdateGroupRequest(*iam.UpdateGroupInput) (*request.Request, *iam.UpdateGroupOutput)

	UpdateGroup(*iam.UpdateGroupInput) (*iam.UpdateGroupOutput, error)

	UpdateLoginProfileRequest(*iam.UpdateLoginProfileInput) (*request.Request, *iam.UpdateLoginProfileOutput)

	UpdateLoginProfile(*iam.UpdateLoginProfileInput) (*iam.UpdateLoginProfileOutput, error)

	UpdateOpenIDConnectProviderThumbprintRequest(*iam.UpdateOpenIDConnectProviderThumbprintInput) (*request.Request, *iam.UpdateOpenIDConnectProviderThumbprintOutput)

	UpdateOpenIDConnectProviderThumbprint(*iam.UpdateOpenIDConnectProviderThumbprintInput) (*iam.UpdateOpenIDConnectProviderThumbprintOutput, error)

	UpdateSAMLProviderRequest(*iam.UpdateSAMLProviderInput) (*request.Request, *iam.UpdateSAMLProviderOutput)

	UpdateSAMLProvider(*iam.UpdateSAMLProviderInput) (*iam.UpdateSAMLProviderOutput, error)

	UpdateSSHPublicKeyRequest(*iam.UpdateSSHPublicKeyInput) (*request.Request, *iam.UpdateSSHPublicKeyOutput)

	UpdateSSHPublicKey(*iam.UpdateSSHPublicKeyInput) (*iam.UpdateSSHPublicKeyOutput, error)

	UpdateServerCertificateRequest(*iam.UpdateServerCertificateInput) (*request.Request, *iam.UpdateServerCertificateOutput)

	UpdateServerCertificate(*iam.UpdateServerCertificateInput) (*iam.UpdateServerCertificateOutput, error)

	UpdateSigningCertificateRequest(*iam.UpdateSigningCertificateInput) (*request.Request, *iam.UpdateSigningCertificateOutput)

	UpdateSigningCertificate(*iam.UpdateSigningCertificateInput) (*iam.UpdateSigningCertificateOutput, error)

	UpdateUserRequest(*iam.UpdateUserInput) (*request.Request, *iam.UpdateUserOutput)

	UpdateUser(*iam.UpdateUserInput) (*iam.UpdateUserOutput, error)

	UploadSSHPublicKeyRequest(*iam.UploadSSHPublicKeyInput) (*request.Request, *iam.UploadSSHPublicKeyOutput)

	UploadSSHPublicKey(*iam.UploadSSHPublicKeyInput) (*iam.UploadSSHPublicKeyOutput, error)

	UploadServerCertificateRequest(*iam.UploadServerCertificateInput) (*request.Request, *iam.UploadServerCertificateOutput)

	UploadServerCertificate(*iam.UploadServerCertificateInput) (*iam.UploadServerCertificateOutput, error)

	UploadSigningCertificateRequest(*iam.UploadSigningCertificateInput) (*request.Request, *iam.UploadSigningCertificateOutput)

	UploadSigningCertificate(*iam.UploadSigningCertificateInput) (*iam.UploadSigningCertificateOutput, error)
}

var _ IAMAPI = (*iam.IAM)(nil)
