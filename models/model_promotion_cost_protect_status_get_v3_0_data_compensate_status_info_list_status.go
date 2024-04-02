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

// PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus
type PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus string

// List of promotion_cost_protect_status_get_v3.0_data_compensate_status_info_list_status
const (
	SUCCESS_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus = "SUCCESS"
	FAILED_PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus  PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus = "FAILED"
)

// All allowed values of PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus enum
var AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatusEnumValues = []PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus{
	"SUCCESS",
	"FAILED",
}

// NewPromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatusFromValue returns a pointer to a valid PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatusFromValue(v string) (*PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus, error) {
	ev := PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus: valid values are %v", v, AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus) IsValid() bool {
	for _, existing := range AllowedPromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_cost_protect_status_get_v3.0_data_compensate_status_info_list_status value
func (v PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus) Ptr() *PromotionCostProtectStatusGetV30DataCompensateStatusInfoListStatus {
	return &v
}
