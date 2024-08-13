/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2CarrierRegionOptimize
type AudiencePackageCreateV2CarrierRegionOptimize string

// List of audience_package_create_v2_carrier_region_optimize
const (
	ON_AudiencePackageCreateV2CarrierRegionOptimize  AudiencePackageCreateV2CarrierRegionOptimize = "ON"
	OFF_AudiencePackageCreateV2CarrierRegionOptimize AudiencePackageCreateV2CarrierRegionOptimize = "OFF"
)

// All allowed values of AudiencePackageCreateV2CarrierRegionOptimize enum
var AllowedAudiencePackageCreateV2CarrierRegionOptimizeEnumValues = []AudiencePackageCreateV2CarrierRegionOptimize{
	"ON",
	"OFF",
}

// NewAudiencePackageCreateV2CarrierRegionOptimizeFromValue returns a pointer to a valid AudiencePackageCreateV2CarrierRegionOptimize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2CarrierRegionOptimizeFromValue(v string) (*AudiencePackageCreateV2CarrierRegionOptimize, error) {
	ev := AudiencePackageCreateV2CarrierRegionOptimize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2CarrierRegionOptimize: valid values are %v", v, AllowedAudiencePackageCreateV2CarrierRegionOptimizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2CarrierRegionOptimize) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2CarrierRegionOptimizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_carrier_region_optimize value
func (v AudiencePackageCreateV2CarrierRegionOptimize) Ptr() *AudiencePackageCreateV2CarrierRegionOptimize {
	return &v
}
