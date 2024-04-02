/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseDataAudiencePackagesInner struct for AudiencePackageGetV2ResponseDataAudiencePackagesInner
type AudiencePackageGetV2ResponseDataAudiencePackagesInner struct {
	//
	AdType *string `json:"ad_type,omitempty"`
	//
	AdvertiserId      *int64                                                         `json:"advertiser_id,omitempty"`
	AppBehaviorTarget *AudiencePackageGetV2DataAudiencePackagesAppBehaviorTarget     `json:"app_behavior_target,omitempty"`
	Audience          *AudiencePackageGetV2ResponseDataAudiencePackagesInnerAudience `json:"audience,omitempty"`
	//
	AudiencePackageId *int64 `json:"audience_package_id,omitempty"`
	//
	BindInfo              []*AudiencePackageGetV2ResponseDataAudiencePackagesInnerBindInfoInner `json:"bind_info,omitempty"`
	ConvertedTimeDuration *AudiencePackageGetV2DataAudiencePackagesConvertedTimeDuration        `json:"converted_time_duration,omitempty"`
	DeliveryRange         *AudiencePackageGetV2DataAudiencePackagesDeliveryRange                `json:"delivery_range,omitempty"`
	//
	Description     *string                                                  `json:"description,omitempty"`
	HideIfConverted *AudiencePackageGetV2DataAudiencePackagesHideIfConverted `json:"hide_if_converted,omitempty"`
	//
	HideIfExists  *int64                                                 `json:"hide_if_exists,omitempty"`
	LandingType   *AudiencePackageGetV2DataAudiencePackagesLandingType   `json:"landing_type,omitempty"`
	MarketingGoal *AudiencePackageGetV2DataAudiencePackagesMarketingGoal `json:"marketing_goal,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SmartTarget []string `json:"smart_target,omitempty"`
}
