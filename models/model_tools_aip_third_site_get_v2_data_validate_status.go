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

// ToolsAipThirdSiteGetV2DataValidateStatus
type ToolsAipThirdSiteGetV2DataValidateStatus string

// List of tools_aip_third_site_get_v2_data_validate_status
const (
	VALIDATE_SUCCESS_ToolsAipThirdSiteGetV2DataValidateStatus                 ToolsAipThirdSiteGetV2DataValidateStatus = "VALIDATE_SUCCESS"
	VALIDATE_FAIL_ToolsAipThirdSiteGetV2DataValidateStatus                    ToolsAipThirdSiteGetV2DataValidateStatus = "VALIDATE_FAIL"
	HIGH_RISK_INDUSTRY_VALIDATE_FAIL_ToolsAipThirdSiteGetV2DataValidateStatus ToolsAipThirdSiteGetV2DataValidateStatus = "HIGH_RISK_INDUSTRY_VALIDATE_FAIL"
	UN_VALIDATE_ToolsAipThirdSiteGetV2DataValidateStatus                      ToolsAipThirdSiteGetV2DataValidateStatus = "UN_VALIDATE"
	PUBLISH_ToolsAipThirdSiteGetV2DataValidateStatus                          ToolsAipThirdSiteGetV2DataValidateStatus = "PUBLISH"
)

// All allowed values of ToolsAipThirdSiteGetV2DataValidateStatus enum
var AllowedToolsAipThirdSiteGetV2DataValidateStatusEnumValues = []ToolsAipThirdSiteGetV2DataValidateStatus{
	"VALIDATE_SUCCESS",
	"VALIDATE_FAIL",
	"HIGH_RISK_INDUSTRY_VALIDATE_FAIL",
	"UN_VALIDATE",
	"PUBLISH",
}

// NewToolsAipThirdSiteGetV2DataValidateStatusFromValue returns a pointer to a valid ToolsAipThirdSiteGetV2DataValidateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAipThirdSiteGetV2DataValidateStatusFromValue(v string) (*ToolsAipThirdSiteGetV2DataValidateStatus, error) {
	ev := ToolsAipThirdSiteGetV2DataValidateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAipThirdSiteGetV2DataValidateStatus: valid values are %v", v, AllowedToolsAipThirdSiteGetV2DataValidateStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAipThirdSiteGetV2DataValidateStatus) IsValid() bool {
	for _, existing := range AllowedToolsAipThirdSiteGetV2DataValidateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aip_third_site_get_v2_data_validate_status value
func (v ToolsAipThirdSiteGetV2DataValidateStatus) Ptr() *ToolsAipThirdSiteGetV2DataValidateStatus {
	return &v
}
