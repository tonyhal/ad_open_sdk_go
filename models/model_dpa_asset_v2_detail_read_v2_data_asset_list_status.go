/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaAssetV2DetailReadV2DataAssetListStatus
type DpaAssetV2DetailReadV2DataAssetListStatus string

// List of dpa_asset_v2_detail_read_v2_data_asset_list_status
const (
	DISABLE_DpaAssetV2DetailReadV2DataAssetListStatus DpaAssetV2DetailReadV2DataAssetListStatus = "DISABLE"
	ENABLE_DpaAssetV2DetailReadV2DataAssetListStatus  DpaAssetV2DetailReadV2DataAssetListStatus = "ENABLE"
)

// Ptr returns reference to dpa_asset_v2_detail_read_v2_data_asset_list_status value
func (v DpaAssetV2DetailReadV2DataAssetListStatus) Ptr() *DpaAssetV2DetailReadV2DataAssetListStatus {
	return &v
}
