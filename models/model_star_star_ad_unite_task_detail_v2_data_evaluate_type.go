/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarStarAdUniteTaskDetailV2DataEvaluateType
type StarStarAdUniteTaskDetailV2DataEvaluateType string

// List of star_star_ad_unite_task_detail_v2_data_evaluate_type
const (
	ACTIVE_StarStarAdUniteTaskDetailV2DataEvaluateType         StarStarAdUniteTaskDetailV2DataEvaluateType = "ACTIVE"
	ACTIVE_PAY_StarStarAdUniteTaskDetailV2DataEvaluateType     StarStarAdUniteTaskDetailV2DataEvaluateType = "ACTIVE_PAY"
	DEEP_PURCHASE_StarStarAdUniteTaskDetailV2DataEvaluateType  StarStarAdUniteTaskDetailV2DataEvaluateType = "DEEP_PURCHASE"
	INSTALL_FINISH_StarStarAdUniteTaskDetailV2DataEvaluateType StarStarAdUniteTaskDetailV2DataEvaluateType = "INSTALL_FINISH"
	REGISTER_StarStarAdUniteTaskDetailV2DataEvaluateType       StarStarAdUniteTaskDetailV2DataEvaluateType = "REGISTER"
)

// All allowed values of StarStarAdUniteTaskDetailV2DataEvaluateType enum
var AllowedStarStarAdUniteTaskDetailV2DataEvaluateTypeEnumValues = []StarStarAdUniteTaskDetailV2DataEvaluateType{
	"ACTIVE",
	"ACTIVE_PAY",
	"DEEP_PURCHASE",
	"INSTALL_FINISH",
	"REGISTER",
}

// NewStarStarAdUniteTaskDetailV2DataEvaluateTypeFromValue returns a pointer to a valid StarStarAdUniteTaskDetailV2DataEvaluateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarStarAdUniteTaskDetailV2DataEvaluateTypeFromValue(v string) (*StarStarAdUniteTaskDetailV2DataEvaluateType, error) {
	ev := StarStarAdUniteTaskDetailV2DataEvaluateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarStarAdUniteTaskDetailV2DataEvaluateType: valid values are %v", v, AllowedStarStarAdUniteTaskDetailV2DataEvaluateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarStarAdUniteTaskDetailV2DataEvaluateType) IsValid() bool {
	for _, existing := range AllowedStarStarAdUniteTaskDetailV2DataEvaluateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_star_ad_unite_task_detail_v2_data_evaluate_type value
func (v StarStarAdUniteTaskDetailV2DataEvaluateType) Ptr() *StarStarAdUniteTaskDetailV2DataEvaluateType {
	return &v
}