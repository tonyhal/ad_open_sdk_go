/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupCreateV2ResponseData
type ToolsLandingGroupCreateV2ResponseData struct {
	GroupFlowType *ToolsLandingGroupCreateV2DataGroupFlowType `json:"group_flow_type,omitempty"`
	//
	GroupId     *string                                   `json:"group_id,omitempty"`
	GroupStatus *ToolsLandingGroupCreateV2DataGroupStatus `json:"group_status,omitempty"`
	//
	GroupTitle *string `json:"group_title,omitempty"`
	//
	GroupUrl *string `json:"group_url,omitempty"`
	//
	Sites []*ToolsLandingGroupCreateV2ResponseDataSitesInner `json:"sites,omitempty"`
}
