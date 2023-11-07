/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsOrangeSiteGetV30ResponseDataListInner struct for ToolsOrangeSiteGetV30ResponseDataListInner
type ToolsOrangeSiteGetV30ResponseDataListInner struct {
	FunctionType *ToolsOrangeSiteGetV30DataListFunctionType `json:"function_type,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SiteId   *string                                `json:"site_id,omitempty"`
	SiteType *ToolsOrangeSiteGetV30DataListSiteType `json:"site_type,omitempty"`
	Status   *ToolsOrangeSiteGetV30DataListStatus   `json:"status,omitempty"`
	//
	Thumbnail *string `json:"thumbnail,omitempty"`
	//
	Url *string `json:"url,omitempty"`
}
