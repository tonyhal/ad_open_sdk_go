/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30Operation
type PromotionCreateV30Operation string

// List of promotion_create_v3.0_operation
const (
	DISABLE_PromotionCreateV30Operation PromotionCreateV30Operation = "DISABLE"
	ENABLE_PromotionCreateV30Operation  PromotionCreateV30Operation = "ENABLE"
)

// All allowed values of PromotionCreateV30Operation enum
var AllowedPromotionCreateV30OperationEnumValues = []PromotionCreateV30Operation{
	"DISABLE",
	"ENABLE",
}

// NewPromotionCreateV30OperationFromValue returns a pointer to a valid PromotionCreateV30Operation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30OperationFromValue(v string) (*PromotionCreateV30Operation, error) {
	ev := PromotionCreateV30Operation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30Operation: valid values are %v", v, AllowedPromotionCreateV30OperationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30Operation) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30OperationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_operation value
func (v PromotionCreateV30Operation) Ptr() *PromotionCreateV30Operation {
	return &v
}
