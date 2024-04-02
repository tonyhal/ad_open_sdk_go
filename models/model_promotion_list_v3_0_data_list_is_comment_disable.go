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

// PromotionListV30DataListIsCommentDisable
type PromotionListV30DataListIsCommentDisable string

// List of promotion_list_v3.0_data_list_is_comment_disable
const (
	OFF_PromotionListV30DataListIsCommentDisable PromotionListV30DataListIsCommentDisable = "OFF"
	ON_PromotionListV30DataListIsCommentDisable  PromotionListV30DataListIsCommentDisable = "ON"
)

// All allowed values of PromotionListV30DataListIsCommentDisable enum
var AllowedPromotionListV30DataListIsCommentDisableEnumValues = []PromotionListV30DataListIsCommentDisable{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListIsCommentDisableFromValue returns a pointer to a valid PromotionListV30DataListIsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListIsCommentDisableFromValue(v string) (*PromotionListV30DataListIsCommentDisable, error) {
	ev := PromotionListV30DataListIsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListIsCommentDisable: valid values are %v", v, AllowedPromotionListV30DataListIsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListIsCommentDisable) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListIsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_is_comment_disable value
func (v PromotionListV30DataListIsCommentDisable) Ptr() *PromotionListV30DataListIsCommentDisable {
	return &v
}
