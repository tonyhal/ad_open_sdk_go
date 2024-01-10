/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BudgetGroupCreateV30CampaignType
type BudgetGroupCreateV30CampaignType string

// List of budget_group_create_v3.0_campaign_type
const (
	ALL_BudgetGroupCreateV30CampaignType    BudgetGroupCreateV30CampaignType = "ALL"
	SEARCH_BudgetGroupCreateV30CampaignType BudgetGroupCreateV30CampaignType = "SEARCH"
)

// All allowed values of BudgetGroupCreateV30CampaignType enum
var AllowedBudgetGroupCreateV30CampaignTypeEnumValues = []BudgetGroupCreateV30CampaignType{
	"ALL",
	"SEARCH",
}

// NewBudgetGroupCreateV30CampaignTypeFromValue returns a pointer to a valid BudgetGroupCreateV30CampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupCreateV30CampaignTypeFromValue(v string) (*BudgetGroupCreateV30CampaignType, error) {
	ev := BudgetGroupCreateV30CampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupCreateV30CampaignType: valid values are %v", v, AllowedBudgetGroupCreateV30CampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupCreateV30CampaignType) IsValid() bool {
	for _, existing := range AllowedBudgetGroupCreateV30CampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_create_v3.0_campaign_type value
func (v BudgetGroupCreateV30CampaignType) Ptr() *BudgetGroupCreateV30CampaignType {
	return &v
}
