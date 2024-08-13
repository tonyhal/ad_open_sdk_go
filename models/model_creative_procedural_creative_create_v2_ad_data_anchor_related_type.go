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

// CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType
type CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType string

// List of creative_procedural_creative_create_v2_ad_data_anchor_related_type
const (
	SELECT_CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType = "SELECT"
	AUTO_CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType   CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType = "AUTO"
	OFF_CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType    CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType = "OFF"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType enum
var AllowedCreativeProceduralCreativeCreateV2AdDataAnchorRelatedTypeEnumValues = []CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType{
	"SELECT",
	"AUTO",
	"OFF",
}

// NewCreativeProceduralCreativeCreateV2AdDataAnchorRelatedTypeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataAnchorRelatedTypeFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_anchor_related_type value
func (v CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType) Ptr() *CreativeProceduralCreativeCreateV2AdDataAnchorRelatedType {
	return &v
}
