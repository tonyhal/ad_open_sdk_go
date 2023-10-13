/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2FilteringCampaignBudgetOptimization
type CampaignGetV2FilteringCampaignBudgetOptimization string

// List of campaign_get_v2_filtering_campaign_budget_optimization
const (
	ON_CampaignGetV2FilteringCampaignBudgetOptimization  CampaignGetV2FilteringCampaignBudgetOptimization = "ON"
	OFF_CampaignGetV2FilteringCampaignBudgetOptimization CampaignGetV2FilteringCampaignBudgetOptimization = "OFF"
)

// All allowed values of CampaignGetV2FilteringCampaignBudgetOptimization enum
var AllowedCampaignGetV2FilteringCampaignBudgetOptimizationEnumValues = []CampaignGetV2FilteringCampaignBudgetOptimization{
	"ON",
	"OFF",
}

// NewCampaignGetV2FilteringCampaignBudgetOptimizationFromValue returns a pointer to a valid CampaignGetV2FilteringCampaignBudgetOptimization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2FilteringCampaignBudgetOptimizationFromValue(v string) (*CampaignGetV2FilteringCampaignBudgetOptimization, error) {
	ev := CampaignGetV2FilteringCampaignBudgetOptimization(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2FilteringCampaignBudgetOptimization: valid values are %v", v, AllowedCampaignGetV2FilteringCampaignBudgetOptimizationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2FilteringCampaignBudgetOptimization) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2FilteringCampaignBudgetOptimizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_filtering_campaign_budget_optimization value
func (v CampaignGetV2FilteringCampaignBudgetOptimization) Ptr() *CampaignGetV2FilteringCampaignBudgetOptimization {
	return &v
}
