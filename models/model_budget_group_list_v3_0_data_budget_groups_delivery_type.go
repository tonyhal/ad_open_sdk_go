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

// BudgetGroupListV30DataBudgetGroupsDeliveryType
type BudgetGroupListV30DataBudgetGroupsDeliveryType string

// List of budget_group_list_v3.0_data_budget_groups_delivery_type
const (
	MANUAL_BudgetGroupListV30DataBudgetGroupsDeliveryType     BudgetGroupListV30DataBudgetGroupsDeliveryType = "MANUAL"
	PROCEDURAL_BudgetGroupListV30DataBudgetGroupsDeliveryType BudgetGroupListV30DataBudgetGroupsDeliveryType = "PROCEDURAL"
)

// All allowed values of BudgetGroupListV30DataBudgetGroupsDeliveryType enum
var AllowedBudgetGroupListV30DataBudgetGroupsDeliveryTypeEnumValues = []BudgetGroupListV30DataBudgetGroupsDeliveryType{
	"MANUAL",
	"PROCEDURAL",
}

// NewBudgetGroupListV30DataBudgetGroupsDeliveryTypeFromValue returns a pointer to a valid BudgetGroupListV30DataBudgetGroupsDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30DataBudgetGroupsDeliveryTypeFromValue(v string) (*BudgetGroupListV30DataBudgetGroupsDeliveryType, error) {
	ev := BudgetGroupListV30DataBudgetGroupsDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30DataBudgetGroupsDeliveryType: valid values are %v", v, AllowedBudgetGroupListV30DataBudgetGroupsDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30DataBudgetGroupsDeliveryType) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30DataBudgetGroupsDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_data_budget_groups_delivery_type value
func (v BudgetGroupListV30DataBudgetGroupsDeliveryType) Ptr() *BudgetGroupListV30DataBudgetGroupsDeliveryType {
	return &v
}
