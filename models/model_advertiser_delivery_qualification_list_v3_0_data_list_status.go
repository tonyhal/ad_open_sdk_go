/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserDeliveryQualificationListV30DataListStatus
type AdvertiserDeliveryQualificationListV30DataListStatus string

// List of advertiser_delivery_qualification_list_v3.0_data_list_status
const (
	STATUS_CONFIRM_AdvertiserDeliveryQualificationListV30DataListStatus         AdvertiserDeliveryQualificationListV30DataListStatus = "STATUS_CONFIRM"
	STATUS_CONFIRM_FAIL_AdvertiserDeliveryQualificationListV30DataListStatus    AdvertiserDeliveryQualificationListV30DataListStatus = "STATUS_CONFIRM_FAIL"
	STATUS_NOT_SUBMIT_AdvertiserDeliveryQualificationListV30DataListStatus      AdvertiserDeliveryQualificationListV30DataListStatus = "STATUS_NOT_SUBMIT"
	STATUS_PENDING_CONFIRM_AdvertiserDeliveryQualificationListV30DataListStatus AdvertiserDeliveryQualificationListV30DataListStatus = "STATUS_PENDING_CONFIRM"
	STATUS_WAIT_CONFIRM_AdvertiserDeliveryQualificationListV30DataListStatus    AdvertiserDeliveryQualificationListV30DataListStatus = "STATUS_WAIT_CONFIRM"
)

// All allowed values of AdvertiserDeliveryQualificationListV30DataListStatus enum
var AllowedAdvertiserDeliveryQualificationListV30DataListStatusEnumValues = []AdvertiserDeliveryQualificationListV30DataListStatus{
	"STATUS_CONFIRM",
	"STATUS_CONFIRM_FAIL",
	"STATUS_NOT_SUBMIT",
	"STATUS_PENDING_CONFIRM",
	"STATUS_WAIT_CONFIRM",
}

// NewAdvertiserDeliveryQualificationListV30DataListStatusFromValue returns a pointer to a valid AdvertiserDeliveryQualificationListV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryQualificationListV30DataListStatusFromValue(v string) (*AdvertiserDeliveryQualificationListV30DataListStatus, error) {
	ev := AdvertiserDeliveryQualificationListV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryQualificationListV30DataListStatus: valid values are %v", v, AllowedAdvertiserDeliveryQualificationListV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryQualificationListV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryQualificationListV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_qualification_list_v3.0_data_list_status value
func (v AdvertiserDeliveryQualificationListV30DataListStatus) Ptr() *AdvertiserDeliveryQualificationListV30DataListStatus {
	return &v
}
