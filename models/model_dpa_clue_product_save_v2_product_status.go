/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductSaveV2ProductStatus
type DpaClueProductSaveV2ProductStatus string

// List of dpa_clue_product_save_v2_product_status
const (
	STATUS_OFFLINE_DpaClueProductSaveV2ProductStatus DpaClueProductSaveV2ProductStatus = "STATUS_OFFLINE"
	STATUS_ONLINE_DpaClueProductSaveV2ProductStatus  DpaClueProductSaveV2ProductStatus = "STATUS_ONLINE"
)

// Ptr returns reference to dpa_clue_product_save_v2_product_status value
func (v DpaClueProductSaveV2ProductStatus) Ptr() *DpaClueProductSaveV2ProductStatus {
	return &v
}
