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

// AudiencePackageUpdateV2CarrierRegionOptimize
type AudiencePackageUpdateV2CarrierRegionOptimize string

// List of audience_package_update_v2_carrier_region_optimize
const (
	ON_AudiencePackageUpdateV2CarrierRegionOptimize  AudiencePackageUpdateV2CarrierRegionOptimize = "ON"
	OFF_AudiencePackageUpdateV2CarrierRegionOptimize AudiencePackageUpdateV2CarrierRegionOptimize = "OFF"
)

// All allowed values of AudiencePackageUpdateV2CarrierRegionOptimize enum
var AllowedAudiencePackageUpdateV2CarrierRegionOptimizeEnumValues = []AudiencePackageUpdateV2CarrierRegionOptimize{
	"ON",
	"OFF",
}

// NewAudiencePackageUpdateV2CarrierRegionOptimizeFromValue returns a pointer to a valid AudiencePackageUpdateV2CarrierRegionOptimize
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2CarrierRegionOptimizeFromValue(v string) (*AudiencePackageUpdateV2CarrierRegionOptimize, error) {
	ev := AudiencePackageUpdateV2CarrierRegionOptimize(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2CarrierRegionOptimize: valid values are %v", v, AllowedAudiencePackageUpdateV2CarrierRegionOptimizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2CarrierRegionOptimize) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2CarrierRegionOptimizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_carrier_region_optimize value
func (v AudiencePackageUpdateV2CarrierRegionOptimize) Ptr() *AudiencePackageUpdateV2CarrierRegionOptimize {
	return &v
}
