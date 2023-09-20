/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserQualificationGetV30DataSubjectCompanyType
type AdvertiserQualificationGetV30DataSubjectCompanyType string

// List of advertiser_qualification_get_v3.0_data_subject_company_type
const (
	COMPANY_AdvertiserQualificationGetV30DataSubjectCompanyType    AdvertiserQualificationGetV30DataSubjectCompanyType = "COMPANY"
	INDIVIDUAL_AdvertiserQualificationGetV30DataSubjectCompanyType AdvertiserQualificationGetV30DataSubjectCompanyType = "INDIVIDUAL"
)

// All allowed values of AdvertiserQualificationGetV30DataSubjectCompanyType enum
var AllowedAdvertiserQualificationGetV30DataSubjectCompanyTypeEnumValues = []AdvertiserQualificationGetV30DataSubjectCompanyType{
	"COMPANY",
	"INDIVIDUAL",
}

// NewAdvertiserQualificationGetV30DataSubjectCompanyTypeFromValue returns a pointer to a valid AdvertiserQualificationGetV30DataSubjectCompanyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserQualificationGetV30DataSubjectCompanyTypeFromValue(v string) (*AdvertiserQualificationGetV30DataSubjectCompanyType, error) {
	ev := AdvertiserQualificationGetV30DataSubjectCompanyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserQualificationGetV30DataSubjectCompanyType: valid values are %v", v, AllowedAdvertiserQualificationGetV30DataSubjectCompanyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserQualificationGetV30DataSubjectCompanyType) IsValid() bool {
	for _, existing := range AllowedAdvertiserQualificationGetV30DataSubjectCompanyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_qualification_get_v3.0_data_subject_company_type value
func (v AdvertiserQualificationGetV30DataSubjectCompanyType) Ptr() *AdvertiserQualificationGetV30DataSubjectCompanyType {
	return &v
}
