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

// ToolsEventAllAssetsDetailV2DataAssetListAssetType
type ToolsEventAllAssetsDetailV2DataAssetListAssetType string

// List of tools_event_all_assets_detail_v2_data_asset_list_asset_type
const (
	THIRD_EXTERNAL_ToolsEventAllAssetsDetailV2DataAssetListAssetType  ToolsEventAllAssetsDetailV2DataAssetListAssetType = "THIRD_EXTERNAL"
	TETRIS_EXTERNAL_ToolsEventAllAssetsDetailV2DataAssetListAssetType ToolsEventAllAssetsDetailV2DataAssetListAssetType = "TETRIS_EXTERNAL"
	APP_ToolsEventAllAssetsDetailV2DataAssetListAssetType             ToolsEventAllAssetsDetailV2DataAssetListAssetType = "APP"
	QUICK_APP_ToolsEventAllAssetsDetailV2DataAssetListAssetType       ToolsEventAllAssetsDetailV2DataAssetListAssetType = "QUICK_APP"
	MINI_PROGRAME_ToolsEventAllAssetsDetailV2DataAssetListAssetType   ToolsEventAllAssetsDetailV2DataAssetListAssetType = "MINI_PROGRAME"
)

// All allowed values of ToolsEventAllAssetsDetailV2DataAssetListAssetType enum
var AllowedToolsEventAllAssetsDetailV2DataAssetListAssetTypeEnumValues = []ToolsEventAllAssetsDetailV2DataAssetListAssetType{
	"THIRD_EXTERNAL",
	"TETRIS_EXTERNAL",
	"APP",
	"QUICK_APP",
	"MINI_PROGRAME",
}

// NewToolsEventAllAssetsDetailV2DataAssetListAssetTypeFromValue returns a pointer to a valid ToolsEventAllAssetsDetailV2DataAssetListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsEventAllAssetsDetailV2DataAssetListAssetTypeFromValue(v string) (*ToolsEventAllAssetsDetailV2DataAssetListAssetType, error) {
	ev := ToolsEventAllAssetsDetailV2DataAssetListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsEventAllAssetsDetailV2DataAssetListAssetType: valid values are %v", v, AllowedToolsEventAllAssetsDetailV2DataAssetListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsEventAllAssetsDetailV2DataAssetListAssetType) IsValid() bool {
	for _, existing := range AllowedToolsEventAllAssetsDetailV2DataAssetListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_event_all_assets_detail_v2_data_asset_list_asset_type value
func (v ToolsEventAllAssetsDetailV2DataAssetListAssetType) Ptr() *ToolsEventAllAssetsDetailV2DataAssetListAssetType {
	return &v
}
