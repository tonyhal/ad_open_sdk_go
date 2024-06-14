/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataEnableSmartSource
type CreativeCustomCreativeUpdateV2AdDataEnableSmartSource string

// List of creative_custom_creative_update_v2_ad_data_enable_smart_source
const (
	OFF_CreativeCustomCreativeUpdateV2AdDataEnableSmartSource CreativeCustomCreativeUpdateV2AdDataEnableSmartSource = "OFF"
	ON_CreativeCustomCreativeUpdateV2AdDataEnableSmartSource  CreativeCustomCreativeUpdateV2AdDataEnableSmartSource = "ON"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataEnableSmartSource enum
var AllowedCreativeCustomCreativeUpdateV2AdDataEnableSmartSourceEnumValues = []CreativeCustomCreativeUpdateV2AdDataEnableSmartSource{
	"OFF",
	"ON",
}

// NewCreativeCustomCreativeUpdateV2AdDataEnableSmartSourceFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataEnableSmartSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataEnableSmartSourceFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataEnableSmartSource, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataEnableSmartSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataEnableSmartSource: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataEnableSmartSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataEnableSmartSource) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataEnableSmartSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_enable_smart_source value
func (v CreativeCustomCreativeUpdateV2AdDataEnableSmartSource) Ptr() *CreativeCustomCreativeUpdateV2AdDataEnableSmartSource {
	return &v
}
