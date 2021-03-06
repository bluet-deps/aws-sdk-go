// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package iotdataplaneiface provides an interface for the AWS IoT Data Plane.
package iotdataplaneiface

import (
	"github.com/bluet-deps/aws-sdk-go/aws/request"
	"github.com/bluet-deps/aws-sdk-go/service/iotdataplane"
)

// IoTDataPlaneAPI is the interface type for iotdataplane.IoTDataPlane.
type IoTDataPlaneAPI interface {
	DeleteThingShadowRequest(*iotdataplane.DeleteThingShadowInput) (*request.Request, *iotdataplane.DeleteThingShadowOutput)

	DeleteThingShadow(*iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error)

	GetThingShadowRequest(*iotdataplane.GetThingShadowInput) (*request.Request, *iotdataplane.GetThingShadowOutput)

	GetThingShadow(*iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error)

	PublishRequest(*iotdataplane.PublishInput) (*request.Request, *iotdataplane.PublishOutput)

	Publish(*iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error)

	UpdateThingShadowRequest(*iotdataplane.UpdateThingShadowInput) (*request.Request, *iotdataplane.UpdateThingShadowOutput)

	UpdateThingShadow(*iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error)
}

var _ IoTDataPlaneAPI = (*iotdataplane.IoTDataPlane)(nil)
