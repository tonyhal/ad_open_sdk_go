/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV30Filtering
type AudiencePackageGetV30Filtering struct {
	AdType        *AudiencePackageGetV30FilteringAdType        `json:"ad_type,omitempty"`
	DeliveryRange *AudiencePackageGetV30FilteringDeliveryRange `json:"delivery_range,omitempty"`
	LandingType   *AudiencePackageGetV30FilteringLandingType   `json:"landing_type,omitempty"`
	MarketingGoal *AudiencePackageGetV30FilteringMarketingGoal `json:"marketing_goal,omitempty"`
	//
	Name *string `json:"name,omitempty"`
}
