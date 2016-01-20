// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package directconnectiface provides an interface for the AWS Direct Connect.
package directconnectiface

import (
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/service/directconnect"
)

// DirectConnectAPI is the interface type for directconnect.DirectConnect.
type DirectConnectAPI interface {
	AllocateConnectionOnInterconnectRequest(*directconnect.AllocateConnectionOnInterconnectInput) (*request.Request, *directconnect.Connection)

	AllocateConnectionOnInterconnect(*directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error)

	AllocatePrivateVirtualInterfaceRequest(*directconnect.AllocatePrivateVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	AllocatePrivateVirtualInterface(*directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	AllocatePublicVirtualInterfaceRequest(*directconnect.AllocatePublicVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	AllocatePublicVirtualInterface(*directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	ConfirmConnectionRequest(*directconnect.ConfirmConnectionInput) (*request.Request, *directconnect.ConfirmConnectionOutput)

	ConfirmConnection(*directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error)

	ConfirmPrivateVirtualInterfaceRequest(*directconnect.ConfirmPrivateVirtualInterfaceInput) (*request.Request, *directconnect.ConfirmPrivateVirtualInterfaceOutput)

	ConfirmPrivateVirtualInterface(*directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error)

	ConfirmPublicVirtualInterfaceRequest(*directconnect.ConfirmPublicVirtualInterfaceInput) (*request.Request, *directconnect.ConfirmPublicVirtualInterfaceOutput)

	ConfirmPublicVirtualInterface(*directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error)

	CreateConnectionRequest(*directconnect.CreateConnectionInput) (*request.Request, *directconnect.Connection)

	CreateConnection(*directconnect.CreateConnectionInput) (*directconnect.Connection, error)

	CreateInterconnectRequest(*directconnect.CreateInterconnectInput) (*request.Request, *directconnect.Interconnect)

	CreateInterconnect(*directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error)

	CreatePrivateVirtualInterfaceRequest(*directconnect.CreatePrivateVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	CreatePrivateVirtualInterface(*directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	CreatePublicVirtualInterfaceRequest(*directconnect.CreatePublicVirtualInterfaceInput) (*request.Request, *directconnect.VirtualInterface)

	CreatePublicVirtualInterface(*directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error)

	DeleteConnectionRequest(*directconnect.DeleteConnectionInput) (*request.Request, *directconnect.Connection)

	DeleteConnection(*directconnect.DeleteConnectionInput) (*directconnect.Connection, error)

	DeleteInterconnectRequest(*directconnect.DeleteInterconnectInput) (*request.Request, *directconnect.DeleteInterconnectOutput)

	DeleteInterconnect(*directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error)

	DeleteVirtualInterfaceRequest(*directconnect.DeleteVirtualInterfaceInput) (*request.Request, *directconnect.DeleteVirtualInterfaceOutput)

	DeleteVirtualInterface(*directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error)

	DescribeConnectionsRequest(*directconnect.DescribeConnectionsInput) (*request.Request, *directconnect.Connections)

	DescribeConnections(*directconnect.DescribeConnectionsInput) (*directconnect.Connections, error)

	DescribeConnectionsOnInterconnectRequest(*directconnect.DescribeConnectionsOnInterconnectInput) (*request.Request, *directconnect.Connections)

	DescribeConnectionsOnInterconnect(*directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error)

	DescribeInterconnectsRequest(*directconnect.DescribeInterconnectsInput) (*request.Request, *directconnect.DescribeInterconnectsOutput)

	DescribeInterconnects(*directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error)

	DescribeLocationsRequest(*directconnect.DescribeLocationsInput) (*request.Request, *directconnect.DescribeLocationsOutput)

	DescribeLocations(*directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error)

	DescribeVirtualGatewaysRequest(*directconnect.DescribeVirtualGatewaysInput) (*request.Request, *directconnect.DescribeVirtualGatewaysOutput)

	DescribeVirtualGateways(*directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error)

	DescribeVirtualInterfacesRequest(*directconnect.DescribeVirtualInterfacesInput) (*request.Request, *directconnect.DescribeVirtualInterfacesOutput)

	DescribeVirtualInterfaces(*directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error)
}

var _ DirectConnectAPI = (*directconnect.DirectConnect)(nil)
