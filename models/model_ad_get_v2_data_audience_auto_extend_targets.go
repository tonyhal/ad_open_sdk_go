/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceAutoExtendTargets
type AdGetV2DataAudienceAutoExtendTargets string

// List of ad_get_v2_data_audience_auto_extend_targets
const (
	AGE_AdGetV2DataAudienceAutoExtendTargets             AdGetV2DataAudienceAutoExtendTargets = "AGE"
	REGION_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "REGION"
	INTEREST_TAG_AdGetV2DataAudienceAutoExtendTargets    AdGetV2DataAudienceAutoExtendTargets = "INTEREST_TAG"
	GENDER_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "GENDER"
	CUSTOM_AUDIENCE_AdGetV2DataAudienceAutoExtendTargets AdGetV2DataAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	AD_TAG_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "AD_TAG"
)

// Ptr returns reference to ad_get_v2_data_audience_auto_extend_targets value
func (v AdGetV2DataAudienceAutoExtendTargets) Ptr() *AdGetV2DataAudienceAutoExtendTargets {
	return &v
}
