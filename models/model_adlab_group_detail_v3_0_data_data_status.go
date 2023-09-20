/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataStatus
type AdlabGroupDetailV30DataDataStatus string

// List of adlab_group_detail_v3.0_data_data_status
const (
	ADV_BUDGET_EXCEED_AdlabGroupDetailV30DataDataStatus      AdlabGroupDetailV30DataDataStatus = "ADV_BUDGET_EXCEED"
	AUDIT_FAILED_AdlabGroupDetailV30DataDataStatus           AdlabGroupDetailV30DataDataStatus = "AUDIT_FAILED"
	AUTO_STOP_AdlabGroupDetailV30DataDataStatus              AdlabGroupDetailV30DataDataStatus = "AUTO_STOP"
	BALANCE_EXCEED_AdlabGroupDetailV30DataDataStatus         AdlabGroupDetailV30DataDataStatus = "BALANCE_EXCEED"
	CAMPAIGN_DELETED_AdlabGroupDetailV30DataDataStatus       AdlabGroupDetailV30DataDataStatus = "CAMPAIGN_DELETED"
	CAMPAIGN_DISABLE_AdlabGroupDetailV30DataDataStatus       AdlabGroupDetailV30DataDataStatus = "CAMPAIGN_DISABLE"
	CREATING_AdlabGroupDetailV30DataDataStatus               AdlabGroupDetailV30DataDataStatus = "CREATING"
	DELETED_AdlabGroupDetailV30DataDataStatus                AdlabGroupDetailV30DataDataStatus = "DELETED"
	DELIVERY_OK_AdlabGroupDetailV30DataDataStatus            AdlabGroupDetailV30DataDataStatus = "DELIVERY_OK"
	DISABLE_AdlabGroupDetailV30DataDataStatus                AdlabGroupDetailV30DataDataStatus = "DISABLE"
	GROUP_GOODS_CLOSE_AdlabGroupDetailV30DataDataStatus      AdlabGroupDetailV30DataDataStatus = "GROUP_GOODS_CLOSE"
	INITIALIZING_AdlabGroupDetailV30DataDataStatus           AdlabGroupDetailV30DataDataStatus = "INITIALIZING"
	INIT_FAILED_AdlabGroupDetailV30DataDataStatus            AdlabGroupDetailV30DataDataStatus = "INIT_FAILED"
	INVALID_STATUS_AdlabGroupDetailV30DataDataStatus         AdlabGroupDetailV30DataDataStatus = "INVALID_STATUS"
	LACK_OF_POSTERIOR_DATA_AdlabGroupDetailV30DataDataStatus AdlabGroupDetailV30DataDataStatus = "LACK_OF_POSTERIOR_DATA"
	LIVE_CAN_NOT_LAUNCH_AdlabGroupDetailV30DataDataStatus    AdlabGroupDetailV30DataDataStatus = "LIVE_CAN_NOT_LAUNCH"
	NO_SCHEDULE_AdlabGroupDetailV30DataDataStatus            AdlabGroupDetailV30DataDataStatus = "NO_SCHEDULE"
	OUT_OF_BUDGET_AdlabGroupDetailV30DataDataStatus          AdlabGroupDetailV30DataDataStatus = "OUT_OF_BUDGET"
	OUT_OF_CREATIVE_AdlabGroupDetailV30DataDataStatus        AdlabGroupDetailV30DataDataStatus = "OUT_OF_CREATIVE"
	OUT_OF_DAILY_BUDGET_AdlabGroupDetailV30DataDataStatus    AdlabGroupDetailV30DataDataStatus = "OUT_OF_DAILY_BUDGET"
	OUT_OF_DAILY_TIME_AdlabGroupDetailV30DataDataStatus      AdlabGroupDetailV30DataDataStatus = "OUT_OF_DAILY_TIME"
	OUT_OF_QUOTA_AdlabGroupDetailV30DataDataStatus           AdlabGroupDetailV30DataDataStatus = "OUT_OF_QUOTA"
	OUT_OF_TIME_AdlabGroupDetailV30DataDataStatus            AdlabGroupDetailV30DataDataStatus = "OUT_OF_TIME"
	READY_AdlabGroupDetailV30DataDataStatus                  AdlabGroupDetailV30DataDataStatus = "READY"
	RECOVERABLE_ERROR_AdlabGroupDetailV30DataDataStatus      AdlabGroupDetailV30DataDataStatus = "RECOVERABLE_ERROR"
	SYSTEM_ERROR_AdlabGroupDetailV30DataDataStatus           AdlabGroupDetailV30DataDataStatus = "SYSTEM_ERROR"
)

// All allowed values of AdlabGroupDetailV30DataDataStatus enum
var AllowedAdlabGroupDetailV30DataDataStatusEnumValues = []AdlabGroupDetailV30DataDataStatus{
	"ADV_BUDGET_EXCEED",
	"AUDIT_FAILED",
	"AUTO_STOP",
	"BALANCE_EXCEED",
	"CAMPAIGN_DELETED",
	"CAMPAIGN_DISABLE",
	"CREATING",
	"DELETED",
	"DELIVERY_OK",
	"DISABLE",
	"GROUP_GOODS_CLOSE",
	"INITIALIZING",
	"INIT_FAILED",
	"INVALID_STATUS",
	"LACK_OF_POSTERIOR_DATA",
	"LIVE_CAN_NOT_LAUNCH",
	"NO_SCHEDULE",
	"OUT_OF_BUDGET",
	"OUT_OF_CREATIVE",
	"OUT_OF_DAILY_BUDGET",
	"OUT_OF_DAILY_TIME",
	"OUT_OF_QUOTA",
	"OUT_OF_TIME",
	"READY",
	"RECOVERABLE_ERROR",
	"SYSTEM_ERROR",
}

// NewAdlabGroupDetailV30DataDataStatusFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataStatusFromValue(v string) (*AdlabGroupDetailV30DataDataStatus, error) {
	ev := AdlabGroupDetailV30DataDataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataStatus: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataStatus) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_status value
func (v AdlabGroupDetailV30DataDataStatus) Ptr() *AdlabGroupDetailV30DataDataStatus {
	return &v
}