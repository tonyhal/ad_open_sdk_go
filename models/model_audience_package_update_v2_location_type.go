/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2LocationType
type AudiencePackageUpdateV2LocationType string

// List of audience_package_update_v2_location_type
const (
	CURRENT_AudiencePackageUpdateV2LocationType AudiencePackageUpdateV2LocationType = "CURRENT"
	ALL_AudiencePackageUpdateV2LocationType     AudiencePackageUpdateV2LocationType = "ALL"
	HOME_AudiencePackageUpdateV2LocationType    AudiencePackageUpdateV2LocationType = "HOME"
	TRAVEL_AudiencePackageUpdateV2LocationType  AudiencePackageUpdateV2LocationType = "TRAVEL"
)

// Ptr returns reference to audience_package_update_v2_location_type value
func (v AudiencePackageUpdateV2LocationType) Ptr() *AudiencePackageUpdateV2LocationType {
	return &v
}
