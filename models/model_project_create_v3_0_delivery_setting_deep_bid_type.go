/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliverySettingDeepBidType
type ProjectCreateV30DeliverySettingDeepBidType string

// List of project_create_v3.0_delivery_setting_deep_bid_type
const (
	ALI_FL_ProjectCreateV30DeliverySettingDeepBidType                       ProjectCreateV30DeliverySettingDeepBidType = "ALI_FL"
	AUTO_MIN_SECOND_STAGE_ProjectCreateV30DeliverySettingDeepBidType        ProjectCreateV30DeliverySettingDeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_ProjectCreateV30DeliverySettingDeepBidType               ProjectCreateV30DeliverySettingDeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_ProjectCreateV30DeliverySettingDeepBidType             ProjectCreateV30DeliverySettingDeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_ProjectCreateV30DeliverySettingDeepBidType                 ProjectCreateV30DeliverySettingDeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_ProjectCreateV30DeliverySettingDeepBidType              ProjectCreateV30DeliverySettingDeepBidType = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_ProjectCreateV30DeliverySettingDeepBidType ProjectCreateV30DeliverySettingDeepBidType = "DEEP_BID_TYPE_RETENTION_DAYS"
	MIN_SECOND_STAGE_ProjectCreateV30DeliverySettingDeepBidType             ProjectCreateV30DeliverySettingDeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_ProjectCreateV30DeliverySettingDeepBidType          ProjectCreateV30DeliverySettingDeepBidType = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_ProjectCreateV30DeliverySettingDeepBidType              ProjectCreateV30DeliverySettingDeepBidType = "ROI_COEFFICIENT"
	ROI_PACING_ProjectCreateV30DeliverySettingDeepBidType                   ProjectCreateV30DeliverySettingDeepBidType = "ROI_PACING"
	SMARTBID_ProjectCreateV30DeliverySettingDeepBidType                     ProjectCreateV30DeliverySettingDeepBidType = "SMARTBID"
	SOCIAL_ROI_ProjectCreateV30DeliverySettingDeepBidType                   ProjectCreateV30DeliverySettingDeepBidType = "SOCIAL_ROI"
	FORM_BID_ProjectCreateV30DeliverySettingDeepBidType                     ProjectCreateV30DeliverySettingDeepBidType = "FORM_BID"
	PHONE_CONNECT_BID_ProjectCreateV30DeliverySettingDeepBidType            ProjectCreateV30DeliverySettingDeepBidType = "PHONE_CONNECT_BID"
	ROI_DIRECT_MAIL_ProjectCreateV30DeliverySettingDeepBidType              ProjectCreateV30DeliverySettingDeepBidType = "ROI_DIRECT_MAIL"
)

// All allowed values of ProjectCreateV30DeliverySettingDeepBidType enum
var AllowedProjectCreateV30DeliverySettingDeepBidTypeEnumValues = []ProjectCreateV30DeliverySettingDeepBidType{
	"ALI_FL",
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"ROI_COEFFICIENT",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
	"FORM_BID",
	"PHONE_CONNECT_BID",
	"ROI_DIRECT_MAIL",
}

// NewProjectCreateV30DeliverySettingDeepBidTypeFromValue returns a pointer to a valid ProjectCreateV30DeliverySettingDeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliverySettingDeepBidTypeFromValue(v string) (*ProjectCreateV30DeliverySettingDeepBidType, error) {
	ev := ProjectCreateV30DeliverySettingDeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliverySettingDeepBidType: valid values are %v", v, AllowedProjectCreateV30DeliverySettingDeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliverySettingDeepBidType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliverySettingDeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_setting_deep_bid_type value
func (v ProjectCreateV30DeliverySettingDeepBidType) Ptr() *ProjectCreateV30DeliverySettingDeepBidType {
	return &v
}