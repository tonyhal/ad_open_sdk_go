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

// QianchuanAudienceListGetV10DataAudiencesListAudienceSource
type QianchuanAudienceListGetV10DataAudiencesListAudienceSource int64

// List of qianchuan_audience_list_get_v1.0_data_audiences_list_audience_source
const (
	Enum_0_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 0
	Enum_1_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 1
	Enum_2_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 2
	Enum_3_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 3
	Enum_4_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 4
	Enum_5_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 5
	Enum_6_QianchuanAudienceListGetV10DataAudiencesListAudienceSource QianchuanAudienceListGetV10DataAudiencesListAudienceSource = 6
)

// All allowed values of QianchuanAudienceListGetV10DataAudiencesListAudienceSource enum
var AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceSourceEnumValues = []QianchuanAudienceListGetV10DataAudiencesListAudienceSource{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

// NewQianchuanAudienceListGetV10DataAudiencesListAudienceSourceFromValue returns a pointer to a valid QianchuanAudienceListGetV10DataAudiencesListAudienceSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAudienceListGetV10DataAudiencesListAudienceSourceFromValue(v int64) (*QianchuanAudienceListGetV10DataAudiencesListAudienceSource, error) {
	ev := QianchuanAudienceListGetV10DataAudiencesListAudienceSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAudienceListGetV10DataAudiencesListAudienceSource: valid values are %v", v, AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAudienceListGetV10DataAudiencesListAudienceSource) IsValid() bool {
	for _, existing := range AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_audience_list_get_v1.0_data_audiences_list_audience_source value
func (v QianchuanAudienceListGetV10DataAudiencesListAudienceSource) Ptr() *QianchuanAudienceListGetV10DataAudiencesListAudienceSource {
	return &v
}
