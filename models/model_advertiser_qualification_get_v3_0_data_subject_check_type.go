/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationGetV30DataSubjectCheckType
type AdvertiserQualificationGetV30DataSubjectCheckType string

// List of advertiser_qualification_get_v3.0_data_subject_check_type
const (
	COMPANY_AdvertiserQualificationGetV30DataSubjectCheckType         AdvertiserQualificationGetV30DataSubjectCheckType = "COMPANY"
	GOVERNMENT_AdvertiserQualificationGetV30DataSubjectCheckType      AdvertiserQualificationGetV30DataSubjectCheckType = "GOVERNMENT"
	HK_MACAO_TAIWAN_AdvertiserQualificationGetV30DataSubjectCheckType AdvertiserQualificationGetV30DataSubjectCheckType = "HK_MACAO_TAIWAN"
	INDIVIDUAL_AdvertiserQualificationGetV30DataSubjectCheckType      AdvertiserQualificationGetV30DataSubjectCheckType = "INDIVIDUAL"
	OTHERS_AdvertiserQualificationGetV30DataSubjectCheckType          AdvertiserQualificationGetV30DataSubjectCheckType = "OTHERS"
	OVERSEA_AdvertiserQualificationGetV30DataSubjectCheckType         AdvertiserQualificationGetV30DataSubjectCheckType = "OVERSEA"
	SELF_EMPLOY_AdvertiserQualificationGetV30DataSubjectCheckType     AdvertiserQualificationGetV30DataSubjectCheckType = "SELF_EMPLOY"
)

// All allowed values of AdvertiserQualificationGetV30DataSubjectCheckType enum
var AllowedAdvertiserQualificationGetV30DataSubjectCheckTypeEnumValues = []AdvertiserQualificationGetV30DataSubjectCheckType{
	"COMPANY",
	"GOVERNMENT",
	"HK_MACAO_TAIWAN",
	"INDIVIDUAL",
	"OTHERS",
	"OVERSEA",
	"SELF_EMPLOY",
}

// NewAdvertiserQualificationGetV30DataSubjectCheckTypeFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataSubjectCheckType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataSubjectCheckTypeFromValue(v string) (*AdvertiserQualificationGetV30DataSubjectCheckType, error) {
	ev := AdvertiserQualificationGetV30DataSubjectCheckType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataSubjectCheckType: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataSubjectCheckTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataSubjectCheckType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataSubjectCheckTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_subject_check_type value
func (v AdvertiserQualificationGetV30DataSubjectCheckType) Ptr() *AdvertiserQualificationGetV30DataSubjectCheckType {
	return &v
}
