/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2Ac
type AudiencePackageUpdateV2Ac string

// List of audience_package_update_v2_ac
const (
	Enum_4_G_AudiencePackageUpdateV2Ac AudiencePackageUpdateV2Ac = "4G"
	WIFI_AudiencePackageUpdateV2Ac     AudiencePackageUpdateV2Ac = "WIFI"
	Enum_2_G_AudiencePackageUpdateV2Ac AudiencePackageUpdateV2Ac = "2G"
	Enum_3_G_AudiencePackageUpdateV2Ac AudiencePackageUpdateV2Ac = "3G"
)

// Ptr returns reference to audience_package_update_v2_ac value
func (v AudiencePackageUpdateV2Ac) Ptr() *AudiencePackageUpdateV2Ac {
	return &v
}
