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

// SpiTaskGetV2Status
type SpiTaskGetV2Status string

// List of spi_task_get_v2_status
const (
	ALL_SpiTaskGetV2Status     SpiTaskGetV2Status = "ALL"
	WAITING_SpiTaskGetV2Status SpiTaskGetV2Status = "WAITING"
	FAIL_SpiTaskGetV2Status    SpiTaskGetV2Status = "FAIL"
	SUCCESS_SpiTaskGetV2Status SpiTaskGetV2Status = "SUCCESS"
)

// All allowed values of SpiTaskGetV2Status enum
var AllowedSpiTaskGetV2StatusEnumValues = []SpiTaskGetV2Status{
	"ALL",
	"WAITING",
	"FAIL",
	"SUCCESS",
}

// NewSpiTaskGetV2StatusFromValue returns a pointer to a valid SpiTaskGetV2Status
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSpiTaskGetV2StatusFromValue(v string) (*SpiTaskGetV2Status, error) {
	ev := SpiTaskGetV2Status(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SpiTaskGetV2Status: valid values are %v", v, AllowedSpiTaskGetV2StatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SpiTaskGetV2Status) IsValid() bool {
	for _, existing := range AllowedSpiTaskGetV2StatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to spi_task_get_v2_status value
func (v SpiTaskGetV2Status) Ptr() *SpiTaskGetV2Status {
	return &v
}
