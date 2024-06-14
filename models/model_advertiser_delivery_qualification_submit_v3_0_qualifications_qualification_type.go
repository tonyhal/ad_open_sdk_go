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

// AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType
type AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType string

// List of advertiser_delivery_qualification_submit_v3.0_qualifications_qualification_type
const (
	AGENT_ARRANGEMENT_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                           AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "AGENT_ARRANGEMENT"
	AUTHORIZATION_CONTRACT_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                      AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "AUTHORIZATION_CONTRACT"
	DISTRIBUTION_AUTHORIZATION_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                  AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "DISTRIBUTION_AUTHORIZATION"
	ICP_RECORD_AUTHORIZATION_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                    AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "ICP_RECORD_AUTHORIZATION"
	OTHER_CERTIFICATION_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                         AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "OTHER_CERTIFICATION"
	PATENT_CERTIFICATE_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                          AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "PATENT_CERTIFICATE"
	PORTRAIT_AUTHORIZATION_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                      AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "PORTRAIT_AUTHORIZATION"
	QUALITY_REPORT_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType                              AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "QUALITY_REPORT"
	SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE"
	TRADEMARK_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType          AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "TRADEMARK_REGISTRATION_CERTIFICATE"
	VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION_AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType  AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType = "VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION"
)

// All allowed values of AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType enum
var AllowedAdvertiserDeliveryQualificationSubmitV30QualificationsQualificationTypeEnumValues = []AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType{
	"AGENT_ARRANGEMENT",
	"AUTHORIZATION_CONTRACT",
	"DISTRIBUTION_AUTHORIZATION",
	"ICP_RECORD_AUTHORIZATION",
	"OTHER_CERTIFICATION",
	"PATENT_CERTIFICATE",
	"PORTRAIT_AUTHORIZATION",
	"QUALITY_REPORT",
	"SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE",
	"TRADEMARK_REGISTRATION_CERTIFICATE",
	"VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION",
}

// NewAdvertiserDeliveryQualificationSubmitV30QualificationsQualificationTypeFromValue returns a pointer to a valid AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryQualificationSubmitV30QualificationsQualificationTypeFromValue(v string) (*AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType, error) {
	ev := AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType: valid values are %v", v, AllowedAdvertiserDeliveryQualificationSubmitV30QualificationsQualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryQualificationSubmitV30QualificationsQualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_qualification_submit_v3.0_qualifications_qualification_type value
func (v AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType) Ptr() *AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType {
	return &v
}
