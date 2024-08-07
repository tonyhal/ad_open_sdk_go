/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaAssetV2DetailReadV2DataAssetListSource
type DpaAssetV2DetailReadV2DataAssetListSource string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_source
const (
	AUTO_DpaAssetV2DetailReadV2DataAssetListSource   DpaAssetV2DetailReadV2DataAssetListSource = "AUTO"
	MANUAL_DpaAssetV2DetailReadV2DataAssetListSource DpaAssetV2DetailReadV2DataAssetListSource = "MANUAL"
)

// All allowed values of DpaAssetV2DetailReadV2DataAssetListSource enum
var AllowedDpaAssetV2DetailReadV2DataAssetListSourceEnumValues = []DpaAssetV2DetailReadV2DataAssetListSource{
	"AUTO",
	"MANUAL",
}

// NewDpaAssetV2DetailReadV2DataAssetListSourceFromValue returns a pointer to a valid DpaAssetV2DetailReadV2DataAssetListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaAssetV2DetailReadV2DataAssetListSourceFromValue(v string) (*DpaAssetV2DetailReadV2DataAssetListSource, error) {
	ev := DpaAssetV2DetailReadV2DataAssetListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaAssetV2DetailReadV2DataAssetListSource: valid values are %v", v, AllowedDpaAssetV2DetailReadV2DataAssetListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaAssetV2DetailReadV2DataAssetListSource) IsValid() bool {
	for _, existing := range AllowedDpaAssetV2DetailReadV2DataAssetListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_source value
func (v DpaAssetV2DetailReadV2DataAssetListSource) Ptr() *DpaAssetV2DetailReadV2DataAssetListSource {
	return &v
}
