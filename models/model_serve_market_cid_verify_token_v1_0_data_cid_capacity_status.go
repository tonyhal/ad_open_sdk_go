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

// ServeMarketCidVerifyTokenV10DataCidCapacityStatus
type ServeMarketCidVerifyTokenV10DataCidCapacityStatus int64

// List of serve_market_cid_verify_token_v1.0_data_cid_capacity_status
const (
	Enum_1_ServeMarketCidVerifyTokenV10DataCidCapacityStatus   ServeMarketCidVerifyTokenV10DataCidCapacityStatus = 1
	Enum_2_ServeMarketCidVerifyTokenV10DataCidCapacityStatus   ServeMarketCidVerifyTokenV10DataCidCapacityStatus = 2
	Enum_3_ServeMarketCidVerifyTokenV10DataCidCapacityStatus   ServeMarketCidVerifyTokenV10DataCidCapacityStatus = 3
	Enum_6_ServeMarketCidVerifyTokenV10DataCidCapacityStatus   ServeMarketCidVerifyTokenV10DataCidCapacityStatus = 6
	Enum_100_ServeMarketCidVerifyTokenV10DataCidCapacityStatus ServeMarketCidVerifyTokenV10DataCidCapacityStatus = 100
)

// All allowed values of ServeMarketCidVerifyTokenV10DataCidCapacityStatus enum
var AllowedServeMarketCidVerifyTokenV10DataCidCapacityStatusEnumValues = []ServeMarketCidVerifyTokenV10DataCidCapacityStatus{
	1,
	2,
	3,
	6,
	100,
}

// NewServeMarketCidVerifyTokenV10DataCidCapacityStatusFromValue returns a pointer to a valid ServeMarketCidVerifyTokenV10DataCidCapacityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServeMarketCidVerifyTokenV10DataCidCapacityStatusFromValue(v int64) (*ServeMarketCidVerifyTokenV10DataCidCapacityStatus, error) {
	ev := ServeMarketCidVerifyTokenV10DataCidCapacityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServeMarketCidVerifyTokenV10DataCidCapacityStatus: valid values are %v", v, AllowedServeMarketCidVerifyTokenV10DataCidCapacityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServeMarketCidVerifyTokenV10DataCidCapacityStatus) IsValid() bool {
	for _, existing := range AllowedServeMarketCidVerifyTokenV10DataCidCapacityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to serve_market_cid_verify_token_v1.0_data_cid_capacity_status value
func (v ServeMarketCidVerifyTokenV10DataCidCapacityStatus) Ptr() *ServeMarketCidVerifyTokenV10DataCidCapacityStatus {
	return &v
}
