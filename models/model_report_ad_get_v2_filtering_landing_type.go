/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringLandingType
type ReportAdGetV2FilteringLandingType string

// List of report_ad_get_v2_filtering_landing_type
const (
	DPA_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "DPA"
	LINK_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LINK"
	ARTICLE_ReportAdGetV2FilteringLandingType   ReportAdGetV2FilteringLandingType = "ARTICLE"
	APP_ReportAdGetV2FilteringLandingType       ReportAdGetV2FilteringLandingType = "APP"
	STORE_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "STORE"
	QUICK_APP_ReportAdGetV2FilteringLandingType ReportAdGetV2FilteringLandingType = "QUICK_APP"
	SHOP_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "SHOP"
	AWEME_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "AWEME"
	LIVE_ReportAdGetV2FilteringLandingType      ReportAdGetV2FilteringLandingType = "LIVE"
	GOODS_ReportAdGetV2FilteringLandingType     ReportAdGetV2FilteringLandingType = "GOODS"
)

// Ptr returns reference to report_ad_get_v2_filtering_landing_type value
func (v ReportAdGetV2FilteringLandingType) Ptr() *ReportAdGetV2FilteringLandingType {
	return &v
}
