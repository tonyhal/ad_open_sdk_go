/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsStarTaskMaterialTypeV2ResponseDataMaterialListInner struct for ToolsStarTaskMaterialTypeV2ResponseDataMaterialListInner
type ToolsStarTaskMaterialTypeV2ResponseDataMaterialListInner struct {
	//
	Children []*ToolsStarTaskMaterialTypeV2ResponseDataMaterialListInnerChildrenInner `json:"children,omitempty"`
	//
	FirstMaterialName string `json:"first_material_name"`
	//
	FirstMaterialType int64 `json:"first_material_type"`
}
