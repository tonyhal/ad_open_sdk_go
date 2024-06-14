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

// StardeliveryTaskCreateResultGetV30DataStatus
type StardeliveryTaskCreateResultGetV30DataStatus string

// List of stardelivery_task_create_result_get_v3.0_data_status
const (
	FAILED_StardeliveryTaskCreateResultGetV30DataStatus  StardeliveryTaskCreateResultGetV30DataStatus = "FAILED"
	RUNNING_StardeliveryTaskCreateResultGetV30DataStatus StardeliveryTaskCreateResultGetV30DataStatus = "RUNNING"
	SUCCESS_StardeliveryTaskCreateResultGetV30DataStatus StardeliveryTaskCreateResultGetV30DataStatus = "SUCCESS"
)

// All allowed values of StardeliveryTaskCreateResultGetV30DataStatus enum
var AllowedStardeliveryTaskCreateResultGetV30DataStatusEnumValues = []StardeliveryTaskCreateResultGetV30DataStatus{
	"FAILED",
	"RUNNING",
	"SUCCESS",
}

// NewStardeliveryTaskCreateResultGetV30DataStatusFromValue returns a pointer to a valid StardeliveryTaskCreateResultGetV30DataStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStardeliveryTaskCreateResultGetV30DataStatusFromValue(v string) (*StardeliveryTaskCreateResultGetV30DataStatus, error) {
	ev := StardeliveryTaskCreateResultGetV30DataStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StardeliveryTaskCreateResultGetV30DataStatus: valid values are %v", v, AllowedStardeliveryTaskCreateResultGetV30DataStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StardeliveryTaskCreateResultGetV30DataStatus) IsValid() bool {
	for _, existing := range AllowedStardeliveryTaskCreateResultGetV30DataStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stardelivery_task_create_result_get_v3.0_data_status value
func (v StardeliveryTaskCreateResultGetV30DataStatus) Ptr() *StardeliveryTaskCreateResultGetV30DataStatus {
	return &v
}
