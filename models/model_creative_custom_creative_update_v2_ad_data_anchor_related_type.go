/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType
type CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType string

// List of creative_custom_creative_update_v2_ad_data_anchor_related_type
const (
	OFF_CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType    CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType = "OFF"
	SELECT_CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType = "SELECT"
	AUTO_CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType   CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType = "AUTO"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType enum
var AllowedCreativeCustomCreativeUpdateV2AdDataAnchorRelatedTypeEnumValues = []CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType{
	"OFF",
	"SELECT",
	"AUTO",
}

// NewCreativeCustomCreativeUpdateV2AdDataAnchorRelatedTypeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataAnchorRelatedTypeFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_anchor_related_type value
func (v CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType) Ptr() *CreativeCustomCreativeUpdateV2AdDataAnchorRelatedType {
	return &v
}
