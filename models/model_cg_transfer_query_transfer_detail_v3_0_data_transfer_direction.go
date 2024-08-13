/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferQueryTransferDetailV30DataTransferDirection
type CgTransferQueryTransferDetailV30DataTransferDirection string

// List of cg_transfer_query_transfer_detail_v3.0_data_transfer_direction
const (
	TRANSFER_IN_CgTransferQueryTransferDetailV30DataTransferDirection  CgTransferQueryTransferDetailV30DataTransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferQueryTransferDetailV30DataTransferDirection CgTransferQueryTransferDetailV30DataTransferDirection = "TRANSFER_OUT"
)

// All allowed values of CgTransferQueryTransferDetailV30DataTransferDirection enum
var AllowedCgTransferQueryTransferDetailV30DataTransferDirectionEnumValues = []CgTransferQueryTransferDetailV30DataTransferDirection{
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewCgTransferQueryTransferDetailV30DataTransferDirectionFromValue returns a pointer to a valid CgTransferQueryTransferDetailV30DataTransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferDetailV30DataTransferDirectionFromValue(v string) (*CgTransferQueryTransferDetailV30DataTransferDirection, error) {
	ev := CgTransferQueryTransferDetailV30DataTransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferDetailV30DataTransferDirection: valid values are %v", v, AllowedCgTransferQueryTransferDetailV30DataTransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferDetailV30DataTransferDirection) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferDetailV30DataTransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_detail_v3.0_data_transfer_direction value
func (v CgTransferQueryTransferDetailV30DataTransferDirection) Ptr() *CgTransferQueryTransferDetailV30DataTransferDirection {
	return &v
}
