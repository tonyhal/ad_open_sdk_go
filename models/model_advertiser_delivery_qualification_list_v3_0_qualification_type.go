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

// AdvertiserDeliveryQualificationListV30QualificationType
type AdvertiserDeliveryQualificationListV30QualificationType string

// List of advertiser_delivery_qualification_list_v3.0_qualification_type
const (
	AGENT_ARRANGEMENT_AdvertiserDeliveryQualificationListV30QualificationType                           AdvertiserDeliveryQualificationListV30QualificationType = "AGENT_ARRANGEMENT"
	AUTHORIZATION_CONTRACT_AdvertiserDeliveryQualificationListV30QualificationType                      AdvertiserDeliveryQualificationListV30QualificationType = "AUTHORIZATION_CONTRACT"
	DISTRIBUTION_AUTHORIZATION_AdvertiserDeliveryQualificationListV30QualificationType                  AdvertiserDeliveryQualificationListV30QualificationType = "DISTRIBUTION_AUTHORIZATION"
	ICP_RECORD_AUTHORIZATION_AdvertiserDeliveryQualificationListV30QualificationType                    AdvertiserDeliveryQualificationListV30QualificationType = "ICP_RECORD_AUTHORIZATION"
	OTHER_CERTIFICATION_AdvertiserDeliveryQualificationListV30QualificationType                         AdvertiserDeliveryQualificationListV30QualificationType = "OTHER_CERTIFICATION"
	PATENT_CERTIFICATE_AdvertiserDeliveryQualificationListV30QualificationType                          AdvertiserDeliveryQualificationListV30QualificationType = "PATENT_CERTIFICATE"
	PORTRAIT_AUTHORIZATION_AdvertiserDeliveryQualificationListV30QualificationType                      AdvertiserDeliveryQualificationListV30QualificationType = "PORTRAIT_AUTHORIZATION"
	QUALITY_REPORT_AdvertiserDeliveryQualificationListV30QualificationType                              AdvertiserDeliveryQualificationListV30QualificationType = "QUALITY_REPORT"
	SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationListV30QualificationType AdvertiserDeliveryQualificationListV30QualificationType = "SOFTWARE_COPYRIGHT_REGISTRATION_CERTIFICATE"
	TRADEMARK_REGISTRATION_CERTIFICATE_AdvertiserDeliveryQualificationListV30QualificationType          AdvertiserDeliveryQualificationListV30QualificationType = "TRADEMARK_REGISTRATION_CERTIFICATE"
	VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION_AdvertiserDeliveryQualificationListV30QualificationType  AdvertiserDeliveryQualificationListV30QualificationType = "VIDEO_MATERIAL_PRODUCTION_IP_AUTHORIZATION"
)

// All allowed values of AdvertiserDeliveryQualificationListV30QualificationType enum
var AllowedAdvertiserDeliveryQualificationListV30QualificationTypeEnumValues = []AdvertiserDeliveryQualificationListV30QualificationType{
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

// NewAdvertiserDeliveryQualificationListV30QualificationTypeFromValue returns a pointer to a valid AdvertiserDeliveryQualificationListV30QualificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserDeliveryQualificationListV30QualificationTypeFromValue(v string) (*AdvertiserDeliveryQualificationListV30QualificationType, error) {
	ev := AdvertiserDeliveryQualificationListV30QualificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserDeliveryQualificationListV30QualificationType: valid values are %v", v, AllowedAdvertiserDeliveryQualificationListV30QualificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserDeliveryQualificationListV30QualificationType) IsValid() bool {
	for _, existing := range AllowedAdvertiserDeliveryQualificationListV30QualificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_delivery_qualification_list_v3.0_qualification_type value
func (v AdvertiserDeliveryQualificationListV30QualificationType) Ptr() *AdvertiserDeliveryQualificationListV30QualificationType {
	return &v
}