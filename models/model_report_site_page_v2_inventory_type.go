/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportSitePageV2InventoryType
type ReportSitePageV2InventoryType string

// List of report_site_page_v2_inventory_type
const (
	AWEME_ReportSitePageV2InventoryType      ReportSitePageV2InventoryType = "AWEME"
	TOUTIAO_ReportSitePageV2InventoryType    ReportSitePageV2InventoryType = "TOUTIAO"
	UNION_SLOT_ReportSitePageV2InventoryType ReportSitePageV2InventoryType = "UNION_SLOT"
	WATERMELON_ReportSitePageV2InventoryType ReportSitePageV2InventoryType = "WATERMELON"
	HOTSOON_ReportSitePageV2InventoryType    ReportSitePageV2InventoryType = "HOTSOON"
)

// Ptr returns reference to report_site_page_v2_inventory_type value
func (v ReportSitePageV2InventoryType) Ptr() *ReportSitePageV2InventoryType {
	return &v
}
