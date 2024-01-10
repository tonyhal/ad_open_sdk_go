/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2FilterAwemeAbnormalActive
type AudiencePackageCreateV2FilterAwemeAbnormalActive int64

// List of audience_package_create_v2_filter_aweme_abnormal_active
const (
	Enum_0_AudiencePackageCreateV2FilterAwemeAbnormalActive AudiencePackageCreateV2FilterAwemeAbnormalActive = 0
	Enum_1_AudiencePackageCreateV2FilterAwemeAbnormalActive AudiencePackageCreateV2FilterAwemeAbnormalActive = 1
)

// All allowed values of AudiencePackageCreateV2FilterAwemeAbnormalActive enum
var AllowedAudiencePackageCreateV2FilterAwemeAbnormalActiveEnumValues = []AudiencePackageCreateV2FilterAwemeAbnormalActive{
	0,
	1,
}

// NewAudiencePackageCreateV2FilterAwemeAbnormalActiveFromValue returns a pointer to a valid AudiencePackageCreateV2FilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2FilterAwemeAbnormalActiveFromValue(v int64) (*AudiencePackageCreateV2FilterAwemeAbnormalActive, error) {
	ev := AudiencePackageCreateV2FilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2FilterAwemeAbnormalActive: valid values are %v", v, AllowedAudiencePackageCreateV2FilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2FilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2FilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_filter_aweme_abnormal_active value
func (v AudiencePackageCreateV2FilterAwemeAbnormalActive) Ptr() *AudiencePackageCreateV2FilterAwemeAbnormalActive {
	return &v
}
