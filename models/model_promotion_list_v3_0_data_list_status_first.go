/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListStatusFirst
type PromotionListV30DataListStatusFirst string

// List of promotion_list_v3.0_data_list_status_first
const (
	PROMOTION_STATUS_DELETED_PromotionListV30DataListStatusFirst PromotionListV30DataListStatusFirst = "PROMOTION_STATUS_DELETED"
	PROMOTION_STATUS_DISABLE_PromotionListV30DataListStatusFirst PromotionListV30DataListStatusFirst = "PROMOTION_STATUS_DISABLE"
	PROMOTION_STATUS_DONE_PromotionListV30DataListStatusFirst    PromotionListV30DataListStatusFirst = "PROMOTION_STATUS_DONE"
	PROMOTION_STATUS_ENABLE_PromotionListV30DataListStatusFirst  PromotionListV30DataListStatusFirst = "PROMOTION_STATUS_ENABLE"
	PROMOTION_STATUS_FROZEN_PromotionListV30DataListStatusFirst  PromotionListV30DataListStatusFirst = "PROMOTION_STATUS_FROZEN"
)

// All allowed values of PromotionListV30DataListStatusFirst enum
var AllowedPromotionListV30DataListStatusFirstEnumValues = []PromotionListV30DataListStatusFirst{
	"PROMOTION_STATUS_DELETED",
	"PROMOTION_STATUS_DISABLE",
	"PROMOTION_STATUS_DONE",
	"PROMOTION_STATUS_ENABLE",
	"PROMOTION_STATUS_FROZEN",
}

// NewPromotionListV30DataListStatusFirstFromValue returns a pointer to a valid PromotionListV30DataListStatusFirst
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListStatusFirstFromValue(v string) (*PromotionListV30DataListStatusFirst, error) {
	ev := PromotionListV30DataListStatusFirst(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListStatusFirst: valid values are %v", v, AllowedPromotionListV30DataListStatusFirstEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListStatusFirst) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListStatusFirstEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_status_first value
func (v PromotionListV30DataListStatusFirst) Ptr() *PromotionListV30DataListStatusFirst {
	return &v
}
