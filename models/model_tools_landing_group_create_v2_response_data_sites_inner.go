/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupCreateV2ResponseDataSitesInner struct for ToolsLandingGroupCreateV2ResponseDataSitesInner
type ToolsLandingGroupCreateV2ResponseDataSitesInner struct {
	//
	MemberId *int64 `json:"member_id,omitempty"`
	//
	SiteId        *int64                                           `json:"site_id,omitempty"`
	SiteOptStatus *ToolsLandingGroupCreateV2DataSitesSiteOptStatus `json:"site_opt_status,omitempty"`
	//
	SiteUrl *string `json:"site_url,omitempty"`
}