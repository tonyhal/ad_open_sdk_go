/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCreativeGetV2Filtering
type ReportCreativeGetV2Filtering struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	AdName *string `json:"ad_name,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	CampaignTypes []*ReportCreativeGetV2FilteringCampaignTypes `json:"campaign_types,omitempty"`
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
	//
	CreativeMaterialModes []*ReportCreativeGetV2FilteringCreativeMaterialModes `json:"creative_material_modes,omitempty"`
	//
	CreativeTitle *string `json:"creative_title,omitempty"`
	//
	DeliveryMode []*ReportCreativeGetV2FilteringDeliveryMode `json:"delivery_mode,omitempty"`
	ImageMode    *ReportCreativeGetV2FilteringImageMode      `json:"image_mode,omitempty"`
	//
	ImageModes []*ReportCreativeGetV2FilteringImageModes `json:"image_modes,omitempty"`
	//
	InventoryTypes []*ReportCreativeGetV2FilteringInventoryTypes `json:"inventory_types,omitempty"`
	LandingType    *ReportCreativeGetV2FilteringLandingType      `json:"landing_type,omitempty"`
	//
	LandingTypes []*ReportCreativeGetV2FilteringLandingTypes `json:"landing_types,omitempty"`
	Pricing      *ReportCreativeGetV2FilteringPricing        `json:"pricing,omitempty"`
	//
	PricingCategories []*ReportCreativeGetV2FilteringPricingCategories `json:"pricing_categories,omitempty"`
	//
	Pricings []*ReportCreativeGetV2FilteringPricings `json:"pricings,omitempty"`
	Status   *ReportCreativeGetV2FilteringStatus     `json:"status,omitempty"`
}
