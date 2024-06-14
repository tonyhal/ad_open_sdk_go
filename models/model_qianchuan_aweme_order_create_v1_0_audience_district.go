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

// QianchuanAwemeOrderCreateV10AudienceDistrict
type QianchuanAwemeOrderCreateV10AudienceDistrict string

// List of qianchuan_aweme_order_create_v1.0_audience_district
const (
	CITY_QianchuanAwemeOrderCreateV10AudienceDistrict   QianchuanAwemeOrderCreateV10AudienceDistrict = "CITY"
	COUNTY_QianchuanAwemeOrderCreateV10AudienceDistrict QianchuanAwemeOrderCreateV10AudienceDistrict = "COUNTY"
)

// All allowed values of QianchuanAwemeOrderCreateV10AudienceDistrict enum
var AllowedQianchuanAwemeOrderCreateV10AudienceDistrictEnumValues = []QianchuanAwemeOrderCreateV10AudienceDistrict{
	"CITY",
	"COUNTY",
}

// NewQianchuanAwemeOrderCreateV10AudienceDistrictFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10AudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10AudienceDistrictFromValue(v string) (*QianchuanAwemeOrderCreateV10AudienceDistrict, error) {
	ev := QianchuanAwemeOrderCreateV10AudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10AudienceDistrict: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10AudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10AudienceDistrict) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10AudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_audience_district value
func (v QianchuanAwemeOrderCreateV10AudienceDistrict) Ptr() *QianchuanAwemeOrderCreateV10AudienceDistrict {
	return &v
}
