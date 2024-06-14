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

// PromotionListV30FilteringStatus
type PromotionListV30FilteringStatus string

// List of promotion_list_v3.0_filtering_status
const (
	ALL_PromotionListV30FilteringStatus                       PromotionListV30FilteringStatus = "ALL"
	AUDIT_PromotionListV30FilteringStatus                     PromotionListV30FilteringStatus = "AUDIT"
	AUDIT_DENY_PromotionListV30FilteringStatus                PromotionListV30FilteringStatus = "AUDIT_DENY"
	AWEME_ACCOUNT_DISABLED_PromotionListV30FilteringStatus    PromotionListV30FilteringStatus = "AWEME_ACCOUNT_DISABLED"
	AWEME_ANCHOR_DISABLED_PromotionListV30FilteringStatus     PromotionListV30FilteringStatus = "AWEME_ANCHOR_DISABLED"
	DELETED_PromotionListV30FilteringStatus                   PromotionListV30FilteringStatus = "DELETED"
	DISABLED_PromotionListV30FilteringStatus                  PromotionListV30FilteringStatus = "DISABLED"
	DISABLE_BY_QUOTA_PromotionListV30FilteringStatus          PromotionListV30FilteringStatus = "DISABLE_BY_QUOTA"
	FROZEN_PromotionListV30FilteringStatus                    PromotionListV30FilteringStatus = "FROZEN"
	LIVE_ROOM_OFF_PromotionListV30FilteringStatus             PromotionListV30FilteringStatus = "LIVE_ROOM_OFF"
	NOT_DELETED_PromotionListV30FilteringStatus               PromotionListV30FilteringStatus = "NOT_DELETED"
	NO_SCHEDULE_PromotionListV30FilteringStatus               PromotionListV30FilteringStatus = "NO_SCHEDULE"
	OFFLINE_BALANCE_PromotionListV30FilteringStatus           PromotionListV30FilteringStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_PromotionListV30FilteringStatus            PromotionListV30FilteringStatus = "OFFLINE_BUDGET"
	OK_PromotionListV30FilteringStatus                        PromotionListV30FilteringStatus = "OK"
	PREOFFLINE_BUDGET_PromotionListV30FilteringStatus         PromotionListV30FilteringStatus = "PREOFFLINE_BUDGET"
	PRODUCT_OFFLINE_PromotionListV30FilteringStatus           PromotionListV30FilteringStatus = "PRODUCT_OFFLINE"
	PROJECT_DISABLED_PromotionListV30FilteringStatus          PromotionListV30FilteringStatus = "PROJECT_DISABLED"
	PROJECT_OFFLINE_BUDGET_PromotionListV30FilteringStatus    PromotionListV30FilteringStatus = "PROJECT_OFFLINE_BUDGET"
	PROJECT_PREOFFLINE_BUDGET_PromotionListV30FilteringStatus PromotionListV30FilteringStatus = "PROJECT_PREOFFLINE_BUDGET"
	REAUDIT_PromotionListV30FilteringStatus                   PromotionListV30FilteringStatus = "REAUDIT"
	TIME_DONE_PromotionListV30FilteringStatus                 PromotionListV30FilteringStatus = "TIME_DONE"
	TIME_NO_REACH_PromotionListV30FilteringStatus             PromotionListV30FilteringStatus = "TIME_NO_REACH"
)

// All allowed values of PromotionListV30FilteringStatus enum
var AllowedPromotionListV30FilteringStatusEnumValues = []PromotionListV30FilteringStatus{
	"ALL",
	"AUDIT",
	"AUDIT_DENY",
	"AWEME_ACCOUNT_DISABLED",
	"AWEME_ANCHOR_DISABLED",
	"DELETED",
	"DISABLED",
	"DISABLE_BY_QUOTA",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NOT_DELETED",
	"NO_SCHEDULE",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"OK",
	"PREOFFLINE_BUDGET",
	"PRODUCT_OFFLINE",
	"PROJECT_DISABLED",
	"PROJECT_OFFLINE_BUDGET",
	"PROJECT_PREOFFLINE_BUDGET",
	"REAUDIT",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewPromotionListV30FilteringStatusFromValue returns a pointer to a valid PromotionListV30FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringStatusFromValue(v string) (*PromotionListV30FilteringStatus, error) {
	ev := PromotionListV30FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringStatus: valid values are %v", v, AllowedPromotionListV30FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_status value
func (v PromotionListV30FilteringStatus) Ptr() *PromotionListV30FilteringStatus {
	return &v
}
