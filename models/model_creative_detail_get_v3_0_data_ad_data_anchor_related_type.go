/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataAdDataAnchorRelatedType
type CreativeDetailGetV30DataAdDataAnchorRelatedType string

// List of creative_detail_get_v3.0_data_ad_data_anchor_related_type
const (
	AUTO_CreativeDetailGetV30DataAdDataAnchorRelatedType   CreativeDetailGetV30DataAdDataAnchorRelatedType = "AUTO"
	OFF_CreativeDetailGetV30DataAdDataAnchorRelatedType    CreativeDetailGetV30DataAdDataAnchorRelatedType = "OFF"
	SELECT_CreativeDetailGetV30DataAdDataAnchorRelatedType CreativeDetailGetV30DataAdDataAnchorRelatedType = "SELECT"
)

// All allowed values of CreativeDetailGetV30DataAdDataAnchorRelatedType enum
var AllowedCreativeDetailGetV30DataAdDataAnchorRelatedTypeEnumValues = []CreativeDetailGetV30DataAdDataAnchorRelatedType{
	"AUTO",
	"OFF",
	"SELECT",
}

// NewCreativeDetailGetV30DataAdDataAnchorRelatedTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataAnchorRelatedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataAnchorRelatedTypeFromValue(v string) (*CreativeDetailGetV30DataAdDataAnchorRelatedType, error) {
	ev := CreativeDetailGetV30DataAdDataAnchorRelatedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataAnchorRelatedType: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataAnchorRelatedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataAnchorRelatedType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataAnchorRelatedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_anchor_related_type value
func (v CreativeDetailGetV30DataAdDataAnchorRelatedType) Ptr() *CreativeDetailGetV30DataAdDataAnchorRelatedType {
	return &v
}
