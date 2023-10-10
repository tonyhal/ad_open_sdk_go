/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus
type DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus string

// List of dmp_custom_audience_read_v2_data_custom_audience_list_delivery_status
const (
	CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus  DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus    DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE"
	CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH_DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus = "CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH"
)

// All allowed values of DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus enum
var AllowedDmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatusEnumValues = []DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus{
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_UNAVAILABLE",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUSH",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_AVAILABLE",
	"CUSTOM_AUDIENCE_DELIVERY_STATUS_NEED_PUBLISH",
}

// NewDmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatusFromValue returns a pointer to a valid DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatusFromValue(v string) (*DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus, error) {
	ev := DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus: valid values are %v", v, AllowedDmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus) IsValid() bool {
	for _, existing := range AllowedDmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_custom_audience_read_v2_data_custom_audience_list_delivery_status value
func (v DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus) Ptr() *DmpCustomAudienceReadV2DataCustomAudienceListDeliveryStatus {
	return &v
}
