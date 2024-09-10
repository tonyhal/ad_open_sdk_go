/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2AutoExtendTargets
type AudiencePackageCreateV2AutoExtendTargets string

// List of audience_package_create_v2_auto_extend_targets
const (
	AGE_AudiencePackageCreateV2AutoExtendTargets             AudiencePackageCreateV2AutoExtendTargets = "AGE"
	REGION_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "REGION"
	INTEREST_TAG_AudiencePackageCreateV2AutoExtendTargets    AudiencePackageCreateV2AutoExtendTargets = "INTEREST_TAG"
	GENDER_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "GENDER"
	CUSTOM_AUDIENCE_AudiencePackageCreateV2AutoExtendTargets AudiencePackageCreateV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	AD_TAG_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "AD_TAG"
)

// Ptr returns reference to audience_package_create_v2_auto_extend_targets value
func (v AudiencePackageCreateV2AutoExtendTargets) Ptr() *AudiencePackageCreateV2AutoExtendTargets {
	return &v
}
