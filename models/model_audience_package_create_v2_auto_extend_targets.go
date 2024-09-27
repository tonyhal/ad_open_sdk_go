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
	GENDER_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "GENDER"
	AD_TAG_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "AD_TAG"
	CUSTOM_AUDIENCE_AudiencePackageCreateV2AutoExtendTargets AudiencePackageCreateV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	REGION_AudiencePackageCreateV2AutoExtendTargets          AudiencePackageCreateV2AutoExtendTargets = "REGION"
	INTEREST_TAG_AudiencePackageCreateV2AutoExtendTargets    AudiencePackageCreateV2AutoExtendTargets = "INTEREST_TAG"
	AGE_AudiencePackageCreateV2AutoExtendTargets             AudiencePackageCreateV2AutoExtendTargets = "AGE"
)

// Ptr returns reference to audience_package_create_v2_auto_extend_targets value
func (v AudiencePackageCreateV2AutoExtendTargets) Ptr() *AudiencePackageCreateV2AutoExtendTargets {
	return &v
}
