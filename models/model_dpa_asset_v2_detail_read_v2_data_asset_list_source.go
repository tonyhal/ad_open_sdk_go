/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2DetailReadV2DataAssetListSource
type DpaAssetV2DetailReadV2DataAssetListSource string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_source
const (
	AUTO_DpaAssetV2DetailReadV2DataAssetListSource   DpaAssetV2DetailReadV2DataAssetListSource = "AUTO"
	MANUAL_DpaAssetV2DetailReadV2DataAssetListSource DpaAssetV2DetailReadV2DataAssetListSource = "MANUAL"
)

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_source value
func (v DpaAssetV2DetailReadV2DataAssetListSource) Ptr() *DpaAssetV2DetailReadV2DataAssetListSource {
	return &v
}
