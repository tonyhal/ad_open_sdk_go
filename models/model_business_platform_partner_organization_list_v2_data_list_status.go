/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BusinessPlatformPartnerOrganizationListV2DataListStatus
type BusinessPlatformPartnerOrganizationListV2DataListStatus string

// List of business_platform_partner_organization_list_v2_data_list_status
const (
	BOUND_BusinessPlatformPartnerOrganizationListV2DataListStatus   BusinessPlatformPartnerOrganizationListV2DataListStatus = "BOUND"
	BINDING_BusinessPlatformPartnerOrganizationListV2DataListStatus BusinessPlatformPartnerOrganizationListV2DataListStatus = "BINDING"
	INVALID_BusinessPlatformPartnerOrganizationListV2DataListStatus BusinessPlatformPartnerOrganizationListV2DataListStatus = "INVALID"
	DELETED_BusinessPlatformPartnerOrganizationListV2DataListStatus BusinessPlatformPartnerOrganizationListV2DataListStatus = "DELETED"
)

// All allowed values of BusinessPlatformPartnerOrganizationListV2DataListStatus enum
var AllowedBusinessPlatformPartnerOrganizationListV2DataListStatusEnumValues = []BusinessPlatformPartnerOrganizationListV2DataListStatus{
	"BOUND",
	"BINDING",
	"INVALID",
	"DELETED",
}

// NewBusinessPlatformPartnerOrganizationListV2DataListStatusFromValue returns a pointer to a valid BusinessPlatformPartnerOrganizationListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformPartnerOrganizationListV2DataListStatusFromValue(v string) (*BusinessPlatformPartnerOrganizationListV2DataListStatus, error) {
	ev := BusinessPlatformPartnerOrganizationListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformPartnerOrganizationListV2DataListStatus: valid values are %v", v, AllowedBusinessPlatformPartnerOrganizationListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformPartnerOrganizationListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformPartnerOrganizationListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_partner_organization_list_v2_data_list_status value
func (v BusinessPlatformPartnerOrganizationListV2DataListStatus) Ptr() *BusinessPlatformPartnerOrganizationListV2DataListStatus {
	return &v
}
