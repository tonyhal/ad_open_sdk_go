/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2AutoExtendTargets
type AudiencePackageUpdateV2AutoExtendTargets string

// List of audience_package_update_v2_auto_extend_targets
const (
	AD_TAG_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "AD_TAG"
	AGE_AudiencePackageUpdateV2AutoExtendTargets             AudiencePackageUpdateV2AutoExtendTargets = "AGE"
	GENDER_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "GENDER"
	INTEREST_TAG_AudiencePackageUpdateV2AutoExtendTargets    AudiencePackageUpdateV2AutoExtendTargets = "INTEREST_TAG"
	REGION_AudiencePackageUpdateV2AutoExtendTargets          AudiencePackageUpdateV2AutoExtendTargets = "REGION"
	CUSTOM_AUDIENCE_AudiencePackageUpdateV2AutoExtendTargets AudiencePackageUpdateV2AutoExtendTargets = "CUSTOM_AUDIENCE"
)

// Ptr returns reference to audience_package_update_v2_auto_extend_targets value
func (v AudiencePackageUpdateV2AutoExtendTargets) Ptr() *AudiencePackageUpdateV2AutoExtendTargets {
	return &v
}
