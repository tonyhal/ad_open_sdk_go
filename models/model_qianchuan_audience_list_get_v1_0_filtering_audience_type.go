/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAudienceListGetV10FilteringAudienceType
type QianchuanAudienceListGetV10FilteringAudienceType string

// List of qianchuan_audience_list_get_v1.0_filtering_audience_type
const (
	BASIC_QianchuanAudienceListGetV10FilteringAudienceType  QianchuanAudienceListGetV10FilteringAudienceType = "BASIC"
	EXTEND_QianchuanAudienceListGetV10FilteringAudienceType QianchuanAudienceListGetV10FilteringAudienceType = "EXTEND"
	UPLOAD_QianchuanAudienceListGetV10FilteringAudienceType QianchuanAudienceListGetV10FilteringAudienceType = "UPLOAD"
)

// All allowed values of QianchuanAudienceListGetV10FilteringAudienceType enum
var AllowedQianchuanAudienceListGetV10FilteringAudienceTypeEnumValues = []QianchuanAudienceListGetV10FilteringAudienceType{
	"BASIC",
	"EXTEND",
	"UPLOAD",
}

// NewQianchuanAudienceListGetV10FilteringAudienceTypeFromValue returns a pointer to a valid QianchuanAudienceListGetV10FilteringAudienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAudienceListGetV10FilteringAudienceTypeFromValue(v string) (*QianchuanAudienceListGetV10FilteringAudienceType, error) {
	ev := QianchuanAudienceListGetV10FilteringAudienceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAudienceListGetV10FilteringAudienceType: valid values are %v", v, AllowedQianchuanAudienceListGetV10FilteringAudienceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAudienceListGetV10FilteringAudienceType) IsValid() bool {
	for _, existing := range AllowedQianchuanAudienceListGetV10FilteringAudienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_audience_list_get_v1.0_filtering_audience_type value
func (v QianchuanAudienceListGetV10FilteringAudienceType) Ptr() *QianchuanAudienceListGetV10FilteringAudienceType {
	return &v
}
