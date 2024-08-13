/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandAdGetV30ResponseDataDataReportsInner struct for ReportBrandAdGetV30ResponseDataDataReportsInner
type ReportBrandAdGetV30ResponseDataDataReportsInner struct {
	// 计划ID
	AdId *string `json:"ad_id,omitempty"`
	// 计划名称
	AdName *string `json:"ad_name,omitempty"`
	// 广告组ID
	CampaignId *string `json:"campaign_id,omitempty"`
	// 广告组名称
	CampaignName *string                                                    `json:"campaign_name,omitempty"`
	DataReport   *ReportBrandAdGetV30ResponseDataDataReportsInnerDataReport `json:"data_report,omitempty"`
	// 日期
	Date *string `json:"date,omitempty"`
}
