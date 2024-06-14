/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceDistrict
type QianchuanAdUpdateV10AudienceDistrict string

// List of qianchuan_ad_update_v1.0_audience_district
const (
	CITY_QianchuanAdUpdateV10AudienceDistrict   QianchuanAdUpdateV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAdUpdateV10AudienceDistrict QianchuanAdUpdateV10AudienceDistrict = "COUNTY"
	NONE_QianchuanAdUpdateV10AudienceDistrict   QianchuanAdUpdateV10AudienceDistrict = "NONE"
)

// All allowed values of QianchuanAdUpdateV10AudienceDistrict enum
var AllowedQianchuanAdUpdateV10AudienceDistrictEnumValues = []QianchuanAdUpdateV10AudienceDistrict{
	"CITY",
	"COUNTY",
	"NONE",
}

// NewQianchuanAdUpdateV10AudienceDistrictFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceDistrictFromValue(v string) (*QianchuanAdUpdateV10AudienceDistrict, error) {
	ev := QianchuanAdUpdateV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceDistrict: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_district value
func (v QianchuanAdUpdateV10AudienceDistrict) Ptr() *QianchuanAdUpdateV10AudienceDistrict {
	return &v
}
