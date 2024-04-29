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

// StardeliveryTaskCreateV30StarTaskExternalAction
type StardeliveryTaskCreateV30StarTaskExternalAction string

// List of stardelivery_task_create_v3.0_star_task_external_action
const (
	AD_CONVERT_TYPE_ACTIVE_StardeliveryTaskCreateV30StarTaskExternalAction          StardeliveryTaskCreateV30StarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_StardeliveryTaskCreateV30StarTaskExternalAction StardeliveryTaskCreateV30StarTaskExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_PAY_StardeliveryTaskCreateV30StarTaskExternalAction             StardeliveryTaskCreateV30StarTaskExternalAction = "AD_CONVERT_TYPE_PAY"
)

// All allowed values of StardeliveryTaskCreateV30StarTaskExternalAction enum
var AllowedStardeliveryTaskCreateV30StarTaskExternalActionEnumValues = []StardeliveryTaskCreateV30StarTaskExternalAction{
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_PAY",
}

// NewStardeliveryTaskCreateV30StarTaskExternalActionFromValue returns a pointer to a valid StardeliveryTaskCreateV30StarTaskExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskCreateV30StarTaskExternalActionFromValue(v string) (*StardeliveryTaskCreateV30StarTaskExternalAction, error) {
	ev := StardeliveryTaskCreateV30StarTaskExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskCreateV30StarTaskExternalAction: valid values are %v", v, AllowedStardeliveryTaskCreateV30StarTaskExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskCreateV30StarTaskExternalAction) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskCreateV30StarTaskExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_create_v3.0_star_task_external_action value
func (v StardeliveryTaskCreateV30StarTaskExternalAction) Ptr() *StardeliveryTaskCreateV30StarTaskExternalAction {
	return &v
}