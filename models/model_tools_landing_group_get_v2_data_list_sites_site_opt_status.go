/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsLandingGroupGetV2DataListSitesSiteOptStatus
type ToolsLandingGroupGetV2DataListSitesSiteOptStatus string

// List of tools_landing_group_get_v2_data_list_sites_site_opt_status
const (
	SITE_OPT_STATUS_DISABLE_ToolsLandingGroupGetV2DataListSitesSiteOptStatus ToolsLandingGroupGetV2DataListSitesSiteOptStatus = "SITE_OPT_STATUS_DISABLE"
	SITE_OPT_STATUS_ENABLE_ToolsLandingGroupGetV2DataListSitesSiteOptStatus  ToolsLandingGroupGetV2DataListSitesSiteOptStatus = "SITE_OPT_STATUS_ENABLE"
)

// All allowed values of ToolsLandingGroupGetV2DataListSitesSiteOptStatus enum
var AllowedToolsLandingGroupGetV2DataListSitesSiteOptStatusEnumValues = []ToolsLandingGroupGetV2DataListSitesSiteOptStatus{
	"SITE_OPT_STATUS_DISABLE",
	"SITE_OPT_STATUS_ENABLE",
}

// NewToolsLandingGroupGetV2DataListSitesSiteOptStatusFromValue returns a pointer to a valid ToolsLandingGroupGetV2DataListSitesSiteOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsLandingGroupGetV2DataListSitesSiteOptStatusFromValue(v string) (*ToolsLandingGroupGetV2DataListSitesSiteOptStatus, error) {
	ev := ToolsLandingGroupGetV2DataListSitesSiteOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsLandingGroupGetV2DataListSitesSiteOptStatus: valid values are %v", v, AllowedToolsLandingGroupGetV2DataListSitesSiteOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsLandingGroupGetV2DataListSitesSiteOptStatus) IsValid() bool {
	for _, existing := range AllowedToolsLandingGroupGetV2DataListSitesSiteOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_landing_group_get_v2_data_list_sites_site_opt_status value
func (v ToolsLandingGroupGetV2DataListSitesSiteOptStatus) Ptr() *ToolsLandingGroupGetV2DataListSitesSiteOptStatus {
	return &v
}
