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

// AdCostProtectStatusGetV2DataListStatus
type AdCostProtectStatusGetV2DataListStatus string

// List of ad_cost_protect_status_get_v2_data_list_status
const (
	Enum_当前计划还未产生首次send或者不在赔付范围内_AdCostProtectStatusGetV2DataListStatus AdCostProtectStatusGetV2DataListStatus = "当前计划还未产生首次send或者不在赔付范围内"
	Enum_成本保障生效中_AdCostProtectStatusGetV2DataListStatus                 AdCostProtectStatusGetV2DataListStatus = "成本保障生效中"
	Enum_成本保障已失效_AdCostProtectStatusGetV2DataListStatus                 AdCostProtectStatusGetV2DataListStatus = "成本保障已失效"
	Enum_成本保障确认中_AdCostProtectStatusGetV2DataListStatus                 AdCostProtectStatusGetV2DataListStatus = "成本保障确认中"
	Enum_成本保障已到账_AdCostProtectStatusGetV2DataListStatus                 AdCostProtectStatusGetV2DataListStatus = "成本保障已到账"
	Enum_成本保障已结束_AdCostProtectStatusGetV2DataListStatus                 AdCostProtectStatusGetV2DataListStatus = "成本保障已结束"
)

// All allowed values of AdCostProtectStatusGetV2DataListStatus enum
var AllowedAdCostProtectStatusGetV2DataListStatusEnumValues = []AdCostProtectStatusGetV2DataListStatus{
	"当前计划还未产生首次send或者不在赔付范围内",
	"成本保障生效中",
	"成本保障已失效",
	"成本保障确认中",
	"成本保障已到账",
	"成本保障已结束",
}

// NewAdCostProtectStatusGetV2DataListStatusFromValue returns a pointer to a valid AdCostProtectStatusGetV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdCostProtectStatusGetV2DataListStatusFromValue(v string) (*AdCostProtectStatusGetV2DataListStatus, error) {
	ev := AdCostProtectStatusGetV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdCostProtectStatusGetV2DataListStatus: valid values are %v", v, AllowedAdCostProtectStatusGetV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdCostProtectStatusGetV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedAdCostProtectStatusGetV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_cost_protect_status_get_v2_data_list_status value
func (v AdCostProtectStatusGetV2DataListStatus) Ptr() *AdCostProtectStatusGetV2DataListStatus {
	return &v
}
