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

// DpaAssetsListV2DataListAssetType
type DpaAssetsListV2DataListAssetType string

// List of dpa_assets_list_v2_data_list_asset_type
const (
	AUTO_DpaAssetsListV2DataListAssetType DpaAssetsListV2DataListAssetType = "AUTO"
)

// All allowed values of DpaAssetsListV2DataListAssetType enum
var AllowedDpaAssetsListV2DataListAssetTypeEnumValues = []DpaAssetsListV2DataListAssetType{
	"AUTO",
}

// NewDpaAssetsListV2DataListAssetTypeFromValue returns a pointer to a valid DpaAssetsListV2DataListAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetsListV2DataListAssetTypeFromValue(v string) (*DpaAssetsListV2DataListAssetType, error) {
	ev := DpaAssetsListV2DataListAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetsListV2DataListAssetType: valid values are %v", v, AllowedDpaAssetsListV2DataListAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetsListV2DataListAssetType) IsValid() bool {
	for _, existing := range AllowedDpaAssetsListV2DataListAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_assets_list_v2_data_list_asset_type value
func (v DpaAssetsListV2DataListAssetType) Ptr() *DpaAssetsListV2DataListAssetType {
	return &v
}
