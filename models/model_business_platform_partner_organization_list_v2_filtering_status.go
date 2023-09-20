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

// BusinessPlatformPartnerOrganizationListV2FilteringStatus
type BusinessPlatformPartnerOrganizationListV2FilteringStatus string

// List of business_platform_partner_organization_list_v2_filtering_status
const (
	BOUND_BusinessPlatformPartnerOrganizationListV2FilteringStatus   BusinessPlatformPartnerOrganizationListV2FilteringStatus = "BOUND"
	BINDING_BusinessPlatformPartnerOrganizationListV2FilteringStatus BusinessPlatformPartnerOrganizationListV2FilteringStatus = "BINDING"
	INVALID_BusinessPlatformPartnerOrganizationListV2FilteringStatus BusinessPlatformPartnerOrganizationListV2FilteringStatus = "INVALID"
	DELETED_BusinessPlatformPartnerOrganizationListV2FilteringStatus BusinessPlatformPartnerOrganizationListV2FilteringStatus = "DELETED"
)

// All allowed values of BusinessPlatformPartnerOrganizationListV2FilteringStatus enum
var AllowedBusinessPlatformPartnerOrganizationListV2FilteringStatusEnumValues = []BusinessPlatformPartnerOrganizationListV2FilteringStatus{
	"BOUND",
	"BINDING",
	"INVALID",
	"DELETED",
}

// NewBusinessPlatformPartnerOrganizationListV2FilteringStatusFromValue returns a pointer to a valid BusinessPlatformPartnerOrganizationListV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformPartnerOrganizationListV2FilteringStatusFromValue(v string) (*BusinessPlatformPartnerOrganizationListV2FilteringStatus, error) {
	ev := BusinessPlatformPartnerOrganizationListV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformPartnerOrganizationListV2FilteringStatus: valid values are %v", v, AllowedBusinessPlatformPartnerOrganizationListV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformPartnerOrganizationListV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformPartnerOrganizationListV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_partner_organization_list_v2_filtering_status value
func (v BusinessPlatformPartnerOrganizationListV2FilteringStatus) Ptr() *BusinessPlatformPartnerOrganizationListV2FilteringStatus {
	return &v
}
