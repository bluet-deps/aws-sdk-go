// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mobileanalyticsiface provides an interface for the Amazon Mobile Analytics.
package mobileanalyticsiface

import (
	"bluet-deps/aws-sdk-go/aws/request"
	"bluet-deps/aws-sdk-go/service/mobileanalytics"
)

// MobileAnalyticsAPI is the interface type for mobileanalytics.MobileAnalytics.
type MobileAnalyticsAPI interface {
	PutEventsRequest(*mobileanalytics.PutEventsInput) (*request.Request, *mobileanalytics.PutEventsOutput)

	PutEvents(*mobileanalytics.PutEventsInput) (*mobileanalytics.PutEventsOutput, error)
}

var _ MobileAnalyticsAPI = (*mobileanalytics.MobileAnalytics)(nil)
