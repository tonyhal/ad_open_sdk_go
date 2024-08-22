/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageCreateV2District
type AudiencePackageCreateV2District string

// List of audience_package_create_v2_district
const (
	NONE_AudiencePackageCreateV2District              AudiencePackageCreateV2District = "NONE"
	BUSINESS_DISTRICT_AudiencePackageCreateV2District AudiencePackageCreateV2District = "BUSINESS_DISTRICT"
	REGION_AudiencePackageCreateV2District            AudiencePackageCreateV2District = "REGION"
	OVERSEA_AudiencePackageCreateV2District           AudiencePackageCreateV2District = "OVERSEA"
)

// Ptr returns reference to audience_package_create_v2_district value
func (v AudiencePackageCreateV2District) Ptr() *AudiencePackageCreateV2District {
	return &v
}
