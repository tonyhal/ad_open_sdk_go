/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListStatusSecond
type PromotionListV30DataListStatusSecond string

// List of promotion_list_v3.0_data_list_status_second
const (
	AUDIT_PromotionListV30DataListStatusSecond                    PromotionListV30DataListStatusSecond = "AUDIT"
	AUDIT_DENY_PromotionListV30DataListStatusSecond               PromotionListV30DataListStatusSecond = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_PromotionListV30DataListStatusSecond   PromotionListV30DataListStatusSecond = "AWEME_ACCOUNT_DISABLED"
	AWEME_ANCHOR_DISABLED_PromotionListV30DataListStatusSecond    PromotionListV30DataListStatusSecond = "AWEME_ANCHOR_DISABLED"
	BALANCE_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond   PromotionListV30DataListStatusSecond = "BALANCE_OFFLINE_BUDGET"
	DISABLED_PromotionListV30DataListStatusSecond                 PromotionListV30DataListStatusSecond = "DISABLED"
	DISABLE_BY_QUOTA_PromotionListV30DataListStatusSecond         PromotionListV30DataListStatusSecond = "DISABLE_BY_QUOTA"
	LIVE_ROOM_OFF_PromotionListV30DataListStatusSecond            PromotionListV30DataListStatusSecond = "LIVE_ROOM_OFF"
	NO_SCHEDULE_PromotionListV30DataListStatusSecond              PromotionListV30DataListStatusSecond = "NO_SCHEDULE"
	OFFLINE_BALANCE_PromotionListV30DataListStatusSecond          PromotionListV30DataListStatusSecond = "OFFLINE_BALANCE"
	PRODUCT_OFFLINE_PromotionListV30DataListStatusSecond          PromotionListV30DataListStatusSecond = "PRODUCT_OFFLINE"
	PROJECT_DISABLED_PromotionListV30DataListStatusSecond         PromotionListV30DataListStatusSecond = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond   PromotionListV30DataListStatusSecond = "PROJECT_OFFLINE_BUDGET"
	PROMOTION_OFFLINE_BUDGET_PromotionListV30DataListStatusSecond PromotionListV30DataListStatusSecond = "PROMOTION_OFFLINE_BUDGET"
	REAUDIT_PromotionListV30DataListStatusSecond                  PromotionListV30DataListStatusSecond = "REAUDIT"
	TIME_NO_REACH_PromotionListV30DataListStatusSecond            PromotionListV30DataListStatusSecond = "TIME_NO_REACH"
)

// All allowed values of PromotionListV30DataListStatusSecond enum
var AllowedPromotionListV30DataListStatusSecondEnumValues = []PromotionListV30DataListStatusSecond{
	"AUDIT",
	"AUDIT_DENY",
	"AWEME_ACCOUNT_DISABLED",
	"AWEME_ANCHOR_DISABLED",
	"BALANCE_OFFLINE_BUDGET",
	"DISABLED",
	"DISABLE_BY_QUOTA",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_BALANCE",
	"PRODUCT_OFFLINE",
	"PROJECT_DISABLED",
	"PROJECT_OFFLINE_BUDGET",
	"PROMOTION_OFFLINE_BUDGET",
	"REAUDIT",
	"TIME_NO_REACH",
}

// NewPromotionListV30DataListStatusSecondFromValue returns a pointer to a valid PromotionListV30DataListStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListStatusSecondFromValue(v string) (*PromotionListV30DataListStatusSecond, error) {
	ev := PromotionListV30DataListStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListStatusSecond: valid values are %v", v, AllowedPromotionListV30DataListStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListStatusSecond) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_status_second value
func (v PromotionListV30DataListStatusSecond) Ptr() *PromotionListV30DataListStatusSecond {
	return &v
}
