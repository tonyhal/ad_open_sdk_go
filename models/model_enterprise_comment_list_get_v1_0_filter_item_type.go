/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EnterpriseCommentListGetV10FilterItemType
type EnterpriseCommentListGetV10FilterItemType string

// List of enterprise_comment_list_get_v1.0_filter_item_type
const (
	ITEM_AD_EnterpriseCommentListGetV10FilterItemType      EnterpriseCommentListGetV10FilterItemType = "ITEM_AD"
	ITEM_CONTENT_EnterpriseCommentListGetV10FilterItemType EnterpriseCommentListGetV10FilterItemType = "ITEM_CONTENT"
)

// All allowed values of EnterpriseCommentListGetV10FilterItemType enum
var AllowedEnterpriseCommentListGetV10FilterItemTypeEnumValues = []EnterpriseCommentListGetV10FilterItemType{
	"ITEM_AD",
	"ITEM_CONTENT",
}

// NewEnterpriseCommentListGetV10FilterItemTypeFromValue returns a pointer to a valid EnterpriseCommentListGetV10FilterItemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseCommentListGetV10FilterItemTypeFromValue(v string) (*EnterpriseCommentListGetV10FilterItemType, error) {
	ev := EnterpriseCommentListGetV10FilterItemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseCommentListGetV10FilterItemType: valid values are %v", v, AllowedEnterpriseCommentListGetV10FilterItemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseCommentListGetV10FilterItemType) IsValid() bool {
	for _, existing := range AllowedEnterpriseCommentListGetV10FilterItemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_comment_list_get_v1.0_filter_item_type value
func (v EnterpriseCommentListGetV10FilterItemType) Ptr() *EnterpriseCommentListGetV10FilterItemType {
	return &v
}
