/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsEventAllAssetsListV2DataAssetListShareType
type ToolsEventAllAssetsListV2DataAssetListShareType string

// List of tools_event_all_assets_list_v2_data_asset_list_share_type
const (
	MY_CREATIONS_ToolsEventAllAssetsListV2DataAssetListShareType  ToolsEventAllAssetsListV2DataAssetListShareType = "MY_CREATIONS"
	SHARING_ToolsEventAllAssetsListV2DataAssetListShareType       ToolsEventAllAssetsListV2DataAssetListShareType = "SHARING"
	SHATE_EXPIRED_ToolsEventAllAssetsListV2DataAssetListShareType ToolsEventAllAssetsListV2DataAssetListShareType = "SHATE_EXPIRED"
)

// All allowed values of ToolsEventAllAssetsListV2DataAssetListShareType enum
var AllowedToolsEventAllAssetsListV2DataAssetListShareTypeEnumValues = []ToolsEventAllAssetsListV2DataAssetListShareType{
	"MY_CREATIONS",
	"SHARING",
	"SHATE_EXPIRED",
}

// NewToolsEventAllAssetsListV2DataAssetListShareTypeFromValue returns a pointer to a valid ToolsEventAllAssetsListV2DataAssetListShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAllAssetsListV2DataAssetListShareTypeFromValue(v string) (*ToolsEventAllAssetsListV2DataAssetListShareType, error) {
	ev := ToolsEventAllAssetsListV2DataAssetListShareType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAllAssetsListV2DataAssetListShareType: valid values are %v", v, AllowedToolsEventAllAssetsListV2DataAssetListShareTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAllAssetsListV2DataAssetListShareType) IsValid() bool {
	for _, existing := range AllowedToolsEventAllAssetsListV2DataAssetListShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_all_assets_list_v2_data_asset_list_share_type value
func (v ToolsEventAllAssetsListV2DataAssetListShareType) Ptr() *ToolsEventAllAssetsListV2DataAssetListShareType {
	return &v
}