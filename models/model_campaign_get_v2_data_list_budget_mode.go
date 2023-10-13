/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2DataListBudgetMode
type CampaignGetV2DataListBudgetMode string

// List of campaign_get_v2_data_list_budget_mode
const (
	BUDGET_MODE_DAY_CampaignGetV2DataListBudgetMode      CampaignGetV2DataListBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_CampaignGetV2DataListBudgetMode    CampaignGetV2DataListBudgetMode = "BUDGET_MODE_TOTAL"
	BUDGET_MODE_INFINITE_CampaignGetV2DataListBudgetMode CampaignGetV2DataListBudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of CampaignGetV2DataListBudgetMode enum
var AllowedCampaignGetV2DataListBudgetModeEnumValues = []CampaignGetV2DataListBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
	"BUDGET_MODE_INFINITE",
}

// NewCampaignGetV2DataListBudgetModeFromValue returns a pointer to a valid CampaignGetV2DataListBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListBudgetModeFromValue(v string) (*CampaignGetV2DataListBudgetMode, error) {
	ev := CampaignGetV2DataListBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListBudgetMode: valid values are %v", v, AllowedCampaignGetV2DataListBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListBudgetMode) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_budget_mode value
func (v CampaignGetV2DataListBudgetMode) Ptr() *CampaignGetV2DataListBudgetMode {
	return &v
}