/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageUpdateV2DeviceType
type AudiencePackageUpdateV2DeviceType string

// List of audience_package_update_v2_device_type
const (
	PAD_AudiencePackageUpdateV2DeviceType    AudiencePackageUpdateV2DeviceType = "PAD"
	MOBILE_AudiencePackageUpdateV2DeviceType AudiencePackageUpdateV2DeviceType = "MOBILE"
)

// Ptr returns reference to audience_package_update_v2_device_type value
func (v AudiencePackageUpdateV2DeviceType) Ptr() *AudiencePackageUpdateV2DeviceType {
	return &v
}
