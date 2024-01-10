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

// AdGetV2DataAudienceDpaRtaRecommendType
type AdGetV2DataAudienceDpaRtaRecommendType string

// List of ad_get_v2_data_audience_dpa_rta_recommend_type
const (
	MORE_AdGetV2DataAudienceDpaRtaRecommendType AdGetV2DataAudienceDpaRtaRecommendType = "MORE"
	ONLY_AdGetV2DataAudienceDpaRtaRecommendType AdGetV2DataAudienceDpaRtaRecommendType = "ONLY"
)

// All allowed values of AdGetV2DataAudienceDpaRtaRecommendType enum
var AllowedAdGetV2DataAudienceDpaRtaRecommendTypeEnumValues = []AdGetV2DataAudienceDpaRtaRecommendType{
	"MORE",
	"ONLY",
}

// NewAdGetV2DataAudienceDpaRtaRecommendTypeFromValue returns a pointer to a valid AdGetV2DataAudienceDpaRtaRecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceDpaRtaRecommendTypeFromValue(v string) (*AdGetV2DataAudienceDpaRtaRecommendType, error) {
	ev := AdGetV2DataAudienceDpaRtaRecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceDpaRtaRecommendType: valid values are %v", v, AllowedAdGetV2DataAudienceDpaRtaRecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceDpaRtaRecommendType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceDpaRtaRecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_dpa_rta_recommend_type value
func (v AdGetV2DataAudienceDpaRtaRecommendType) Ptr() *AdGetV2DataAudienceDpaRtaRecommendType {
	return &v
}
