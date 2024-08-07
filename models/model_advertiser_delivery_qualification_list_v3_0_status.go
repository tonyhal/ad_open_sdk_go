/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserDeliveryQualificationListV30Status
type AdvertiserDeliveryQualificationListV30Status string

// List of advertiser_delivery_qualification_list_v3.0_status
const (
	STATUS_CONFIRM_AdvertiserDeliveryQualificationListV30Status         AdvertiserDeliveryQualificationListV30Status = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserDeliveryQualificationListV30Status    AdvertiserDeliveryQualificationListV30Status = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserDeliveryQualificationListV30Status      AdvertiserDeliveryQualificationListV30Status = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserDeliveryQualificationListV30Status AdvertiserDeliveryQualificationListV30Status = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserDeliveryQualificationListV30Status    AdvertiserDeliveryQualificationListV30Status = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserDeliveryQualificationListV30Status enum
var AllowedAdvertiserDeliveryQualificationListV30StatusEnumValues = []AdvertiserDeliveryQualificationListV30Status{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserDeliveryQualificationListV30StatusFromValue returns a pointer to a valid AdvertiserDeliveryQualificationListV30Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryQualificationListV30StatusFromValue(v string) (*AdvertiserDeliveryQualificationListV30Status, error) {
	ev := AdvertiserDeliveryQualificationListV30Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryQualificationListV30Status: valid values are %v", v, AllowedAdvertiserDeliveryQualificationListV30StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryQualificationListV30Status) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryQualificationListV30StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_qualification_list_v3.0_status value
func (v AdvertiserDeliveryQualificationListV30Status) Ptr() *AdvertiserDeliveryQualificationListV30Status {
	return &v
}
