/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCountryInfoV2ResponseDataDistrictsInner struct for ToolsCountryInfoV2ResponseDataDistrictsInner
type ToolsCountryInfoV2ResponseDataDistrictsInner struct {
	// 行政区域code
	Code        *string                                     `json:"code,omitempty"`
	Description *ToolsCountryInfoV2DataDistrictsDescription `json:"description,omitempty"`
	// 行政区域名称
	Name *string `json:"name,omitempty"`
}
