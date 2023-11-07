/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataStoreType
type AdGetV2DataStoreType string

// List of ad_get_v2_data_store_type
const (
	STORE_LANDING_AdGetV2DataStoreType     AdGetV2DataStoreType = "STORE_LANDING"
	STORE_NORMAL_AdGetV2DataStoreType      AdGetV2DataStoreType = "STORE_NORMAL"
	STORE_THIRT_PARTY_AdGetV2DataStoreType AdGetV2DataStoreType = "STORE_THIRT_PARTY"
	STORE_DOUYIN_AdGetV2DataStoreType      AdGetV2DataStoreType = "STORE_DOUYIN"
)

// All allowed values of AdGetV2DataStoreType enum
var AllowedAdGetV2DataStoreTypeEnumValues = []AdGetV2DataStoreType{
	"STORE_LANDING",
	"STORE_NORMAL",
	"STORE_THIRT_PARTY",
	"STORE_DOUYIN",
}

// NewAdGetV2DataStoreTypeFromValue returns a pointer to a valid AdGetV2DataStoreType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataStoreTypeFromValue(v string) (*AdGetV2DataStoreType, error) {
	ev := AdGetV2DataStoreType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataStoreType: valid values are %v", v, AllowedAdGetV2DataStoreTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataStoreType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataStoreTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_store_type value
func (v AdGetV2DataStoreType) Ptr() *AdGetV2DataStoreType {
	return &v
}
